package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	got := Add(5,6)
	want := 11

	assert.Equal(t, got, want)
}