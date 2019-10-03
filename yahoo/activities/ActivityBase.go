package activities

import (
	"github.com/mvinturis/mailbox-automation/activity"
	"fmt"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

// ActivityBase class extends the Activity struct with common reusable methods
type ActivityBase struct {
	activity.Activity
}

// GetSearchKeyword tests the current search string
func (self *ActivityBase) GetSearchKeyword() (keyword, filter string) {

	chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//div[@id="mail-search"]/descendant::div[@data-test-id="pill"]')[0].innerText`, &filter))
	chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//div[@id="mail-search"]/descendant::input[contains(@aria-label, "Search box")]')[0].value`, &keyword),
	)

	fmt.Println("[DEBUG] GetSearchKeyword(): filter = '%s', keyword = '%s'", filter, keyword)
	return
}

// SetSearchKeyword sets the specified keyword to the search box
func (self *ActivityBase) SetSearchKeyword(keyword string, filter string) {

	localKeyword, localFilter := self.GetSearchKeyword()

	if filter != localFilter || keyword != localKeyword {
		chromedp.Run(self.Context,
			// Click Search box
			chromedp.Click(`//input[contains(@aria-label, "Search box")]`, chromedp.NodeVisible), self.RandomSleep(),
			// Click Filters
			chromedp.Click(`//button[@title="Toggle advanced search pane"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if filter != localFilter {
			chromedp.Run(self.Context,
				// Select search Inbox folder
				chromedp.Click(`//span[.="Search in"]/parent::div/descendant::div[@data-test-id="selectbox-input"]/div/span`, chromedp.NodeVisible), self.RandomSleep(),
				chromedp.Click(`//*[@title="`+filter+`"][@role="menuitem"]`, chromedp.NodeVisible), self.RandomSleep(),
			)
		}
		if keyword != localKeyword {
			chromedp.Run(self.Context,
				// Click search keyword box
				chromedp.DoubleClick(`#adv-search-keyword-input`, chromedp.NodeVisible), self.RandomSleep(),
				chromedp.KeyEvent("\b\b", chromedp.KeyModifiers(input.ModifierNone)), self.RandomSleep(),
				// Input search keyword
				chromedp.SendKeys(`#adv-search-keyword-input`, keyword), self.RandomSleep(),
			)
		}
		chromedp.Run(self.Context,
			// Click Search button
			chromedp.Click(`//button[@title="Search"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
	}
}

func (self *ActivityBase) IsAvailableMailActionByName(name, dual string) bool {
	var value string
	errName := chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//li[@title="`+name+`"]')[0].type`, &value))
	errDual := chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//li[@title="`+dual+`"]')[0].type`, &value))

	if errName != nil && errDual != nil {
		// Check if the More mail actions menu is disabled
		err := chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//button[@title="More"][@disabled]')[0].type`, &value),
		)
		if err == nil {
			return false
		}
		// Open More mail actions menu
		chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//button[@title="More"]')[0].type`, &value),
			chromedp.Click(`//button[@title="More"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
	}

	errName = chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//li[@title="`+name+`"]')[0].type`, &value))

	if errName != nil {
		return false
	}

	return true
}

func (self *ActivityBase) SetMailActionByName(name, dual string) {
	if self.IsAvailableMailActionByName(name, dual) {
		chromedp.Run(self.Context,
			chromedp.Click(`//li[@title="`+name+`"]/a`, chromedp.NodeVisible), self.RandomSleep(),
		)
	}
}

func (self *ActivityBase) IsSetOptionMarkAsRead() bool {
	return false
}

func (self *ActivityBase) IsSetOptionCategorize() bool {
	return false
}
