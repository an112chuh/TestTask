package work

import (
	"encoding/json"
	"healthcheck/config"
	"healthcheck/errors"
	"net/http"
	"net/url"
	"strings"
)

func sendAPIData(site string, status string) {
	data := url.Values{
		"site":   {site},
		"status": {status},
	}
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", "http://httpbin.org/post", strings.NewReader(data.Encode()))
	if err != nil {
		errors.Print(err.Error())
		return
	}
	_, err = httpClient.Do(req)
	if err != nil {
		errors.Print(err.Error())
		return
	}
}

func GetErrorsHandler(w http.ResponseWriter, r *http.Request) {
	var errs ErrorsStruct
	db := config.ConnectDB()
	query := `SELECT site, fails, created_at FROM actual_statuses WHERE status = 'fail' ORDER BY id DESC LIMIT 10`
	rows, err := db.Query(query)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var e ErrorStruct
		err := rows.Scan(&e.Site, &e.Fails, &e.Time)
		if err != nil {
			errors.Print(err.Error())
			return
		}
		errs.Errors = append(errs.Errors, e)
	}
	json.NewEncoder(w).Encode(errs)

}
