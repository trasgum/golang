package main

import (
	"flag"
	"os"

	"github.com/trasgum/mesos_operator_stream/app"
	"github.com/go-kit/kit/log"
)

func main() {
	cfg := app.NewConfig()
	fs := flag.NewFlagSet("operator", flag.ExitOnError)
	cfg.AddFlags(fs)
	fs.Parse(os.Args[1:])

	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)


	if err := app.Run(cfg); err != nil {
		logger.Log(err)
	}
}
