package emitter

import (
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/metric"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusEmitter struct {
	jobsScheduled  prometheus.Counter
	jobsScheduling prometheus.Gauge

	buildsStarted prometheus.Counter
	buildsRunning prometheus.Gauge

	checkBuildsStarted prometheus.Counter
	checkBuildsRunning prometheus.Gauge

	concurrentRequestsLimitHit *prometheus.CounterVec
	concurrentRequests         *prometheus.GaugeVec

	stepsWaiting         *prometheus.GaugeVec
	stepsWaitingDuration *prometheus.HistogramVec

	buildDurationsVec *prometheus.HistogramVec
	buildsAborted     prometheus.Counter
	buildsErrored     prometheus.Counter
	buildsFailed      prometheus.Counter
	buildsFinished    prometheus.Counter
	buildsFinishedVec *prometheus.CounterVec
	buildsSucceeded   prometheus.Counter

	checkBuildsAborted   prometheus.Counter
	checkBuildsErrored   prometheus.Counter
	checkBuildsFailed    prometheus.Counter
	checkBuildsFinished  prometheus.Counter
	checkBuildsSucceeded prometheus.Counter

	dbConnections  *prometheus.GaugeVec
	dbQueriesTotal prometheus.Counter

	errorLogs *prometheus.CounterVec

	httpRequestsDuration *prometheus.HistogramVec

	locksHeld *prometheus.GaugeVec

	// TODO: deprecate
	checksFinished *prometheus.CounterVec
	checksStarted  prometheus.Counter

	checksEnqueued prometheus.Counter

	volumesStreamed prometheus.Counter

	workerContainers        *prometheus.GaugeVec
	workerUnknownContainers *prometheus.GaugeVec
	workerVolumes           *prometheus.GaugeVec
	workerUnknownVolumes    *prometheus.GaugeVec
	workerTasks             *prometheus.GaugeVec
	workersRegistered       *prometheus.GaugeVec

	workerContainersLabels map[string]map[string]prometheus.Labels
	workerVolumesLabels    map[string]map[string]prometheus.Labels
	workerTasksLabels      map[string]map[string]prometheus.Labels
	workerLastSeen         map[string]time.Time
	mu                     sync.Mutex
}

type PrometheusConfig struct {
	BindIP   string `long:"prometheus-bind-ip" description:"IP to listen on to expose Prometheus metrics."`
	BindPort string `long:"prometheus-bind-port" description:"Port to listen on to expose Prometheus metrics."`
}

// The most natural data type to hold the labels is a set because each worker can have multiple but
// unique sets of labels. A set in Go is represented by a map[T]struct{}. Unfortunately, we cannot
// put prometheus.Labels inside a map[prometheus.Labels]struct{} because prometheus.Labels are not
// hashable. To work around this, we compute a string from the labels and use this as the keys of
// the map.
func serializeLabels(labels *prometheus.Labels) string {
	var (
		key   string
		names []string
	)
	for _, v := range *labels {
		names = append(names, v)
	}
	sort.Strings(names)
	key = strings.Join(names, "_")

	return key
}

func init() {
	metric.Metrics.RegisterEmitter(&PrometheusConfig{})
}

func (config *PrometheusConfig) Description() string { return "Prometheus" }
func (config *PrometheusConfig) IsConfigured() bool {
	return config.BindPort != "" && config.BindIP != ""
}
func (config *PrometheusConfig) bind() string {
	return fmt.Sprintf("%s:%s", config.BindIP, config.BindPort)
}

