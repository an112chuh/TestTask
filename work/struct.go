package work

type ErrorsStruct struct {
	Errors []ErrorStruct `json:"errors"`
}

type ErrorStruct struct {
	Site  string `json:"site"`
	Fails string `json:"fails"`
	Time  string `json:"time"`
}
