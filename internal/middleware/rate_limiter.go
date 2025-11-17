package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(rdb redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rate" + c.ClientIP()
		now := time.Now().Unix()
		windowStart := now - int64(window.Seconds())

		rdb.ZRemRangeByScore(c.Request.Context(), key, "0", fmt.Sprintf("%d", windowStart))

		count, err := rdb.ZCard(c.Request.Context(), key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "rate limiter error"})
			return
		}

		if count >= int64(limit) {
			c.Header("Retry-After", fmt.Sprintf("%d", int(window.Seconds())))
			c.AbortWithStatusJSON(429, gin.H{
				"error":       "rate limit exceeded",
				"retry_after": int(window.Seconds()),
			})
			return
		}

		rdb.ZAdd(c.Request.Context(), key, redis.Z{
			Score:  float64(now),
			Member: fmt.Sprintf("%d", now),
		})
		rdb.Expire(c.Request.Context(), key, window)

		c.Next()
	}
}
