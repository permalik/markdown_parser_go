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

type HorizontalRuleHyphenNode struct {
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
	VisitHorizontalRuleHyphen(n *HorizontalRuleHyphenNode)
	VisitList(n *ListNode)
	VisitTaskList(n *TaskListNode)
	VisitParagraph(n *ParagraphNode)
}

func (n *TreeNode) Accept(v Visitor)                 { v.VisitTree(n) }
func (n *HorizontalRuleHyphenNode) Accept(v Visitor) { v.VisitHorizontalRuleHyphen(n) }
func (n *ListNode) Accept(v Visitor)                 { v.VisitList(n) }
func (n *TaskListNode) Accept(v Visitor)             { v.VisitTaskList(n) }
func (n *ParagraphNode) Accept(v Visitor)            { v.VisitParagraph(n) }

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
		case literal.HorizontalRuleHyphen:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			if len(currTaskList) > 0 {
				tree.Children = append(tree.Children, &TaskListNode{Items: currTaskList})
				currTaskList = nil
			}
			tree.Children = append(tree.Children, &HorizontalRuleHyphenNode{
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
		case literal.ListItemHyphen:
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
