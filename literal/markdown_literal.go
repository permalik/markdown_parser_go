package literal

type MarkdownLiteral interface {
	isLiteral()
}

type HorizontalRuleHyphen struct {
	Text string
}

type ListItem struct {
	Text string
}

type Paragraph struct {
	Text string
}

type BlankLine struct{}

func (h HorizontalRuleHyphen) isLiteral() {}
func (l ListItem) isLiteral()             {}
func (p Paragraph) isLiteral()            {}
func (b BlankLine) isLiteral()            {}
