package Handlers

import (
	"encoding/json"
	"sunsend/internals/Data"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

// var (
// 	ipRateLimiter *limiter.Limiter
// 	store         limiter.Store
// )

func GetLimiterMiddleWare() *limiter.Limiter {
	limiter := tollbooth.NewLimiter(0.4, nil) // 1 request in 3 seconds
	// limiter.SetBurst(2) // e.g burst requests in max seconds
	limiter.SetMessageContentType("application/json")
	response, _ := Data.NewResponse(nil, 18, "", nil)
	jsonMessage, _ := json.Marshal(response)
	// limiter.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
	// 	limiter.ExecOnLimitReached(w, r)
	// 	fmt.Println("user reached the limit of the server")
	// })
	limiter.SetMessage(string(jsonMessage))
	error_obj := Data.GetErrorByResult(18)
	limiter.SetStatusCode(error_obj.StatusCode)
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
