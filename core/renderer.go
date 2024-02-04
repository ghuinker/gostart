package core

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type contextKey string

var echoContextKey contextKey = "echo"

func GetEcho(ctx context.Context) echo.Context {
	if echoContext, ok := ctx.Value(echoContextKey).(echo.Context); ok {
		return echoContext
	}
	newCtx := echo.New().AcquireContext()
	defer echo.New().ReleaseContext(newCtx)
	return newCtx
}

func Render(ctx echo.Context, t templ.Component, statusCode ...int) error {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	ctx.Response().Writer.WriteHeader(code)

	newCtx := context.WithValue(context.Background(), echoContextKey, ctx)

	err := t.Render(newCtx, ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
