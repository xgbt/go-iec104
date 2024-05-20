package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xgbt/go-iec104"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	iec104.SetLogger(logger)

	server := iec104.NewServer(":2404", nil)
	if err := server.Serve(); err != nil {
		panic(any(err))
	}
}
