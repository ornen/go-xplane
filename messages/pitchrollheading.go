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

type PitchRollHeadingMessage struct {
	Pitch           float64
	Roll            float64
	HeadingTrue     float64
	HeadingMagnetic float64
}

func NewPitchRollHeadingMessage(data []float32) PitchRollHeadingMessage {
	return PitchRollHeadingMessage{
		Pitch:           float64(data[0]),
		Roll:            float64(data[1]),
		HeadingTrue:     float64(data[2]),
		HeadingMagnetic: float64(data[3]),
	}
}

func (m PitchRollHeadingMessage) Type() uint {
	return PitchRollHeadingMessageType
}
