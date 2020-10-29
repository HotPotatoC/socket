package ws

// Actor struct containing information of actor
type Actor struct {
	id *string
}

// ID return id information
func (actor *Actor) ID() string {
	return *actor.id
}
