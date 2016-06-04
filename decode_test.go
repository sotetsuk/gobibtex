package gobibtex

import (
	"reflect"
	"testing"
)

const bibstring = "@inproceedings{martens2010deep, title={Deep learning via Hessian-free optimization}, author={Martens, James}, booktitle={Proceedings of the 27th International Conference on Machine Learning (ICML-10)}, pages={735--742}, year={2010}}"

func TestParseEntry(t *testing.T) {
	expected := "inproceedings"

	if actual, _ := parseEntry(bibstring); expected != actual {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestParseFields(t *testing.T) {
	expected := map[string]string{
		"title":     "Deep learning via Hessian-free optimization",
		"author":    "Martens, James",
		"booktitle": "Proceedings of the 27th International Conference on Machine Learning (ICML-10)",
		"pages":     "735--742",
		"year":      "2010",
	}

	if actual := parseFields(bibstring); !reflect.DeepEqual(expected, actual) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestParseName(t *testing.T) {
	expected := "martens2010deep"

	if actual := parseName(bibstring); expected != actual {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestParseAuthors(t *testing.T) {
	authorString := "LeCun, Yann and Bengio, Yoshua and Hinton, Geoffrey"
	expected := []string{"LeCun, Yann", "Bengio, Yoshua", "Hinton, Geoffrey"}

	if actual := parseAuthors(authorString); !reflect.DeepEqual(expected, actual) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}

func TestDecode(t *testing.T) {
	expected := map[string]interface{}{
		"entry":     "inproceedings",
		"name":      "martens2010deep",
		"title":     "Deep learning via Hessian-free optimization",
		"author":    []string{"Martens, James"},
		"booktitle": "Proceedings of the 27th International Conference on Machine Learning (ICML-10)",
		"pages":     "735--742",
		"year":      "2010",
	}

	if actual, _ := Decode(bibstring); !reflect.DeepEqual(expected, actual) {
		t.Errorf("\ngot  %v\nwant %v", actual, expected)
	}
}
