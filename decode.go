package gobibtex

import (
	"errors"
	"regexp"
	"strings"
)

// Parse entry name
func parseEntry(bibstring string) (entry string, err error) {
	listEntry := []string{"article", "book", "booklet", "conference", "inbook", "incollection", "inproceedings", "manual", "mastersthesis", "misc", "phdthesis", "proceedings", "techreport", "unpublished"}

	for _, entry := range listEntry {
		r := regexp.MustCompile("^@" + entry + "*{")
		if r.MatchString(bibstring) {
			return entry, nil
		}
	}
	return "", errors.New("Entry not found")
}

// Parse fields
func parseFields(bibstring string) (fields map[string]string) {
	listField := []string{"address", "annote", "author", "booktitle", "chapter", "crossref", "edition", "editor", "howpublished", "institution", "journal", "key", "month", "note", "number", "organization", "pages", "publisher", "school", "series", "title", "type", "volume", "year"}

	fields = make(map[string]string)
	for _, field := range listField {
		r := regexp.MustCompile(field + `[\s]*=[\s]*[{"][^}]*["}][},]`)

		if !r.MatchString(bibstring) {
			continue
		}

		val := r.FindString(bibstring)

		// trim
		val = val[len(field):]
		val = strings.TrimSpace(val)
		val = strings.Trim(val, "=")
		val = strings.TrimSpace(val)
		val = strings.Trim(val, ",")
		val = strings.Trim(val, "}")
		val = strings.Trim(val, "}")
		val = strings.Trim(val, "{")
		val = strings.Trim(val, "{")
		val = strings.Trim(val, `"`)

		fields[field] = val
	}

	return fields
}

// Parse bibtex name
func parseName(bibstring string) (name string) {
	ix1 := strings.Index(bibstring, "{")
	ix2 := strings.Index(bibstring, ",")

	name = bibstring[ix1+1 : ix2]

	return name
}

// Parse authors
func parseAuthors(authorString string) (authors []string) {
	s := strings.Split(authorString, " and ")
	for i, v := range s {
		s[i] = strings.TrimSpace(v)
	}

	return s
}

// Decode bibtex string to map.
func Decode(bibstring string) (bibmap map[string]interface{}, err error) {
	entry, err := parseEntry(bibstring)
	if err != nil {
		return nil, err
	}

	name := parseName(bibstring)
	fields := parseFields(bibstring)

	bibmap = make(map[string]interface{})
	bibmap["entry"] = entry
	bibmap["name"] = name
	for k, v := range fields {
		bibmap[k] = v
	}

	if _, ok := bibmap["author"]; ok {
		bibmap["author"] = parseAuthors(bibmap["author"].(string))
	}

	return bibmap, err
}