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
		chromedp.EvaluateAsDevTools(`$x('((//a[@data-test-folder-name="Inbox"])[1]/span)[2]/span/text()')[0].data`, &value),
	)
	if err != nil {
		return false
	}

	return true
}

func (self *OpenMessages) Run() {
	fmt.Println("[INFO] open message... ")
	var value string

	self.ActivityBase.SetSearchKeyword(self.SearchKeyword, "Inbox")

	// xpath to search for the first unread message
	selectorXPath := `//a[@data-test-id="message-list-item"]/descendant::button[@data-test-id="icon-btn-checkbox"][@aria-label="Select message"][1]`

	alreadySelectedXPath := `$x('//a[@data-test-id="message-list-item"]/descendant::button[@data-test-id="icon-btn-checkbox"][@aria-label="Deselect message"]')[0].type`
	err := chromedp.Run(self.Tasks,
		// Check if a message is already selected
		chromedp.EvaluateAsDevTools(alreadySelectedXPath, &value),
	)
	if err != nil {
		chromedp.Run(self.Tasks,
			// Open message
			chromedp.Click(selectorXPath, chromedp.NodeVisible), self.RandomSleep(),
			// Set as read with Shift+k
			// chromedp.KeyEvent("k", chromedp.KeyModifiers(input.ModifierShift)), self.RandomSleep(),
		)
		// Workaround, since Shft+k does not work
		self.SetMailActionByName("Mark as read", "Mark as unread")
	}

	fmt.Println("[INFO] done")
}
