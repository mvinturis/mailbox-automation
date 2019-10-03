package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
)

// PinMessages describes the activity to pin an opened message in Hotmail mailbox
type PinMessages struct {
	ActivityBase
}

// NewPinMessages creates a new PinMessages object
func NewPinMessages(tasksContext context.Context, weight int) activity.Activity {
	a := PinMessages{
		ActivityBase{
			activity.Activity{
				Weight: weight, Context: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *PinMessages) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *PinMessages) IsAvailable() bool {
	if self.IsAvailableMailActionByName("Pin", "Unpin") {
		fmt.Println("[INFO] PinMessages() is available")
		return true
	}
	fmt.Println("[WARN] PinMessages() is not available")
	return false
}

func (self *PinMessages) Run() {
	fmt.Println("[DEBUG] PinMessages() running")

	self.SetMailActionByName("Mark as read", "Mark as unread")
	self.SetMailActionByName("Pin", "Unpin")

	fmt.Println("[INFO] PinMessages() done")
}
