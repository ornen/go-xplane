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

type BatteryAmperageMessage struct {
	MessageData
	Amperage float64
}

func NewBatteryAmperageMessage(data []float32) BatteryAmperageMessage {
	return BatteryAmperageMessage{
		MessageData: MessageData{Name: "BatteryAmperage"},
		Amperage:    float64(data[0]),
	}
}

func (m BatteryAmperageMessage) Type() uint {
	return BatteryAmperageMessageType
}
func (m BatteryAmperageMessage) GetName() string { return m.Name }
