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
				Weight: weight, Context: tasksContext,
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

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//*[@thefn="Spam"]/following-sibling::div/div[@class="unseen"]')[0].innerText`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] MarkAllNotSpam() not available: %s", err.Error())
		return false
	}
	intValue, _ := strconv.Atoi(value)
	if intValue > 0 {
		fmt.Println("[INFO] MarkAllNotSpam() available")
		return true
	}
	fmt.Println("[WARN] MarkAllNotSpam() not available")
	return false
}

func (self *MarkAllNotSpam) Run() {
	fmt.Println("[INFO] MarkAllNotSpam() running")
	chromedp.Run(self.Context,
		// Click Spam button
		chromedp.Click(`//*[@thefn="Spam"]/following-sibling::div/span`, chromedp.NodeVisible),
		self.RandomSleep(),
	)
	chromedp.Run(self.Context,
		// Select all messages
		chromedp.Click(`//div[@dojoattachpoint="headerContentNode"]/table/tbody/tr/th[contains(@class,"dojoxGrid-cell")][1]`, chromedp.NodeVisible),
		self.RandomSleep(),
	)

	self.SetMailActionByName("Inbox", "Spam")

	fmt.Println("[INFO] MarkAllNotSpam() done")
}
