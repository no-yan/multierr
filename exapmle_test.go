package multierr_test

import (
	"errors"
	"fmt"

	"github.com/no-yan/multierr"
)

func ExampleNew() {
	m := multierr.New()

	reasons := []string{"1st error", "2nd error"}
	for _, reason := range reasons {
		err := errors.New(reason)
		m.Add(err)
	}

	fmt.Println(m.Err())
	// Output:
	// 1st error
	// 2nd error
}
