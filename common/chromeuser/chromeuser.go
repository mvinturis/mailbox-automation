package chromeuser

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mvinturis/mailbox-automation/common/models"
)

func SetProfile(seed *models.Seed) (profile_path string) {

	fmt.Println("[INFO] SetProfile user %s", seed.Email)

	profile_path = "profiles/" + seed.Email

	if _, err := os.Stat(profile_path); os.IsNotExist(err) {
		fmt.Println("[INFO] Profile not found, creating...")
		cmd := exec.Command("cp", "-R", "-f", "profiles/first", profile_path+"/")
		err := cmd.Run()
		if err != nil {
			fmt.Println("[ERROR] Copy profile for user %s returned: %v", seed.Email, err)
			return
		}
	}

	return
}
