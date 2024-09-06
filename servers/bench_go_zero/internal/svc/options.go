package svc

import "time"

type ServerOptions struct {
	Target    int
	CpuBound  bool
	SleepTime time.Duration
	Address   string
}
