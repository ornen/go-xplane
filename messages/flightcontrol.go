package messages

type FlightControlMessage struct {
	Elevator float32
	Aileron  float32
	Rudder   float32
}

func (c FlightControlMessage) Type() uint {
	return FlightControlMessageType
}

func (c FlightControlMessage) Data() [8]float32 {
	return [8]float32{
		c.Elevator,
		c.Aileron,
		c.Rudder,
		-999,
		c.Rudder,
		-999,
		-999,
		-999,
	}
}
