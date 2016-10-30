package messages

type SpeedMessage struct {
	IndicatedSpeed float32
	Airspeed       float32
	TrueAirspeed   float32
	GroundSpeed    float32
}

func NewSpeedMessage(data []float32) SpeedMessage {
	return SpeedMessage{
		IndicatedSpeed: data[0] * knotsToMetersPerSecond,
		Airspeed:       data[1] * knotsToMetersPerSecond,
		TrueAirspeed:   data[2] * knotsToMetersPerSecond,
		GroundSpeed:    data[3] * knotsToMetersPerSecond,
	}
}

func (m SpeedMessage) Type() uint {
	return SpeedMessageType
}
