package notify

type Event uint8

const (
	Create Event = 1 << iota
	Write
	Remove
	Rename
	Recursive
)

type EventInfo interface {
	Name() string
	Event() Event
	Sys() interface{}
}

func Notify(name string, c chan<- EventInfo, events ...Event) {
	impl.Notify(name, c, events...)
}

func Stop(c chan<- EventInfo) {
	impl.Stop(c)
}

var impl interface {
	Notify(string, chan<- EventInfo, ...Event)
	Stop(chan<- EventInfo)
}
