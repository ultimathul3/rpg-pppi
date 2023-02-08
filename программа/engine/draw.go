package engine

type Drawer interface {
	Draw()
}

func Draw(d Drawer) {
	d.Draw()
}
