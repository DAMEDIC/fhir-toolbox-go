package model

type Resource interface {
	Element
	ResourceType() string
	ResourceId() (string, bool)
}
type Element interface {
	MemSize() int
}