func (config *PrometheusConfig) NewEmitter() (metric.Emitter, error) {
	// error log metrics
	errorLogs := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "error",
			Name:      "logs",
			Help:      "Number of error logged",
		}, []string{"message"},
	)
	prometheus.MustRegister(errorLogs)

	// lock metrics
	locksHeld := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "concourse",
		Subsystem: "locks",
		Name:      "held",
		Help:      "Database locks held",
	}, []string{"type"})
	prometheus.MustRegister(locksHeld)

	// job metrics
	jobsScheduled := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "jobs",
		Name:      "scheduled_total",
		Help:      "Total number of Concourse jobs scheduled.",
	})
	prometheus.MustRegister(jobsScheduled)

	jobsScheduling := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "concourse",
		Subsystem: "jobs",
		Name:      "scheduling",
		Help:      "Number of Concourse jobs currently being scheduled.",
	})
	prometheus.MustRegister(jobsScheduling)

	// build metrics
	buildsStarted := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "started_total",
		Help:      "Total number of Concourse builds started.",
	})
	prometheus.MustRegister(buildsStarted)

	buildsRunning := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "running",
		Help:      "Number of Concourse builds currently running.",
	})
	prometheus.MustRegister(buildsRunning)

	checkBuildsStarted := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_started_total",
		Help:      "Total number of Concourse check builds started.",
	})
	prometheus.MustRegister(checkBuildsStarted)

	checkBuildsRunning := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_running",
		Help:      "Number of Concourse check builds currently running.",
	})
	prometheus.MustRegister(checkBuildsRunning)

	concurrentRequestsLimitHit := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "concurrent_requests",
		Name:      "limit_hit_total",
		Help:      "Total number of requests rejected because the server was already serving too many concurrent requests.",
	}, []string{"action"})
	prometheus.MustRegister(concurrentRequestsLimitHit)

	concurrentRequests := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "concourse",
		Name:      "concurrent_requests",
		Help:      "Number of concurrent requests being served by endpoints that have a specified limit of concurrent requests.",
	}, []string{"action"})
	prometheus.MustRegister(concurrentRequests)

	stepsWaiting := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "concourse",
		Subsystem: "steps",
		Name:      "waiting",
		Help:      "Number of Concourse build steps currently waiting.",
	}, []string{"platform", "teamId", "type", "workerTags"})
	prometheus.MustRegister(stepsWaiting)

	stepsWaitingDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "concourse",
		Subsystem: "steps",
		Name:      "wait_duration",
		Help:      "Elapsed time waiting for execution",
		Buckets:   []float64{10, 30, 60, 120, 300, 600, 1800, 2400, 3000, 3600},
	}, []string{"platform", "teamId", "type", "workerTags"})
	prometheus.MustRegister(stepsWaitingDuration)

	buildsFinished := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "finished_total",
		Help:      "Total number of Concourse builds finished.",
	})
	prometheus.MustRegister(buildsFinished)

	buildsSucceeded := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "succeeded_total",
		Help:      "Total number of Concourse builds succeeded.",
	})
	prometheus.MustRegister(buildsSucceeded)

	buildsErrored := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "errored_total",
		Help:      "Total number of Concourse builds errored.",
	})
	prometheus.MustRegister(buildsErrored)

	buildsFailed := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "failed_total",
		Help:      "Total number of Concourse builds failed.",
	})
	prometheus.MustRegister(buildsFailed)

	buildsAborted := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "aborted_total",
		Help:      "Total number of Concourse builds aborted.",
	})
	prometheus.MustRegister(buildsAborted)

	buildsFinishedVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "builds",
			Name:      "finished",
			Help:      "Count of builds finished across various dimensions.",
		},
		[]string{"team", "pipeline", "job", "status"},
	)
	prometheus.MustRegister(buildsFinishedVec)

	buildDurationsVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "concourse",
			Subsystem: "builds",
			Name:      "duration_seconds",
			Help:      "Build time in seconds",
			Buckets:   []float64{1, 60, 180, 300, 600, 900, 1200, 1800, 2700, 3600, 7200, 18000, 36000},
		},
		[]string{"team", "pipeline", "job"},
	)
	prometheus.MustRegister(buildDurationsVec)

	checkBuildsFinished := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_finished_total",
		Help:      "Total number of Concourse check builds finished.",
	})
	prometheus.MustRegister(checkBuildsFinished)

	checkBuildsSucceeded := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_succeeded_total",
		Help:      "Total number of Concourse check builds succeeded.",
	})
	prometheus.MustRegister(checkBuildsSucceeded)

	checkBuildsErrored := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_errored_total",
		Help:      "Total number of Concourse check builds errored.",
	})
	prometheus.MustRegister(checkBuildsErrored)

	checkBuildsFailed := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_failed_total",
		Help:      "Total number of Concourse check builds failed.",
	})
	prometheus.MustRegister(checkBuildsFailed)

	checkBuildsAborted := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "builds",
		Name:      "check_aborted_total",
		Help:      "Total number of Concourse check builds aborted.",
	})
	prometheus.MustRegister(checkBuildsAborted)

	// worker metrics
	workerContainers := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "containers",
			Help:      "Number of containers per worker",
		},
		[]string{"worker", "platform", "team", "tags"},
	)
	prometheus.MustRegister(workerContainers)

	workerUnknownContainers := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "unknown_containers",
			Help:      "Number of unknown containers found on worker",
		},
		[]string{"worker"},
	)
	prometheus.MustRegister(workerUnknownContainers)

	workerVolumes := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "volumes",
			Help:      "Number of volumes per worker",
		},
		[]string{"worker", "platform", "team", "tags"},
	)
	prometheus.MustRegister(workerVolumes)

	workerUnknownVolumes := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "unknown_volumes",
			Help:      "Number of unknown volumes found on worker",
		},
		[]string{"worker"},
	)
	prometheus.MustRegister(workerUnknownVolumes)

	workerTasks := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "tasks",
			Help:      "Number of active tasks per worker",
		},
		[]string{"worker", "platform"},
	)
	prometheus.MustRegister(workerTasks)

	workersRegistered := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "workers",
			Name:      "registered",
			Help:      "Number of workers per state as seen by the database",
		},
		[]string{"state"},
	)
	prometheus.MustRegister(workersRegistered)

	// http metrics
	httpRequestsDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "concourse",
			Subsystem: "http_responses",
			Name:      "duration_seconds",
			Help:      "Response time in seconds",
		},
		[]string{"method", "route", "status"},
	)
	prometheus.MustRegister(httpRequestsDuration)

	dbQueriesTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "concourse",
		Subsystem: "db",
		Name:      "queries_total",
		Help:      "Total number of database Concourse database queries",
	})
	prometheus.MustRegister(dbQueriesTotal)

	dbConnections := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "concourse",
			Subsystem: "db",
			Name:      "connections",
			Help:      "Current number of concourse database connections",
		},
		[]string{"dbname"},
	)
	prometheus.MustRegister(dbConnections)

	checksFinished := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "lidar",
			Name:      "checks_finished_total",
			Help:      "Total number of checks finished. DEPRECATED: use concourse_builds_check_finished_total instead.",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(checksFinished)

	checksStarted := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "lidar",
			Name:      "checks_started_total",
			Help:      "Total number of checks started. DEPRECATED: use concourse_builds_check_started_total instead.",
		},
	)
	prometheus.MustRegister(checksStarted)

	checksEnqueued := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "lidar",
			Name:      "checks_enqueued_total",
			Help:      "Total number of checks enqueued",
		},
	)
	prometheus.MustRegister(checksEnqueued)

	volumesStreamed := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "concourse",
			Subsystem: "volumes",
			Name:      "volumes_streamed",
			Help:      "Total number of volumes streamed from one worker to the other",
		},
	)
	prometheus.MustRegister(volumesStreamed)

	listener, err := net.Listen("tcp", config.bind())
	if err != nil {
		return nil, err
	}

	go http.Serve(listener, promhttp.Handler())

	emitter := &PrometheusEmitter{
		jobsScheduled:  jobsScheduled,
		jobsScheduling: jobsScheduling,

		buildsStarted: buildsStarted,
		buildsRunning: buildsRunning,

		checkBuildsStarted: checkBuildsStarted,
		checkBuildsRunning: checkBuildsRunning,

		concurrentRequestsLimitHit: concurrentRequestsLimitHit,
		concurrentRequests:         concurrentRequests,

		stepsWaiting:         stepsWaiting,
		stepsWaitingDuration: stepsWaitingDuration,

		buildDurationsVec: buildDurationsVec,
		buildsAborted:     buildsAborted,
		buildsErrored:     buildsErrored,
		buildsFailed:      buildsFailed,
		buildsFinished:    buildsFinished,
		buildsFinishedVec: buildsFinishedVec,
		buildsSucceeded:   buildsSucceeded,

		checkBuildsAborted:   checkBuildsAborted,
		checkBuildsErrored:   checkBuildsErrored,
		checkBuildsFailed:    checkBuildsFailed,
		checkBuildsFinished:  checkBuildsFinished,
		checkBuildsSucceeded: checkBuildsSucceeded,

		dbConnections:  dbConnections,
		dbQueriesTotal: dbQueriesTotal,

		errorLogs: errorLogs,

		httpRequestsDuration: httpRequestsDuration,

		locksHeld: locksHeld,

		// TODO: deprecate
		checksFinished: checksFinished,
		checksStarted:  checksStarted,

		checksEnqueued: checksEnqueued,

		workerContainers:        workerContainers,
		workersRegistered:       workersRegistered,
		workerContainersLabels:  map[string]map[string]prometheus.Labels{},
		workerVolumesLabels:     map[string]map[string]prometheus.Labels{},
		workerTasksLabels:       map[string]map[string]prometheus.Labels{},
		workerLastSeen:          map[string]time.Time{},
		workerVolumes:           workerVolumes,
		workerTasks:             workerTasks,
		workerUnknownContainers: workerUnknownContainers,
		workerUnknownVolumes:    workerUnknownVolumes,

		volumesStreamed: volumesStreamed,
	}
	go emitter.periodicMetricGC()

	return emitter, nil
}

