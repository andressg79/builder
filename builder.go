package builder

func Fill[T any](t T, opts ...fillerOpt[T]) (T, error) {
	f := &filler[T]{
		stringMax: 10,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f.fill(t)
}

type filler[T any] struct {
	arrayCount   int
	sliceCount   int
	stringMin    int
	stringMax    int
	respectEmpty bool
}

func (f filler[T]) fill(t T) (T, error) {
	return t, nil
}
