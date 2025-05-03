package game

import "time"

const (
	nsPerSecond       = 1000000000
	maxNsPerUpdate    = 1000000000
	maxTicksPerUpdate = 100
)

type Timer struct {
	ticksPerSecond float32
	lastTime       int64

	TimeScale    float32
	FPS          float32
	PassedTime   float32
	Ticks        int
	PartialTicks float32
}

func NewTimer(ticksPerSecond float32) *Timer {
	timer := new(Timer)

	timer.ticksPerSecond = ticksPerSecond
	timer.lastTime = time.Now().UnixNano()

	timer.TimeScale = 1.0
	timer.FPS = 0.0
	timer.PassedTime = 0.0

	return timer
}

func (timer *Timer) AdvanceTime() {
	now := time.Now().UnixNano()
	passedNs := now - timer.lastTime

	timer.lastTime = now

	passedNs = max(0, passedNs)
	passedNs = min(maxNsPerUpdate, passedNs)

	timer.FPS = nsPerSecond / float32(passedNs)

	timer.PassedTime += float32(passedNs) * timer.TimeScale * timer.ticksPerSecond / nsPerSecond
	timer.Ticks = int(timer.PassedTime)

	timer.Ticks = min(maxTicksPerUpdate, timer.Ticks)

	timer.PassedTime -= float32(timer.Ticks)
	timer.PartialTicks = timer.PassedTime
}
