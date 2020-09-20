package models

type options struct {
	storable    bool
	actorizable bool
	list        *List
}

func Options() *options {
	options := &options{}

	return options
}

func (options *options) WithStorable(storable bool) *options {
	options.storable = storable

	return options
}

func (options *options) WithShouldStore(storable bool) *options {
	options.storable = storable

	return options
}

func (options *options) WithActoriazable(actorizable bool) *options {
	options.actorizable = actorizable

	return options
}

func (options *options) WithList(list *List) *options {
	options.list = list

	return options
}
