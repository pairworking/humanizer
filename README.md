# humanizer
Humanizing in human ways

[![Build Status](https://travis-ci.org/pairworking/humanizer.svg?branch=master)](https://travis-ci.org/pairworking/humanizer)
[![Codecov](https://codecov.io/gh/pairworking/humanizer/branch/master/graph/badge.svg)](https://codecov.io/gh/pairworking/humanizer)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/pairworking/humanizer/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/pairworking/humanizer?status.svg)](https://godoc.org/github.com/pairworking/humanizer)

## How to use

```go

import "github.com/pairworking/humanizer"

func ExampleNumberToWords() {
		res := NumberToWords(24)
		fmt.Println(res)
		// Output: twenty-four
}
```