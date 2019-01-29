package person

// Person component of world
type Person struct {
	Identify uint64
	Channel  chan []byte
}
