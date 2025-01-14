package multierr

import "errors"

type Builder interface {
	Add(error)
	Err() error
}

type builder struct {
	errs []error
}

func (b *builder) Add(err error) {
	b.errs = append(b.errs, err)
}

func (b *builder) Err() error {
	return errors.Join(b.errs...)
}

func New() Builder {
	return &builder{
		errs: make([]error, 0),
	}
}
