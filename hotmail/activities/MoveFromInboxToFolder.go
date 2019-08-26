package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

type MoveFromInboxToFolder struct {
	activity.Activity
}

func NewMoveFromInboxToFolder(tasksContext context.Context, weight int) activity.Activity {
	a := MoveFromInboxToFolder{
		activity.Activity{
			Weight: weight, Tasks: tasksContext,
		},
	}

	a.init()

	return a.Activity
}

func (self *MoveFromInboxToFolder) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MoveFromInboxToFolder) IsAvailable() bool {
	return false
}

func (self *MoveFromInboxToFolder) Run() {
	fmt.Println("[INFO] move from inbox to folder")
}
