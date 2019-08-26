package activity

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

type Activity struct {
	/* How often the activity to be selected
	 * value = 1 (default), means equal probability to be selected
	 * value > 1, means the activity is more likely to be selected
	 */
	Weight int

	/* The Chromedp context for tasks */
	Tasks context.Context

	/* "Virtual" functions */
	VirtualIsAvailable func() bool
	VirtualRun         func()
}

// Verify if the current Activity is available
func (self *Activity) IsAvailable() bool {
	return self.VirtualIsAvailable()
}

// Runs the activity
func (self *Activity) Run() {
	self.VirtualRun()
}

func (self *Activity) RandomSleep() chromedp.Action {
	return chromedp.Sleep(time.Millisecond * time.Duration(rand.Intn(4000)+1000))
}
