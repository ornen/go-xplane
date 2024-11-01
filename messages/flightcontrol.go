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

type FlightControlMessage struct {
	MessageData
	Elevator float64
	Aileron  float64
	Rudder   float64
}

func NewFlightControlMessage(data []float32) FlightControlMessage {
	return FlightControlMessage{
		MessageData: MessageData{Name: "FlightControl"},
		Elevator:    float64(data[0]),
		Aileron:     float64(data[1]),
		Rudder:      float64(data[2]),
	}
}

func (c FlightControlMessage) Type() uint {
	return FlightControlMessageType
}

func (c FlightControlMessage) GetName() string { return c.Name }

func (c FlightControlMessage) Data() [8]float32 {
	return [8]float32{
		float32(c.Elevator),
		float32(c.Aileron),
		float32(c.Rudder),
		-999,
		float32(c.Rudder),
		-999,
		-999,
		-999,
	}
}
