package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

// StarMessages describes the activity to Star an opened message in Hotmail mailbox
type StarMessages struct {
	ActivityBase
}

// NewStarMessages creates a new StarMessages object
func NewStarMessages(tasksContext context.Context, weight int) activity.Activity {
	a := StarMessages{
		ActivityBase{
			activity.Activity{
				Weight: weight, Context: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *StarMessages) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *StarMessages) IsAvailable() bool {
	if self.IsAvailableMailActionByName("Star", "Clear star") {
		fmt.Println("[INFO] StarMessages() available")
		return true
	}
	fmt.Println("[WARN] StarMessages() not available")
	return false
}

func (self *StarMessages) Run() {
	fmt.Println("[INFO] StarMessages() running")
	self.SetMailActionByName("Star", "Clear star")
	fmt.Println("[INFO] StarMessages() done")
}
