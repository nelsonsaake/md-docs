package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
	"github.com/nelsonsaake/md-docs/bootstrap"
	"github.com/nelsonsaake/md-docs/src/env"
)

var (
	ls         = []string{}
	buf        = bytes.Buffer{}
	ignorelist = []string{
		"key.md",
	}
)

func init() {
	err := bootstrap.Boot()
	do.Die(err)
}

func walkDoc(path string, d fs.DirEntry, err error) error {
	ls = append(ls, filepath.ToSlash(path))
	return nil
}

func writeln(v ...any) {
	buf.WriteString(fmt.Sprintln(v...))
}

func shouldIgnore(path string) bool {
	for _, ignorePattern := range ignorelist {
		if strings.HasSuffix(path, ignorePattern) {
			return true
		}
	}
	return false
}

func main() {

	writeln("---")
	writeln("title: msc thesis")

	err := filepath.WalkDir(env.DocPath(), walkDoc)
	do.Die(err)

	for _, path := range ls {

		if !strings.HasSuffix(path, ".md") {
			continue
		}

		if shouldIgnore(path) {
			continue
		}

		content, err := ufs.ReadFile(path)
		do.Die(err)

		content = tr0(content)
		content = tr1(content)
		content = tr2(content)

		writeln("---")
		writeln()

		writeln(path)
		writeln()

		writeln(content)
		writeln()
	}

	doc := buf.String()

	err = ufs.WriteFile("output.md", doc)
	do.Die(err)
}
