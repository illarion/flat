# Flat 

Minimalistic implementation of the `Flatten(map[string]interface{})` and `Unflatten(map[string]interface{})` functions.

## Description

`Flatten` converts the nested map into a flat map, where the keys are the paths to the values in the original map. The path is a dot-separated list of keys. The values are the values of the original map.

`Unflatten` converts the flat map into a nested map, where the keys are the paths to the values in the original map. The path is a dot-separated list of keys. The values are the values of the original map.

Both functions use `javascript` style to represent arrays in the path. For example, the path `a.b[0].c` will be converted into the following structure:

```go
map[string]interface{}{
    "a": map[string]interface{}{
        "b": []interface{}{
            map[string]interface{}{
                "c": <value>,
            },
        },
    },
}
```

## Usage

```go

package main

import (
    "fmt"
    "github.com/illarion/flat"
    "encoding/json"
)

func main() {
    var src map[string]interface{}

    json.Parse([]byte(`{"a": {"b": [{"c": 1}, {"c": 2}]}}`), &src)

    flattened := flat.Flatten(src, &flat.Options{
        Delimiter: ".",
    })
    fmt.Sprintf("%#v", flattened)

    unflat := flat.Unflatten(flattened, &flat.Options{
        Delimiter: ".",
    })
    fmt.Sprintf("%#v", unflattened)
}
```