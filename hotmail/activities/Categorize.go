package activities

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mvinturis/mailbox-automation/activity"

	"github.com/chromedp/chromedp"
)

type Categorize struct {
	ActivityBase
}

func NewCategorize(tasksContext context.Context, weight int) activity.Activity {
	a := Categorize{
		ActivityBase{
			activity.Activity{
				Weight: weight, Tasks: tasksContext,
			},
		},
	}

	a.init()

	return a.Activity
}

func (self *Categorize) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *Categorize) IsAvailable() bool {
	var value string

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('(//*[contains(@aria-label, "Opens Profile Card")])//text()')[0].data`, &value),
	)
	if err != nil {
		return false
	}

	categorySelector := "$x(`(//div[@aria-label='Content pane']/descendant::div[" +
		"contains(@title, 'category') " +
		"and (contains(@title, 'Purple') " +
		"or contains(@title, 'Blue')" +
		"or contains(@title, 'Green')" +
		"or contains(@title, 'Yellow')" +
		"or contains(@title, 'Orange')" +
		"or contains(@title, 'Red')" +
		")])`)[0].title"

	err = chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(categorySelector, &value),
	)
	if err != nil {
		// No categories were found
		// Activity is available
		return true
	}

	// One or more categories were found
	// Activity is not available

	return false
}

func (self *Categorize) Run() {
	fmt.Println("[INFO] Categorize Mail Randomly")

	// Get random Category for email

	rand.Seed(time.Now().Unix())

	// Build slice with xpath selectors for all categories
	categorySelectors := []string{
		`(//*[contains(@title, "Purple category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
		`(//*[contains(@title, "Blue category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
		`(//*[contains(@title, "Green category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
		`(//*[contains(@title, "Yellow category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
		`(//*[contains(@title, "Orange category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
		`(//*[contains(@title, "Red category")])[contains(@role, "menuitemcheckbox")]/div[2]`,
	}
	// get random int
	n := rand.Int() % len(categorySelectors)
	category := categorySelectors[n]
	fmt.Println("[INFO] Will categorize email as %s", category)

	chromedp.Run(self.Tasks,
		// Click Categorize
		chromedp.Click(`//*[@name="Categorize"]`, chromedp.NodeVisible),
		self.RandomSleep(),

		// Click Category
		chromedp.Click(category, chromedp.NodeVisible),
		self.RandomSleep(),
	)

	fmt.Println("[INFO] done")
}
