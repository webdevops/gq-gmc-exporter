GQ GMC exporter (Geiger–Muller counter)
=========================================

[![license](https://img.shields.io/github/license/webdevops/gq-gmc-exporter.svg)](https://github.com/webdevops/gq-gmc-exporter/blob/master/LICENSE)
[![DockerHub](https://img.shields.io/badge/DockerHub-webdevops%2Fgq--gmc--exporter-blue)](https://hub.docker.com/r/webdevops/gq-gmc-exporter/)
[![Quay.io](https://img.shields.io/badge/Quay.io-webdevops%2Fgq--gmc--exporter-blue)](https://quay.io/repository/webdevops/gq-gmc-exporter)

Prometheus exporter for Geiger–Muller counter from GQ GMC with serial interfaces (eg. USB)

Serial documentation: https://www.gqelectronicsllc.com/download/GQ-RFC1201.txt

Usage
-----

```
Usage:
  gq-gmc-exporter [OPTIONS]

Application Options:
      --debug                         debug mode [$DEBUG]
  -v, --verbose                       verbose mode [$VERBOSE]
      --log.json                      Switch log output to json format [$LOG_JSON]
      --serial.port=                  Serial port device (eg. /dev/ttyUSB1) [$SERIAL_PORT]
      --serial.baudrate=              Serial bound rate (eg. 57600) [$SERIAL_BAUDRATE]
      --serial.databits=              Serial data bits (eg. 8) [$SERIAL_DATABITS]
      --serial.stopbits=              Serial stop bits (eg. 1) [$SERIAL_STOPBITS]
      --serial.intercharactertimeout= An inter-character timeout value, in milliseconds, see
                                      https://github.com/jacobsa/go-serial/blob/master/serial/serial.go#L91 (default:
                                      1000) [$SERIAL_INTERCHARACTERTIMEOUT]
      --serial.minimumreadsize=       Minimum read size, see
                                      https://github.com/jacobsa/go-serial/blob/master/serial/serial.go#L91 (default:
                                      0) [$SERIAL_MINIMUMREADSIZE]
      --bind=                         Server address (default: :8080) [$SERVER_BIND]

Help Options:
  -h, --help                          Show this help message
```

HTTP Endpoints
--------------

| Endpoint                       | Description                                                                         |
|--------------------------------|-------------------------------------------------------------------------------------|
| `/metrics`                     | Default prometheus golang metrics                                                   |

Metrics
-------

| Metric                               | Description                                                                    |
|--------------------------------------|--------------------------------------------------------------------------------|
| `gqgmc_info`                         | Device information                                                             |
| `gqgmc_cpm`                          | Detected counts per minute from Geiger–Muller counter                          |
| `gqgmc_voltage`                      | Current device voltage                                                         |
| `gqgmc_temperature`                  | Current device temperature (if supported)                                      |

Tested devices
--------------

This exporter is tested with following devices:
- GMC-300 plus

If you need support for other devices I ready to add support other devices but I need access to it (eg. via SSH or physical).
