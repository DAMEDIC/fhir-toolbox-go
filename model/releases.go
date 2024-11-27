package model

type Release interface {
	isRelease()
}

type R4 struct{}

func (r R4) isRelease() {}

func ReleaseName[R Release]() string {
	var r R
	switch any(r).(type) {
	case R4:
		return "R4"
	default:
		panic("unsupported release")
	}
}

func ReleaseVersion[R Release]() string {
	var r R
	switch any(r).(type) {
	case R4:
		return "4.0.1"
	default:
		panic("unsupported release")
	}
}
