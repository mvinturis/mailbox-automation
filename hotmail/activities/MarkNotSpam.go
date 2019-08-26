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
				Weight: weight, Tasks: tasksContext,
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

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('//*[@title="Junk Email"]/span/span/text()')[0].data`, &value),
	)
	if err != nil {
		return false
	}

	return true
}

func (self *MarkNotSpam) Run() {
	fmt.Println("[INFO] mark not spam")

	chromedp.Run(self.Tasks,
		// Click Junk Email button
		chromedp.Click(`//*[@title="Junk Email"]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	// selector for unread mesages
	selectorXPath := `//*[contains(@aria-label, "Unread")][1]`
	chromedp.Run(self.Tasks,
		// Open message
		// Set as read with ctrl+q
		chromedp.Click(selectorXPath, chromedp.NodeVisible), self.RandomSleep(),

		// Click Not Junk
		chromedp.Click(`//*[contains(@name, "Not junk")][1]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`//*[contains(@name, "Not junk")][contains(@class, "ms-ContextualMenu-link")]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	fmt.Println("[INFO] done")
}
