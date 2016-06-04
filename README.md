[![Build Status](https://travis-ci.org/sotetsuk/gobibtex.svg?branch=master)](https://travis-ci.org/sotetsuk/gobibtex)

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
