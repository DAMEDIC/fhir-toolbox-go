package model

import "fmt"

type Release interface {
	fmt.Stringer
	Version() string

	// this seals the interface, i.e. prevents other implementations of this interface from outside the package
	// necessary to make sure the assumption in `Generic[R model.Release](api any)` holds true
	isRelease()
}

type R4 struct{}

func (r R4) String() string {
	return "R4"
}

func (r R4) Version() string {
	return "4.0"
}

func (r R4) isRelease() {}

type R4B struct{}

func (r R4B) String() string {
	return "R4B"
}

func (r R4B) Version() string {
	return "4.3"
}

func (r R4B) isRelease() {}

type R5 struct{}

func (r R5) String() string {
	return "R5"
}

func (r R5) Version() string {
	return "5.0"
}
func (r R5) isRelease() {}

func ReleaseName[R Release]() string {
	return (*new(R)).String()
}

func ReleaseVersion[R Release]() string {
	return (*new(R)).Version()
}
