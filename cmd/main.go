package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/pmoieni/gopx/internal/config"
	"github.com/pmoieni/gopx/internal/proxy"
	"github.com/pmoieni/gopx/internal/shell"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gopx",
		Usage: "wrap your shell with a proxy",
		Action: func(*cli.Context) error {
			cfg, err := config.Read()
			if err != nil {
				return err
			}

			// TODO: avoid parsing URL twice
			origin, err := url.Parse(cfg.Origin)
			if err != nil {
				return err
			}

			px := proxy.New(origin)

			if err := shell.Init("http://localhost" + cfg.Port + "/"); err != nil {
				return err
			}

			http.HandleFunc("/", px.Handler())
			return http.ListenAndServe(cfg.Port, nil)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
