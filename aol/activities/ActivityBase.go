package activities

import (
	"fmt"

	"github.com/mvinturis/mailbox-automation/activity"
	"github.com/chromedp/chromedp"
)

// ActivityBase class extends the Activity struct with common reusable methods
type ActivityBase struct {
	activity.Activity
}

func (self *ActivityBase) IsAvailableMailActionByName(name, dual string) bool {
	var value string
	errName := chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//div[not(contains(@style, "display: none"))]/table/descendant::tr[not(contains(@style, "display: none"))]/td[.="`+name+`"]')[0].className`, &value))
	errDual := chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//div[not(contains(@style, "display: none"))]/table/descendant::tr[not(contains(@style, "display: none"))]/td[.="`+dual+`"]')[0].className`, &value))

	if errName != nil && errDual != nil {
		// Open More mail actions menu
		err := chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//div[@role="tabpanel"][contains(@class, "dijitVisible")]/descendant::div[@class="containerNode"][not(contains(@style, "width: 0px") or contains(@style, "width: 2px"))]/descendant::div[.="More"]')[0].className`, &value),
			chromedp.Click(`//div[@role="tabpanel"][contains(@class, "dijitVisible")]/descendant::div[@class="containerNode"][not(contains(@style, "width: 0px") or contains(@style, "width: 2px"))]/descendant::div[.="More"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			return false
		}
	}

	errName = chromedp.Run(self.Context, chromedp.EvaluateAsDevTools(`$x('//div[not(contains(@style, "display: none"))]/table/descendant::tr[not(contains(@style, "display: none"))]/td[.="`+name+`"]')[0].className`, &value))
	if errName != nil {
		return false
	}

	return true
}

func (self *ActivityBase) SetMailActionByName(name, dual string) {
	if self.IsAvailableMailActionByName(name, dual) {
		chromedp.Run(self.Context,
			chromedp.Click(`//div[not(contains(@style, "display: none"))]/table/descendant::tr[not(contains(@style, "display: none"))]/td[.="`+name+`"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
	}
}
