package main

import (
	"fmt"
	"os"

	"github.com/kevinsantana/gosolve-recruitment-task/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.New()
	lvl, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		fmt.Println("Invalid log level specified:", err)
		lvl = log.DebugLevel
	}
	fmt.Println("log level specified:", lvl)
	log.SetLevel(lvl)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
}

func main() {
	cmd.Execute()
}
