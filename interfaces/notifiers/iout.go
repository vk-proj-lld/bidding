package notifiers

type IOut interface {
	Write(contents ...interface{})
}
