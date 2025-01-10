package main

import "strings"

func tr0(doc string) string {

	tr := map[string]string{
		"# q": "# Question",
		"# t": "## Title",
		"# s": "### Summary",
		"# r": "### Reference: APA reference to the actual paper",
		// "# l": "### Link",
		// "# l": "### Link: consensus.app link",
		"# m": "### Method",
		"# o": "### Outcome",
		"# a": "### Results",
		"# b": "### Related Searches",
		"# c": "Highly Cited***",
		"# p": "### Pointer: Points to Local Extracted Version",
	}

	for k, v := range tr {
		doc = strings.ReplaceAll(doc, k, v)
	}

	return doc
}
