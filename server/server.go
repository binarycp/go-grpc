package server

type Server interface {
	Run() error
	Close() error
}

type parent struct{}

func (parent) Run() error {
	return nil
}

func (parent) Close() error {
	return nil
}
