package model

import "fmt"

type Release interface {
	fmt.Stringer
	Version() string
}

type R4 struct{}

func (r R4) String() string {
	return "R4"
}

func (r R4) Version() string {
	return "4.0"
}

type R4B struct{}

func (r R4B) String() string {
	return "R4B"
}

func (r R4B) Version() string {
	return "4.3"
}

type R5 struct{}

func (r R5) String() string {
	return "R5"
}

func (r R5) Version() string {
	return "5.0"
}

func ReleaseName[R Release]() string {
	return (*new(R)).String()
}

func ReleaseVersion[R Release]() string {
	return (*new(R)).Version()
}
