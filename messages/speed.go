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

type SpeedMessage struct {
	IndicatedSpeed float32
	Airspeed       float32
	TrueAirspeed   float32
	GroundSpeed    float32
}

func NewSpeedMessage(data []float32) SpeedMessage {
	return SpeedMessage{
		IndicatedSpeed: data[0] * knotsToMetersPerSecond,
		Airspeed:       data[1] * knotsToMetersPerSecond,
		TrueAirspeed:   data[2] * knotsToMetersPerSecond,
		GroundSpeed:    data[3] * knotsToMetersPerSecond,
	}
}

func (m SpeedMessage) Type() uint {
	return SpeedMessageType
}
