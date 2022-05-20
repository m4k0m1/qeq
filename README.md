# github.com/m4k0m1/gitter/go/qeq

## Usage

### Initialize your module

```
$ go mod init qeq-demo
```

### Get the qeq module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/m4k0m1/gitter/go/qeq@v0.1.0
```

```go
package main

import (
    "fmt"
    "os"
    "github.com/m4k0m1/gitter/go/qeq"
)

func main() {

    s1, s2, err := qeq.Solveqeq(1, -6, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Solutions calculated: x = %f or x = %f\n", s1, s2)
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```