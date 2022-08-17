package repository

func (hr healthRepository) DatabaseHealth() (status bool, err error) {
	if hr.database.DB != nil {
		status = true
	} else {
		return
	}

	_, err = hr.database.DB.Query("SELECT 1")
	return
}
