package main

import (
	"context"
	"github.com/AndreySuhonosov/calendar/pkg/logger"
	"io"
	"log"
	"net/http"
)

type Config struct {
	Host     string `config:"host"`
	Port     uint32 `config:"port"`
	LogLevel string `config:"logLevel"`
}

func main() {
	var config Config
	loader := logger.NewConfig("config/config.json")
	err := loader.GetConfig(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}
	newLogger, err := logger.NewLogger("debug", []string{"stdout"})
	if err != nil {
		panic(err)
	}

	newLogger.Debug("srfgasfWEFw")
	newLogger.Info("wsrdfEFwds")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
