![Coverage](https://img.shields.io/badge/Coverage-66.7%25-yellow)

# randomselector

Randomly select objects in golang
Current version: v0.0.2

## Install

```console
go get  github.com/zeroboo/randomselector
```

## Usage 
```golang
package main

import (
	"fmt"

	"github.com/zeroboo/randomselector"
)

func main() {
	
	//Select values randomly with equally rate for each value, 1/3 chance for each value
	value, errSelect := randomselector.SelectValues("1", "2", "3")
	fmt.Println("Select values: ", value, errSelect)

	//Select one of 2 string: "hello", "world" with equal rate (50% chance for each)
	weightValue, errSelect := randomselector.SelectWithWeight(
		randomselector.WeightValue{Value: "1", Weight: 1},
		randomselector.WeightValue{Value: "2", Weight: 1},
	)
	fmt.Println("Select weight value: ", weightValue, errSelect)

}


```
## Test

```console
go test -timeout 60s github.com/zeroboo/randomselector/test -v
```

## Publish

```console
.\publish.ps1
```
