package game

import "time"

const (
	nsPerSecond       = 1000000000
	maxNsPerUpdate    = 1000000000
	maxTicksPerUpdate = 100
)

type Timer struct {
	Ticks        int
	PartialTicks float32

	tps        float32
	lastTime   int64
	timeScale  float32
	fps        float32
	passedTime float32
}

func NewTimer(tps float32) *Timer {
	timer := new(Timer)

	timer.tps = tps
	timer.lastTime = time.Now().UnixNano()
	timer.timeScale = 1.0
	timer.fps = 0.0
	timer.passedTime = 0.0

	return timer
}

func (timer *Timer) AdvanceTime() {
	now := time.Now().UnixNano()
	passedNs := now - timer.lastTime

	timer.lastTime = now

	passedNs = max(0, passedNs)
	passedNs = min(maxNsPerUpdate, passedNs)

	timer.fps = float32(nsPerSecond) / float32(passedNs)

	timer.passedTime += float32(passedNs) * timer.timeScale * timer.tps / nsPerSecond
	timer.Ticks = int(timer.passedTime)

	timer.Ticks = min(maxTicksPerUpdate, timer.Ticks)

	timer.passedTime -= float32(timer.Ticks)
	timer.PartialTicks = timer.passedTime
}
