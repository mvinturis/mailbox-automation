package runner

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mvinturis/mailbox-automation/activity"
	"github.com/mvinturis/mailbox-automation/common/models"

	"github.com/chromedp/chromedp"
)

type Runner struct {
	/* User profile */
	Profile *models.Seed

	/* The Chromedp context for tasks */
	Context context.Context

	/* "Virtual" functions */
	VirtualGetAvailableActivities func() []activity.Activity
	VirtualIsSignedIn             func() bool
	VirtualLogin                  func() bool
	VirtualLogout                 func()
	VirtualInitActivities         func(*models.TaskParams)
	VirtualReadMessages           func(*models.TaskParams)

	// metoda principala pentru a "porni" un runner cu behaviour-ul dorit + parametrii
	VirtualStart func(behaviour string, params *models.TaskParams)
}

func (self *Runner) Start(behaviour string, params *models.TaskParams) {
	self.VirtualStart(behaviour, params)
}

func (self *Runner) Login() bool {
	return self.VirtualLogin()
}

func (self *Runner) IsSignedIn() bool {
	return self.VirtualIsSignedIn()
}

func (self *Runner) Logout() {
	self.VirtualLogout()
}

func (self *Runner) InitActivities(params *models.TaskParams) {
	self.VirtualInitActivities(params)
}

func (self *Runner) GetAvailableActivities() []activity.Activity {
	return self.VirtualGetAvailableActivities()
}

func (self *Runner) RandomSleep() chromedp.Action {
	return chromedp.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)+500))
}

// ReadMessages() generic care ar putea functiona pentru orice runner de chromedp, pe baza de activitati + ponderi
func (self *Runner) ReadMessages(params *models.TaskParams) {
	fmt.Println("[INFO] read messages... ")

	self.InitActivities(params)

	for retryIndex, retryMaxCount := 0, 2; retryIndex < retryMaxCount; {

		if !self.IsSignedIn() {
			self.Login()
			retryIndex++
			continue
		}
		retryIndex = 0

		activities := self.GetAvailableActivities()
		if len(activities) <= 0 {
			fmt.Println("[INFO] finished... no more activities")
			break
		}

		// Compute total weights amount
		totalWeights := 1
		for _, activity := range activities {
			totalWeights += activity.Weight
		}

		// Pick one random activity and run it
		randomWeight := rand.Intn(totalWeights)
		for _, activity := range activities {
			if randomWeight > activity.Weight {
				randomWeight -= activity.Weight
				continue
			}
			activity.Run()
			break
		}

		chromedp.Run(self.Context,
			self.RandomSleep(),
		)
	}

	fmt.Println("[INFO] read messages done")
}
