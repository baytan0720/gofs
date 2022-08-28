package leasemanager

import (
	"errors"
	"time"
)

type Lease struct {
	status int //0 释放/1 占有
	timer  *time.Timer
}

func MakeLease() *Lease {
	timer := time.NewTimer(10 * time.Minute)
	timer.Stop()
	l := &Lease{
		status: 0,
		timer:  timer,
	}
	return l
}

func (l *Lease) Get() error {
	if l.status == 1 {
		return errors.New("Lease possession")
	}
	l.status = 1
	l.timer.Reset(10 * time.Minute)
	go func() {
		<-l.timer.C
		l.status = 0
	}()
	return nil
}

func (l *Lease) Release() {
	l.timer.Reset(0)
}

func (l *Lease) Renew() {
	l.timer.Reset(10 * time.Minute)
}
