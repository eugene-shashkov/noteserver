package handlers

type Handlers struct{}

type HandlersLoad interface {
}

func (h Handlers) CreateHandlers(hd *Handlers) HandlersLoad {
	return hd
}
