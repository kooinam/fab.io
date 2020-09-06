package views

import "github.com/kooinam/fabio/helpers"

type options struct {
	shouldIncludeRootKey bool
	rootKey              string
	params               helpers.H
}

func Options() *options {
	options := &options{
		params: helpers.H{},
	}

	return options
}

func (options *options) WithRootKey(rootKey string) *options {
	options.shouldIncludeRootKey = true
	options.rootKey = rootKey

	return options
}

func (options *options) WithParams(params helpers.H) *options {
	options.params = params

	return options
}
