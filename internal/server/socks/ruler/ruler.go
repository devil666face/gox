package ruler

import (
	"context"

	"github.com/things-go/go-socks5"
)

type Ruler struct{}

func (r *Ruler) Allow(
	ctx context.Context,
	req *socks5.Request,
) (context.Context, bool) {
	return ctx, true
}

func New() *Ruler {
	return &Ruler{}
}
