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

type SpeedMessage struct {
	IndicatedSpeed float64
	Airspeed       float64
	TrueAirspeed   float64
	GroundSpeed    float64
}

func NewSpeedMessage(data []float32) SpeedMessage {
	return SpeedMessage{
		IndicatedSpeed: float64(data[0]) * knotsToMetersPerSecond,
		Airspeed:       float64(data[1]) * knotsToMetersPerSecond,
		TrueAirspeed:   float64(data[2]) * knotsToMetersPerSecond,
		GroundSpeed:    float64(data[3]) * knotsToMetersPerSecond,
	}
}

func (m SpeedMessage) Type() uint {
	return SpeedMessageType
}
