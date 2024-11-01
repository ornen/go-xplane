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

type WeightAndBalance struct {
	MessageData
	Empty   float64 //lb
	Payload float64
	Fuel    float64
	Jetti   float64
	Ice     float64
	Total   float64
	Gross   float64
	CG      float64 //ftref
}

func NewWeightAndBalance(data []float32) WeightAndBalance {
	return WeightAndBalance{
		MessageData: MessageData{Name: "WeightAndBalance"},
		Empty:       float64(data[0]),
		Payload:     float64(data[1]),
		Fuel:        float64(data[2]),
		Jetti:       float64(data[3]),
		Ice:         float64(data[4]),
		Total:       float64(data[5]),
		Gross:       float64(data[6]),
		CG:          float64(data[7]),
	}
}

func (m WeightAndBalance) Type() uint { return WeightAndBalanceType }

func (m WeightAndBalance) GetName() string { return m.Name }
