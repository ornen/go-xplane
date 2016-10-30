package messages

type BatteryAmperageMessage struct {
	Amperage float32
}

func NewBatteryAmperageMessage(data []float32) BatteryAmperageMessage {
	return BatteryAmperageMessage{
		Amperage: data[0],
	}
}

func (m BatteryAmperageMessage) Type() uint {
	return BatteryAmperageMessageType
}
