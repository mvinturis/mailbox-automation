package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

type MarkAsRead struct {
	ActivityBase
}

func NewMarkAsRead(tasksContext context.Context, weight int) activity.Activity {
	a := MarkAsRead{
		ActivityBase{
			activity.Activity{
				Weight: weight, Context: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *MarkAsRead) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MarkAsRead) IsAvailable() bool {
	if self.IsAvailableMailActionByName("Mark Read", "Mark Unread") {
		fmt.Println("[INFO] MarkAsRead() available")
		return true
	}

	fmt.Println("[WARN] MarkAsRead() not available")
	return false
}

func (self *MarkAsRead) Run() {
	fmt.Println("[INFO] MarkAsRead() running")
	self.SetMailActionByName("Mark Read", "Mark Unread")
	fmt.Println("[INFO] MarkAsRead() done")
}
