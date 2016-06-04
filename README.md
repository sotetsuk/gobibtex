[![GoDoc](https://godoc.org/github.com/sotetsuk/gobibtex?status.svg)](https://godoc.org/github.com/sotetsuk/gobibtex)
[![Build Status](https://travis-ci.org/sotetsuk/gobibtex.svg?branch=master)](https://travis-ci.org/sotetsuk/gobibtex)
[![Coverage Status](https://coveralls.io/repos/github/sotetsuk/gobibtex/badge.svg?branch=master)](https://coveralls.io/github/sotetsuk/gobibtex?branch=master)
[![GitHub version](https://badge.fury.io/gh/sotetsuk%2Fgobibtex.svg)](https://badge.fury.io/gh/sotetsuk%2Fgobibtex)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()

# gobibtex
BibTeX parser written in Go

# install

```
$ go get github.com/sotetsuk/gobibtex
```

# Usage

```go
import "gobibtex"
import "github.com/k0kubun/pp"

bibstring := `@article{lecun2015deep,
                       title={Deep learning},
                       author={LeCun, Yann and Bengio, Yoshua and Hinton, Geoffrey},
                       journal={Nature},
                       volume={521},
                       number={7553},
                       pages={436--444},
                       year={2015},
                       publisher={Nature Publishing Group}
                      }`

bibmap, err := gobibtex.Decode(bibstring)
pp.Println(bibmap)
```

Output: 
```go
{
  "journal": "Nature",
  "number":  "7553",
  "pages":   "436--444",
  "title":   "Deep learning",
  "name":    "lecun2015deep",
  "author":  []string{
    "LeCun, Yann",
    "Bengio, Yoshua",
    "Hinton, Geoffrey",
  },
  "year":   "2015",
  "entry":  "article",
  "volume": "521",
}
```

# License
MIT License
