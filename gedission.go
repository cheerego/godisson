package godisson

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"time"
)

type Godisson struct {
	c               *redis.Client
	watchDogTimeout time.Duration
	uuid            string
}

var DefaultWatchDogTimeout = 30 * time.Second

func NewGodisson(redisClient *redis.Client, opts ...OptionFunc) *Godisson {
	g := &Godisson{
		c:               redisClient,
		uuid:            uuid.New().String(),
		watchDogTimeout: DefaultWatchDogTimeout,
	}
	return g
}

type OptionFunc func(g *Godisson)

func WithWatchDogTimeout(t time.Duration) OptionFunc {
	return func(g *Godisson) {
		if t.Seconds() < 30 {
			t = DefaultWatchDogTimeout
			log.Println("watchDogTimeout is too small, so config default ")
		}
		g.watchDogTimeout = t
	}
}

func (g *Godisson) NewRLock(key string) *RLock {
	return newRLock(key, g)
}
