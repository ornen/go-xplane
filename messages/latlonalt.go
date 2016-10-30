package messages

type LatLonAltMessage struct {
	Latitude          float32
	Longitude         float32
	AltitudeMSL       float32
	AltitudeAGL       float32
	AltitudeIndicated float32
}

func NewLatLonAltMessage(data []float32) LatLonAltMessage {
	return LatLonAltMessage{
		Latitude:          data[0],
		Longitude:         data[1],
		AltitudeMSL:       data[2] * feetToMeters,
		AltitudeAGL:       data[3] * feetToMeters,
		AltitudeIndicated: data[5] * feetToMeters,
	}
}

func (m LatLonAltMessage) Type() uint {
	return LatLonAltMessageType
}
