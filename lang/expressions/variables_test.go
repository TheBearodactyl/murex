package expressions

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type createIndexBlockTestT struct {
	Name     string
	Index    string
	Flags    string
	Expected string
}

func TestCreateIndexBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
		},
		{
			Name:  "",
			Index: "",
		},
		{
			Name:  "a",
			Index: "b",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> [" + test.Index + "]"
		}

		getIorE := &getVarIndexOrElementT{varName: []rune(test.Name), key: []rune(test.Index)}
		block := createIndexBlock(getIorE)
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}

func TestCreateElementBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
		},
		{
			Name:  "",
			Index: "",
		},
		{
			Name:  "a",
			Index: "b",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> [[" + test.Index + "]]"
		}

		getIorE := &getVarIndexOrElementT{varName: []rune(test.Name), key: []rune(test.Index)}
		block := createElementBlock(getIorE)
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}

func TestCreateRangeBlock(t *testing.T) {
	tests := []createIndexBlockTestT{
		{
			Name:  "foobar",
			Index: "hello world",
			Flags: "abc",
		},
		{
			Name:  "",
			Index: "",
			Flags: "",
		},
		{
			Name:  "a",
			Index: "b",
			Flags: "c",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		if test.Expected == "" {
			test.Expected = "$" + test.Name + "-> @[" + test.Index + "]" + test.Flags
		}

		block := createRangeBlock([]rune(test.Name), []rune(test.Index), []rune(test.Flags))
		if string(block) != test.Expected {
			t.Errorf("block does not match expected in test %d", i)
			t.Logf("  Name:     '%s'", test.Name)
			t.Logf("  Index:    '%s'", test.Index)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", string(block))
		}
	}
}
