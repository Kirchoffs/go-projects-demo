package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
    resp, err := http.Get(url)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error status code: %v", resp.StatusCode)
    }

    bodyReader := bufio.NewReader(resp.Body)
    e := DeterminEncoding(bodyReader)
    utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
    return io.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {  
    bytes, err := r.Peek(1024)  
    if err != nil {
        fmt.Printf("fetch error: %v", err)    
        return unicode.UTF8  
    }  
    e, _, _ := charset.DetermineEncoding(bytes, "")  
    return e
}

func main() {
    url := "https://www.thepaper.cn/"
    body, err := Fetch(url)

    if err != nil {
        fmt.Printf("read content failed: %v", err)
        return
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        fmt.Printf("read content failed: %v", err)
    }

    doc.Find("div.small_toplink__GmZhY a h2").Each(func(i int, s *goquery.Selection) {
        title := s.Text()
        fmt.Printf("title %d: %s\n", i, title)
    })
}