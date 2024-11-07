package logger

import (
	"encoding/json"
	"reflect"
)

type ConsoleLogger struct {
}

// Log finds type of msg and logs it to console
func (cl *ConsoleLogger) Log(msg interface{}) {
	r := reflect.TypeOf(msg)

	// print according to type
	switch r.Kind() {
	case reflect.String:
		println(msg.(string))
	case reflect.Int:
		println(msg.(int))
	case reflect.Map:
		// convert to string by json marshall
		toPrint, err := json.Marshal(msg)
		if err != nil {
			println(err)
		}
		println(string(toPrint))
	default:
		toPrint, err := json.Marshal(msg)
		if err != nil {
			println(err)
		}
		println(string(toPrint))
	}

}
