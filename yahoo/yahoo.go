package yahoo

import (
	"context"
	"fmt"
	"time"

	"github.com/mvinturis/mailbox-automation/activity"
	"github.com/mvinturis/mailbox-automation/common/models"
	"github.com/mvinturis/mailbox-automation/runner"
	"github.com/mvinturis/mailbox-automation/yahoo/activities"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

const YAHOO_START_PAGE = "https://mail.yahoo.com"

// VERIFY_SIGNED_IN_TIMEOUT timeout in secunde pentru cat timp asteptam dupa element vizibil care ne zice daca suntem logati in yahoo
const VERIFY_SIGNED_IN_TIMEOUT = 10

type Runner struct {
	runner.Runner
	Activities []activity.Activity
}

func NewRunner(seed *models.Seed, tasksContext context.Context) runner.Runner {
	r := Runner{
		runner.Runner{
			Profile: seed, Context: tasksContext,
		},
		nil,
	}

	r.init()

	return r.Runner
}

func (self *Runner) init() {
	self.Runner.VirtualGetAvailableActivities = self.GetAvailableActivities
	self.Runner.VirtualIsSignedIn = self.IsSignedIn
	self.Runner.VirtualLogin = self.Login
	self.Runner.VirtualLogout = self.Logout
	self.Runner.VirtualInitActivities = self.InitActivities
	self.Runner.VirtualReadMessages = self.ReadMessages
	self.Runner.VirtualStart = self.Start
}

func (self *Runner) Start(behaviour string, params *models.TaskParams) {
	self.Login()
	self.ChangeLanguageEnglish()

	switch behaviour {
	case "readMessages":
		self.ReadMessages(params)
	case "logout":
		self.Logout()
	}
}

func (self *Runner) Login() bool {
	self.navigateToYahoo()
	if self.IsSignedIn() {
		fmt.Println("[DEBUG] already signed in")
		return true
	}

	for retryIndex, retryMaxCount := 0, 3; retryIndex < retryMaxCount; {
		if self.verifyLoginMode() {
			self.mailboxLogin()
			if self.IsSignedIn() {
				return true
			}
			continue
		}
		if self.IsSignedIn() {
			return true
		}
		retryIndex++
		chromedp.Run(self.Context,
			self.RandomSleep(),
		)
	}

	return false
}

// IsSignedIn verifica daca suntem autentificati in casuta de yahoo dupa elementul #searchBoxId
// asteptam 10 secunde sa vedem daca e vizibil, daca nu -> timeout
func (self *Runner) IsSignedIn() bool {
	fmt.Println("[DEBUG] IsSignedIn(): start")	
	var value string

	timeout := time.After(time.Second * VERIFY_SIGNED_IN_TIMEOUT)
	result := make(chan error)

	go func() {
		err := chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//div[@id="mail-search"]')[0].innerText`, &value),
			// chromedp.WaitVisible("#mail-search"),
		)

		result <- err
	}()

	select {
	case res := <-result:
		if res == nil {
			fmt.Println("[INFO] IsSignedIn(): success, we're logged in")
			return true
		}
		fmt.Println("[ERROR] error while waiting signed-in verification: %s\n", res.Error())
		return false

	case <-timeout:
		fmt.Println("[INFO] timeout while verifying if we're signed in, we're probably not logged in..")
		return false
	}
}

func (self *Runner) Logout() {
	fmt.Println("[DEBUG] logging out")
	var value string

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="ybarAccountMenu"]')[0].value`, &value),
		chromedp.Click(`//input[@id="ybarAccountMenu"]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//a[contains(@data-ylk, "sign out")]')[0].href`, &value),
		chromedp.Click(`//a[contains(@data-ylk, "sign out")]`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] Logout() failed: %s", err.Error())
	} else {
		fmt.Println("[INFO] logged out")
	}
}

func (self *Runner) GetAvailableActivities() []activity.Activity {
	var availableActivities []activity.Activity

	for _, a := range self.Activities {
		if a.IsAvailable() {
			availableActivities = append(availableActivities, a)
		}
	}

	return availableActivities
}

