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

type ThrottleCommandMessage struct {
	MessageData
	Throttle1 float64
	Throttle2 float64
	Throttle3 float64
	Throttle4 float64
}

func NewThrottleCommandMessage(data []float32) ThrottleCommandMessage {
	return ThrottleCommandMessage{
		MessageData: MessageData{Name: "ThrottleCommand"},
		Throttle1:   float64(data[0]),
		Throttle2:   float64(data[1]),
		Throttle3:   float64(data[2]),
		Throttle4:   float64(data[3]),
	}
}

func (c ThrottleCommandMessage) Type() uint {
	return ThrottleCommandMessageType
}

func (c ThrottleCommandMessage) GetName() string { return c.Name }

func (c ThrottleCommandMessage) Data() [8]float32 {
	return [8]float32{
		float32(c.Throttle1),
		float32(c.Throttle2),
		float32(c.Throttle3),
		float32(c.Throttle4),
		-999,
		-999,
		-999,
		-999,
	}
}
