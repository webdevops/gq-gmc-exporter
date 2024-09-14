package config

import (
	"encoding/json"
	"time"
)

type (
	Opts struct {
		Logger struct {
			Debug       bool `long:"log.debug"    env:"LOG_DEBUG"  description:"debug mode"`
			Development bool `long:"log.devel"    env:"LOG_DEVEL"  description:"development mode"`
			Json        bool `long:"log.json"     env:"LOG_JSON"   description:"Switch log output to json format"`
		}

		Serial struct {
			Port                  string `long:"serial.port"  env:"SERIAL_PORT"      description:"Serial port device (eg. /dev/ttyUSB1)" required:"true"`
			BaudRate              uint   `long:"serial.baudrate"  env:"SERIAL_BAUDRATE"      description:"Serial bound rate (eg. 57600)" required:"true"`
			DataBits              uint   `long:"serial.databits"  env:"SERIAL_DATABITS"      description:"Serial data bits (eg. 8)" required:"true"`
			StopBits              uint   `long:"serial.stopbits"  env:"SERIAL_STOPBITS"      description:"Serial stop bits (eg. 1)" required:"true"`
			InterCharacterTimeout uint   `long:"serial.intercharactertimeout"  env:"SERIAL_INTERCHARACTERTIMEOUT"      description:"An inter-character timeout value, in milliseconds, see https://github.com/jacobsa/go-serial/blob/master/serial/serial.go#L91" default:"1000"`
			MinimumReadSize       uint   `long:"serial.minimumreadsize"        env:"SERIAL_MINIMUMREADSIZE"            description:"Minimum read size, see https://github.com/jacobsa/go-serial/blob/master/serial/serial.go#L91" default:"0"`
		}

		// general options
		Server struct {
			// general options
			Bind         string        `long:"server.bind"              env:"SERVER_BIND"           description:"Server address"        default:":8080"`
			ReadTimeout  time.Duration `long:"server.timeout.read"      env:"SERVER_TIMEOUT_READ"   description:"Server read timeout"   default:"5s"`
			WriteTimeout time.Duration `long:"server.timeout.write"     env:"SERVER_TIMEOUT_WRITE"  description:"Server write timeout"  default:"10s"`
		}
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}
