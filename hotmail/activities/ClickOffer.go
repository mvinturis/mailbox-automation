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
				Weight: weight, Context: tasksContext,
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

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//div[@class="wide-content-host"]/descendant::a[4]')[0].href`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] ClickOffer() not available: %s", err.Error())
		return false
	}
	fmt.Println("[INFO] ClickOffer() is available")
	return true
}

func (self *ClickOffer) Run() {
	var value string
	fmt.Println("[DEBUG] ClickOffer() running")

	// Get the first new tab that isn't blank.
	ch := chromedp.WaitNewTarget(self.Context, func(info *target.Info) bool {
		return info.URL != ""
	})

	chromedp.Run(self.Context,
		// Click on first link
		chromedp.EvaluateAsDevTools(`$x('//div[@class="wide-content-host"]/descendant::a[4]')[0].click()`, &value), self.RandomSleep(),
	)
	newTab, cancel := chromedp.NewContext(self.Context, chromedp.WithTargetID(<-ch))

	chromedp.Run(newTab,
		// Do nothing
		self.RandomSleep(),
	)

	// Close newTab
	cancel()

	fmt.Println("[INFO] ClickOffer() done")
}
