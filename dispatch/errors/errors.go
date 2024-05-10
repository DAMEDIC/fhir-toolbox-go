package errors

import "fmt"

type NotImplementedError struct {
	Interaction  string
	ResourceType string
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("%s%s not implemented for resource type %s", e.Interaction, e.ResourceType, e.ResourceType)
}

type UnknownResourceError struct {
	ResourceType string
}

func (e UnknownResourceError) Error() string {
	return fmt.Sprintf("unknown resource type %s", e.ResourceType)
}
