package components

type Player struct {
	Mark string
	Name   string
}

func CreatePlayer(name, mark string) *Player {
	return &Player{Mark: mark, Name: name}
}
