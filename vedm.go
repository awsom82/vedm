package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/go-wordwrap"
)

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Caption string `xml:"description"`
	Category string `xml:"category"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Result struct {
	Channel Channel `xml:"channel"`
}

func fetchRSS(url string) []Item {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var v Result
	err = xml.Unmarshal(body, &v)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return v.Channel.Items
}

func displayRSS(items []Item) {
	for i, item := range items {
		index := i + 1
		color.Set(color.FgRed)
		fmt.Printf("%2d.", index)
		color.Unset()

		cian := color.New(color.FgCyan).SprintFunc()
		fmt.Printf(" %s [%s]\n", item.Title, cian(item.Category))
	}
}

func getArticle(url string) string {
  doc, err := goquery.NewDocument(url) 
  if err != nil {
    log.Fatal(err)
  }

  var para string = ""

  doc.Find(".b-document_news").Each(func(i int, s *goquery.Selection) {
    
    para = s.Find("p").Text()

  })

  return wordwrap.WrapString(para, 120)

}

func bound(i, lower, upper int) int {
	if i < lower {
		return lower
	} else if i > upper {
		return upper
	} else {
		return i
	}
}

func main() {
	var items []Item
	var input string
	refresh := true
	var prompt = color.New(color.FgBlue).PrintfFunc()

	for {
		if refresh {
			items = fetchRSS("http://www.vedomosti.ru/rss/news")
			displayRSS(items)
			refresh = false
		}

		prompt("Наберите номер статьи чтобы прочитать, (r) для обновления, (q) для выхода: ")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "r" {
			refresh = true
			continue
		}

		if strings.ToLower(input) == "q" {
			fmt.Println("Удачи!")
			os.Exit(0)
		}

		i, err := strconv.Atoi(input)
		if err != nil {
			color.Yellow("Попробуйте еще раз")
			continue
		}

		i = bound(i, 1, len(items))
		fmt.Printf("%-9s\n", getArticle(items[i-1].Link)) 

		fmt.Println()

	}
}
