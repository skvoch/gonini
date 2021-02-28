package main

import (
	"context"
	"errors"
	"github.com/skvoch/gonini/internal/ui"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	u := ui.New()

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return u.Run(ctx)
	})

	if err := g.Wait(); err != nil {
		if errors.Is(err, ui.ErrExit) {
			return
		}
		log.Fatal(err)
	}
}
