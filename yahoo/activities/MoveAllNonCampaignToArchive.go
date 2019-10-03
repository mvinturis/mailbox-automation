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
				Weight: weight, Context: tasksContext,
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
		fmt.Println("[ERROR] MoveAllNonCampaignToArchive(): search keyword is empty!")
		return false
	}
	fmt.Println("[INFO] MoveAllNonCampaignToArchive() available")
	return true
}

func (self *MoveAllNonCampaignToArchive) Run() {
	fmt.Println("[INFO] MoveAllNonCampaignToArchive() running")

	self.ActivityBase.SetSearchKeyword("-"+self.SearchKeyword, "Inbox")

	chromedp.Run(self.Context,
		// Select all messages
		chromedp.Click(`//button[@aria-label="Select all messages"]`, chromedp.NodeVisible), self.RandomSleep(),
		// Click Archive
		chromedp.Click(`//button[@title="Archive the selected conversations"]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	// Activity runs once
	self.SearchKeyword = ""

	fmt.Println("[INFO] MoveAllNonCampaignToArchive() done")
}
