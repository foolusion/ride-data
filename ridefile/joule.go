// Copyright 2013 Andrew O'Neill. All rights reserved.  Use of this file is
// governed by a BSD-style license that can be found in the LICENSE file.

package ridefile

import (
	"fmt"
)

type Device struct {
	Version string
	DateTime time.Time
	Km float64 // Total Kilometers
	Minutes float64 // Total minutes
	RPE int
	Tags []string
	Weight float64
	Work int // Kilojoules
	FTP int
	SampleRate int
	DeviceType string
	FirmwareVersion string
	LastUpdated string
	Category1 string
	Category2 string
}

type User struct {
	Name string
	PowerZone1 int
	PowerZone2 int
	PowerZone3 int
	PowerZone4 int
	PowerZone5 int
	PowerZone6 int
	HeartRate1 int
	HeartRate2 int
	HeartRate3 int
	HeartRate4 int
	HeartRate5 int
	CalcPowerA int
	CalcPowerB int
	CalcPowerC int
}

type RideData struct {
	Minutes float64
	Torq int // Newton-meters
	Kmh float64 // Kilometers/hour
	Watts int
	Km float64
	Cadence int
	Hrate int
	ID int
	Altitude int // meters
	Temperature float64 // Celsius
	Grade float64 // percent
	Latitude float64
	Longitude float64
	PowerCalcd float64
	RightPedal float64
	PedalPower float64 // percent
	CadSmooth float64
}

// TODO(foolusion@gmail.com): implement read on joule files.
