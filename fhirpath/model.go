package fhirpath

type Element interface {
	Member(string) Collection
}

type Collection []Element

func (c Collection) Member(name string) Collection {
	var members Collection
	for _, e := range c {
		members = append(members, e.Member(name)...)
	}
	return members
}
