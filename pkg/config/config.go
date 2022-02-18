package config

import "time"

type timeout struct {
	ContextTimeout time.Duration
}

var Timeout timeout
