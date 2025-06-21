/*
Copyright © 2025 creativie <iam@creat.if.ua>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func StartHttpServer(port int) {
	handler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Hello from FastHTTP!")
	}
	addr := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Starting FastHTTP server on %s", addr)
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		log.Error().Err(err).Msg("Error starting FastHTTP server")
		os.Exit(1)
	}
}
