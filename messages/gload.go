// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package messages

type GLoadMessage struct {
	X float32
	Y float32
	Z float32
}

func NewGLoadMessage(data []float32) GLoadMessage {
	return GLoadMessage{
		X: data[5] * gravityInMetersPerSecondSecond,
		Y: data[6] * gravityInMetersPerSecondSecond,
		Z: -data[4] * gravityInMetersPerSecondSecond,
	}
}

func (m GLoadMessage) Type() uint {
	return GLoadMessageType
}
