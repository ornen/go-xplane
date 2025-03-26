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

type LocVelDistTraveledMessage struct {
	MessageData
	X  float64
	Y  float64
	Z  float64
	Xv float64
	Yv float64
	Zv float64
}

func NewLocVelDistTraveledMessage(data []float32) LocVelDistTraveledMessage {
	return LocVelDistTraveledMessage{
		MessageData: MessageData{Name: "LocVelDistTraveled"},
		Y:           float64(data[0]),
		Z:           -float64(data[1]),
		X:           -float64(data[2]),
		Yv:          float64(data[3]),
		Zv:          -float64(data[4]),
		Xv:          -float64(data[5]),
	}
}

func (m LocVelDistTraveledMessage) Type() uint {
	return LocVelDistTraveledMessageType
}

func (e LocVelDistTraveledMessage) GetName() string { return e.Name }
