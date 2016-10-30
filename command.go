package xplane

type Command interface {
	Message
	Data() [8]float32
}
