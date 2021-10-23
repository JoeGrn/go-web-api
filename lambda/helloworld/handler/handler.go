package handler

import (
	"context"
	"fmt"
)

type Event struct {
}

func Handle(ctx context.Context, event Event) (string, error) {
	return fmt.Sprintf("Hello World!"), nil
}
