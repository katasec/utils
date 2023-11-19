package errx

import (
	"fmt"
	"os"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicOnErrorf(err error, format string, args ...interface{}) {
	if nil != err {
		panic(fmt.Sprintf(format, args...) + "\nOriginal Error: " + err.Error())
	}
}

func ReturnError(err error) error {
	if err != nil {
		return err
	}

	return nil
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
