# h0neyp0t: A Simple Honeypot in Go

This project implements a basic honeypot written in Go. It emulates various services to attract and log malicious activity.

## Features

* Supports multiple service emulations (e.g., SSH, HTTP, Telnet).
* Logs connection attempts, commands executed, and other interactions.
* Configurable through `config.json`.

## Getting Started

1. **Configuration:** Modify `config.json` to define the services and ports to emulate.
2. **Build:** `go build`
3. **Run:** `./h0neyp0t`

## Configuration (config.json)

```json
{
  "services": [
    {
      "name": "ssh",
      "port": 2222,
      "banner": "SSH-2.0-OpenSSH_7.6p1 Ubuntu-4ubuntu0.3" // Example SSH banner
    },
    {
      "name": "http",
      "port": 8080
    },
    {
      "name": "telnet",
      "port": 2323
    }
    // Add more services as needed...
  ],
 "logfile": "honeypot.log"
}
