package main

import (
    "fmt"
    "github.com/gocolly/colly"
)
type Article struct {
	Title       string
	Author 		string
	Date     	string
	Text       	string
}
func main() {
	c := colly.NewCollector()
	article := Article{"", "", "", ""}
	//Delay: 1 * time.Second

		// GET TITLE
        c.OnHTML("h1", func(e *colly.HTMLElement) {
            article.Title = e.Text
            fmt.Println(article.Title)
        })

        // GET AUTHOR
        c.OnHTML("span[class=byline__name]", func(e *colly.HTMLElement) {
			article.Author = e.Text
            fmt.Println(article.Author)
        })

        // GET TEXT
        c.OnHTML("div[class=story-body__inner]>p", func(e *colly.HTMLElement) {
            article.Text = e.Text;
            //article.Text = article.Text + "\n";
            // fmt.Println("<p>"+article.Text+"</p>")
        })

        c.Visit("https://www.bbc.com/news/business-47852589")
    }
