package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetMeT struct {
	Ok     bool `json:"ok"`
	Result struct {
		ID                      int    `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	} `json:"result"`
}

const telegramBaseUrl = "https://api.telegram.org/bot"
const telegramToken = "1386957501:AAGMfMjZe7oJ7lJuPIZGrcL4vOLUSXMynDw"

const methodeGetMe = "getMe"

func main() {
	// getMe := GetMeT{}

	// // 0 билд url
	// // 1/ отправить запрос в телеграм
	// // 2 Получить ответ
	// // 3 записать структуру

	// test := []byte(`{
	// 	"ok": true,
	// 	"result": {
	// 		"id": 1386957501,
	// 		"is_bot": true,
	// 		"first_name": "olegGoBot",
	// 		"username": "olegGoBot",
	// 		"can_join_groups": true,
	// 		"can_read_all_group_messages": false,
	// 		"supports_inline_queries": false
	// 	}
	// }`)

	// err := json.Unmarshal(test, &getMe)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Printf("%v", getMe)
	// fmt.Printf(getMe.Result.Username)

	// fmt.Println(getUrlByMethod(methodeGetMe))

	body := getBodyByUrlAndData(getUrlByMethod(methodeGetMe), []byte(""))
	// fmt.Printf("%s", body)

	getMe := GetMeT{}
	err := json.Unmarshal(body, &getMe)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%v", getMe)
}

func getUrlByMethod(methodName string) string {
	return telegramBaseUrl + telegramToken + "/" + methodName
}

func getBodyByUrlAndData(url string, data []byte) []byte {
	r := bytes.NewReader(data)
	response, err := http.Post(url, "application/json", r)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return body
}
