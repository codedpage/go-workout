package main

import (
    "fmt"
    "time"
)

const INTERVAL_PERIOD time.Duration = 24 * time.Hour

const HOUR_TO_TICK int = 07
const MINUTE_TO_TICK int = 36
const SECOND_TO_TICK int = 10

type jobTicker struct {
    t *time.Timer
}

func getNextTickDuration() time.Duration {
    now := time.Now()
    nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
    if nextTick.Before(now) {
        nextTick = nextTick.Add(INTERVAL_PERIOD)
    }
    return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
    fmt.Println("new tick here")
    return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
    fmt.Println("next tick here")
    jt.t.Reset(getNextTickDuration())
}

func main() {
    jt := NewJobTicker()
    for {
        <-jt.t.C
        fmt.Println(time.Now(), "- just ticked")
        jt.updateJobTicker()
    }
}