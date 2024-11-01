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

type AircraftRunwayConditions struct {
	MessageData
	Wet   float64 //ratio
	Puddl float64 //ratio
	Snow  float64 //ratio
	Ice   float64 //ratio
}

func NewAircraftRunwayConditions(data []float32) AircraftRunwayConditions {
	return AircraftRunwayConditions{
		MessageData: MessageData{Name: "AircraftRunwayConditions"},
		Wet:         float64(data[0]),
		Puddl:       float64(data[1]),
		Snow:        float64(data[2]),
		Ice:         float64(data[3]),
	}
}

func (m AircraftRunwayConditions) Type() uint { return AircraftRunwayConditionsType }

func (m AircraftRunwayConditions) GetName() string { return m.Name }
