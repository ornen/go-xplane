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

type GLoadMessage struct {
	MessageData
	X float64
	Y float64
	Z float64
}

func NewGLoadMessage(data []float32) GLoadMessage {
	return GLoadMessage{
		MessageData: MessageData{Name: "GLoad"},
		X:           float64(data[5]) * gravityInMetersPerSecondSecond,
		Y:           float64(data[6]) * gravityInMetersPerSecondSecond,
		Z:           -float64(data[4]) * gravityInMetersPerSecondSecond,
	}
}

func (m GLoadMessage) Type() uint {
	return GLoadMessageType
}

func (m GLoadMessage) GetName() string { return m.Name }
