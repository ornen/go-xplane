package messages

type EngineRPMMessage struct {
	RPM float32
}

func NewEngineRPMMessage(data []float32) EngineRPMMessage {
	return EngineRPMMessage{
		RPM: data[0],
	}
}

func (m EngineRPMMessage) Type() uint {
	return EngineRPMMessageType
}
