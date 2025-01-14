// package multierr provides a simple way to collect errors
// and return them as a single error using errors.Join.
package multierr

import "errors"

// Collector collects multiple errors and returns them as a single error.
// Errors can be added by [multierr.Collector.Add], and the final combined error is obtained by [multierr.Collector.Err].
//
// The resulting error supports [errors.Is] and [errors.As] checks
// for any of the collected errors.
type Collector interface {
	// Add appends an error to the collection.
	Add(error)
	// Err returns all collected errors as a single error, by calling errors.Join.
	// If no errors have been added, Err returns nil.
	Err() error
}

type collector struct {
	errs []error
}

func (c *collector) Add(err error) {
	c.errs = append(c.errs, err)
}

func (c *collector) Err() error {
	return errors.Join(c.errs...)
}

// New creates a new Collector.
//
// This function is not concurrency-safe: if you plan to use the collector
// from multiple goroutines, consider wrapping calls with a mutex or
// using your own concurrency control.
func New() Collector {
	return &collector{
		errs: make([]error, 0),
	}
}
