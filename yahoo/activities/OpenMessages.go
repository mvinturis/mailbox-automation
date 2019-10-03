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
				Weight: weight, Context: tasksContext,
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

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('((//a[@data-test-folder-name="Inbox"])[1]/span)[2]/span/text()')[0].data`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] OpenMessages() not available: %s", err.Error())
		return false
	}
	fmt.Println("[INFO] OpenMessages() available")
	return true
}

func (self *OpenMessages) Run() {
	fmt.Println("[INFO] OpenMessages() running")
	var value string

	self.ActivityBase.SetSearchKeyword(self.SearchKeyword, "Inbox")

	// xpath to search for the first unread message
	selectorXPath := `//a[@data-test-id="message-list-item"]/descendant::button[@data-test-id="icon-btn-checkbox"][@aria-label="Select message"][1]`

	alreadySelectedXPath := `$x('//a[@data-test-id="message-list-item"]/descendant::button[@data-test-id="icon-btn-checkbox"][@aria-label="Deselect message"]')[0].type`
	err := chromedp.Run(self.Context,
		// Check if a message is already selected
		chromedp.EvaluateAsDevTools(alreadySelectedXPath, &value),
	)
	if err != nil {
		chromedp.Run(self.Context,
			// Open message
			chromedp.Click(selectorXPath, chromedp.NodeVisible), self.RandomSleep(),
			// FIXME: this does not work
			// Set as read with Shift+k
			// chromedp.KeyEvent("k", chromedp.KeyModifiers(input.ModifierShift)), self.RandomSleep(),
		)
		// Workaround, since Shft+k does not work
		self.SetMailActionByName("Mark as read", "Mark as unread")
	}

	fmt.Println("[INFO] OpenMessages() done")
}
