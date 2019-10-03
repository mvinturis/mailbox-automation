package chromeuser

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mvinturis/mailbox-automation/common/models"
)

func SetProfile(seed *models.Seed) (profile_path string, err error) {
	profile_path = "profiles/" + seed.Email

	if _, checkErr := os.Stat(profile_path); os.IsNotExist(checkErr) {
		fmt.Println("[DEBUG] profile %s not found, creating new", seed.Email)
		cmd := exec.Command("cp", "-R", "-f", "profiles/first", profile_path+"/")
		err = cmd.Run()
		if err != nil {
			fmt.Println("[ERROR] Copy profile for user %s returned: %v", seed.Email, err)
			return
		}
	} else {
		fmt.Println("[DEBUG] reading profile for %s from %s", seed.Email, profile_path)
	}

	return
}
