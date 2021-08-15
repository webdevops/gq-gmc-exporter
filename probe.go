package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
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
		opts.Serial.Port,
		opts.Serial.BaudRate,
		opts.Serial.DataBits,
		opts.Serial.StopBits,
	)
	gmcDevice.Connect()

	go func() {
		defer gmcDevice.Close()

		// get model details
		hwModelName, hwModelVersion := gmcDevice.GetHardwareModel()

		prometheusGmcInfo.WithLabelValues(
			opts.Serial.Port,
			hwModelName,
			hwModelVersion,
		).Set(1)

		time.Sleep(10 * time.Second)

		for {
			go func() {
				prometheusGmcCpm.WithLabelValues(opts.Serial.Port).Set(gmcDevice.GetCpm())
				prometheusGmcVoltage.WithLabelValues(opts.Serial.Port).Set(gmcDevice.GetVoltage())
				prometheusGmcTemperature.WithLabelValues(opts.Serial.Port).Set(gmcDevice.GetTemperature())
			}()
			time.Sleep(30 * time.Second)
		}
	}()
}
