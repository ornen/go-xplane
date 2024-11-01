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

type PropellerTorque struct {
	MessageData
	Prop1 float64 //ft-lb
	Prop2 float64
	Prop3 float64
	Prop4 float64
}

func NewPropellerTorque(data []float32) PropellerTorque {
	return PropellerTorque{
		MessageData: MessageData{Name: "PropellerTorque"},
		Prop1:       float64(data[0]),
		Prop2:       float64(data[1]),
		Prop3:       float64(data[2]),
		Prop4:       float64(data[3]),
	}
}

func (m PropellerTorque) Type() uint { return PropellerTorqueType }

func (m PropellerTorque) GetName() string { return m.Name }
