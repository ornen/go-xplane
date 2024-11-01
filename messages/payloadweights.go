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

type PayloadWeights struct {
	MessageData
	Load1 float64 //lb
	Load2 float64 //lb
	Load3 float64 //lb
	Load4 float64
	Load5 float64
	Load6 float64
	Load7 float64
	Load8 float64
}

func NewPayloadWeights(data []float32) PayloadWeights {
	return PayloadWeights{
		MessageData: MessageData{Name: "PayloadWeights"},
		Load1:       float64(data[0]),
		Load2:       float64(data[1]),
		Load3:       float64(data[2]),
		Load4:       float64(data[3]),
		Load5:       float64(data[4]),
		Load6:       float64(data[5]),
		Load7:       float64(data[6]),
		Load8:       float64(data[7]),
	}
}

func (m PayloadWeights) Type() uint { return PayloadWeightsType }

func (m PayloadWeights) GetName() string { return m.Name }
