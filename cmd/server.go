/*
Copyright © 2025 creativie <iam@creat.if.ua>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func StartHttpServer(port int) {
	handler := func(ctx *fasthttp.RequestCtx) {
		request_uid := uuid.New().String()
		ctx.Response.Header.SetServer("k8s-controller")
		if string(ctx.RequestURI()) == "/api/liveness" || string(ctx.RequestURI()) == "/api/readiness" {
			ctx.SetContentType("application/json")

			fmt.Fprintf(ctx, "{\"health\":\"ok\"}")
		} else {
			fmt.Fprintf(ctx, "Hello from FastHTTP!\nYour IP is %q\n", ctx.RemoteIP())
		}
		log.Info().
			Str("SERVER_ADDR", ctx.LocalAddr().String()).
			Str("REMOTE_ADDR", ctx.RemoteIP().String()).
			Str("TIME_ISO8601", ctx.Time().UTC().Format("2006-01-02T15:04:05Z07:00")).
			Str("REQUEST", fmt.Sprintf("%s %s %s", ctx.Method(), ctx.RequestURI(), ctx.Request.Header.Protocol())).
			Int("STATUS", ctx.Response.StatusCode()).
			Int("BODY_BYTES_SEND", ctx.Response.Header.ContentLength()).
			Str("REFERER", string(ctx.Request.Header.Referer())).
			Str("USER_AGENT", string(ctx.Request.Header.UserAgent())).
			Str("COOKIE", string(ctx.Request.Header.Cookie("cookie-name"))).
			Str("HTTP_HOST", string(ctx.Host())).
			Str("REQUEST_UID", request_uid).
			Str("XFF", string(ctx.Request.Header.Peek("X-Forwarded-For"))).
			Msg("HTTP log > ")
	}
	addr := fmt.Sprintf(":%d", port)
	log.Info().Msgf("Starting FastHTTP server on %s", addr)
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		log.Error().Err(err).Msg("Error starting FastHTTP server")
		os.Exit(1)
	}
}
