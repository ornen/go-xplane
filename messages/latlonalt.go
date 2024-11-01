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

type LatLonAltMessage struct {
	MessageData
	Latitude          float64
	Longitude         float64
	AltitudeMSL       float64
	AltitudeAGL       float64
	AltitudeIndicated float64
}

func NewLatLonAltMessage(data []float32) LatLonAltMessage {
	return LatLonAltMessage{
		MessageData:       MessageData{Name: "LatLonAlt"},
		Latitude:          float64(data[0]),
		Longitude:         float64(data[1]),
		AltitudeMSL:       float64(data[2]) * feetToMeters,
		AltitudeAGL:       float64(data[3]) * feetToMeters,
		AltitudeIndicated: float64(data[5]) * feetToMeters,
	}
}

func (m LatLonAltMessage) Type() uint {
	return LatLonAltMessageType
}

func (m LatLonAltMessage) GetName() string { return m.Name }
