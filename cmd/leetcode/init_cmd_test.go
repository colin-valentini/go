package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformCase(t *testing.T) {
	testCases := []struct {
		slug, camel, title string
	}{
		{slug: "", camel: "", title: ""},
		{slug: "foo", camel: "foo", title: "Foo"},
		{slug: "foo-bar-baz", camel: "fooBarBaz", title: "FooBarBaz"},
		{slug: "foobarbaz", camel: "foobarbaz", title: "Foobarbaz"},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.camel, transformCase(tc.slug, camelCase))
		assert.Equal(t, tc.title, transformCase(tc.slug, titleCase))
	}
}
