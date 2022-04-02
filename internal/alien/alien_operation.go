package alien

func New(Id int) Alien {
	var e Alien = Alien{Id}
	return e
}

func (a *Alien) Move() {
	a.Id = 90
}
