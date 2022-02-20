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

const (
	AngleOfAttackSideslipMessageType = 18
)

type AngleOfAttackSideslipMessage struct {
	Alpha float64
	Beta float64
	HPath float64
	VPath float64
	Slip float64
}

func NewAngleOfAttackSideslipMessage(data []float32) AngleOfAttackSideslipMessage {
	return AngleOfAttackSideslipMessage{
		Alpha: float64(data[0]),
		Beta: float64(data[1]),
		HPath: float64(data[2]),
		VPath: float64(data[3]),
		Slip: float64(data[4]),
	}
}

func (m AngleOfAttackSideslipMessage) Type() uint {
	return AngleOfAttackSideslipMessageType
}
