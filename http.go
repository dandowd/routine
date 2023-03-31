package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting gin server")
			go func() {
				err := r.Run()
				if err != nil {
					panic(err)
				}
			}()

			return nil
		},
	})

	return r
}
