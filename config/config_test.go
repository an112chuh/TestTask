package config

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("conf_bad.json")
	if err != nil {
		t.Errorf("got %q, wanted nil", err)
	}
	got := conf.parse(file).Error()
	want := `json: Unmarshal(nil *config.JSONStruct)`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
