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

type FrameRate struct {
	MessageData
	FPS       float64 //Frames per second
	UPS       float64 //Simulation Update per second
	FrameTime float64 //ms
	CPUUsage  float64 //time
	GPUUsage  float64 //time
	GRND      float64 //ratio
	Flit      float64 //ratio
}

func NewFrameRate(data []float32) FrameRate {
	return FrameRate{
		MessageData: MessageData{
			Name: "FrameRate",
		},
		FPS:       float64(data[0]),
		UPS:       float64(data[1]),
		FrameTime: float64(data[2]),
		CPUUsage:  float64(data[3]),
		GPUUsage:  float64(data[4]),
		GRND:      float64(data[5]),
		Flit:      float64(data[6]),
	}
}

func (m FrameRate) Type() uint { return FrameRateType }

func (m FrameRate) GetName() string { return m.Name }
