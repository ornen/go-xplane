// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package messages

type LatLonAltMessage struct {
	Latitude          float32
	Longitude         float32
	AltitudeMSL       float32
	AltitudeAGL       float32
	AltitudeIndicated float32
}

func NewLatLonAltMessage(data []float32) LatLonAltMessage {
	return LatLonAltMessage{
		Latitude:          data[0],
		Longitude:         data[1],
		AltitudeMSL:       data[2] * feetToMeters,
		AltitudeAGL:       data[3] * feetToMeters,
		AltitudeIndicated: data[5] * feetToMeters,
	}
}

func (m LatLonAltMessage) Type() uint {
	return LatLonAltMessageType
}
