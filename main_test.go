package multierr_test

import (
	"errors"
	"testing"

	"github.com/no-yan/multierr"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		errs []string
		want string
	}{
		{"multi err", []string{"a", "b", "c"}, "a\nb\nc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := multierr.New()
			for i := range tt.errs {
				m.Add(errors.New(tt.errs[i]))
			}

			got := m.Err().Error()
			if got != tt.want {
				t.Errorf("New() = %v, want %v", m, tt.want)
			}
		})
	}
}

func TestNew_Nil(t *testing.T) {
	m := multierr.New()
	if got := m.Err(); got != nil {
		t.Errorf("should return nil if error is empty, got: %v", got)
	}
}

func TestNew_Is(t *testing.T) {
	m := multierr.New()
	subErr := errors.New("suberr")
	m.Add(subErr)

	merr := m.Err()
	if !errors.Is(merr, subErr) {
		t.Errorf("should contain subErr in err tree: %#v", merr)
	}
}