// Emit processes incoming metrics.
// In order to provide idiomatic Prometheus metrics, we'll have to convert the various
// Event types (differentiated by the less-than-ideal string Name field) into different
// Prometheus metrics.
func (emitter *PrometheusEmitter) Emit(logger lager.Logger, event metric.Event) {

	//update last seen counters, used to gc stale timeseries
	emitter.updateLastSeen(event)

	switch event.Name {
	case "error log":
		emitter.errorLogsMetric(logger, event)
	case "lock held":
		emitter.lock(logger, event)
	case "jobs scheduled":
		emitter.jobsScheduled.Add(event.Value)
	case "jobs scheduling":
		emitter.jobsScheduling.Set(event.Value)
	case "builds started":
		emitter.buildsStarted.Add(event.Value)
	case "builds running":
		emitter.buildsRunning.Set(event.Value)
	case "check builds started":
		emitter.checkBuildsStarted.Add(event.Value)
	case "check builds running":
		emitter.checkBuildsRunning.Set(event.Value)
	case "concurrent requests limit hit":
		emitter.concurrentRequestsLimitHit.WithLabelValues(event.Attributes["action"]).Add(event.Value)
	case "concurrent requests":
		emitter.concurrentRequests.
			WithLabelValues(event.Attributes["action"]).Set(event.Value)
	case "steps waiting":
		emitter.stepsWaiting.
			WithLabelValues(
				event.Attributes["platform"],
				event.Attributes["teamId"],
				event.Attributes["type"],
				event.Attributes["workerTags"],
			).Set(event.Value)
	case "steps waiting duration":
		emitter.stepsWaitingDuration.
			WithLabelValues(
				event.Attributes["platform"],
				event.Attributes["teamId"],
				event.Attributes["type"],
				event.Attributes["workerTags"],
			).Observe(event.Value)
	case "build finished":
		emitter.buildFinishedMetrics(logger, event)
	case "check build finished":
		emitter.checkBuildFinishedMetrics(logger, event)
	case "worker containers":
		emitter.workerContainersMetric(logger, event)
	case "worker volumes":
		emitter.workerVolumesMetric(logger, event)
	case "worker unknown containers":
		emitter.workerUnknownContainersMetric(logger, event)
	case "worker unknown volumes":
		emitter.workerUnknownVolumesMetric(logger, event)
	case "worker tasks":
		emitter.workerTasksMetric(logger, event)
	case "worker state":
		emitter.workersRegisteredMetric(logger, event)
	case "http response time":
		emitter.httpResponseTimeMetrics(logger, event)
	case "database queries":
		emitter.databaseMetrics(logger, event)
	case "database connections":
		emitter.databaseMetrics(logger, event)
	// TODO: deprecate
	case "checks finished":
		emitter.checksFinished.WithLabelValues(event.Attributes["status"]).Add(event.Value)
	// TODO: deprecate
	case "checks started":
		emitter.checksStarted.Add(event.Value)
	case "checks enqueued":
		emitter.checksEnqueued.Add(event.Value)
	case "volumes streamed":
		emitter.volumesStreamed.Add(event.Value)
	default:
		// unless we have a specific metric, we do nothing
	}
}

