package mistakes

import (
	"errors"
	"reflect"
	"runtime"
)

var (
	ErrEnv = func(key string) error {
		return errors.New("Cannot get data from .env file! Key - " + key)
	}
	ErrDatabase = func(err error, functionName ...interface{}) error {
		row := err.Error() + " \n"

		for _, v := range functionName {
			row += getFunctionName(v) + " \n"
		}

		return errors.New(row)
	}
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
