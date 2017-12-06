# Go Library for the Alluxio File System API

To get started, fetch the repository into your Go workspace using `go get -d github.com/Alluxio/alluxio-go`.

Your Go applications can then interact with Alluxio as follows:

```
package main

import (
        "fmt"
	"log"

	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
)


func main() {
	fs := alluxio.NewClient(<proxy-host>, <proxy-port>, <timeout>)
	ok, err := fs.Exists("/test_path", &option.Exists{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}
```

For a list of all supported API calls, see https://github.com/Alluxio/alluxio-go/blob/master/paths.go.
