# github.com/MahdadGhasemian/go-rabbitmq-hash

## Usage

### Initialize your module

```
$ go mod init example.com/my-hash-demo
```

### Get the go-rabbitmq-hash module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/MahdadGhasemian/go-rabbitmq-hash@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/MahdadGhasemian/go-rabbitmq-hash"
)

func main() {
    // random salt
    fmt.Println(gorabbitmqhash.Hash("passowrd",0))
    // salt = 1234
    fmt.Println(gorabbitmqhash.Hash("passowrd",uint32(1234)))
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
