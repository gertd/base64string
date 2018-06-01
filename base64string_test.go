package base64string_test

import (
	"encoding/json"
	"fmt"
	"testing"

	b64 "github.com/gertd/base64string"
	"github.com/stretchr/testify/assert"
)

const (
	testValue = "This is a test string"
	b64Value  = "\"VGhpcyBpcyBhIHRlc3Qgc3RyaW5n\""
	masked    = "********"
)

var (
	tests = []struct {
		input    string // test input value
		expected string // test expected result
	}{
		{"", "\"\""},
		{testValue, b64Value},
	}
)

func TestRoundtrip(t *testing.T) {

	for _, test := range tests {

		bStr1 := b64.Base64String(test.input)

		buf, err := json.Marshal(&bStr1)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.expected, string(buf), "marshal buffer contains encode string")
		t.Logf("marshal %s", string(buf))

		var bStr2 b64.Base64String

		err = json.Unmarshal(buf, &bStr2)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.input, bStr2.Get(), "unmarshal")
		t.Logf("unmarshal %s", bStr2.Get())
	}
}

func TestString(t *testing.T) {

	for _, test := range tests {

		bStr1 := b64.Base64String(test.input)

		assert.Equal(t, masked, bStr1.String(), "explicit call to String() should return mashed result")
		assert.Equal(t, masked, fmt.Sprintf("%s", bStr1), "implicit call to String() should return masked")
		t.Logf("string %s -> %s ", test.input, bStr1.String())
	}
}

func TestGet(t *testing.T) {

	for _, test := range tests {

		bStr1 := b64.Base64String(test.input)

		assert.Equal(t, test.input, bStr1.Get(), "explicit call to Get() should return unmashed value")
		t.Logf("get %s -> %s ", test.input, bStr1.Get())
	}
}

func TestMarshalIndent(t *testing.T) {

	for _, test := range tests {

		bStr1 := b64.Base64String(test.input)

		b, err := json.MarshalIndent(bStr1, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.expected, string(b), "marshalindent should return masked value")
		t.Logf("%s -> %s", test.input, string(b))
	}
}
