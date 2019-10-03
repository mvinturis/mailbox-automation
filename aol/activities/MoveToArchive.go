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
	if self.IsAvailableMailActionByName("Archive", "Spam") {
		fmt.Println("[INFO] MoveToArchive() available")
		return true
	}
	fmt.Println("[WARN] MoveToArchive() not available")
	return false
}

func (self *MoveToArchive) Run() {
	fmt.Println("[INFO] MoveToArchive() running")
	self.SetMailActionByName("Archive", "Spam")
	fmt.Println("[INFO] MoveToArchive() done")
}
