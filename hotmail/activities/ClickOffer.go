package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

type ClickOffer struct {
	ActivityBase
}

func NewClickOffer(tasksContext context.Context, weight int) activity.Activity {
	a := ClickOffer{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *ClickOffer) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *ClickOffer) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('//div[@class="wide-content-host"]/descendant::a[4]')[0].href`, &value),
	)
	if err != nil {
		return false
	}

	return true
}

func (self *ClickOffer) Run() {
	var value string
	fmt.Println("[INFO] click offer...")

	// Get the first new tab that isn't blank.
	ch := chromedp.WaitNewTarget(self.Tasks, func(info *target.Info) bool {
		return info.URL != ""
	})

	chromedp.Run(self.Tasks,
		// Click on first link
		chromedp.EvaluateAsDevTools(`$x('//div[@class="wide-content-host"]/descendant::a[4]')[0].click()`, &value), self.RandomSleep(),
	)
	newTab, cancel := chromedp.NewContext(self.Tasks, chromedp.WithTargetID(<-ch))

	chromedp.Run(newTab,
		// Do nothing
		self.RandomSleep(),
	)

	// Close newTab
	cancel()

	fmt.Println("[INFO] done")
}
