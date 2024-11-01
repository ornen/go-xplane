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

type FuelWeights struct {
	MessageData
	Fuel1 float64 //lb or gal
	Fuel2 float64
	Fuel3 float64
	Fuel4 float64
}

func NewFuelWeights(data []float32) FuelWeights {
	return FuelWeights{
		MessageData: MessageData{Name: "FuelWeights"},
		Fuel1:       float64(data[0]),
		Fuel2:       float64(data[1]),
		Fuel3:       float64(data[2]),
		Fuel4:       float64(data[3]),
	}
}

func (m FuelWeights) Type() uint { return FuelWeightsType }

func (m FuelWeights) GetName() string { return m.Name }
