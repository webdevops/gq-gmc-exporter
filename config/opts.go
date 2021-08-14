package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type (
	Opts struct {
		// logger
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		Serial struct {
			Port     string `long:"serial.port"  env:"SERIAL_PORT"      description:"Serial port device (eg. /dev/ttyUSB1)" required:"true"`
			BaudRate uint   `long:"serial.baudrate"  env:"SERIAL_BAUDRATE"      description:"Serial bound rate (eg. 57600)" required:"true"`
			DataBits uint   `long:"serial.databits"  env:"SERIAL_DATABITS"      description:"Serial data bits (eg. 8)" required:"true"`
			StopBits uint   `long:"serial.stopbits"  env:"SERIAL_STOPBITS"      description:"Serial stop bits (eg. 1)" required:"true"`
		}

		// general options
		ServerBind string `long:"bind"     env:"SERVER_BIND"   description:"Server address"     default:":8080"`
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
