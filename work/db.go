package work

import (
	"healthcheck/config"
	"healthcheck/errors"
	"time"
)

func putIntoDB(URL string, status string, fails string) {
	db := config.ConnectDB()
	query := `INSERT INTO actual_statuses (site, status, fails, created_at) VALUES ($1, $2, $3, $4)`
	params := []any{URL, status, fails, time.Now()}
	_, err := db.Exec(query, params...)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	var count int
	query = `SELECT COUNT(*) FROM (SELECT DISTINCT status FROM (SELECT * FROM actual_statuses WHERE site = $1 ORDER BY id DESC LIMIT 2))`
	params = []any{URL}
	err = db.QueryRow(query, params...).Scan(&count)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	if count > 1 {
		sendAPIData(URL, status)
	}
}
