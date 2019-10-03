package activities

import (
	"context"
	"strings"

	"github.com/chromedp/cdproto/runtime"

	"github.com/chromedp/chromedp"

	"git.sparta.email/smash/grindeanu-common/models"
	"git.sparta.email/smash/grindeanu-runner/activity"
)

// OpenMessage describes the activity to open a message in Hotmail mailbox
type OpenMessage struct {
	ActivityBase

	SearchKeyword string
}

// NewOpenMessage creates a new OpenMessage object
func NewOpenMessage(tasksContext context.Context, task *models.Task, weight int) activity.Activity {
	a := OpenMessage{
		ActivityBase{
			activity.Activity{
				Context: tasksContext,
				Task:    task,
				Weight:  weight,
			},
		},
		task.Params.Keyword,
	}

	a.init()

	return a.Activity
}

func (self *OpenMessage) init() {
	self.Activity.VirtualIsAvailable = self.IsAvailable
	self.Activity.VirtualRun = self.Run
}

func (self *OpenMessage) IsAvailable() bool {
	var inboxLinkClasses string

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//div[*[name()="svg" and @thefn="Inbox"]]')[0].className`, &inboxLinkClasses),
	)
	if err != nil {
		fmt.Println("[WARN] OpenMessage() not available: %s", err.Error())
		return false
	} else {
		if strings.Contains(inboxLinkClasses, "hasUnseen") {
			fmt.Println("[INFO] OpenMessage() available")
			return true
		} else {
			fmt.Println("[INFO] OpenMessage() not available: no unseen messages in Inbox")
			return false
		}
	}
}

func (self *OpenMessage) Run() {
	var result *runtime.RemoteObject
	fmt.Println("[INFO] OpenMessage() running")

	// switch to inbox
	err := chromedp.Run(self.Context,
		chromedp.Click(`//div[*[name()="svg" and @thefn="Inbox"]]`, chromedp.NodeVisible),
		self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] OpenMessage() error: %s", err.Error())
		return
	}

	chromedp.Run(self.Context,
		// Open message first message
		chromedp.EvaluateAsDevTools(`$x('//td[@class="dojoxGrid-cell gridColFrom "]/span[@class="fromAddress"]')[0].click()`, result),
		self.RandomSleep(),
	)
	fmt.Println("[INFO] OpenMessage() done")
}
