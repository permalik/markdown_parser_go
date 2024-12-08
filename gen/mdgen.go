package gen

import (
	"fmt"
	"io"

	"github.com/permalik/markdown_parser_go/parse"
)

type MDGen struct {
	writer io.Writer
}

func NewMDGen(writer io.Writer) *MDGen {
	return &MDGen{writer: writer}
}

func (g *MDGen) VisitTree(n *parse.TreeNode) {
	for _, child := range n.Children {
		child.Accept(g)
	}
}

func (g *MDGen) VisitHeadingOne(n *parse.HeadingOneNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitHeadingTwo(n *parse.HeadingTwoNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitHeadingThree(n *parse.HeadingThreeNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitHeadingFour(n *parse.HeadingFourNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitHeadingFive(n *parse.HeadingFiveNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitHeadingSix(n *parse.HeadingSixNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

// TODO: gen max line length
func (g *MDGen) VisitHorizontalRule(n *parse.HorizontalRuleNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitList(n *parse.ListNode) {
	for _, item := range n.Items {
		fmt.Fprintf(g.writer, "- %s\n", item)
	}
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitTaskList(n *parse.TaskListNode) {
	for _, item := range n.Items {
		fmt.Fprintf(g.writer, "- [ ] %s\n", item)
	}
}

func (g *MDGen) VisitParagraph(n *parse.ParagraphNode) {
	fmt.Fprintf(g.writer, "%s\n\n", n.Text)
}
