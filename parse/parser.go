package parse

import (
	"fmt"
	"io"

	"github.com/permalik/markdown_parser_go/lex"
	"github.com/permalik/markdown_parser_go/literal"
)

type Parser struct {
	lexer *lex.Lexer
}

type Node interface {
	Accept(v Visitor)
}

type TreeNode struct {
	Children []Node
}

type HeadingOneNode struct {
	Text string
}

type HeadingTwoNode struct {
	Text string
}

type HeadingThreeNode struct {
	Text string
}

type HeadingFourNode struct {
	Text string
}

type HeadingFiveNode struct {
	Text string
}

type HeadingSixNode struct {
	Text string
}

type HorizontalRuleNode struct {
	Text string
}

type ListNode struct {
	Items []string
}

type TaskListNode struct {
	Items []string
}

type ParagraphNode struct {
	Text string
}

type Visitor interface {
	VisitTree(n *TreeNode)
	VisitHeadingOne(n *HeadingOneNode)
	VisitHeadingTwo(n *HeadingTwoNode)
	VisitHeadingThree(n *HeadingThreeNode)
	VisitHeadingFour(n *HeadingFourNode)
	VisitHeadingFive(n *HeadingFiveNode)
	VisitHeadingSix(n *HeadingSixNode)
	VisitHorizontalRule(n *HorizontalRuleNode)
	VisitList(n *ListNode)
	VisitTaskList(n *TaskListNode)
	VisitParagraph(n *ParagraphNode)
}

func (n *TreeNode) Accept(v Visitor)           { v.VisitTree(n) }
func (n *HeadingOneNode) Accept(v Visitor)     { v.VisitHeadingOne(n) }
func (n *HeadingTwoNode) Accept(v Visitor)     { v.VisitHeadingTwo(n) }
func (n *HeadingThreeNode) Accept(v Visitor)   { v.VisitHeadingThree(n) }
func (n *HeadingFourNode) Accept(v Visitor)    { v.VisitHeadingFour(n) }
func (n *HeadingFiveNode) Accept(v Visitor)    { v.VisitHeadingFive(n) }
func (n *HeadingSixNode) Accept(v Visitor)     { v.VisitHeadingSix(n) }
func (n *HorizontalRuleNode) Accept(v Visitor) { v.VisitHorizontalRule(n) }
func (n *ListNode) Accept(v Visitor)           { v.VisitList(n) }
func (n *TaskListNode) Accept(v Visitor)       { v.VisitTaskList(n) }
func (n *ParagraphNode) Accept(v Visitor)      { v.VisitParagraph(n) }

func NewParser(lexer *lex.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() (Node, error) {
	tree := &TreeNode{}
	var currList []string
	var currTaskList []string

	for {
		token, err := p.lexer.NextToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error parsing: %w", err)
		}

		switch tok := token.Literal.(type) {
		case literal.HeadingOne:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingOneNode{
				Text: tok.Text,
			})
		case literal.HeadingTwo:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingTwoNode{
				Text: tok.Text,
			})
		case literal.HeadingThree:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingThreeNode{
				Text: tok.Text,
			})
		case literal.HeadingFour:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingFourNode{
				Text: tok.Text,
			})
		case literal.HeadingFive:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingFiveNode{
				Text: tok.Text,
			})
		case literal.HeadingSix:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HeadingSixNode{
				Text: tok.Text,
			})
		case literal.HorizontalRule:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			if len(currTaskList) > 0 {
				tree.Children = append(tree.Children, &TaskListNode{Items: currTaskList})
				currTaskList = nil
			}
			tree.Children = append(tree.Children, &HorizontalRuleNode{
				Text: tok.Text,
			})
		case literal.Paragraph:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			if len(currTaskList) > 0 {
				tree.Children = append(tree.Children, &TaskListNode{Items: currTaskList})
				currTaskList = nil
			}
			tree.Children = append(tree.Children, &ParagraphNode{
				Text: tok.Text,
			})
		case literal.ListItem:
			currList = append(currList, tok.Text)
		case literal.TaskList:
			currTaskList = append(currTaskList, tok.Text)
		case literal.BlankLine:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
		}
	}

	if len(currList) > 0 {
		tree.Children = append(tree.Children, &ListNode{Items: currList})
	}

	return tree, nil
}
