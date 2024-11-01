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

type OilTemperature struct {
	MessageData
	OilTemperature1 float64 //deg
	OilTemperature2 float64
}

func NewOilTemperature(data []float32) OilTemperature {
	return OilTemperature{
		MessageData:     MessageData{Name: "OilTemperature"},
		OilTemperature1: float64(data[0]),
		OilTemperature2: float64(data[1]),
	}
}

func (m OilTemperature) Type() uint { return OilTemperatureType }

func (m OilTemperature) GetName() string { return m.Name }
