package ui

import (
	"context"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func commandBox(keys <-chan string) *commandBoxImpl {
	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	return &commandBoxImpl{
		p:    p,
		keys: keys,
	}
}

type commandBoxImpl struct {
	p    *widgets.Paragraph
	keys <-chan string
}

func (c *commandBoxImpl) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case key := <-c.keys:
			c.p.Text = key
		}
	}
}

func (c *commandBoxImpl) Draw() termui.Drawable {
	return c.p
}
