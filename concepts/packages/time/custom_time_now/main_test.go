package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type FakeClock struct {
	t time.Time
}

func (fake FakeClock) Now() time.Time {
	return fake.t
}

func Test_generateClock(t *testing.T) {
	// tests := []struct {
	// 	name string // description of this test case
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		main()
	// 	})
	// }

	fake := FakeClock{
		t: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	res := GenerateClock(fake)
	assert.Equal(t, "haha", res)
}
