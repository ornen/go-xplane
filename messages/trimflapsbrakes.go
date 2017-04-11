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

type TrimFlapsBrakesMessage struct {
	FlapsHandle         float64
	FlapsPosition       float64
	SpeedBrakesHandle   float64
	SpeedBrakesPosition float64
}

func NewTrimFlapsBrakesMessage(data []float32) TrimFlapsBrakesMessage {
	return TrimFlapsBrakesMessage{
		FlapsHandle:         float64(data[3]),
		FlapsPosition:       float64(data[4]),
		SpeedBrakesHandle:   float64(data[6]),
		SpeedBrakesPosition: float64(data[7]),
	}
}

func (m TrimFlapsBrakesMessage) Type() uint {
	return TrimFlapsBrakesMessageType
}


func (c TrimFlapsBrakesMessage) Data() [8]float32 {
	return [8]float32{
		-999,
		-999,
		-999,
		float32(c.FlapsHandle),
		-999,
		-999,
		float32(c.SpeedBrakesHandle),
		-999,
	}
}