package ui

import (
	"context"
	"errors"
	"fmt"

	"github.com/gizak/termui/v3"
)

var (
	ErrExit = errors.New("exixt")
)

func New() *UI {
	return &UI{}
}

type UI struct {
	keyBoardChannels []chan string
}

func (u *UI) SubscribeKeyboard() <-chan string {
	out := make(chan string, 1024)
	u.keyBoardChannels = append(u.keyBoardChannels, out)
	return out
}

func (u *UI) Run(ctx context.Context) error {
	if err := termui.Init(); err != nil {
		return fmt.Errorf("failed to init termui: %w", err)
	}
	defer termui.Close()

	commandBox := commandBox(u.SubscribeKeyboard())
	go commandBox.Run(ctx)
	termui.Render(commandBox.Draw())

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			for e := range termui.PollEvents() {
				switch e.ID {
				case "q", "<C-c>":
					return ErrExit
				}

				switch e.Type {
				case termui.KeyboardEvent: // handle all key presses
					for _, c := range u.keyBoardChannels {
						c <- e.ID
					}
				}
			}
		}
	}
}
