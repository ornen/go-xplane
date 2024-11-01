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

type PropPitchMessage struct {
	MessageData
	MainRotorPitch float64
	TailRotorPitch float64
}

func NewPropPitchMessage(data []float32) PropPitchMessage {
	return PropPitchMessage{
		MessageData:    MessageData{Name: "PropPitch"},
		MainRotorPitch: float64(data[0]),
		TailRotorPitch: float64(data[1]),
	}
}

func (m PropPitchMessage) Type() uint {
	return PropPitchMessageType
}

func (m PropPitchMessage) GetName() string { return m.Name }

func (c PropPitchMessage) Data() [8]float32 {
	return [8]float32{
		float32(c.MainRotorPitch),
		float32(c.TailRotorPitch),
		-999,
		-999,
		-999,
		-999,
		-999,
		-999,
	}
}
