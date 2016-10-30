package messages

const (
	AngularVelocitiesMessageType = 16
)

type AngularVelocitiesMessage struct {
	X float32
	Y float32
	Z float32
}

func NewAngularVelocitiesMessage(data []float32) AngularVelocitiesMessage {
	return AngularVelocitiesMessage{
		X: data[1],
		Y: data[0],
		Z: data[2],
	}
}

func (m AngularVelocitiesMessage) Type() uint {
	return AngularVelocitiesMessageType
}
