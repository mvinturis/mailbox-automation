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
				Weight: weight, Context: tasksContext,
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
	if self.IsAvailableMailActionByName("Mark Unread", "Mark Read") {
		fmt.Println("[INFO] MarkAsUnread() available")
		return true
	}
	fmt.Println("[WARN] MarkAsUnread() not available")
	return false
}

func (self *MarkAsUnread) Run() {
	fmt.Println("[INFO] MarkAsUnread() running")
	self.SetMailActionByName("Mark Unread", "Mark Read")
	fmt.Println("[INFO] MarkAsUnread() done")
}
