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
      --serial.intercharactertimeout= An inter-character timeout value, in milliseconds (default: 1000)
                                      [$SERIAL_INTERCHARACTERTIMEOUT]
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

