package builder

type fillerOpt func(*filler)

func WithArrayCount(count int) fillerOpt {
	return func(f *filler) {
		f.arrayCount = count
	}
}

func WithSliceCount(count int) fillerOpt {
	return func(f *filler) {
		f.sliceCount = count
	}
}

func WithStringMinMax(min, max int) fillerOpt {
	return func(f *filler) {
		f.stringMin = min
		f.stringMax = max
	}
}

func WithRespectEmpty() fillerOpt {
	return func(f *filler) {
		f.respectEmpty = true
	}
}
