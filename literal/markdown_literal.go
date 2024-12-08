package literal

/*
## Block
- [x] UnorderedList (-, *, +)
- [ ] Nested UnorderedList (-, *, +)
- [ ] Ordered Lists (1., 2., 3.)
- [ ] Task Lists (- [ ] Task item)
- [ ] Nested Task Lists (- [ ] Task item)
- [ ] Definition Lists (Term, followed by a new line : Definition)
- [ ] Blockquotes (>)
- [ ] Code Blocks (triple-backspace, four-space indentation)
- [ ] Tables (using | for columns in a table row)

## Line
- [x] Heading (#, ##, ###, ####, #####, ######)
- [x] Horizontal Rule (---, ___, ***)
- [ ] Footnotes ([^1]: Definition text)

## Inline
- [ ] Bold (**bold**, __bold__)
- [ ] Italic (*italic*, _italic_)
- [ ] Bold and Italic (***bold italic***)
- [ ] Inline Code (`code`)
- [ ] Links ([link text](URL "optional title"))
- [ ] Images (![alt text](URL "optional title"))
- [ ] Strikethrough (~~strikethrough~~)
- [ ] Footnote References ([^1] Inline Text)
- [ ] HTML Elements (<div>content</div>)
- [ ] Comments (<!-- comment -->)

## Separators
- [ ] Line Breaks (two or more spaces at the end of a line)
- [ ] Paragraph (line-separated text)
*/

type MarkdownLiteral interface {
	isLiteral()
}

type HeadingOne struct {
	Text string
}

type HeadingTwo struct {
	Text string
}

type HeadingThree struct {
	Text string
}

type HeadingFour struct {
	Text string
}

type HeadingFive struct {
	Text string
}

type HeadingSix struct {
	Text string
}

type HorizontalRule struct {
	Text string
}

type ListItem struct {
	Text string
}

type TaskList struct {
	Text string
}

type Paragraph struct {
	Text string
}

type BlankLine struct{}

func (h HeadingOne) isLiteral()     {}
func (h HeadingTwo) isLiteral()     {}
func (h HeadingThree) isLiteral()   {}
func (h HeadingFour) isLiteral()    {}
func (h HeadingFive) isLiteral()    {}
func (h HeadingSix) isLiteral()     {}
func (h HorizontalRule) isLiteral() {}
func (l ListItem) isLiteral()       {}
func (t TaskList) isLiteral()       {}
func (p Paragraph) isLiteral()      {}
func (b BlankLine) isLiteral()      {}
