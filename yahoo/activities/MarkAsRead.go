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
				Weight: weight, Tasks: tasksContext,
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
		return true
	}

	return false
}

func (self *MarkAsRead) Run() {
	fmt.Println("[INFO] mark as read")

	self.SetMailActionByName("Mark as read", "Mark as unread")

	fmt.Println("[INFO] done")
}
