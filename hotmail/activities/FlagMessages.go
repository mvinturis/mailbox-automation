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
				Weight: weight, Context: tasksContext,
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
		fmt.Println("[INFO] FlagMessages() is available")
		return true
	}
	fmt.Println("[WARN] FlagMessages() is not available")
	return false
}

func (self *FlagMessages) Run() {
	fmt.Println("[DEBUG] FlagMessages() running")

	self.SetMailActionByName("Flag", "Unflag")
	
	fmt.Println("[INFO] FlagMessages() done")
}
