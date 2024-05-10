package model

type Release interface {
	Release() string
}

type R4 struct{}

func (r R4) Release() string {
	return "R4"
}
