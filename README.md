# string

```bash
$ task module:init
task: [module:init] 
go mod init github.com/golang-oop/boolean

go: creating new go.mod: module github.com/golang-oop/boolean
go: to add module requirements and sums:
        go mod tidy
```

```bash
$ task module:requirements
```

## Usage

```golang
package main

import (
        "fmt"
        "github.com/golang-oop/string"
)

func main {
        s := String.New(`Hello World!`)
        fmt.Println(s.ToGoString())
        if r:=s.Split(` `); r.HasError() {
                fmt.Prints(
                        "%s\n",
                        r.Error().Message(),
                )
        } else {
                fmt.Prints(
                        "%+v",
                        r.Payload(),
                ) 
        }
}
```

