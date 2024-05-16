package model

type Resource interface {
	ResourceType() string
	ResourceId() (string, bool)
}
type Element interface{}
