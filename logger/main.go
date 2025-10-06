package main

import logger "go_tutorial/logger/custom_logger"

func main() {
	logger.Init()

	logger.Info("this is an info message")
	logger.Warn("this is an warning message")
	logger.SetLevel(logger.ErrorLevel)
	logger.Error("this is an error message")
}
