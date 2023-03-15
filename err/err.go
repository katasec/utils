package err

import "fmt"

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
