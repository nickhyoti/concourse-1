package accessor_test

import (
	"errors"

	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/api/accessor/accessorfakes"
	"github.com/concourse/concourse/atc/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClaimsCacher", func() {
	var (
		fakeAccessTokenFetcher *accessorfakes.FakeAccessTokenFetcher
		maxCacheSizeBytes      int

		claimsCacher accessor.AccessTokenFetcher
	)

	BeforeEach(func() {
		fakeAccessTokenFetcher = new(accessorfakes.FakeAccessTokenFetcher)
		maxCacheSizeBytes = 1000
	})

	JustBeforeEach(func() {
		claimsCacher = accessor.NewClaimsCacher(fakeAccessTokenFetcher, maxCacheSizeBytes)
	})

	It("fetches claims from the DB", func() {
		expectedToken := db.AccessToken{Claims: db.Claims{Username: "foo"}}
		fakeAccessTokenFetcher.GetAccessTokenReturns(expectedToken, true, nil)

		token, found, err := claimsCacher.GetAccessToken("token")
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(BeTrue())
		Expect(token).To(Equal(expectedToken))

		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(1), "did not fetch from DB")
	})

	It("returns not found if the token isn't found", func() {
		fakeAccessTokenFetcher.GetAccessTokenReturns(db.AccessToken{}, false, nil)
		_, found, err := claimsCacher.GetAccessToken("token")
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(BeFalse())
	})

	It("doesn't fetch from the DB when the result is cached", func() {
		fakeAccessTokenFetcher.GetAccessTokenReturns(db.AccessToken{}, true, nil)
		claimsCacher.GetAccessToken("token")
		claimsCacher.GetAccessToken("token")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(1), "did not cache claims")
	})

	It("doesn't cache claims when cache size is exceeded", func() {
		fakeAccessTokenFetcher.GetAccessTokenReturns(db.AccessToken{
			Claims: db.Claims{RawClaims: map[string]interface{}{"a": stringWithLen(2000)}},
		}, true, nil)
		claimsCacher.GetAccessToken("token")
		claimsCacher.GetAccessToken("token")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(2), "cached claims that exceed length")
	})

	It("evicts the least recently used access token when size limit exceeded", func() {
		fakeAccessTokenFetcher.GetAccessTokenReturns(db.AccessToken{
			Claims: db.Claims{RawClaims: map[string]interface{}{"a": stringWithLen(400)}},
		}, true, nil)

		By("filling the cache")
		claimsCacher.GetAccessToken("token1")
		claimsCacher.GetAccessToken("token2")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(2))

		By("overflowing the cache")
		claimsCacher.GetAccessToken("token3")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(3))

		By("fetching the least recently used token")
		claimsCacher.GetAccessToken("token1")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(4), "did not evict least recently used")

		By("ensuring the latest token was not evicted")
		claimsCacher.GetAccessToken("token3")
		Expect(fakeAccessTokenFetcher.GetAccessTokenCallCount()).To(Equal(4), "evicted the latest token")
	})

	It("errors when the DB fails", func() {
		fakeAccessTokenFetcher.GetAccessTokenReturns(db.AccessToken{}, false, errors.New("error"))
		_, _, err := claimsCacher.GetAccessToken("token")
		Expect(err).To(HaveOccurred())
	})

	It("fetches claims from the DB concurrently", func() {
		// this is designed purely to trigger the race detector
		go claimsCacher.GetAccessToken("token1")
		go claimsCacher.GetAccessToken("token2")
		go claimsCacher.GetAccessToken("token3")
		Eventually(fakeAccessTokenFetcher.GetAccessTokenCallCount).Should(Equal(3))
	})
})

func stringWithLen(l int) string {
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		b[i] = 'a'
	}
	return string(b)
}
