# Go Snowflake

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/gopkg.in/shaxbee/go-snowflake.v1) [![license](http://img.shields.io/badge/license-apache_2.0-red.svg?style=flat)](https://raw.githubusercontent.com/shaxbee/go-snowflake/master/LICENSE) [![build](https://travis-ci.org/shaxbee/go-snowflake.svg?branch=master)](https://travis-ci.org/shaxbee/go-snowflake) [![coverage](https://coveralls.io/repos/github/shaxbee/go-snowflake/badge.svg?branch=master)](https://coveralls.io/r/shaxbee/go-snowflake)

Go Snowflake is Twitter Snowflake inspired monotonic ID generator written in Golang.
When using generator in distributed environment library user has to ensure that worker ID remains unique.

Example:
```go
package main

import "gopkg.in/shaxbee/go-snowflake.v1"

func main() {
  sf, err := snowflake.New(42)
  if err != nil {
    panic(err)
  }
  
  for i := 0; i < 10; i++ {
    fmt.Println(<- sf)
  }
}
```