func (emitter *PrometheusEmitter) lock(logger lager.Logger, event metric.Event) {
	lockType, exists := event.Attributes["type"]
	if !exists {
		logger.Error("failed-to-find-type-in-event", fmt.Errorf("expected type to exist in event.Attributes"))
		return
	}

	if event.Value == 1 {
		emitter.locksHeld.WithLabelValues(lockType).Inc()
	} else {
		emitter.locksHeld.WithLabelValues(lockType).Dec()
	}
}

func (emitter *PrometheusEmitter) errorLogsMetric(logger lager.Logger, event metric.Event) {
	message, exists := event.Attributes["message"]
	if !exists {
		logger.Error("failed-to-find-message-in-event",
			fmt.Errorf("expected team_name to exist in event.Attributes"))
		return
	}

	emitter.errorLogs.WithLabelValues(message).Inc()
}

func (emitter *PrometheusEmitter) buildFinishedMetrics(logger lager.Logger, event metric.Event) {
	// concourse_builds_finished_total
	emitter.buildsFinished.Inc()

	// concourse_builds_finished
	team, exists := event.Attributes["team_name"]
	if !exists {
		logger.Error("failed-to-find-team-name-in-event", fmt.Errorf("expected team_name to exist in event.Attributes"))
		return
	}

	pipeline, exists := event.Attributes["pipeline"]
	if !exists {
		logger.Error("failed-to-find-pipeline-in-event", fmt.Errorf("expected pipeline to exist in event.Attributes"))
		return
	}

	job, exists := event.Attributes["job"]
	if !exists {
		logger.Error("failed-to-find-job-in-event", fmt.Errorf("expected job to exist in event.Attributes"))
		return
	}

	buildStatus, exists := event.Attributes["build_status"]
	if !exists {
		logger.Error("failed-to-find-build_status-in-event", fmt.Errorf("expected build_status to exist in event.Attributes"))
		return
	}
	emitter.buildsFinishedVec.WithLabelValues(team, pipeline, job, buildStatus).Inc()

	// concourse_builds_(aborted|succeeded|failed|errored)_total
	switch buildStatus {
	case string(db.BuildStatusAborted):
		// concourse_builds_aborted_total
		emitter.buildsAborted.Inc()
	case string(db.BuildStatusSucceeded):
		// concourse_builds_succeeded_total
		emitter.buildsSucceeded.Inc()
	case string(db.BuildStatusFailed):
		// concourse_builds_failed_total
		emitter.buildsFailed.Inc()
	case string(db.BuildStatusErrored):
		// concourse_builds_errored_total
		emitter.buildsErrored.Inc()
	}

	// seconds are the standard prometheus base unit for time
	duration := event.Value / 1000
	emitter.buildDurationsVec.WithLabelValues(team, pipeline, job).Observe(duration)
}

