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

type apiResult struct {
	url string
	status string
}

func main() {
	results := make(map[string] string)
	c := make(chan apiResult)
	urls := [] string {
		"https://www.airbnb.com/",
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		// fmt.Println(<- c)
		result := <- c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
 }

func hitURL(url string, c chan<- apiResult) {
	// fmt.Println("Checking: ", url)
	status := "OK" 
	resp, err := http.Get(url)
	 if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	 }
	 c <- apiResult{ url: url, status: status}

}

// // goroutine. 프로그램이 작동하는 동안에만 유효
// // 메인함수가 goroutine을 기다려주지 않음
// func main() {
// 	c := make(chan string)
// 	people := [4]string {"dean", "aaaaa", "kango", "sellen"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	fmt.Println("Waiting for  messages")
// 	// blocking operation
// 	for i:=0; i < len(people); i++ {
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + "is sexy"
// }