package messages

type PitchRollHeadingMessage struct {
	Pitch           float32
	Roll            float32
	HeadingTrue     float32
	HeadingMagnetic float32
}

func NewPitchRollHeadingMessage(data []float32) PitchRollHeadingMessage {
	return PitchRollHeadingMessage{
		Pitch:           data[0],
		Roll:            data[1],
		HeadingTrue:     data[2],
		HeadingMagnetic: data[3],
	}
}

func (m PitchRollHeadingMessage) Type() uint {
	return PitchRollHeadingMessageType
}
