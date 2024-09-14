package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	prometheusGmcInfo        *prometheus.GaugeVec
	prometheusGmcCpm         *prometheus.GaugeVec
	prometheusGmcVoltage     *prometheus.GaugeVec
	prometheusGmcTemperature *prometheus.GaugeVec
)

func setupMetrics() {
	prometheusGmcInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gqgmc_info",
			Help: "GQ GMC counts per minute",
		},
		[]string{
			"port",
			"model",
			"version",
		},
	)
	prometheus.MustRegister(prometheusGmcInfo)

	prometheusGmcCpm = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gqgmc_cpm",
			Help: "GQ GMC counts per minute",
		},
		[]string{
			"port",
		},
	)
	prometheus.MustRegister(prometheusGmcCpm)

	prometheusGmcVoltage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gqgmc_voltage",
			Help: "GQ GMC battery voltage",
		},
		[]string{
			"port",
		},
	)
	prometheus.MustRegister(prometheusGmcVoltage)

	prometheusGmcTemperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gqgmc_temperature",
			Help: "GQ GMC device temperature",
		},
		[]string{
			"port",
		},
	)
	prometheus.MustRegister(prometheusGmcTemperature)
}

func runProbes() {
	setupMetrics()

	gmcDevice := NewGqGmcDevice(
		Opts.Serial.Port,
		Opts.Serial.BaudRate,
		Opts.Serial.DataBits,
		Opts.Serial.StopBits,
		Opts.Serial.InterCharacterTimeout,
		Opts.Serial.MinimumReadSize,
	)
	gmcDevice.Connect()

	go func() {
		runProbeLoop(gmcDevice)
	}()
}

func runProbeLoop(gmcDevice *GqGmcDevice) {
	defer gmcDevice.Close()

	//

	// get model details
	hwModelName, hwModelVersion := gmcDevice.GetHardwareModel()
	if hwModelName == "" || hwModelVersion == "" {
		logger.Panic("no model or version detected, please check serial settings or device support")
	} else {
		logger.Infof(
			"detected device model \"%s\" with version \"%s\"\n",
			hwModelName,
			hwModelVersion,
		)
	}

	prometheusGmcInfo.WithLabelValues(
		Opts.Serial.Port,
		hwModelName,
		hwModelVersion,
	).Set(1)

	hwModelNumber := 0
	hwModelNameLowercase := strings.ToLower(hwModelName)
	if strings.HasPrefix(hwModelNameLowercase, "gmc-") {
		hwModelNumberString := strings.TrimPrefix(strings.ToLower(hwModelName), "gmc-")
		if v, err := strconv.Atoi(hwModelNumberString); err == nil {
			hwModelNumber = v
			logger.Infof("detected model number \"%v\"", hwModelNumber)
		}
	}

	time.Sleep(5 * time.Second)

	for {
		gmcDevice.ClearSerialConsole()

		if val := gmcDevice.GetCpm(); val != nil {
			prometheusGmcCpm.WithLabelValues(Opts.Serial.Port).Set(*val)
		}

		if val := gmcDevice.GetVoltage(); val != nil {
			prometheusGmcVoltage.WithLabelValues(Opts.Serial.Port).Set(*val)
		}

		if hwModelNumber > 320 {
			if val := gmcDevice.GetTemperature(); val != nil {
				prometheusGmcTemperature.WithLabelValues(Opts.Serial.Port).Set(*val)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
