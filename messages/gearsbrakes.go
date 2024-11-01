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

type GearsBrakesMessage struct {
	MessageData
	Gears  float64
	Brakes float64
}

func NewGearsBrakesMessage(data []float32) GearsBrakesMessage {
	return GearsBrakesMessage{
		MessageData: MessageData{Name: "GearsBrakes"},
		Gears:       float64(data[0]),
		Brakes:      float64(data[0]),
	}
}

func (c GearsBrakesMessage) Type() uint {
	return GearsBrakesMessageType
}

func (c GearsBrakesMessage) GetName() string { return c.Name }

func (c GearsBrakesMessage) Data() [8]float32 {
	return [8]float32{
		float32(c.Gears),
		float32(c.Brakes),
		-999,
		-999,
		-999,
		-999,
		-999,
		-999,
	}
}
