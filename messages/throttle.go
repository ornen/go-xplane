package messages

const (
	ThrottleCommandMessageType = 25
)

type ThrottleCommandMessage struct {
	Throttle float32
}

func (c ThrottleCommandMessage) Type() uint {
	return ThrottleCommandMessageType
}

func (c ThrottleCommandMessage) Data() [8]float32 {
	return [8]float32{
		c.Throttle,
		c.Throttle,
		c.Throttle,
		c.Throttle,
		-999,
		-999,
		-999,
		-999,
	}
}