func (self *Runner) navigateToYahoo() error {
	fmt.Println("[DEBUG] navigating to yahoo start page")

	return chromedp.Run(self.Context, chromedp.Navigate(YAHOO_START_PAGE))
}

func (self *Runner) verifyLoginMode() bool {
	var title string

	fmt.Println("[INFO] verify login mode... ")

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="login-username"]')[0].value`, &title),
	)
	if err != nil {
		fmt.Println("[ERROR] verifyLoginMode(): %s", err.Error())
		return false
	}

	fmt.Println("[INFO] verifyLoginMode(): success")

	return true
}

func (self *Runner) mailboxLogin() error {

	fmt.Println("[INFO] login... ")

	err := chromedp.Run(self.Context,
		chromedp.DoubleClick(`#login-username`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.KeyEvent("\b", chromedp.KeyModifiers(input.ModifierNone)), self.RandomSleep(),
		chromedp.SendKeys(`#login-username`, self.Profile.Email+"\n"), self.RandomSleep(),
		chromedp.WaitVisible(`#login-passwd`, chromedp.ByID),
		chromedp.SendKeys(`#login-passwd`, self.Profile.Password+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] Can't navigate to yahoo: ", err)
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) InitActivities(params *models.TaskParams) {
	searchKeyword := params.Keyword

	// set weights for activities
	self.Activities = []activity.Activity{
		// activities.NewMarkAllNotSpam(self.Context, 100),
		activities.NewOpenMessages(self.Context, 5, searchKeyword),
		// activities.NewClickOffer(self.Context, 10),
		activities.NewMarkAsRead(self.Context, 100),
		activities.NewMarkAsUnread(self.Context, 1),
		activities.NewStarMessages(self.Context, 100),
		activities.NewMoveToArchive(self.Context, 20),
	}
}

func (self *Runner) ChangeLanguageEnglish() {
	var title, value, language string
	elements := []string{"Inbox", "Archive", "Sent", "Trash"}
	changeLanguage := false
	expectedLanguage := "en-US"
	titleLanguage := "english - united states"

	for _, expectedTitle := range elements {
		err := chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//span[@data-test-folder-name="`+expectedTitle+`"]')[0].innerText`, &title),
		)
		if err != nil {
			changeLanguage = true
			break
		}
		if title != expectedTitle {
			changeLanguage = true
			break
		}
	}
	if !changeLanguage {
		return
	}

	fmt.Println("[INFO] change language to English - United States")

	// Get the first new tab that isn't blank.
	ch := chromedp.WaitNewTarget(self.Context, func(info *target.Info) bool {
		return info.URL != ""
	})

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="ybarAccountMenu"]')[0].value`, &value),
		chromedp.Click(`//input[@id="ybarAccountMenu"]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//a[contains(@data-ylk, "Account Info")]')[0].href`, &value),
		chromedp.Click(`//a[contains(@data-ylk, "Account Info")]`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		return
	}

	newTab, cancel := chromedp.NewContext(self.Context, chromedp.WithTargetID(<-ch))

	chromedp.Run(newTab,
		chromedp.EvaluateAsDevTools(`$x('//li[contains(@class,"li-preferences")]')[0].innerText`, &value),
		chromedp.Click(`//li[contains(@class,"li-preferences")]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//select[@id="language-dropdown"]')[0].value`, &language),
	)
	if language != expectedLanguage {
		chromedp.Run(newTab,
			chromedp.EvaluateAsDevTools(`$x('//select[@id="language-dropdown"]')[0].value`, &value),
			chromedp.Click(`//select[@id="language-dropdown"]`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.SendKeys(`//select[@id="language-dropdown"]`, titleLanguage+"\n"), self.RandomSleep(), self.RandomSleep(),
		)
	}

	// Close newTab
	cancel()

	// Wait for Language change to take effect
	chromedp.Run(self.Context,
		self.RandomSleep(), self.RandomSleep(),
	)

	fmt.Println("[INFO] done")
}
