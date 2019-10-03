package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

// MoveToArchive describes the activity to move an opened message to Archive folder
type MoveToArchive struct {
	ActivityBase
}

// NewMoveToArchive creates a new MoveToArchive object
func NewMoveToArchive(tasksContext context.Context, weight int) activity.Activity {
	a := MoveToArchive{
		ActivityBase{
			activity.Activity{
				Weight: weight, Context: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *MoveToArchive) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MoveToArchive) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//button[@title="Archive the selected conversations"]')[0].type`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] MoveToArchive() not available: %s", err.Error())
		return false
	}
	
	// Check if the Archive button is disabled
	err = chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//button[@title="Archive the selected conversations"][@disabled]')[0].type`, &value),
	)
	if err != nil {
		fmt.Println("[WARN] MoveToArchive() not available: %s", err.Error())
		return false
	}

	fmt.Println("[INFO] MoveToArchive() available")
	return true
}

func (self *MoveToArchive) Run() {
	fmt.Println("[INFO] MoveToArchive() running")

	chromedp.Run(self.Context,
		// Click Archive
		chromedp.Click(`//button[@title="Archive the selected conversations"]`, chromedp.NodeVisible), self.RandomSleep(),
	)
	
	fmt.Println("[INFO] MoveToArchive() done")
}
