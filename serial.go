package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
)

const (
	// GqGmcClearConsoleInputChars try to read that many chars to get input buffer cleared
	GqGmcClearConsoleInputChars = 128
)

type (
	GqGmcDevice struct {
		serialPort     string
		serialBaudRate uint
		serialDataBits uint
		serialStopBits uint

		interCharacterTimeout uint
		minimumReadSize       uint

		port io.ReadWriteCloser
	}
)

func NewGqGmcDevice(port string, baudRate, dataBits, stopBits, InterCharacterTimeout, minimumReadSize uint) *GqGmcDevice {
	return &GqGmcDevice{
		serialPort:            port,
		serialBaudRate:        baudRate,
		serialDataBits:        dataBits,
		serialStopBits:        stopBits,
		interCharacterTimeout: InterCharacterTimeout,
		minimumReadSize:       minimumReadSize,
	}
}

func (d *GqGmcDevice) Connect() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:              d.serialPort,
		BaudRate:              d.serialBaudRate,
		DataBits:              d.serialDataBits,
		StopBits:              d.serialStopBits,
		ParityMode:            serial.PARITY_NONE,
		InterCharacterTimeout: d.interCharacterTimeout,
		MinimumReadSize:       d.minimumReadSize,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Panicf("cannot open %v: %v", d.serialPort, err)
	}

	d.port = port
}

func (d *GqGmcDevice) Close() error {
	return d.port.Close()
}

func (d *GqGmcDevice) write(command string) error {
	log.Debugf("sending %v command", command)
	command = fmt.Sprintf("<%s>>", command)
	_, err := d.port.Write([]byte(command))
	return err
}

func (d *GqGmcDevice) read(chars uint) ([]byte, error) {
	log.Debugf("reading %v bytes", chars)

	buf := make([]byte, chars)
	n, err := d.port.Read(buf)
	if opts.Logger.Debug {
		log.Debugf("fetched %v bytes:\n%v", len(buf), hex.Dump(buf))
	}
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return buf, err
		}
	} else {
		return buf[:n], nil
	}

	return buf, nil
}

func (d *GqGmcDevice) readString(chars uint) (string, error) {
	if buf, err := d.read(chars); err == nil {
		buf = bytes.Trim(buf, "\x00")
		ret := strings.TrimSpace(string(buf))
		return ret, nil
	} else {
		return "", err
	}
}

func (d *GqGmcDevice) ClearSerialConsole() {
	log.Debug("clear console input")
	/* #nosec G104 -- we dont care about errors here */
	d.read(GqGmcClearConsoleInputChars) //nolint
}

func (d *GqGmcDevice) GetHardwareModel() (hwModelName string, hwModelVersion string) {
	if err := d.write("GETVER"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if val, err := d.readString(7); err == nil {
		hwModelName = val
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	if val, err := d.readString(7); err == nil {
		hwModelVersion = val
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetHardwareSerial() (hwSerial string) {
	if err := d.write("GETSERIAL"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if val, err := d.readString(7); err == nil {
		hwSerial = val
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetCpm() (cpm *float64) {
	if err := d.write("GETCPM"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(2); err == nil {
		if len(buf) == 2 {
			val := float64(binary.BigEndian.Uint16(buf))
			cpm = &val
		}
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetVoltage() (voltage *float64) {
	if err := d.write("GETVOLT"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(1); err == nil {
		if len(buf) == 1 {
			val := float64(uint(buf[0]))
			voltage = &val
		}
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetTemperature() (temp *float64) {
	if err := d.write("GETTEMP"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(4); err == nil {
		if len(buf) == 4 {
			tempInt := int(buf[0])
			tempDec := int(buf[1])
			tempSign := int(buf[2])

			// if sign is 0, temp is greater 0 and so positive
			// if sign != 0, temp is below 0 and so negative
			if tempSign != 0 {
				tempSign = -1
			}

			calcTemp := float64(tempSign*(tempInt*1000) + tempDec)
			temp = &calcTemp
		}
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}
