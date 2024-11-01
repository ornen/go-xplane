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

type AircraftPointWeatherPrecipWind struct {
	MessageData
	Immr          float64
	Turbulence    float64
	Rain          float64
	Snow          float64
	Hail          float64
	WindSpeed     float64 //knots
	WindDirection float64 //dir
	WindVelocity  float64 //fpm
}

func NewAircraftPointWeatherPrecipWind(data []float32) AircraftPointWeatherPrecipWind {
	return AircraftPointWeatherPrecipWind{
		MessageData:   MessageData{Name: "AircraftPointWeatherPrecipWind"},
		Immr:          float64(data[0]),
		Turbulence:    float64(data[1]),
		Rain:          float64(data[2]),
		Snow:          float64(data[3]),
		Hail:          float64(data[4]),
		WindSpeed:     float64(data[5]),
		WindDirection: float64(data[6]),
		WindVelocity:  float64(data[7]),
	}
}

func (m AircraftPointWeatherPrecipWind) Type() uint { return AircraftPointWeatherPrecipWindType }

func (m AircraftPointWeatherPrecipWind) GetName() string { return m.Name }
