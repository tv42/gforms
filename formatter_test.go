package gforms_test

import (
	. "launchpad.net/gocheck"

	"github.com/vmihailenco/gforms"
)

type FormatterTest struct{}

var _ = Suite(&FormatterTest{})

func (t *FormatterTest) TestSplitWords(c *C) {
	table := []struct {
		s     string
		words []string
	}{
		{"FooBar", []string{"Foo", "Bar"}},
		{"HTTP", []string{"HTTP"}},
		{"HTTPReq", []string{"HTTP", "Req"}},
		{"HTTPReqX", []string{"HTTP", "Req", "X"}},
	}

	for _, row := range table {
		c.Assert(gforms.SplitWords(row.s), DeepEquals, row.words)
	}
}