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

package xplane

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"net"
	"time"

	"github.com/ornen/go-xplane/messages"
)

const (
	datagramPrefixLength = 5
	messageLength        = 36
)

type XPlane struct {
	RemoteAddress string
	LocalAddress  string
	Messages      chan Message
	connection    *net.UDPConn
	receivePeriod *time.Duration
}

type Opt func(*XPlane)

func New(remoteAddress, localAddress string, opts ...Opt) XPlane {
	x := XPlane{
		RemoteAddress: remoteAddress,
		LocalAddress:  localAddress,
		Messages:      make(chan Message),
	}

	for _, o := range opts {
		o(&x)
	}

	return x
}

func ReceiveEvery(t time.Duration) Opt {
	return func(x *XPlane) {
		x.receivePeriod = &t
	}
}

func (x *XPlane) Receive() error {
	serverAddr, err := net.ResolveUDPAddr("udp", x.LocalAddress)
	serverConn, err := net.ListenUDP("udp", serverAddr)

	if err != nil {
		return err
	}

	defer serverConn.Close()

	buf := make([]byte, 1024)
	if x.receivePeriod != nil {
		t := time.NewTicker(*x.receivePeriod)
		for {
			select {
			case <-t.C:
				x.readBuf(serverConn, buf)
			}
		}
	} else {
		for {
			x.readBuf(serverConn, buf)
		}
	}

	return nil
}

func (x *XPlane) readBuf(c *net.UDPConn, b []byte) {
	n, _, _ := c.ReadFromUDP(b)
	m := (n - datagramPrefixLength) / messageLength

	for i := 0; i < m; i++ {
		sentence := b[datagramPrefixLength+i*messageLength : datagramPrefixLength+(i+1)*messageLength]
		x.parse(sentence)
	}
}

func (x *XPlane) parse(sentence []byte) {
	messageType := sentence[0]
	messageBuffer := bytes.NewBuffer(sentence[4:])

	messageData := make([]float32, 8)
	binary.Read(messageBuffer, binary.LittleEndian, &messageData)

	switch messageType {
	case messages.AircraftPointWeatherPrecipWindType:
		x.Messages <- messages.NewAircraftPointWeatherPrecipWind(messageData)
	case messages.AircraftPointWeatherTempPressType:
		x.Messages <- messages.NewAircraftPointWeatherTempPress(messageData)
	case messages.AircraftRunwayConditionsType:
		x.Messages <- messages.NewAircraftRunwayConditions(messageData)
	case messages.EnginePowerType:
		x.Messages <- messages.NewEnginePower(messageData)
	case messages.EngineThrustType:
		x.Messages <- messages.NewEngineThrust(messageData)
	case messages.EngineTorqueType:
		x.Messages <- messages.NewEngineTorque(messageData)
	case messages.FrameRateType:
		x.Messages <- messages.NewFrameRate(messageData)
	case messages.FuelFlowType:
		x.Messages <- messages.NewFuelFlow(messageData)
	case messages.FuelPressureType:
		x.Messages <- messages.NewFuelPressure(messageData)
	case messages.FuelWeightsType:
		x.Messages <- messages.NewFuelWeights(messageData)
	case messages.N1Type:
		x.Messages <- messages.NewN1(messageData)
	case messages.N2Type:
		x.Messages <- messages.NewN2(messageData)
	case messages.OilPressureType:
		x.Messages <- messages.NewOilPressure(messageData)
	case messages.OilTemperatureType:
		x.Messages <- messages.NewOilTemperature(messageData)
	case messages.PayloadWeightsType:
		x.Messages <- messages.NewPayloadWeights(messageData)
	case messages.SimStatsType:
		x.Messages <- messages.NewSimStats(messageData)
	case messages.ThrottleActualType:
		x.Messages <- messages.NewThrottleActual(messageData)
	case messages.SpeedMessageType:
		x.Messages <- messages.NewSpeedMessage(messageData)
	case messages.GLoadMessageType:
		x.Messages <- messages.NewGLoadMessage(messageData)
	case messages.AngularVelocitiesMessageType:
		x.Messages <- messages.NewAngularVelocitiesMessage(messageData)
	case messages.PitchRollHeadingMessageType:
		x.Messages <- messages.NewPitchRollHeadingMessage(messageData)
	case messages.FlightControlMessageType:
		x.Messages <- messages.NewFlightControlMessage(messageData)
	case messages.GearsBrakesMessageType:
		x.Messages <- messages.NewGearsBrakesMessage(messageData)
	case messages.WeatherMessageType:
		x.Messages <- messages.NewWeatherMessage(messageData)
	case messages.LatLonAltMessageType:
		x.Messages <- messages.NewLatLonAltMessage(messageData)
	case messages.LocVelDistTraveledMessageType:
		x.Messages <- messages.NewLocVelDistTraveledMessage(messageData)
	case messages.BatteryAmperageMessageType:
		x.Messages <- messages.NewBatteryAmperageMessage(messageData)
	case messages.BatteryVoltageMessageType:
		x.Messages <- messages.NewBatteryVoltageMessage(messageData)
	case messages.EngineRPMMessageType:
		x.Messages <- messages.NewEngineRPMMessage(messageData)
	case messages.PropRPMMessageType:
		x.Messages <- messages.NewPropRPMMessage(messageData)
	case messages.PropPitchMessageType:
		x.Messages <- messages.NewPropPitchMessage(messageData)
	case messages.AngleOfAttackSideslipMessageType:
		x.Messages <- messages.NewAngleOfAttackSideslipMessage(messageData)
	default:
		log.Println("Unknown message type: ", messageType)
	}
}

func (x *XPlane) Connect() error {
	udpAddr, err := net.ResolveUDPAddr("udp", x.RemoteAddress)
	if err != nil {
		return errors.New("wrong address!")
	}

	x.connection, err = net.DialUDP("udp", nil, udpAddr)

	return err
}

func (x *XPlane) Send(command Command) error {
	commandData := command.Data()

	buf := new(bytes.Buffer)

	buf.Write([]byte{'D', 'A', 'T', 'A', 0})
	buf.Write([]byte{byte(command.Type()), 0, 0, 0})

	if err := binary.Write(buf, binary.LittleEndian, &commandData); err != nil {
		return err
	}

	_, err := x.connection.Write(buf.Bytes())

	return err
}
