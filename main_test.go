package env2struct

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	StringField  string  `env:"STRING"`
	BoolField    bool    `env:"BOOL"`
	IntField     int     `env:"INT"`
	Int8Field    int8    `env:"INT8"`
	Int16Field   int16   `env:"INT16"`
	Int32Field   int32   `env:"INT32"`
	Int64Field   int64   `env:"INT64"`
	UintField    uint    `env:"UINT"`
	Uint8Field   uint8   `env:"UINT8"`
	Uint16Field  uint16  `env:"UINT16"`
	Uint32Field  uint32  `env:"UINT32"`
	Uint64Field  uint64  `env:"UINT64"`
	Float32Field float32 `env:"FLOAT32"`
	Float64Field float64 `env:"FLOAT64"`
	FirstNested  struct {
		String       string `env:"STRING"`
		SecondNested struct {
			String string `env:"STRING"`
		} `env:"SECOND_NESTED"`
	} `env:"FIRST_NESTED"`
}

var testInt = 8
var testFloat = 4.2

var testData = map[string]string{
	"STRING":                            "test_string",
	"BOOL":                              "true",
	"INT":                               fmt.Sprintf("%d", testInt),
	"INT8":                              fmt.Sprintf("%d", testInt),
	"INT16":                             fmt.Sprintf("%d", testInt),
	"INT32":                             fmt.Sprintf("%d", testInt),
	"INT64":                             fmt.Sprintf("%d", testInt),
	"UINT":                              fmt.Sprintf("%d", testInt),
	"UINT8":                             fmt.Sprintf("%d", testInt),
	"UINT16":                            fmt.Sprintf("%d", testInt),
	"UINT32":                            fmt.Sprintf("%d", testInt),
	"UINT64":                            fmt.Sprintf("%d", testInt),
	"FLOAT32":                           fmt.Sprintf("%f", testFloat),
	"FLOAT64":                           fmt.Sprintf("%f", testFloat),
	"FIRST_NESTED_STRING":               "first nested string",
	"FIRST_NESTED_SECOND_NESTED_STRING": "second nested string",
}

func TestParse(t *testing.T) {
	var s testStruct

	for key, value := range testData {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}

	err := Parse(&s)
	assert.NoError(t, err)

	assert.Equal(t, testData["STRING"], s.StringField)
	assert.True(t, s.BoolField)

	assert.Equal(t, testInt, s.IntField)
	assert.Equal(t, int8(testInt), s.Int8Field)
	assert.Equal(t, int16(testInt), s.Int16Field)
	assert.Equal(t, int32(testInt), s.Int32Field)
	assert.Equal(t, int64(testInt), s.Int64Field)

	assert.Equal(t, uint(testInt), s.UintField)
	assert.Equal(t, uint8(testInt), s.Uint8Field)
	assert.Equal(t, uint16(testInt), s.Uint16Field)
	assert.Equal(t, uint32(testInt), s.Uint32Field)
	assert.Equal(t, uint64(testInt), s.Uint64Field)

	assert.Equal(t, float32(testFloat), s.Float32Field)
	assert.Equal(t, testFloat, s.Float64Field)

	assert.Equal(t, testData["FIRST_NESTED_STRING"], s.FirstNested.String)
	assert.Equal(t, testData["FIRST_NESTED_SECOND_NESTED_STRING"], s.FirstNested.SecondNested.String)
}
