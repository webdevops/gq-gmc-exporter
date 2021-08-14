package main

import (
	"encoding/binary"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	log "github.com/sirupsen/logrus"
	"io"
)

type (
	GqGmcDevice struct {
		serialPort     string
		serialBaudRate uint
		serialDataBits uint
		serialStopBits uint

		port io.ReadWriteCloser
	}
)

func NewGqGmcDevice(port string, baudRate, dataBits, stopBits uint) *GqGmcDevice {
	return &GqGmcDevice{
		serialPort:     port,
		serialBaudRate: baudRate,
		serialDataBits: dataBits,
		serialStopBits: stopBits,
	}
}

func (d *GqGmcDevice) Connect() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        opts.Serial.Port,
		BaudRate:        opts.Serial.BaudRate,
		DataBits:        opts.Serial.DataBits,
		StopBits:        opts.Serial.StopBits,
		ParityMode:      serial.PARITY_NONE,
		MinimumReadSize: 1,
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

func (d *GqGmcDevice) read(bytes uint) ([]byte, error) {
	log.Debugf("reading %v bytes", bytes)

	buf := make([]byte, bytes)
	n, err := d.port.Read(buf)
	if err != nil {
		if err != io.EOF {
			return buf, err
		}
	} else {
		return buf[:n], nil
	}

	return buf, nil
}

func (d *GqGmcDevice) GetHardwareModel() (hwModelName string, hwModelVersion string) {
	if err := d.write("GETVER"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(7); err == nil {
		hwModelName = string(buf)
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	if buf, err := d.read(7); err == nil {
		hwModelVersion = string(buf)
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetHardwareSerial() (hwSerial string) {
	if err := d.write("GETSERIAL"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(7); err == nil {
		hwSerial = string(buf)
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetCpm() (cpm float64) {
	if err := d.write("GETCPM"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(2); err == nil {
		cpm = float64(binary.BigEndian.Uint16(buf))
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}

func (d *GqGmcDevice) GetVoltage() (cpm float64) {
	if err := d.write("GETVOLT"); err != nil {
		log.Panicf("error sending command to serial port: %v", err)
	}

	if buf, err := d.read(1); err == nil {
		cpm = float64(uint(buf[0]))
	} else {
		log.Panicf("error reading from serial port: %v", err)
	}

	return
}
