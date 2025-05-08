package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
)

func main() {
	res,err := http.Get("https://zenn.dev/books/explore")

	if err != nil {
		fmt.Print("url scarapping failed")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("ステータスコード異常:", res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	// ここでタイトルを取得
	title := doc.Find("title").Text()
	fmt.Println(title)

	doc.Find("article").Each(func(i int, s *goquery.Selection){
		author := s.Find(".BookLargeLink_userName__DCG6w").Text()
		book := s.Find(".BookLargeLink_title__7ALlB").Text()

		// if "" == author{
		// 	author = s.Find(".BookLInk_userName__avtjq").Text()
		// }
		// if "" == book{
		// 	book = s.Find(".BookLink_title__b8hGg").Text()
		// }
		fmt.Printf("%d 著者: %v / タイトル: %v\n", i, author, book)
	})

}