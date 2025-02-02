package ruler

import (
	"context"
	"gox/internal/server/socks/request"
	"log"

	"github.com/things-go/go-socks5"
)

type Ruler struct{}

func (r *Ruler) Allow(
	ctx context.Context,
	req *socks5.Request,
) (context.Context, bool) {
	_req := request.New(
		req,
	)
	log.Println(_req)
	log.Println(req)
	// if req.DestAddr.String() == "8.8.8.8:53" {
	// 	return ctx, true
	// }
	// log.Printf("Blocked request to: %s", req.DestAddr.String())
	return ctx, true
}

func New() *Ruler {
	return &Ruler{}
}
