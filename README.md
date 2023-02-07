traffic
==================

[![Build Status](https://travis-ci.org/thedevsaddam/traffic.svg?branch=master)](https://travis-ci.org/thedevsaddam/traffic)
[![Project status](https://img.shields.io/badge/version-1.0-green.svg)](https://github.com/thedevsaddam/traffic/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/traffic)](https://goreportcard.com/report/github.com/thedevsaddam/traffic)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/traffic/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/traffic?branch=master)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/traffic?status.svg)](https://pkg.go.dev/github.com/thedevsaddam/traffic)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/thedevsaddam/traffic/blob/main/LICENSE.md)


Thread safe load-balancer package for Golang

### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/traffic
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/traffic"
```

### Example

```go

package main

import (
	"fmt"

	"github.com/thedevsaddam/traffic"
)

func main() {
	t := traffic.NewWeightedRoundRobin()
	t.Add("a", 5)
	t.Add("b", 2)
	t.Add("c", 3)

	for i := 0; i < 100; i++ {
		fmt.Println(t.Next())
	}
}

```

### **Contribution**
If you are interested to make the package better please send pull requests or create an issue so that others can fix. Read the [contribution guide here](CONTRIBUTING.md). 

### **License**
The **traffic** is an open-source software licensed under the [MIT License](LICENSE.md).
