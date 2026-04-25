package movers

type MoverInfo interface {
	CreateNewMover(name string)
	GetMoverName() string
	GetMoverBatch() string
}

type Mover struct {
	id   string
	name string
}

func (m *Mover) CreateNewMover() (Mover, error) {
	return Mover{}, nil
}

func (m *Mover) GetMoverName() string {
	return m.name
}

func (m *Mover) GetMoverBatch() string {
	return m.id
}