func (emitter *PrometheusEmitter) checkBuildFinishedMetrics(logger lager.Logger, event metric.Event) {
	// concourse_builds_finished_total
	emitter.checkBuildsFinished.Inc()

	buildStatus, exists := event.Attributes["build_status"]
	if !exists {
		logger.Error("failed-to-find-check-build_status-in-event", fmt.Errorf("expected build_status to exist in event.Attributes"))
		return
	}

	// concourse_builds_(aborted|succeeded|failed|errored)_total
	switch buildStatus {
	case string(db.BuildStatusAborted):
		// concourse_builds_check_aborted_total
		emitter.checkBuildsAborted.Inc()
	case string(db.BuildStatusSucceeded):
		// concourse_builds_check_succeeded_total
		emitter.checkBuildsSucceeded.Inc()
	case string(db.BuildStatusFailed):
		// concourse_builds_check_failed_total
		emitter.checkBuildsFailed.Inc()
	case string(db.BuildStatusErrored):
		// concourse_builds_check_errored_total
		emitter.checkBuildsErrored.Inc()
	}
}

func (emitter *PrometheusEmitter) workerContainersMetric(logger lager.Logger, event metric.Event) {
	worker, exists := event.Attributes["worker"]
	if !exists {
		logger.Error("failed-to-find-worker-in-event", fmt.Errorf("expected worker to exist in event.Attributes"))
		return
	}
	platform, exists := event.Attributes["platform"]
	if !exists || platform == "" {
		logger.Error("failed-to-find-platform-in-event", fmt.Errorf("expected platform to exist in event.Attributes"))
		return
	}
	team, exists := event.Attributes["team_name"]
	if !exists {
		logger.Error("failed-to-find-team-name-in-event", fmt.Errorf("expected team_name to exist in event.Attributes"))
		return
	}
	tags := event.Attributes["tags"]

	labels := prometheus.Labels{
		"worker":   worker,
		"platform": platform,
		"team":     team,
		"tags":     tags,
	}
	key := serializeLabels(&labels)
	if emitter.workerContainersLabels[worker] == nil {
		emitter.workerContainersLabels[worker] = make(map[string]prometheus.Labels)
	}
	emitter.workerContainersLabels[worker][key] = labels
	emitter.workerContainers.With(emitter.workerContainersLabels[worker][key]).Set(event.Value)
}

