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

type FuelFlow struct {
	MessageData
	FF1 float64 //lb/h
	FF2 float64 //lb/h
}

func NewFuelFlow(data []float32) FuelFlow {
	return FuelFlow{
		MessageData: MessageData{Name: "FuelFlow"},
		FF1:         float64(data[0]),
		FF2:         float64(data[1]),
	}
}

func (m FuelFlow) Type() uint { return FuelFlowType }

func (m FuelFlow) GetName() string { return m.Name }
