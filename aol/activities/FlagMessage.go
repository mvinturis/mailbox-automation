package activities

import (
	"context"

	"git.sparta.email/smash/grindeanu-common/models"
	"git.sparta.email/smash/grindeanu-runner/activity"
)

type FlagMessage struct {
	ActivityBase
}

func NewFlagMessage(tasksContext context.Context, task *models.Task, weight int) activity.Activity {
	a := FlagMessage{
		ActivityBase{
			activity.Activity{
				Context: tasksContext,
				Task:    task,
				Weight:  weight,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *FlagMessage) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *FlagMessage) IsAvailable() bool {
	if self.IsAvailableMailActionByName("Flag Message", "Clear Flag") {
		fmt.Println("[INFO] FlagMessage() available")
		return true
	} else {
		fmt.Println("[WARN] FlagMessage() not available")
		return false
	}
}

func (self *FlagMessage) Run() {
	fmt.Println("[INFO] FlagMessage() running")
	self.SetMailActionByName("Flag Message", "Clear Flag")
	fmt.Println("[INFO] FlagMessage() done")
}