func (emitter *PrometheusEmitter) workersRegisteredMetric(logger lager.Logger, event metric.Event) {
	state, exists := event.Attributes["state"]
	if !exists {
		logger.Error("failed-to-find-state-in-event", fmt.Errorf("expected state to exist in event.Attributes"))
		return
	}

	emitter.workersRegistered.WithLabelValues(state).Set(event.Value)
}

func (emitter *PrometheusEmitter) workerUnknownContainersMetric(logger lager.Logger, event metric.Event) {
	worker, exists := event.Attributes["worker"]
	if !exists {
		logger.Error("failed-to-find-worker-in-event", fmt.Errorf("expected worker to exist in event.Attributes"))
		return
	}

	labels := prometheus.Labels{
		"worker": worker,
	}

	key := serializeLabels(&labels)
	if emitter.workerContainersLabels[worker] == nil {
		emitter.workerContainersLabels[worker] = make(map[string]prometheus.Labels)
	}
	emitter.workerContainersLabels[worker][key] = labels
	emitter.workerUnknownContainers.With(emitter.workerContainersLabels[worker][key]).Set(event.Value)
}

func (emitter *PrometheusEmitter) workerVolumesMetric(logger lager.Logger, event metric.Event) {
	worker, exists := event.Attributes["worker"]
	if !exists {
		logger.Error("failed-to-find-worker-in-event", fmt.Errorf("expected worker to exist in event.Attributes"))
		return
	}
	platform, exists := event.Attributes["platform"]
	if !exists || platform == "" {
		logger.Error("failed-to-find-platform-in-event", fmt.Errorf("expected platform to exist in event.Attributes"))
		return
	}
	team, exists := event.Attributes["team_name"]
	if !exists {
		logger.Error("failed-to-find-team-name-in-event", fmt.Errorf("expected team_name to exist in event.Attributes"))
		return
	}
	tags := event.Attributes["tags"]

	labels := prometheus.Labels{
		"worker":   worker,
		"platform": platform,
		"team":     team,
		"tags":     tags,
	}
	key := serializeLabels(&labels)
	if emitter.workerVolumesLabels[worker] == nil {
		emitter.workerVolumesLabels[worker] = make(map[string]prometheus.Labels)
	}
	emitter.workerVolumesLabels[worker][key] = labels
	emitter.workerVolumes.With(emitter.workerVolumesLabels[worker][key]).Set(event.Value)
}

func (emitter *PrometheusEmitter) workerUnknownVolumesMetric(logger lager.Logger, event metric.Event) {
	worker, exists := event.Attributes["worker"]
	if !exists {
		logger.Error("failed-to-find-worker-in-event", fmt.Errorf("expected worker to exist in event.Attributes"))
		return
	}

	labels := prometheus.Labels{
		"worker": worker,
	}

	key := serializeLabels(&labels)
	if emitter.workerVolumesLabels[worker] == nil {
		emitter.workerVolumesLabels[worker] = make(map[string]prometheus.Labels)
	}
	emitter.workerVolumesLabels[worker][key] = labels
	emitter.workerUnknownVolumes.With(emitter.workerVolumesLabels[worker][key]).Set(event.Value)
}

func (emitter *PrometheusEmitter) workerTasksMetric(logger lager.Logger, event metric.Event) {
	worker, exists := event.Attributes["worker"]
	if !exists {
		logger.Error("failed-to-find-worker-in-event", fmt.Errorf("expected worker to exist in event.Attributes"))
		return
	}
	platform, exists := event.Attributes["platform"]
	if !exists || platform == "" {
		logger.Error("failed-to-find-platform-in-event", fmt.Errorf("expected platform to exist in event.Attributes"))
		return
	}

	labels := prometheus.Labels{
		"worker":   worker,
		"platform": platform,
	}
	key := serializeLabels(&labels)
	if emitter.workerTasksLabels[worker] == nil {
		emitter.workerTasksLabels[worker] = make(map[string]prometheus.Labels)
	}
	emitter.workerTasksLabels[worker][key] = labels
	emitter.workerTasks.With(emitter.workerTasksLabels[worker][key]).Set(event.Value)
}

