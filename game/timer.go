package game

import "time"

const nsPerSecond = 1e9
const maxNsPerUpdate = 1e9
const maxTicksPerUpdate = 100

type Timer struct {
	ticksPerSecond float32
	lastTime       int64
	Ticks          int
	A              float32
	timeScale      float32
	fps            float32
	passedTime     float32
}

func NewTimer(ticksPerSecond float32) (timer Timer) {
	timer.timeScale = 1.0
	timer.fps = 0.0
	timer.passedTime = 0.0
	timer.ticksPerSecond = ticksPerSecond
	timer.lastTime = time.Now().UnixNano()

	return timer
}

func (t Timer) AdvanceTime() Timer {
	now := time.Now().UnixNano()
	passedNs := now - t.lastTime
	t.lastTime = now

	if passedNs < 0 {
		passedNs = 0
	}
	if passedNs > maxNsPerUpdate {
		passedNs = maxNsPerUpdate
	}

	t.fps = float32(maxNsPerUpdate / passedNs)

	t.passedTime += float32(passedNs) * t.timeScale * t.ticksPerSecond / nsPerSecond

	t.Ticks = int(t.passedTime)
	if t.Ticks > maxTicksPerUpdate {
		t.Ticks = maxTicksPerUpdate
	}

	t.passedTime -= float32(t.Ticks)

	t.A = t.passedTime

	return t
}
