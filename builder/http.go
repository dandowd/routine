package builder

import (
	"context"
	"fmt"
	"routine/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, logger common.Logger) *gin.Engine {
	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting gin server")
			go func() {
				err := r.Run()
				if err != nil {
					logger.Error("Error starting gin server")
				}
			}()

			return nil
		},
	})

	return r
}
