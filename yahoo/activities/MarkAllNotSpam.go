package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

type MarkAllNotSpam struct {
	ActivityBase
}

func NewMarkAllNotSpam(tasksContext context.Context, weight int) activity.Activity {
	a := MarkAllNotSpam{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *MarkAllNotSpam) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MarkAllNotSpam) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('((//a[@data-test-folder-name="Bulk"])[1]/span)[2]/span/text()')[0].data`, &value),
	)
	if err != nil {
		return false
	}

	return true
}

func (self *MarkAllNotSpam) Run() {
	fmt.Println("[INFO] mark all not spam...")

	chromedp.Run(self.Tasks,
		// Click Junk Email button
		chromedp.Click(`//a[@data-test-folder-name="Bulk"]`, chromedp.NodeVisible),
		self.RandomSleep(),
	)

	chromedp.Run(self.Tasks,

		//Click Junk Email
		chromedp.Click(`//div[@title="Junk Email"][1]`, chromedp.NodeVisible), self.RandomSleep(),

		// Select all messages
		chromedp.Click(`//div[@aria-label="Select all messages"]/descendant::i[@data-icon-name="StatusCircleCheckmark"][1]`, chromedp.NodeVisible), self.RandomSleep(),

		// Click Not Junk
		chromedp.Click(`//i[@data-icon-name="ChevronDown"]/ancestor::button[@name="Not junk"][1]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`//span[.="Not junk"]/ancestor::button[@name="Not junk"][1]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	fmt.Println("[INFO] done")
}
