package util

// Unwrap works like the Rust syntax sugar Result<...>.unwrap().
// It panics if err != nil.
func Unwrap[V any](v V, err error) V {
	if err != nil {
		panic(err)
	}
	return v
}

// Whatever is like Unwrap, except it returns a
// zero value if err != nil.
func Whatever[V any](v V, err error) (vret V) {
	if err != nil {
		return
	}
	return v
}
