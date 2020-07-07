package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":3001", "HTTP listen address")
	)
	flag.Parse()

	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "at", kitlog.DefaultCaller)
	}

	mux := http.NewServeMux()

	errorChannel := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errorChannel <- http.ListenAndServe(*httpAddr, mux)
	}()

	logger.Log("exit", <-errorChannel)
}
