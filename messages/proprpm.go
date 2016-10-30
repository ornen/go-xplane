package messages

type PropRPMMessage struct {
	RPM float32
}

func NewPropRPMMessage(data []float32) PropRPMMessage {
	return PropRPMMessage{
		RPM: data[0],
	}
}

func (m PropRPMMessage) Type() uint {
	return PropRPMMessageType
}
