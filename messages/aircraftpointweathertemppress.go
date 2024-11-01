// Copyright 2016 Ornen. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messages

type AircraftPointWeatherTempPress struct {
	MessageData
	AMPressure    float64 //InHG
	AMTemperature float64 //degC
	LETemperature float64 //degC
	Density       float64 //ratio
	A             float64 //ktas
	Q             float64 //psf
	Gravity       float64 //m/ss
}

func NewAircraftPointWeatherTempPress(data []float32) AircraftPointWeatherTempPress {
	return AircraftPointWeatherTempPress{
		MessageData:   MessageData{Name: "AircraftPointWeatherTempPress"},
		AMPressure:    float64(data[0]),
		AMTemperature: float64(data[1]),
		LETemperature: float64(data[2]),
		Density:       float64(data[3]),
		A:             float64(data[4]),
		Q:             float64(data[5]),
		Gravity:       float64(data[6]),
	}
}

func (m AircraftPointWeatherTempPress) Type() uint { return AircraftPointWeatherTempPressType }

func (m AircraftPointWeatherTempPress) GetName() string { return m.Name }
