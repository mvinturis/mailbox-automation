package aol

import (
	"context"
	"fmt"
	"time"

	"github.com/mvinturis/mailbox-automation/activity"
	"github.com/mvinturis/mailbox-automation/aol/activities"
	"github.com/mvinturis/mailbox-automation/common/models"
	"github.com/mvinturis/mailbox-automation/common/smspva"
	"github.com/mvinturis/mailbox-automation/runner"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

const AOL_START_PAGE = "https://mail.aol.com"

// VERIFY_SIGNED_IN_TIMEOUT timeout in secunde pentru cat timp asteptam dupa element vizibil care ne zice daca suntem logati in aol
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

	switch behaviour {
	case "readMessages":
		self.Login()
		self.ChangeLanguageEnglish()
		self.ReadMessages(params)
	case "createNewSeed":
		self.createNewSeed(params)
	case "logout":
		self.Login()
		self.Logout()
	}
}

func (self *Runner) Login() bool {
	self.navigateToAol()
	if self.IsSignedIn() {
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

func (self *Runner) IsSignedIn() bool {
	var value string

	timeout := time.After(time.Second * VERIFY_SIGNED_IN_TIMEOUT)
	result := make(chan error)

	go func() {
		err := chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`document.title`, &pageTitle),
		)

		result <- err
	}()

	select {
	case res := <-result:
		if res == nil {
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
	fmt.Println("[INFO] logout... ")
	var value string

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="ybarAccountMenu"]')[0].value`, &value),
		chromedp.Click(`//input[@id="ybarAccountMenu"]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//a[contains(@data-ylk, "sign out")]')[0].href`, &value),
		chromedp.Click(`//a[contains(@data-ylk, "sign out")]`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s: %s", fn, err.Error())
	} else {
		fmt.Println("[INFO] %s: logged out", fn)
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

func (self *Runner) navigateToAol() error {
	fmt.Println("[INFO] navigating to aol...")

	return chromedp.Run(self.Context, chromedp.Navigate(AOL_START_PAGE))
}

func (self *Runner) verifyLoginMode() bool {
	var title string

	fmt.Println("[INFO] verifyLoginMode()... ")

	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="login-username"]')[0].value`, &title),
	)
	if err != nil {
		fmt.Println("[ERROR] %s: %s", fn, err.Error())
		return false
	}

	fmt.Println("[INFO] verifyLoginMode() success")

	return true
}

func (self *Runner) mailboxLogin() error {

	fmt.Println("[INFO] login... ")

	err := chromedp.Run(self.Context,
		chromedp.DoubleClick(`#login-username`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.KeyEvent("\b", chromedp.KeyModifiers(input.ModifierNone)), self.RandomSleep(),
		chromedp.SendKeys(`#login-username`, self.Profile.Email+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] Can't navigate to aol: ", err)
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) InitActivities() {
	fmt.Println("[INFO] InitActivities")

	// set weights for activities
	self.Activities = []activity.Activity{
		activities.NewMarkAllNotSpam(self.Context, self.Task, 100),
		activities.NewOpenMessage(self.Context, self.Task, 1),
		activities.NewFlagMessage(self.Context, self.Task, 100),
		activities.NewMarkAsRead(self.Context, self.Task, 50),
		activities.NewMarkAsUnread(self.Context, self.Task, 1),
		activities.NewMoveToArchive(self.Context, self.Task, 20),
	}
}

func (self *Runner) ChangeLanguageEnglish() error {
	fn := helpers.FuncName()

	var langLocale, value, language, expectedLanguage, titleLanguage string

	// js -> window.Settings.Locale
	err := chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`window.Settings.Locale`, &langLocale),
	)
	if err != nil {
		return err
	}

	if strings.HasPrefix(langLocale, "en-") {
		fmt.Println("[INFO] %s: language locale is already set to '%s'", fn, langLocale)
		return nil
	}

	fmt.Println("[INFO] %s: language locale set to '%s', changing to english", fn, langLocale)

	// Get the first new tab that isn't blank.
	ch := chromedp.WaitNewTarget(self.Context, func(info *target.Info) bool {
		return info.URL != ""
	})

	err = chromedp.Run(self.Context,
		chromedp.EvaluateAsDevTools(`$x('//input[@id="ybarAccountMenu"]')[0].value`, &value),
		chromedp.Click(`//input[@id="ybarAccountMenu"]`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//a[contains(@data-ylk, "Account Info")]')[0].href`, &value),
		chromedp.Click(`//a[contains(@data-ylk, "Account Info")]`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		return err
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

	return nil
}

func (self *Runner) createNewSeed(params *models.TaskParams) {
	var value string

	fmt.Println("[DEBUG] createNewSeed() starting")

	self.navigateToAol()
	if self.IsSignedIn() {
		fmt.Println("[INFO] createNewSeed(): already signed in, logging out..")
		self.Logout()
		self.navigateToAol()
	}

	birthdate := NewRandDate()
	name := NewRandName()

	err := chromedp.Run(self.Context,
		// Login screen
		chromedp.EvaluateAsDevTools(`$x('//a[@id="createacc"]')[0].href`, &value),
		chromedp.Click(`#createacc`, chromedp.NodeVisible), self.RandomSleep(),
		// Sign up
		chromedp.DoubleClick(`#usernamereg-firstName`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.KeyEvent("\b", chromedp.KeyModifiers(input.ModifierNone)), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-firstName`, name.FirstName), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-lastName`, name.LastName), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-yid`, self.Profile.Email), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-password`, self.Profile.Password), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-month`, birthdate.Month), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-day`, birthdate.Day), self.RandomSleep(),
		chromedp.SendKeys(`#usernamereg-year`, birthdate.Year), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] createNewSeed() usernamereg: %s", err.Error())
		return
	}

	// Generate a new phone number and security code
	for countryCode := range smspva.IsoCountryCodes {
		var phoneNumber, phoneID, countryPrefix string
		for phoneNumber, phoneID, countryPrefix = smspva.GetPhoneNumber(smspva.SERVICEIDAOL, countryCode); phoneNumber == ""; phoneNumber, phoneID, countryPrefix = smspva.GetPhoneNumber(smspva.SERVICEIDAOL, countryCode) {
			fmt.Println("[DEBUG] createNewSeed(): retrying smspva.GetPhoneNumber() in 10 seconds")
			time.Sleep(time.Second * 10)
		}
		self.Profile.Phone = countryPrefix + " " + phoneNumber

		err = chromedp.Run(self.Context,
			// Erase existing phoneNumber
			chromedp.DoubleClick(`#usernamereg-phone`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.KeyEvent("\b", chromedp.KeyModifiers(0)), self.RandomSleep(),
			// Fill in new phoneNumber
			chromedp.SendKeys(`#usernamereg-phone`, phoneNumber), self.RandomSleep(),
			// Select Country code
			chromedp.EvaluateAsDevTools(`$x('//select[@name="shortCountryCode"]')[0].value = "`+countryCode+`"`, &value),
			// Click Continue
			chromedp.Click(`#reg-submit-button`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[ERROR] createNewSeed() reg-submit: %s", err.Error())
			continue
		}

		// Check if the phoneNumber was accepted
		err = chromedp.Run(self.Context,
			chromedp.EvaluateAsDevTools(`$x('//div[@id="reg-error-phone"]')[0].innerText`, &value),
		)
		if err == nil {
			fmt.Println("[ERROR] createNewSeed() phone not accepted: %s", err.Error())
			continue
		}

		err = chromedp.Run(self.Context,
			// Click Send Code
			chromedp.Click(`//button[@name="sendCode"]`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.WaitVisible(`#verification-code-field`, chromedp.ByID),
		)
		if err != nil {
			fmt.Println("[ERROR] createNewSeed() sendCode field: %s", err.Error())
			continue
		}

		var sms string
		for sms = smspva.GetSms(smspva.SERVICEIDAOL, countryCode, phoneID); sms == ""; sms = smspva.GetSms(smspva.SERVICEIDAOL, countryCode, phoneID) {
			fmt.Println("[DEBUG] createNewSeed(): retrying smspva.GetSms() in 10 seconds")
			time.Sleep(time.Second * 10)
		}

		err = chromedp.Run(self.Context,
			chromedp.SendKeys(`#verification-code-field`, sms), self.RandomSleep(),
			chromedp.Click(`#verifyCodeButton`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[DEBUG] createNewSeed() input sms code: %s", err.Error())
			continue
		}
		// All good
		break
	}

	if err != nil {
		fmt.Println("[ERROR] createNewSeed() failed: %s", err.Error())
	} else {
		fmt.Println("[ERROR] createNewSeed() success")
	}
}
