package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		desc      string
		a, b, got int
	}{
		{desc: "add", a: 0, b: 1, got: 1},
		{desc: "add", a: 1, b: 2, got: 3},
		{desc: "add", a: 2, b: 3, got: 5},
		{desc: "add", a: 3, b: 5, got: 8},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, Add(tc.a, tc.b), tc.got)
		})
	}
}
