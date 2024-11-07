package main

import "factory/logger"

func main() {

	// test string log
	logger.CL.Log("test string log")

	// test int log
	logger.CL.Log(123)

	// test interface log
	logger.CL.Log(map[string]string{"test": "test"})

	// test map of interface log
	logger.CL.Log(map[string]interface{}{"test": "test"})

	// test multiple types log
	logger.CL.Log(map[string]interface{}{"test": "test", "test2": 123})

	logger.CL.Log(123.3342)

	// test file logger
	logger.FL.Log("test file log")

	// test file logger with map
	logger.FL.Log(map[string]interface{}{"test": "test"})

	// test file logger with multiple types
	logger.FL.Log(map[string]interface{}{"test": "test", "test2": 123})

	// test file logger with float
	logger.FL.Log(123.3342)

	//    if any panic occurs, write good error message
	defer func() {
		if r := recover(); r != nil {
			logger.CL.Log(r)
		}
	}()
}
