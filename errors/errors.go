package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

func Print(err string) {
	fmt.Printf("\n%s\n%s\n", fileWithLineNum(), err)
}

func New(err string) error {
	return errors.New(err)
}

func fileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ``
}
