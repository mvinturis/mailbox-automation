package hotmail

import (
	"context"
	"fmt"
	"time"

	"github.com/mvinturis/mailbox-automation/activity"
	"github.com/mvinturis/mailbox-automation/common/models"
	"github.com/mvinturis/mailbox-automation/common/smspva"
	"github.com/mvinturis/mailbox-automation/hotmail/activities"
	"github.com/mvinturis/mailbox-automation/runner"

	"github.com/chromedp/chromedp"
)

const HOTMAIL_START_PAGE = "https://outlook.live.com"

// VERIFY_SIGNED_IN_TIMEOUT timeout in secunde pentru cat timp asteptam dupa element vizibil care ne zice daca suntem logati in hotmail
const VERIFY_SIGNED_IN_TIMEOUT = 10

type Runner struct {
	runner.Runner
	Activities []activity.Activity
}

func NewRunner(seed *models.Seed, tasksContext context.Context) runner.Runner {
	r := Runner{
		runner.Runner{
			Profile: seed, Tasks: tasksContext,
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

	switch behaviour {
	case "readMessages":
		self.ReadMessages(params)
	case "recoverAccount":
		// Do nothing more, since Login is already done
	case "logout":
		self.Logout()
	}
}

func (self *Runner) Login() bool {
	self.navigateToHotmail()
	if self.IsSignedIn() {
		return true
	}

	self.startSignIn()
	for retryIndex, retryMaxCount := 0, 3; retryIndex < retryMaxCount; {
		if self.verifyLoginMode() {
			self.mailboxLogin()
			if self.IsSignedIn() {
				return true
			}
			continue
		}
		if self.verifyMailboxUnlockMode() {
			self.executeMailboxUnlock()
			continue
		}
		if self.verifyProtectAccountMode() {
			self.executeProtectAccount()
			continue
		}
		if self.verifyMailboxRecoveryMode() {
			self.executeMailboxRecovery()
			continue
		}
		if self.verifyIdentityMode() {
			self.executeVerifyIdentity()
			continue
		}
		if self.IsSignedIn() {
			return true
		}
		retryIndex++
		chromedp.Run(self.Tasks,
			self.RandomSleep(),
		)
	}

	return false
}

// IsSignedIn verifica daca suntem autentificati in casuta de hotmail dupa elementul #searchBoxId
// asteptam 10 secunde sa vedem daca e vizibil, daca nu -> timeout
func (self *Runner) IsSignedIn() bool {

	timeout := time.After(time.Second * VERIFY_SIGNED_IN_TIMEOUT)
	result := make(chan error)

	go func() {
		err := chromedp.Run(self.Tasks,
			chromedp.WaitVisible("#searchBoxId"),
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

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('//*[@id=O365_MainLink_MePhoto"]/div/div/div/div[2]/img')[0].src`, &value),
		chromedp.Click(`//*[@id="O365_MainLink_MePhoto"]/div/div/div/div[2]/img`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.EvaluateAsDevTools(`$x('//#meControlSignoutLink')[0].innerText`, &value),
		chromedp.Click(`#meControlSignoutLink`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		// Try another logout method
		err = chromedp.Run(self.Tasks,
			chromedp.EvaluateAsDevTools(`$x('//*[contains(@class, "ms-Persona-initials")]/ancestor::button')[0].value`, &value),
			chromedp.Click(`//*[contains(@class, "ms-Persona-initials")]/ancestor::button`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.EvaluateAsDevTools(`$x('//*[.="Sign out"]/ancestor::button')[0].value`, &value),
			chromedp.Click(`//*[.="Sign out"]/ancestor::button`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[INFO] false")
			return
		}
	}

	fmt.Println("[INFO] true")
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

func (self *Runner) navigateToHotmail() error {
	fmt.Println("[INFO] navigating to hotmail...")

	return chromedp.Run(self.Tasks, chromedp.Navigate(HOTMAIL_START_PAGE))
}

func (self *Runner) startSignIn() error {
	var title string

	fmt.Println("[INFO] start sign in... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools(`$x('//a[@class="linkButtonSigninHeader"]')[0].href`, &title),
		chromedp.Click(`//a[@class="linkButtonSigninHeader"]`, chromedp.NodeVisible), self.RandomSleep(),
	)

	if err != nil {
		err = chromedp.Run(self.Tasks,
			chromedp.EvaluateAsDevTools(`$x('//a[contains(@data-m, "SIGNIN")]')[0].click()`, &title),
			// chromedp.Click(`//a[contains(@data-m, "SIGNIN")][1]`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[ERROR] %s", err.Error())
			return err
		}
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) verifyLoginMode() bool {
	var title string

	fmt.Println("[INFO] verify login mode... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools("document.getElementById('i0116').innerHTML", &title),
		chromedp.EvaluateAsDevTools("document.getElementById('i0118').innerHTML", &title),
	)
	if err != nil {
		fmt.Println("[INFO] false")
		return false
	}

	fmt.Println("[INFO] true")

	return true
}

func (self *Runner) mailboxLogin() error {

	fmt.Println("[INFO] login... ")

	err := chromedp.Run(self.Tasks,
		chromedp.WaitVisible(`#i0116`, chromedp.ByID),
		chromedp.SendKeys(`#i0116`, self.Profile.Email+"\n"), self.RandomSleep(),
		chromedp.WaitVisible(`#i0118`, chromedp.ByID),
		chromedp.SendKeys(`#i0118`, self.Profile.Password+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] Can't navigate to hotmail: ", err)
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) verifyMailboxRecoveryMode() bool {
	var title string

	fmt.Println("[INFO] verify mailbox recovery mode... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools("document.getElementById('idDiv_SAOTCC_Title').innerHTML", &title),
	)
	if err != nil {
		fmt.Println("[INFO] false")
		return false
	}

	if title == "Enter code" {
		fmt.Println("[INFO] true")
		return true
	}

	fmt.Println("[INFO] false")

	return false
}

func (self *Runner) executeMailboxRecovery() error {
	fmt.Println("[INFO] execute mailbox recovery... ")
	var new_recovery_code string

	err := chromedp.Run(self.Tasks,
		self.RandomSleep(),
		chromedp.Click(`#signInAnotherWay`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`#idA_SAOTCS_LostProofs`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`#idSubmit_SAOTCS_SendCode`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.WaitVisible(`#iEnterRCField`, chromedp.ByID),
		chromedp.SendKeys(`#iEnterRCField`, self.Profile.RecoveryCode+"\n"), self.RandomSleep(),
		chromedp.WaitVisible(`#EmailAddress`, chromedp.ByID),
		chromedp.SendKeys(`#EmailAddress`, self.Profile.LocalEmail+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	securityCode := self.getMailboxSecurityCode(self.Profile.LocalEmail)

	err = chromedp.Run(self.Tasks,
		chromedp.WaitVisible(`#iOttText`, chromedp.ByID),
		chromedp.SendKeys(`#iOttText`, securityCode+"\n"), self.RandomSleep(),
		chromedp.WaitVisible(`#iRecoveryCodeVal`, chromedp.ByID),
		chromedp.EvaluateAsDevTools("document.getElementById('iRecoveryCodeVal').innerHTML", &new_recovery_code),
		chromedp.Click(`#iRecoveryCodeAction`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.Click(`#iCompleteAction`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	fmt.Println("[INFO] New recovery code: %s\n\n", new_recovery_code)

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) getMailboxSecurityCode(mailbox string) string {

	fmt.Println("[INFO] getMailboxSecurityCode... ")

	securityCode := ""
	// for securityCode == "" {
	// 	msg, err := mailbox_api.FindHotmailSecurityCodeForEmail(mailbox)
	// 	if err != nil {
	// 		fmt.Println("[ERROR] %s", err.Error())
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// 	fmt.Println("[INFO] parsed email message:\n%+v\n\n", msg)

	// 	err = mailbox_api.RemoveOne(msg.ID.Hex())
	// 	if err != nil {
	// 		fmt.Println("[ERROR] %s", err.Error())
	// 	}

	// 	securityCode = msg.Details.(interface{}).(map[string]interface{})["code"].(string)
	// 	if securityCode == "" {
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// }

	fmt.Println("[INFO] done")

	return securityCode
}

func (self *Runner) verifyMailboxUnlockMode() bool {
	var title, description string

	fmt.Println("[INFO] verify mailbox unlock mode... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools("document.getElementById('StartHeader').innerHTML", &title),
		chromedp.EvaluateAsDevTools("document.getElementById('StartQ2').innerHTML", &description),
	)
	if err != nil {
		fmt.Println("[INFO] false")
		return false
	}

	if title == "Your account has been locked" && description == "Unlocking your account" {
		fmt.Println("[INFO] true")
		return true
	}

	fmt.Println("[INFO] false")

	return false
}

func (self *Runner) executeMailboxUnlock() error {
	fmt.Println("[INFO] execute mailbox unlock... ")

	var value string

	err := chromedp.Run(self.Tasks,
		self.RandomSleep(),
		chromedp.Click(`#StartAction`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.WaitVisible(`#idPhonePageTitle`, chromedp.ByID),
		chromedp.WaitVisible(`#wlspispHipChallengeContainer`, chromedp.ByID),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	for countryCode := range smspva.IsoCountryCodes {

		var phoneNumber, phoneID string
		for phoneNumber, phoneID = smspva.GetPhoneNumber(countryCode); phoneNumber == ""; phoneNumber, phoneID = smspva.GetPhoneNumber(countryCode) {
			fmt.Println("[INFO] retry GetPhoneNumber...")
			time.Sleep(time.Second * 10)
		}

		err = chromedp.Run(self.Tasks,
			// Erase existing phoneNumber
			chromedp.DoubleClick(`#wlspispHipChallengeContainer > div:nth-child(2) > input[type=text]:nth-child(2)`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.KeyEvent("\b", chromedp.KeyModifiers(0)), self.RandomSleep(),
			// Fill in new phoneNumber
			chromedp.SendKeys(`#wlspispHipChallengeContainer > div:nth-child(2) > input[type=text]:nth-child(2)`, phoneNumber), self.RandomSleep(),
			// Select Country code
			chromedp.EvaluateAsDevTools(`$x('//select[@aria-label="Country code"]')[0].value = "`+countryCode+`"`, &value),
			// Click Send code
			chromedp.Click(`//a[@title="Send SMS code"]`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[ERROR] %s", err.Error())
			continue
		}
		// Check if the phoneNumber was accepted
		err = chromedp.Run(self.Tasks,
			chromedp.EvaluateAsDevTools(`$x('//div[@class="alert alert-error"][contains(@style, "display: inline")]')[0].innerText`, &value),
		)
		if err == nil {
			fmt.Println("[ERROR] %s", value)
			continue
		}

		err = chromedp.Run(self.Tasks,
			chromedp.WaitVisible(`#wlspispHipSolutionContainer`, chromedp.ByID),
		)
		if err != nil {
			fmt.Println("[ERROR] %s", err.Error())
			continue
		}

		var sms string
		for sms = smspva.GetSms(countryCode, phoneID); sms == ""; sms = smspva.GetSms(countryCode, phoneID) {
			fmt.Println("[INFO] retry GetSms...")
			time.Sleep(time.Second * 10)
		}

		err = chromedp.Run(self.Tasks,
			chromedp.SendKeys(`#wlspispHipSolutionContainer > div > input`, sms), self.RandomSleep(),
			chromedp.Click(`#ProofAction`, chromedp.NodeVisible), self.RandomSleep(),
			chromedp.Click(`#FinishAction`, chromedp.NodeVisible), self.RandomSleep(),
		)
		if err != nil {
			fmt.Println("[ERROR] %s", err.Error())
			continue
		}
		// All good
		break
	}

	if err != nil {
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) verifyIdentityMode() bool {
	var title string

	fmt.Println("[INFO] verify identity mode... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools("document.getElementById('idDiv_SAOTCS_Title').innerHTML", &title),
	)
	if err != nil {
		fmt.Println("[INFO] false")
		return false
	}

	if title == "Verify your identity" {
		fmt.Println("[INFO] true")
		return true
	}

	fmt.Println("[INFO] false")

	return false
}

func (self *Runner) executeVerifyIdentity() error {
	fmt.Println("[INFO] execute verify identity... ")

	err := chromedp.Run(self.Tasks,
		self.RandomSleep(),
		chromedp.Click(`#idDiv_SAOTCS_Proofs > div > div > div`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.WaitVisible(`#idTxtBx_SAOTCS_ProofConfirmation`, chromedp.ByID),
		chromedp.SendKeys(`#idTxtBx_SAOTCS_ProofConfirmation`, self.Profile.LocalEmail), self.RandomSleep(),
		chromedp.Click(`#idSubmit_SAOTCS_SendCode`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	securityCode := self.getMailboxSecurityCode(self.Profile.LocalEmail)

	fmt.Println("[INFO] Please check DOM selector elements... ")

	err = chromedp.Run(self.Tasks,
		chromedp.WaitVisible(`#idTxtBx_SAOTCC_OTC`, chromedp.ByID),
		chromedp.SendKeys(`#idTxtBx_SAOTCC_OTC`, securityCode), self.RandomSleep(),
		chromedp.WaitVisible(`#idSubmit_SAOTCC_Continue`, chromedp.ByID),
		chromedp.Click(`#idSubmit_SAOTCC_Continue`, chromedp.NodeVisible), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) verifyProtectAccountMode() bool {
	var title string

	fmt.Println("[INFO] verify protect account mode... ")

	err := chromedp.Run(self.Tasks,
		chromedp.EvaluateAsDevTools("document.getElementById('iSelectProofTitle').innerHTML", &title),
	)
	if err != nil {
		fmt.Println("[INFO] false")
		return false
	}

	if title == "Help us protect your account" {
		fmt.Println("[INFO] true")
		return true
	}

	fmt.Println("[INFO] false")

	return false
}

func (self *Runner) executeProtectAccount() error {
	fmt.Println("[INFO] execute protect account... ")

	err := chromedp.Run(self.Tasks,
		self.RandomSleep(),
		chromedp.WaitVisible(`#iProofLbl0`, chromedp.ByID),
		chromedp.Click(`#iProofLbl0`, chromedp.NodeVisible), self.RandomSleep(),
		chromedp.WaitVisible(`#iProofEmail`, chromedp.ByID),
		chromedp.SendKeys(`#iProofEmail`, self.Profile.LocalEmail+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	securityCode := self.getMailboxSecurityCode(self.Profile.LocalEmail)

	err = chromedp.Run(self.Tasks,
		chromedp.WaitVisible(`#iOttText`, chromedp.ByID),
		chromedp.SendKeys(`#iOttText`, securityCode+"\n"), self.RandomSleep(),
	)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return err
	}

	fmt.Println("[INFO] done")

	return nil
}

func (self *Runner) InitActivities(params *models.TaskParams) {
	searchKeyword := params.Keyword

	// set weights for activities
	self.Activities = []activity.Activity{
		activities.NewMarkNotSpam(self.Tasks, 1),
		activities.NewMarkAllNotSpam(self.Tasks, 100),
		activities.NewOpenMessages(self.Tasks, 5, searchKeyword),
		activities.NewClickOffer(self.Tasks, 10),
		activities.NewMarkAsRead(self.Tasks, 100),
		activities.NewMarkAsUnread(self.Tasks, 1),
		activities.NewMoveAllNonCampaignToArchive(self.Tasks, 10000, searchKeyword),
		activities.NewCategorize(self.Tasks, 100),
		activities.NewPinMessages(self.Tasks, 100),
		activities.NewFlagMessages(self.Tasks, 100),
		activities.NewMoveToArchive(self.Tasks, 20),
	}
}

// pentru a redeclara ReadMessages() intr-o forma specifica runnerului de hotmail
//func (self *Runner) ReadMessages(params *models.TaskParams) {
//
//	fmt.Println("[INFO] hotmail read messages... ")
//
//	self.Activities = []activity.Activity{
//		activities.NewOpenMessages(self.Tasks, 10, params.Keyword),
//	}
//
//	if !self.IsSignedIn() {
//		self.Login()
//	}
//
//	activities := self.GetAvailableActivities()
//	if len(activities) <= 0 {
//		fmt.Println("[INFO] finished... no more activities")
//	}
//
//	for _, activity := range activities {
//		activity.Run()
//	}
//
//	fmt.Println("[INFO] read messages done")
//}
