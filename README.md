# env2struct

```go
package main

import (
	"os"

	"github.com/github.com/NomNes/env2struct"
)

type EnvStruct struct {
	StringField string `env:"STRING_ENV"`
	NestedField struct {
		IntField string `env:"INT_ENV"`
	} `env:"NESTED"`
}

func Main() {
	_ = os.Setenv("STRING_ENV", "example string")
	_ = os.Setenv("NESTED_INT_ENV", "12")

	var envStruct EnvStruct
	err := env2struct.Parse(&envStruct)
	if err != nil {
		panic(err)
	}
}
```

## Supported types:
* string
* bool
* int
* int8
* int16
* int32
* int64
* uint
* uint8
* uint16
* uint32
* uint64
* float32
* float64
