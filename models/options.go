package models

type options struct {
	shouldStore bool
	list        *List
}

func Options() *options {
	options := &options{}

	return options
}

func (options *options) WithShouldStore(shouldStore bool) *options {
	options.shouldStore = true

	return options
}

func (options *options) WithList(list *List) *options {
	options.list = list

	return options
}
