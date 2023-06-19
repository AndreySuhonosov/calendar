package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AndreySuhonosov/calendar/pkg/logger"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Host     string `config:"host"`
	Port     uint32 `config:"port"`
	LogLevel string `config:"logLevel"`
}

func main() {
	var config Config
	path := flag.String("config", "config/config.json", "path to config file")
	loader := logger.NewConfig(*path)
	err := loader.GetConfig(context.Background(), &config)
	if err != nil {
		log.Fatal(err)
	}
	newLogger, err := logger.NewLogger("debug", []string{"stdout"})
	if err != nil {
		panic(err)
	}

	host := fmt.Sprint("Host:", config.Host, "Port:", config.Port, "LogLevel", config.LogLevel)
	newLogger.Debug(host)

	addr := config.Host + ":" + strconv.Itoa(int(config.Port))
	newLogger.Debug(addr)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Hello, oksana-p!\n")
		if err != nil {
			return
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		newLogger.Info(string(body))
	}
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(addr, nil))

}
