package multierr_test

import (
	"errors"
	"fmt"

	"github.com/no-yan/multierr"
)

func ExampleNew() {
	merr := multierr.New()

	reasons := []string{"1st error", "2nd error"}
	for _, reason := range reasons {
		err := errors.New(reason)
		merr.Add(err)
	}

	fmt.Println(merr.Err())
	// Output:
	// 1st error
	// 2nd error
}
