package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type IPLimiter struct {
	Limiter *Limiter
}

func NewIPLimiter() LimiterInterface {
	l := &IPLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
	return l
}

func (l IPLimiter) Key(c *gin.Context) string {
	ip := c.ClientIP()
	return ip
}

func (l IPLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.Limiter.limiterBuckets[key]
	return bucket, ok
}

func (l IPLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		l.Limiter.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
	}
	return l
}
