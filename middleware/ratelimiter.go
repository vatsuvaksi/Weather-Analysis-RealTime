package middleware

import (
	"errors"
	networkutils "real-time-weather-app/utils/networkUtils"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

// Define a rate limiter for your API
var apiLimiter = rate.NewLimiter(rate.Every(time.Second*10), 10)

// RateLimitMiddleware is a middleware to handle rate limiting
func RateLimitMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if apiLimiter.Allow() {
			next(ctx) // Request allowed, proceed to the handler
		} else {
			err := errors.New("rate limit exceeded")
			networkutils.RenderInternalFailiureResponseForLimitExceed(ctx, err)
			// ctx.Error("Rate limit exceeded ", fasthttp.StatusTooManyRequests)
		}
	}
}
