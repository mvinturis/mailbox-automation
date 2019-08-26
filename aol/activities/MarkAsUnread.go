package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

type MarkAsUnread struct {
	ActivityBase
}

func NewMarkAsUnread(tasksContext context.Context, weight int) activity.Activity {
	a := MarkAsUnread{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *MarkAsUnread) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MarkAsUnread) IsAvailable() bool {

	if self.IsAvailableMailActionByName("Mark as unread", "Mark as read") {
		return true
	}

	return false
}

func (self *MarkAsUnread) Run() {
	fmt.Println("[INFO] mark as read")

	self.SetMailActionByName("Mark as unread", "Mark as read")

	fmt.Println("[INFO] done")
}
