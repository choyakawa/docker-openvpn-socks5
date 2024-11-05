package main

import (
    "log"
    "os"
    "os/exec"

    "github.com/caarlos0/env"
    "github.com/txthinking/socks5"
)

type params struct {
    User     string `env:"SOCKS5_USER" envDefault:""`
    Password string `env:"SOCKS5_PASS" envDefault:""`
    Port     string `env:"SOCKS5_PORT" envDefault:"1080"`
    Up       string `env:"SOCKS5_UP"   envDefault:""`
}

func main() {
    cfg := params{}
    err := env.Parse(&cfg)
    if err != nil {
        log.Printf("%+v\n", err)
    }

    address := ":" + cfg.Port

    var server *socks5.Server
    if cfg.User != "" && cfg.Password != "" {
        server, err = socks5.NewClassicServer(address, cfg.User, cfg.Password, 0)
    } else {
        server, err = socks5.NewClassicServer(address, "", "", 0)
    }
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Start listening proxy service on port %s\n", cfg.Port)

    if cfg.Up != "" {
        err = exec.Command(cfg.Up).Start()
        if err != nil {
            log.Fatal(err)
        }
    }

    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
