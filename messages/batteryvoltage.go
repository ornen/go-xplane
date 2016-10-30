package messages

type BatteryVoltageMessage struct {
	Voltage float32
}

func NewBatteryVoltageMessage(data []float32) BatteryVoltageMessage {
	return BatteryVoltageMessage{
		Voltage: data[0],
	}
}

func (m BatteryVoltageMessage) Type() uint {
	return BatteryVoltageMessageType
}
