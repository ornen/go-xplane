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
