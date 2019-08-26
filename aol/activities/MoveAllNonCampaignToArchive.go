package activities

import (
	"context"
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

// MoveAllNonCampaignToArchive describes the activity to move all messages NOT containing search campaign keyword to Archive folder
type MoveAllNonCampaignToArchive struct {
	ActivityBase

	SearchKeyword string
}

// NewMoveAllNonCampaignToArchive creates a new MoveAllNonCampaignToArchive object
func NewMoveAllNonCampaignToArchive(tasksContext context.Context, weight int, searchKeyword string) activity.Activity {
	a := MoveAllNonCampaignToArchive{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
		searchKeyword,
	}

	a.init()

	return a.Activity
}

func (self *MoveAllNonCampaignToArchive) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *MoveAllNonCampaignToArchive) IsAvailable() bool {

	if self.SearchKeyword == "" {
		return false
	}

	return true
}

func (self *MoveAllNonCampaignToArchive) Run() {
	fmt.Println("[INFO] Move all non campaign to archive")

	self.ActivityBase.SetSearchKeyword("-"+self.SearchKeyword, "Inbox")

	chromedp.Run(self.Tasks,
		// Select all messages
		chromedp.Click(`//button[@aria-label="Select all messages"]`, chromedp.NodeVisible), self.RandomSleep(),
		// Click Archive
		chromedp.Click(`//button[@title="Archive the selected conversations"]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	fmt.Println("[INFO] done")

	// Activity runs once
	self.SearchKeyword = ""
}
