package main

import (
        "log"
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
        // Parse environment variables
        cfg := params{}
        if err := env.Parse(&cfg); err != nil {
                log.Fatalf("Failed to parse environment variables: %+v\n", err)
        }

        // Initialize SOCKS5 server
        // addr: The address to listen on (e.g., ":1080")
        // ip: The external IP address (empty string "" means auto-detect)
        // username and password for authentication (empty strings mean no authentication)
        // tcpTimeout and udpTimeout (set to 0 for no timeout)
        server, err := socks5.NewClassicServer(":"+cfg.Port, "", cfg.User, cfg.Password, 0, 0)
        if err != nil {
                log.Fatalf("Failed to create SOCKS5 server: %v", err)
        }

        log.Printf("Starting SOCKS5 proxy on port %s\n", cfg.Port)

        // Execute the 'Up' command if provided
        if cfg.Up != "" {
                if err := exec.Command(cfg.Up).Start(); err != nil {
                        log.Fatalf("Failed to execute Up command: %v", err)
                }
        }

        // Start the server with the default handler
        if err := server.ListenAndServe(nil); err != nil {
                log.Fatalf("SOCKS5 server error: %v", err)
        }
}
