package ptr

// To returns a pointer to supplied value t.
//
// This is convenience function to get pointers to literals.
//
// I.e. write
//
//	x := ptr.To("hello")
//
// instead of
//
//	s := "hello"
//	x := &s
func To[T any](t T) *T {
	return &t
}
