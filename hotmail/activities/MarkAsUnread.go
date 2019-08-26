package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

type MarkAsUnread struct {
	activity.Activity
}

func NewMarkAsUnread(tasksContext context.Context, weight int) activity.Activity {
	a := MarkAsUnread{
		activity.Activity{
			Weight: weight, Tasks: tasksContext,
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
	return false
}

func (self *MarkAsUnread) Run() {
	fmt.Println("[INFO] mark as unread")
}
