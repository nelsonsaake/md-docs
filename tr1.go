package main

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func tr1(v string) string {
	var (
		md        = goldmark.New()
		src       = []byte(v)
		reader    = text.NewReader(src)
		docParser = md.Parser().Parse(reader)
	)

	// Convenient way of exiting function
	next := func() (ast.WalkStatus, error) {
		return ast.WalkContinue, nil
	}

	ast.Walk(docParser, func(n ast.Node, entering bool) (ast.WalkStatus, error) {

		if heading, ok := n.(*ast.Heading); ok && entering {

			// If not H1; move to next node
			if heading.Level != 1 {
				return next()
			}

			// Extract the heading text
			headingText := string(heading.Text(src))

			// If not # l; move to next node
			if headingText != "l" {
				return next()
			}

			// we pull the actual link from the firstChild, which is a  node
			href := extractLink(heading, src)

			// we create a new link
			link := newLink("link", href)

			// Create a new paragraph node containing the link
			newParagraph := ast.NewParagraph()

			// Add the link to the new paragraph
			newParagraph.AppendChild(newParagraph, link)

			// get heading's parent
			parent := heading.Parent()

			// insert the new link into doc
			parent.InsertBefore(parent, heading, newParagraph)

			parent.RemoveChild(parent, heading)

			return next()
		}

		return next()
	})

	// Render the modified AST back to a string
	var buf bytes.Buffer
	if err := md.Renderer().Render(&buf, src, docParser); err != nil {
		panic(err) // Handle the error appropriately in your code
	}

	return buf.String()
}

func newLink(txt, href string) *ast.Link {

	// Create the link node
	newLink := ast.NewLink()
	newLink.Destination = []byte(href)
	newLink.Title = []byte(txt)

	// For the display text
	textSegment := text.NewSegment(0, 3)

	// Create a text node for the visible part of the link
	linkText := ast.NewText()
	linkText.Segment = textSegment

	// Append the text node to the link
	newLink.AppendChild(newLink, linkText)

	return newLink
}

// extractLink retrieves the content following a heading node
func extractLink(heading *ast.Heading, source []byte) string {
	var content bytes.Buffer

	ls := []ast.Node{}

	for n := heading.NextSibling(); n != nil; n = n.NextSibling() {

		// Stop if we envouter a orizontal like
		if n.Kind() == ast.KindThematicBreak {
			break
		}

		// Stop if we encounter another heading
		if _, ok := n.(*ast.Heading); ok {
			break
		}

		// Check if the node is a paragraph
		if n.Kind() == ast.KindParagraph {
			for c := n.FirstChild(); c != nil; c = c.NextSibling() {
				// Check if the child is a text node or has text-like content
				if textNode, ok := c.(*ast.Text); ok {
					segment := textNode.Segment
					content.Write(segment.Value(source))
				} else {
					// Attempt to render the node to capture its textual representation
					var buf bytes.Buffer
					if err := goldmark.DefaultRenderer().Render(&buf, source, c); err == nil {
						content.WriteString(buf.String())
					}
				}
			}
		}

		ls = append(ls, n)
	}

	parent := heading.Parent()
	for _, v := range ls {
		parent.RemoveChild(parent, v)
	}

	return content.String()
}
