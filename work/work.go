package work

import (
	"bytes"
	"fmt"
	"healthcheck/config"
	"healthcheck/errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func New(config config.JSONStruct) {
	for {
		for _, data := range config.URLs {
			res := makeChecks(data)
			fmt.Println(res)
		}
		duration := time.Duration(config.SleepTime)
		time.Sleep(duration * time.Second)
	}
}

func makeChecks(data config.URLStruct) string {
	checksCount := 0
	errorString := ""
	resp, err := getSiteData(data)
	if err != nil {
		return fmt.Sprintf("%s: error", data.URL)
	}
	defer resp.Body.Close()
	for _, params := range data.Checks {
		isSuccess, err := doChecks(resp, params)
		if err != nil {
			errors.Print(err.Error())
			return fmt.Sprintf("%s: error", data.URL)
		}
		if isSuccess {
			checksCount++
		} else {
			errorString += params.Param
		}
		if checksCount >= data.MinChecks {
			putIntoDB(data.URL, "ok", "")
			return fmt.Sprintf("%s: ok", data.URL)
		}
	}
	putIntoDB(data.URL, "fail", errorString)
	errorString += ")"
	return fmt.Sprintf("%s: fail(", data.URL) + errorString
}

func getSiteData(data config.URLStruct) (resp *http.Response, err error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", data.URL, nil)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		errors.Print(err.Error())
		return
	}
	return
}

func doChecks(resp *http.Response, params config.ParamStruct) (success bool, err error) {
	switch params.Param {
	case `status_code`:
		return checkStatusCode(resp, params.Data)
	case `text`:
		return checkText(resp, params.Data), nil
	}
	return false, errors.New(`unknown check`)
}

func checkStatusCode(resp *http.Response, code string) (success bool, err error) {
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == codeInt, nil
}

func checkText(resp *http.Response, text string) (success bool) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBytes := buf.String()
	respString := string(respBytes)
	return strings.Contains(respString, text)
}
