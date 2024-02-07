package Handlers

import (
	"encoding/json"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

type LimitErrorSample struct {
	Err  string
	Body string
	Code int
}

// var (
// 	ipRateLimiter *limiter.Limiter
// 	store         limiter.Store
// )

func GetLimiterMiddleWare() *limiter.Limiter {
	limiter := tollbooth.NewLimiter(0.4, nil) // 1 request in 3 seconds
	// limiter.SetBurst(2) // e.g burst requests in max seconds
	limiter.SetMessageContentType("application/json")
	message := &LimitErrorSample{
		Err:  "FIALD",
		Body: "you reached the 1 sec request",
		Code: 400,
	}
	jsonMessage, _ := json.Marshal(message)
	limiter.SetMessage(string(jsonMessage))
	limiter.SetStatusCode(message.Code)
	return limiter
}

// func IPRateLimit() echo.MiddlewareFunc {
// 	// 1. Configure
// 	rate := limiter.Rate{
// 		Period: 5 * time.Second,
// 		Limit:  3,
// 	}
// 	store = memory.NewStore()
// 	ipRateLimiter = limiter.New(store, rate)

// 	// 2. Return middleware handler
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) (err error) {
// 			ip := c.RealIP()
// 			limiterCtx, err := ipRateLimiter.Get(c.Request().Context(), ip)
// 			if err != nil {
// 				log.Printf("IPRateLimit - ipRateLimiter.Get - err: %v, %s on %s", err, ip, c.Request().URL)
// 				return c.JSON(http.StatusInternalServerError, echo.Map{
// 					"success": false,
// 					"message": err,
// 				})
// 			}

// 			h := c.Response().Header()
// 			h.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
// 			h.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
// 			h.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

// 			if limiterCtx.Reached {
// 				log.Printf("Too Many Requests from %s on %s", ip, c.Request().URL)
// 				return c.JSON(http.StatusTooManyRequests, echo.Map{
// 					"success": false,
// 					"message": "Too Many Requests on " + c.Request().URL.String(),
// 				})
// 			}

// 			// log.Printf("%s request continue", c.RealIP())
// 			return next(c)
// 		}
// 	}
// }
