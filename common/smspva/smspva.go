package smspva

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ACCOUNTUSER = "testuser"
const APIURL = "http://smspva.com/priemnik.php"
const APIKEY = "apikey=testapykey"
const SERVICEIDHOTMAIL = "service=opt15"
const COUNTRYID = "country="

const APIMETHODGETNUMBER = "metod=get_number" // {"response":"1","number":"9871234567","id":25623}
const APIMETHODGETSMS = "metod=get_sms"       // {"response":"1","number":"8727073721","id":24777001,"text":"Microsoft access code: 9782","extra":"0","karma":61.275000000000055,"pass":"","sms":"9782","balanceOnPhone":0}

type GetNumberResponse struct {
	Response    string `json:"response"`
	PhoneNumber string `json:"number"`
	PhoneID     int    `json:"id"`
}

type GetSmsResponse struct {
	Response    string  `json:"response"`
	PhoneNumber string  `json:"number"`
	PhoneID     int     `json:"id"`
	Text        string  `json:"text"`
	Extra       string  `json:"extra"`
	Karma       float32 `json:"karma"`
	Pass        string  `json:"pass"`
	Sms         string  `json:"sms"`
}

func GetPhoneNumber(country string) (phoneNumber, phoneID string) {
	url := APIURL + "?" + APIKEY +
		"&" + APIMETHODGETNUMBER + "&" + COUNTRYID + country + "&" + SERVICEIDHOTMAIL

	fmt.Println("[INFO] GetPhoneNumber... %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}

	defer resp.Body.Close()
	responseText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}
	fmt.Println("[INFO] %s", string(responseText))

	// parse JSON response
	var response GetNumberResponse
	json.Unmarshal([]byte(responseText), &response)

	//return results
	phoneNumber = response.PhoneNumber
	phoneID = strconv.Itoa(response.PhoneID)

	fmt.Println("[INFO] phone number %s, phone id %s", phoneNumber, phoneID)

	return
}

func GetSms(country, phoneID string) (sms string) {
	url := APIURL + "?" + APIKEY +
		"&" + APIMETHODGETSMS + "&" + COUNTRYID + country + "&" + SERVICEIDHOTMAIL +
		"&id=" + phoneID

	fmt.Println("[INFO] GetSms... %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}

	defer resp.Body.Close()
	responseText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] %s", err.Error())
		return
	}
	fmt.Println("[INFO] %s", string(responseText))

	// parse JSON response
	var response GetSmsResponse
	json.Unmarshal([]byte(responseText), &response)

	//return results
	sms = response.Sms

	fmt.Println("[INFO] sms %s", sms)

	return
}
