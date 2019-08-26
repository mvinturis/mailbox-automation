package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

// OpenMessages describes the activity to open a message in Hotmail mailbox
type OpenMessages struct {
	ActivityBase

	SearchKeyword string
}

// NewOpenMessages creates a new OpenMessage object
func NewOpenMessages(tasksContext context.Context, weight int, searchKeyword string) activity.Activity {
	a := OpenMessages{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
		searchKeyword,
	}

	a.init()

	return a.Activity
}

func (self *OpenMessages) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *OpenMessages) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('((//div[@title="Inbox"])[1]/span)[2]/span/text()')[0].data`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] activity 'OpenMessages' not available: %s", err.Error())
		return false
	}

	return true
}

func (self *OpenMessages) Run() {
	fmt.Println("[INFO] open message... ")

	self.ActivityBase.SetSearchKeyword(self.SearchKeyword, "Inbox")

	// xpath to search for the first unread message
	// selectorXPath := `//div[@data-convid!=""][starts-with(@aria-label,"Unread")][1]`
	selectorXPath := `//div[@data-convid!=""][1]`

	chromedp.Run(self.Tasks,
		// Open message
		chromedp.Click(selectorXPath, chromedp.NodeVisible), self.RandomSleep(),
		// Set as read with ctrl+q
		chromedp.KeyEvent("q", chromedp.KeyModifiers(2)), self.RandomSleep(),
	)

	fmt.Println("[INFO] done")
}
