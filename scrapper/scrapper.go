package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id		string
	location string
	title 	string
	summary string
}

// Scrape saramin.com
func Scrape(term string) {
	var baseUrl string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	var totalJobs []extractedJob
	totalPageNumber := getPageNumber(baseUrl)
	c := make(chan []extractedJob)

	// goroutine
	for i := 0; i < totalPageNumber; i++ {
		go getPageJobData(i, baseUrl, c)
		// totalJobs = append(totalJobs, extractedJob...)
	}
	for i := 0; i < totalPageNumber; i++ {
		totalJobs = append(totalJobs, <- c...)
	}
	
	writeJobs(totalJobs)
	fmt.Println("Done, extracted", len(totalJobs))
}

func writeJobs(jobs []extractedJob){
	file, err := os.Create("jobs.csv")
	checkError(err)
	
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Summary"}

	wError := w.Write(headers)
	checkError(wError)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkError(jwErr)
	}
}

func getPageJobData(page int, baseUrl string, mainC chan []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageUrl := baseUrl + "&recruitPage=" + strconv.Itoa(page+1)
	fmt.Println((pageUrl))
	resp, err := http.Get(pageUrl)
	checkError(err)
	checkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection){
		go extractJob(card, c)
	})
	
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".area_job>h2>a").Text())
	location := CleanString(card.Find(".job_condition span a").Text())
	summary := CleanString(card.Find(".job_sector").Text())
	c <- extractedJob {
		id: id,
		title: title,
		location: location,
		summary: summary,
	}
	// fmt.Println(id, title, location, summary)
}

func getPageNumber(baseUrl string) int {
	pages := 0
	resp, err := http.Get(baseUrl)
	checkError(err)
	checkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkError(err)

	doc.Find(".pagination").Each(func (i int, s *goquery.Selection){
		pages = s.Find("a").Length()
	})
	return pages
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with StatusCode", res.StatusCode)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}