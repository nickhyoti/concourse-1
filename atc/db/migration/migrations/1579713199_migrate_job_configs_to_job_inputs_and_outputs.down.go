package migrations

func (m *migrations) Down_1574452410() error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("TRUNCATE TABLE job_inputs")
	if err != nil {
		return err
	}

	_, err = tx.Exec("TRUNCATE TABLE job_outputs")
	if err != nil {
		return err
	}

	return tx.Commit()
}
