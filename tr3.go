package main

import "strings"

func tr2(doc string) string {

	tr := map[string]string{
		"##":  "Concensus Link",
		"# Q": "Concensus Link",
	}

	for k, v := range tr {
		doc = strings.ReplaceAll(doc, k, v)
	}

	return doc
}
