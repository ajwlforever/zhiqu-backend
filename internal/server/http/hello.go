package http1

import "github.com/go-kratos/kratos/v2/transport/http"

func Helloworld(ctx http.Context) error {
	ctx.JSON(1, "wocaac")
	return nil
}
