package main

import (
	"errors"
	"fmt"
	"net/http"
)

// func main() {
// 	account := accounts.NewAccount("dean")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// }

// Dictionary
// func main() {
// 	dictionary := myDict.Dictionary {}
// 	baseWord := "hello"
// 	dictionary.Add(baseWord, "First")
// 	dictionary.Search(baseWord)
// 	dictionary.Delete(baseWord)
// 	word, err := dictionary.Search(baseWord)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(word)
// }

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = map[string] string {}
	urls := [] string {
		"https://www.airbnb.com/",
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}


func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	 resp, err := http.Get(url)
	 if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	 }
	 return nil
}