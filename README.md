# Go CORS handler 

CORS is a `net/http` handler implementing [Cross Origin Resource Sharing W3 specification](http://www.w3.org/TR/cors/) in Golang.

This is a fork of https://github.com/rs/cors with some minor performance improvements and more opinionated default configuration.

The main changes
- Cleaner API, kept only one option for allowing an origin dynamically: `AllowOriginVaryRequestFunc`
- Updated default allowed HTTP methods: `GET`, `POST`, `PATCH`, `PUT`, `DELETE`
- Updated default allowed headers: `Accept`, `Content-Type`
- Removed support for wildcard `*` in `AllowedOrigins` to avoid accidental insecure configurations

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file. We'll call it `server.go`.

```go
package main

import (
    "net/http"

    "github.com/stfsy/go-cors"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("{\"hello\": \"world\"}"))
    })

    // cors.Default() sets up the middleware with sensible defaults (simple
    // methods and default allowed headers). Note: by design the library does
    // not enable a match-all origin by default; you must set explicit
    // `AllowedOrigins` or provide an `AllowOriginVaryRequestFunc`.
    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)
}
```

Install `cors`:

    go get github.com/stfsy/go-cors

Then run your server:

    go run server.go

The server now runs on `localhost:8080`:

    $ curl -D - -H 'Origin: http://foo.com' http://localhost:8080/
    HTTP/1.1 200 OK
    Access-Control-Allow-Origin: foo.com
    Content-Type: application/json
    Date: Sat, 25 Oct 2014 03:43:57 GMT
    Content-Length: 18

    {"hello": "world"}

### Allowing all origins intentionally

The library does not treat `"*"` in `AllowedOrigins` as a special match-all token anymore. If you intentionally want to allow every origin, provide an explicit function such as `AllowOriginVaryRequestFunc`:

```go
AllowOriginVaryRequestFunc: func(r *http.Request, origin string) (bool, []string) { return true, nil },
```

This makes the behavior explicit and avoids accidental insecure configurations (for example, pairing `"*"` with `AllowCredentials: true`).

### More Examples

* `net/http`: [examples/nethttp/server.go](https://github.com/stfsy/go-cors/blob/master/examples/nethttp/server.go)
* [Goji](https://goji.io): [examples/goji/server.go](https://github.com/stfsy/go-cors/blob/master/examples/goji/server.go)
* [Martini](http://martini.codegangsta.io): [examples/martini/server.go](https://github.com/stfsy/go-cors/blob/master/examples/martini/server.go)
* [Negroni](https://github.com/codegangsta/negroni): [examples/negroni/server.go](https://github.com/stfsy/go-cors/blob/master/examples/negroni/server.go)
* [Alice](https://github.com/justinas/alice): [examples/alice/server.go](https://github.com/stfsy/go-cors/blob/master/examples/alice/server.go)
* [HttpRouter](https://github.com/julienschmidt/httprouter): [examples/httprouter/server.go](https://github.com/stfsy/go-cors/blob/master/examples/httprouter/server.go)
* [Gorilla](http://www.gorillatoolkit.org/pkg/mux): [examples/gorilla/server.go](https://github.com/stfsy/go-cors/blob/master/examples/gorilla/server.go)
* [Buffalo](https://gobuffalo.io): [examples/buffalo/server.go](https://github.com/stfsy/go-cors/blob/master/examples/buffalo/server.go)
* [Gin](https://gin-gonic.github.io/gin): [examples/gin/server.go](https://github.com/stfsy/go-cors/blob/master/examples/gin/server.go)
* [Chi](https://github.com/go-chi/chi): [examples/chi/server.go](https://github.com/stfsy/go-cors/blob/master/examples/chi/server.go)

## Parameters

Parameters are passed to the middleware thru the `cors.New` method as follow:

```go
c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://foo.com", "http://foo.com:8080"},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
})

// Insert the middleware
handler = c.Handler(handler)
```

* **AllowedOrigins** `[]string`: A list of origins a cross-domain request can be executed from. An origin may contain a wildcard (`*`) to replace 0 or more characters (i.e.: `http://*.domain.com`). Usage of wildcards implies a small performance penality. Only one wildcard can be used per origin. By design the library does not treat the literal `"*"` as a special match-all token; if you need to allow all origins, provide an explicit `AllowOriginVaryRequestFunc`.
* **AllowOriginVaryRequestFunc** `func(r *http.Request, origin string) (bool, []string)`: A custom function to validate the origin. It takes the HTTP Request object and the origin as argument and returns true if allowed or false otherwise with a list of headers used to take that decision if any so they can be added to the Vary header. If this option is set, the contents of `AllowedOrigins` are ignored.
* **AllowedMethods** `[]string`: A list of methods the client is allowed to use with cross-domain requests. Default value is simple methods (`GET` and `POST`).
* **AllowedHeaders** `[]string`: A list of non simple headers the client is allowed to use with cross-domain requests.
* **ExposedHeaders** `[]string`: Indicates which headers are safe to expose to the API of a CORS API specification.
* **AllowCredentials** `bool`: Indicates whether the request can include user credentials like cookies, HTTP authentication or client side SSL certificates. The default is `false`.
* **AllowPrivateNetwork** `bool`: Indicates whether to accept cross-origin requests over a private network.
* **MaxAge** `int`: Indicates how long (in seconds) the results of a preflight request can be cached. The default is `0` which stands for no max age.
* **OptionsPassthrough** `bool`: Instructs preflight to let other potential next handlers to process the `OPTIONS` method. Turn this on if your application handles `OPTIONS`.
* **OptionsSuccessStatus** `int`: Provides a status code to use for successful OPTIONS requests. Default value is `http.StatusNoContent` (`204`).
* **Debug** `bool`: Debugging flag adds additional output to debug server side CORS issues.

See [API documentation](http://godoc.org/github.com/stfsy/go-cors) for more info.

## Benchmarks

```
goos: darwin
goarch: arm64
pkg: github.com/stfsy/go-cors
BenchmarkWithout-10            	135325480	         8.124 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefault-10            	24082140	        51.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkAllowedOrigin-10      	16424518	        88.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkPreflight-10          	 8010259	       147.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkPreflightHeader-10    	 6850962	       175.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkWildcard/match-10     	253275342	         4.714 ns/op	       0 B/op	       0 allocs/op
BenchmarkWildcard/too_short-10 	1000000000	         0.6235 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/stfsy/go-cors	99.131s
```

## Licenses

All source code is licensed under the [MIT License](https://raw.github.com/stfsy/go-cors/master/LICENSE).