func (emitter *PrometheusEmitter) httpResponseTimeMetrics(logger lager.Logger, event metric.Event) {
	route, exists := event.Attributes["route"]
	if !exists {
		logger.Error("failed-to-find-route-in-event", fmt.Errorf("expected method to exist in event.Attributes"))
		return
	}

	method, exists := event.Attributes["method"]
	if !exists {
		logger.Error("failed-to-find-method-in-event", fmt.Errorf("expected method to exist in event.Attributes"))
		return
	}

	status, exists := event.Attributes["status"]
	if !exists {
		logger.Error("failed-to-find-status-in-event", fmt.Errorf("expected status to exist in event.Attributes"))
		return
	}

	emitter.httpRequestsDuration.WithLabelValues(method, route, status).Observe(event.Value / 1000)
}

func (emitter *PrometheusEmitter) databaseMetrics(logger lager.Logger, event metric.Event) {
	switch event.Name {
	case "database queries":
		emitter.dbQueriesTotal.Add(event.Value)
	case "database connections":
		connectionName, exists := event.Attributes["ConnectionName"]
		if !exists {
			logger.Error("failed-to-connection-name-in-event", fmt.Errorf("expected ConnectionName to exist in event.Attributes"))
			return
		}
		emitter.dbConnections.WithLabelValues(connectionName).Set(event.Value)
	default:
	}

}

// updateLastSeen tracks for each worker when it last received a metric event.
func (emitter *PrometheusEmitter) updateLastSeen(event metric.Event) {
	emitter.mu.Lock()
	defer emitter.mu.Unlock()
	if worker, exists := event.Attributes["worker"]; exists {
		emitter.workerLastSeen[worker] = time.Now()
	}
}

//periodically remove stale metrics for workers
func (emitter *PrometheusEmitter) periodicMetricGC() {
	for {
		emitter.mu.Lock()
		now := time.Now()
		for worker, lastSeen := range emitter.workerLastSeen {
			if now.Sub(lastSeen) > 5*time.Minute {
				DoGarbageCollection(emitter, worker)
				delete(emitter.workerLastSeen, worker)
			}
		}
		emitter.mu.Unlock()
		time.Sleep(60 * time.Second)
	}
}

// DoGarbageCollection retrieves and deletes stale metrics by their labels.
func DoGarbageCollection(emitter PrometheusGarbageCollectable, worker string) {
	for _, labels := range emitter.WorkerContainersLabels()[worker] {
		emitter.WorkerContainers().Delete(labels)
	}

	for _, labels := range emitter.WorkerVolumesLabels()[worker] {
		emitter.WorkerVolumes().Delete(labels)
	}

	for _, labels := range emitter.WorkerTasksLabels()[worker] {
		emitter.WorkerTasks().Delete(labels)
	}

	delete(emitter.WorkerContainersLabels(), worker)
	delete(emitter.WorkerVolumesLabels(), worker)
	delete(emitter.WorkerTasksLabels(), worker)
}

//go:generate counterfeiter . PrometheusGarbageCollectable
type PrometheusGarbageCollectable interface {
	WorkerContainers() *prometheus.GaugeVec
	WorkerVolumes() *prometheus.GaugeVec
	WorkerTasks() *prometheus.GaugeVec

	WorkerContainersLabels() map[string]map[string]prometheus.Labels
	WorkerVolumesLabels() map[string]map[string]prometheus.Labels
	WorkerTasksLabels() map[string]map[string]prometheus.Labels
}

func (emitter *PrometheusEmitter) WorkerContainers() *prometheus.GaugeVec {
	return emitter.workerContainers
}

func (emitter *PrometheusEmitter) WorkerVolumes() *prometheus.GaugeVec {
	return emitter.workerVolumes
}

func (emitter *PrometheusEmitter) WorkerTasks() *prometheus.GaugeVec {
	return emitter.workerTasks
}

func (emitter *PrometheusEmitter) WorkerContainersLabels() map[string]map[string]prometheus.Labels {
	return emitter.workerContainersLabels
}

func (emitter *PrometheusEmitter) WorkerVolumesLabels() map[string]map[string]prometheus.Labels {
	return emitter.workerVolumesLabels
}

func (emitter *PrometheusEmitter) WorkerTasksLabels() map[string]map[string]prometheus.Labels {
	return emitter.workerTasksLabels
}
