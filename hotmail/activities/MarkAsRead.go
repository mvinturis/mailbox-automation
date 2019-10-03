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
	if self.IsAvailableMailActionByName("Mark as read", "Mark as unread") {
		fmt.Println("[INFO] MarkAsRead() is available")
		return true
	}
	fmt.Println("[WARN] MarkAsRead() is not available")
	return false
}

func (self *MarkAsRead) Run() {
	fmt.Println("[DEBUG] MarkAsRead() running")

	self.SetMailActionByName("Mark as read", "Mark as unread")

	fmt.Println("[INFO] MarkAsRead() done")
}
