package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kenkoii/go-starter-k8s/server"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		DSN string
	}
	Server struct {
		Addr         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
	}
}

func load(file string) (cfg Config, err error) {
	//Set defaults in case no config file can be read
	cfg.Server.Addr = ":8080"
	cfg.Server.ReadTimeout = 5 * time.Second
	cfg.Server.WriteTimeout = 5 * time.Second
	cfg.Server.IdleTimeout = 5 * time.Second

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return
	}

	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := "We're online boiz"
	fmt.Fprintf(w, "%s!", msg)
}

func main() {
	filename := flag.String("config", "config.yml", "Configuration file")
	flag.Parse()

	config, err := load(*filename)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	//Define http handler
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler)

	logger := log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)

	//Define http/server
	httpSrv := &http.Server{
		ReadHeaderTimeout: config.Server.ReadTimeout,
		IdleTimeout:       config.Server.IdleTimeout,
		ReadTimeout:       config.Server.ReadTimeout,
		WriteTimeout:      config.Server.WriteTimeout,
		Addr:              config.Server.Addr,
		Handler:           r,
	}

	//Create the server(our wrapper with std out logging)
	srv := &server.Server{
		Srv:    httpSrv,
		Logger: logger,
	}

	//Start Server
	srv.Start()
}
