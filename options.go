package builder

type fillerOpt[T any] func(*filler[T])

func WithArrayCount(count int) fillerOpt[any] {
	return func(f *filler[any]) {
		f.arrayCount = count
	}
}

func WithSliceCount(count int) fillerOpt[any] {
	return func(f *filler[any]) {
		f.sliceCount = count
	}
}

func WithStringMinMax(min, max int) fillerOpt[any] {
	return func(f *filler[any]) {
		f.stringMin = min
		f.stringMax = max
	}
}

func WithRespectEmpty() fillerOpt[any] {
	return func(f *filler[any]) {
		f.respectEmpty = true
	}
}
