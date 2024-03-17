# 🛰️ GHTTP — HTTP Request Simplified 🛰️

[![CI](https://github.com/institute-atri/ghttp/actions/workflows/ci.yml/badge.svg)](https://github.com/institute-atri/ghttp/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/institute-atri/ghttp/graph/badge.svg?token=nR2sLEINBZ)](https://codecov.io/gh/institute-atri/ghttp)

Making simplified web requests, meeting ATRI needs.

# 🛰️ Usage

Example GET:

```go
package main

import (
	"github.com/institute-atri/ghttp"
)

func main() {
	var response = ghttp.GET("http://httpbin.org/get")

	// To obtain and print the entire Source Code of the site.
	println(response.BRaw)

	// To obtain and print the method used in the request.
	println(response.Method)
}
```

Example POST:

```go
package main

import (
	"github.com/institute-atri/ghttp"
)

func main() {
	var response = ghttp.POST("http://httpbin.org/get", "user=admin&password=123")

	// To obtain and print the entire Source Code of the site.
	println(response.BRaw)

	// To obtain and print the method used in the request.
	println(response.Method)
}
```

# 🛰️ License

This project is licensed under the **MIT License**—see the [LICENSE](LICENSE) file.

```text
MIT License

Copyright (c) 2024 ATRI - Advanced Technology Research Institute

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```