package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html"
)


type Link struct {
	Href string
	Text string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the file name: ")
	fileName, _ := reader.ReadString('\n')
	fileName = fileName[:len(fileName)-1]
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	links := make([]Link, 0)
	z := html.NewTokenizer(file)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			goto done
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			curLink := Link{}
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						curLink.Href = attr.Val
					}
				}
				tt = z.Next()
				if tt == html.TextToken {
					curLink.Text = z.Token().Data
				}
				links = append(links, curLink)
				// Move back to continue processing
				z.NextIsNotRawText()
			
			}
		}
	}
	done:
		for _, link := range links {
			fmt.Printf("Link: %s, Text: %s\n", link.Href, link.Text)
		}
	
}
	

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}