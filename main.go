package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/usrname/learngo/scrapper"
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

// var errRequestFailed = errors.New("Request Failed")

// type apiResult struct {
// 	url string
// 	status string
// }

// func main() {
// 	results := make(map[string] string)
// 	c := make(chan apiResult)
// 	urls := [] string {
// 		"https://www.airbnb.com/",
// 		"https://www.naver.com/",
// 		"https://www.google.com/",
// 		"https://www.reddit.com/",
// 		"https://www.facebook.com/",
// 	}
// 	for _, url := range urls {
// 		go hitURL(url, c)
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		// fmt.Println(<- c)
// 		result := <- c
// 		results[result.url] = result.status
// 	}
// 	for url, status := range results {
// 		fmt.Println(url, status)
// 	}
//  }

// func hitURL(url string, c chan<- apiResult) {
// 	// fmt.Println("Checking: ", url)
// 	status := "OK"
// 	resp, err := http.Get(url)
// 	 if err != nil || resp.StatusCode >= 400 {
// 		status = "FAILED"
// 	 }
// 	 c <- apiResult{ url: url, status: status}
// }

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

// JOB SCRAPPER

// type extractedJob struct {
// 	id		string
// 	location string
// 	title 	string
// 	summary string
// }

// var baseUrl string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

// func main() {
// 	var totalJobs []extractedJob
// 	totalPageNumber := getPageNumber()
// 	c := make(chan []extractedJob)

// 	// goroutine
// 	for i := 0; i < totalPageNumber; i++ {
// 		go getPageJobData(i, c)
// 		// totalJobs = append(totalJobs, extractedJob...)
// 	}
// 	for i := 0; i < totalPageNumber; i++ {
// 		totalJobs = append(totalJobs, <- c...)
// 	}

// 	writeJobs(totalJobs)
// 	fmt.Println("Done, extracted", len(totalJobs))
// }

// func writeJobs(jobs []extractedJob){
// 	file, err := os.Create("jobs.csv")
// 	checkError(err)

// 	w := csv.NewWriter(file)
// 	defer w.Flush()

// 	headers := []string{"ID", "Title", "Location", "Summary"}

// 	wError := w.Write(headers)
// 	checkError(wError)

// 	for _, job := range jobs {
// 		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.summary}
// 		jwErr := w.Write(jobSlice)
// 		checkError(jwErr)
// 	}
// }

// func getPageJobData(page int, mainC chan []extractedJob) {
// 	var jobs []extractedJob
// 	c := make(chan extractedJob)
// 	pageUrl := baseUrl + "&recruitPage=" + strconv.Itoa(page+1)
// 	fmt.Println((pageUrl))
// 	resp, err := http.Get(pageUrl)
// 	checkError(err)
// 	checkCode(resp)

// 	defer resp.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	checkError(err)

// 	searchCards := doc.Find(".item_recruit")

// 	searchCards.Each(func(i int, card *goquery.Selection){
// 		go extractJob(card, c)
// 	})

// 	for i := 0; i < searchCards.Length(); i++ {
// 		job := <-c
// 		jobs = append(jobs, job)
// 	}
// 	mainC <- jobs
// }

// func extractJob(card *goquery.Selection, c chan<- extractedJob) {
// 	id, _ := card.Attr("value")
// 	title := cleanString(card.Find(".area_job>h2>a").Text())
// 	location := cleanString(card.Find(".job_condition span a").Text())
// 	summary := cleanString(card.Find(".job_sector").Text())
// 	c <- extractedJob {
// 		id: id,
// 		title: title,
// 		location: location,
// 		summary: summary,
// 	}
// 	// fmt.Println(id, title, location, summary)
// }

// func getPageNumber() int {
// 	pages := 0
// 	resp, err := http.Get(baseUrl)
// 	checkError(err)
// 	checkCode(resp)

// 	defer resp.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	checkError(err)

// 	doc.Find(".pagination").Each(func (i int, s *goquery.Selection){
// 		pages = s.Find("a").Length()
// 	})
// 	return pages
// }

// func checkError(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func checkCode(res *http.Response) {
// 	if res.StatusCode != 200 {
// 		log.Fatalln("Request failed with StatusCode", res.StatusCode)
// 	}
// }

// func cleanString(str string) string {
// 	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
// }

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main(){
	// scrapper.Scrape("term")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}