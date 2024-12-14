package utils

// Ptr returns a pointer to supplied value t.
//
// This is convenience function to get pointers to literals.
//
// I.e. write
//
//	x := utils.Ptr("hello")
//
// instead of
//
//	s := "hello"
//	x := &s
func Ptr[T any](t T) *T {
	return &t
}
