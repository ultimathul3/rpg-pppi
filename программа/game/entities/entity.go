package entities

type Entity struct {
	HP            int
	MP            int
	X, Y          int
	Width, Height int
	Speed         int
	Damage        int
	Armor         int
}

func (e *Entity) Move() {

}

func (e *Entity) Attack() {

}

func (e *Entity) Draw() {

}
