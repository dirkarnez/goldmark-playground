package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

func main() {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	source := []byte(`# TITLE
### Topic 1
- 123
- 456
`)
	doc := markdown.Parser().Parse(text.NewReader(source))
	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if n.Kind() == ast.KindHeading {
				if string(n.Text(source)) == "Topic 1" {
					//.AppendChild(ast.NewListItem())
					targetList := n.NextSibling()

					fmt.Println(targetList.ChildCount())

					listItem := ast.NewListItem(targetList.ChildCount())
					targetList.AppendChild(targetList, listItem)

					textBlock := ast.NewTextBlock()
					listItem.AppendChild(listItem, textBlock)

					text := ast.NewText()
					textBlock.AppendChild(textBlock, text)

					str := ast.NewString([]byte("789"))
					text.AppendChild(text, str)

					fmt.Println(targetList.ChildCount())
					//fmt.Println(targetList.FirstChild().FirstChild().FirstChild().Kind())

				}
			}
		}
		// fmt.Println(entering, n.Type(), string(n.Text(source)))
		return ast.WalkContinue, nil
	})

	//fmt.Println(string(doc.Text(source)))
	var d = []byte{}
	doc.Dump(d, 0)

	var buf bytes.Buffer
	markdown.Renderer().Render(&buf, source, doc)
	fmt.Println(buf.String())
}

// "github.com/rhinoman/go-commonmark"

// // https://github.com/rhinoman/go-commonmark/blob/master/commonmark_test.go
// // heading text
// //
// // CMARK_NODE_LIST
// // CMARK_NODE_ITEM
// // CMARK_NODE_PARAGRAPH
// // CMARK_NODE_TEXT
// type List struct {
// 	topic *commonmark.CMarkNode
// }

// const (
// 	FileContent = `TITLE
// =====
// ### Topic 1
// - list item 1
// - list item 2
// `
// )

// func main() {
// 	cMarkNode := commonmark.ParseDocument(FileContent, 0)

// 	cMarkIter := commonmark.NewCMarkIter(cMarkNode)

// 	var prev *commonmark.CMarkNode = nil
// 	var selectedTopicNode *commonmark.CMarkNode = nil

// 	for cMarkIter.Next() != commonmark.CMARK_EVENT_DONE {
// 		node := cMarkIter.GetNode()

// 		if prev != nil && prev.GetNodeType() == commonmark.CMARK_NODE_HEADING &&
// 			node.GetNodeType() == commonmark.CMARK_NODE_TEXT {
// 			if node.GetLiteral() == "Topic 1" {
// 				selectedTopicNode = node
// 				newNode := commonmark.NewCMarkNode(commonmark.CMARK_NODE_TEXT)
// 				newNode.SetLiteral("hi")

// 				selectedTopicNode.FirstChild().AppendChild(newNode)
// 				//fmt.Println(newNode.GetNodeTypeString(), ":", newNode.GetLiteral())
// 				//fmt.Println(selectedTopicNode.Next().Next().FirstChild().GetNodeTypeString(), ":", selectedTopicNode.Next().GetLiteral())
// 			}
// 		}
// 		fmt.Println(node.GetNodeTypeString(), ":", node.GetLiteral())

// 		prev = node
// 	}

// 	// if selectedTopicNode != nil {
// 	// 	fmt.Println(selectedTopicNode.FirstChild().Next().GetNodeTypeString(), ":", selectedTopicNode.FirstChild().Next().GetLiteral())

// 	// 	//		selectedTopicNode.GetListStart().GetNode().AppendChild(newNode)
// 	// }

// 	fmt.Println("============")

// 	md := cMarkNode.RenderCMark(commonmark.CMARK_OPT_DEFAULT, 0)
// 	fmt.Println(md)

// 	//var event commonmark.CMarkEvent = cMarkIter.Next()

// 	//
// 	// 	event = cMarkIter.Next()
// 	// }

// 	// document2 := commonmark.ParseDocument("Foobar\n------", 0)
// 	// htmlText = document2
// 	// document2.RenderXML(commonmark.CMARK_OPT_DEFAULT)
// 	// if htmlText != "<h2>Foobar</h2>\n" {
// 	// 	t.Error("Html text 2 is not as expected :(")
// 	// }
// 	// t.Logf("Html Text2: %v", htmlText)
// 	cMarkNode.Free()
// 	cMarkNode.Free()
// }
