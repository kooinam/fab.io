package actors

type options struct {
	shouldUpdate bool
}

func Options() *options {
	options := &options{}

	return options
}

func (options *options) WithShouldUpdate(shouldUpdate bool) *options {
	options.shouldUpdate = shouldUpdate

	return options
}
