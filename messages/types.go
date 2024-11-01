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

const (
	FrameRateType                      = 0
	SimStatsType                       = 2
	SpeedMessageType                   = 3
	GLoadMessageType                   = 4
	WeatherMessageType                 = 6
	FlightControlMessageType           = 11
	TrimFlapsBrakesMessageType         = 13
	GearsBrakesMessageType             = 14
	AngularVelocitiesMessageType       = 16
	PitchRollHeadingMessageType        = 17
	AngleOfAttackSideslipMessageType   = 18
	PayloadWeightsType                 = 19
	LatLonAltMessageType               = 20
	LocVelDistTraveledMessageType      = 21
	ThrottleCommandMessageType         = 25
	ThrottleActualType                 = 26
	EnginePowerType                    = 34
	EngineThrustType                   = 35
	EngineTorqueType                   = 36
	EngineRPMMessageType               = 37
	PropRPMMessageType                 = 38
	PropPitchMessageType               = 39
	N1Type                             = 41
	N2Type                             = 42
	FuelFlowType                       = 45
	OilPressureType                    = 49
	OilTemperatureType                 = 50
	FuelPressureType                   = 51
	BatteryAmperageMessageType         = 53
	BatteryVoltageMessageType          = 54
	FuelWeightsType                    = 62
	WeightAndBalanceType               = 63
	PropellerTorqueType                = 145
	AircraftPointWeatherTempPressType  = 151
	AircraftPointWeatherPrecipWindType = 152
	AircraftRunwayConditionsType       = 166
)
