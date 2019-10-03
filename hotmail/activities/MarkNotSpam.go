package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

type MarkNotSpam struct {
	ActivityBase
}

func NewMarkNotSpam(tasksContext context.Context, weight int) activity.Activity {
	a := MarkNotSpam{
		ActivityBase{
			activity.Activity{
				Weight: weight, Context: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *MarkNotSpam) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MarkNotSpam) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//*[@title="Junk Email"]/span/span/text()')[0].data`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] MarkNotSpam() not available: %s", err.Error())
		return false
	}
	fmt.Println("[INFO] MarkNotSpam() is available")
	return true
}

func (self *MarkNotSpam) Run() {
	fmt.Println("[DEBUG] MarkNotSpam() running")

	chromedp.Run(self.Context,
		// Click Junk Email button
		chromedp.Click(`//*[@title="Junk Email"]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	// selector for unread mesages
	selectorXPath := `//*[contains(@aria-label, "Unread")][1]`
	chromedp.Run(self.Context,
		// Open message
		// Set as read with ctrl+q
		chromedp.Click(selectorXPath, chromedp.NodeVisible), self.RandomSleep(),

		// Click Not Junk
		chromedp.Click(`//*[contains(@name, "Not junk")][1]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`//*[contains(@name, "Not junk")][contains(@class, "ms-ContextualMenu-link")]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	fmt.Println("[INFO] MarkNotSpam() done")
}
