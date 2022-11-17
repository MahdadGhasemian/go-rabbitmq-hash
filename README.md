# github.com/MahdadGhasemian/go-rabbitmq-hash

## Usage

### Initialize your module

```
$ go mod init example.com/my-hash-demo
```

### Get the go-rabbitmq-hash module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/MahdadGhasemian/go-rabbitmq-hash@v0.1.1
```

```go
package main

import (
    "fmt"

    "github.com/MahdadGhasemian/go-rabbitmq-hash"
)

func main() {
    // random salt
	fmt.Println(rabbitmqHash.Hash("passowrd", 0))
	// salt = 1234
	fmt.Println(rabbitmqHash.Hash("passowrd", 1234))
	// salt = 0x908DC60A
	fmt.Println(rabbitmqHash.Hash("test12", 0x908DC60A))
}
```

## Testing

```
$ go test
```
