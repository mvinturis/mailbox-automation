package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

// FlagMessages describes the activity to flag an opened message in Hotmail mailbox
type FlagMessages struct {
	ActivityBase
}

// NewFlagMessages creates a new FlagMessage object
func NewFlagMessages(tasksContext context.Context, weight int) activity.Activity {
	a := FlagMessages{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *FlagMessages) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *FlagMessages) IsAvailable() bool {

	if self.IsAvailableMailActionByName("Flag", "Unflag") {
		return true
	}

	return false
}

func (self *FlagMessages) Run() {
	fmt.Println("[INFO] Flag messages")

	self.SetMailActionByName("Flag", "Unflag")
	fmt.Println("[INFO] done")
}
