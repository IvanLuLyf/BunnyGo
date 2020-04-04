# BunnyGo

a lightweight Go web framework.

![BunnyGo](BunnyGo.png)

[![Sourcegraph](https://sourcegraph.com/github.com/ivanlulyf/bunnygo/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/ivanlulyf/bunnygo?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ivanlulyf/bunnygo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivanlulyf/bunnygo?style=flat-square)](https://goreportcard.com/report/github.com/ivanlulyf/bunnygo)
[![GitHub](https://img.shields.io/github/license/ivanlulyf/bunnygo?color=blue&style=flat-square)](LICENSE)

## Install 

```shell script
go get github.com/ivanlulyf/bunnygo
```

## Simple Usage

a Hello World Application

```go
package main

import (
	"github.com/ivanlulyf/bunnygo"
)

type IndexController struct {
}

func (c IndexController) AcIndex() string {
	return "Hello World"
}

func main() {
	app := bunnygo.Bunny{}
	app.Init()
	app.Controller(IndexController{})
	app.Run()
}
```