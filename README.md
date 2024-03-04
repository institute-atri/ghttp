<h1 align="center">GNET</h1>
<h4 align="center">Golang package for making simplified web requests, meeting ATRI's needs.</h4>

Example GET:

```go
package main

import (
    "github.com/institute-atri/gnet"
    fmt
)

func main() {
    var response = gnet.GET("http://httpbin.org/get")	
	
    // To obtain and print the entire Source Code of the site.
    fmt.Println(response.BRaw)

    // To obtain and print the method used in the request.
    fmt.Println(response.Method)
}
```
Example POST:
```go
package main

import (
    "github.com/institute-atri/gnet"
    fmt
)

func main() {
    var response = gnet.POST("http://httpbin.org/get", "user=admin&password=123")
	
    // To obtain and print the entire Source Code of the site.
    fmt.Println(response.BRaw)

    // To obtain and print the method used in the request.
    fmt.Println(response.Method)
}
```


<h5 align="center">ATRI@2024 | License <a href="https://github.com/institute-atri/gnet/blob/main/LICENSE">MIT</a>.</h5>