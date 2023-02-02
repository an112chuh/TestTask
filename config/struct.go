package config

type JSONStruct struct {
	SleepTime int         `json:"sleep_time"`
	URLs      []URLStruct `json:"urls"`
}

type URLStruct struct {
	URL       string        `json:"url"`
	Checks    []ParamStruct `json:"checks"`
	MinChecks int           `json:"min_checks_cnt"`
}

type ParamStruct struct {
	Param string `json:"param"`
	Data  string `json:"data"`
}
