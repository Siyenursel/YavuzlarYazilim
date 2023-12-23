//not : apiye gidecek versiyon x.x.x şeklinden düzeltilecek. api bağlantısı tamamlanacak /strconv.ATOİ(DEĞİŞKEN ADI)

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http isteği başarısız", err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	//fmt.Println(doc.Text())
	var content1 string
	doc.Find("meta[name=generator]").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			fmt.Printf("WordPress Sürümü: %s\n", content)
			content1 = content
		}
	})

	//versiyonu halletme
	versiyon := strings.Index(content1, "WordPress")
	//fmt.Println(versiyon)
	//fmt.Println(content1)
	trimmedVersion := content1[versiyon+len("WordPress"):]
	trimmedVersion = strings.ReplaceAll(trimmedVersion, "WordPress ", "")
	//fmt.Println("trimmed tekrar", trimmedVersion)
	trimmedVersion = strings.TrimSpace(trimmedVersion)
	//fmt.Println("boşluk silme", trimmedVersion)
	trimmedVersion = strings.ReplaceAll(trimmedVersion, ".", "")
	//fmt.Println("noktaları boşluk yap", trimmedVersion)
	lenght := len(trimmedVersion)
	if lenght < 3 {
		trimmedVersion = trimmedVersion + "0"
	}
	fmt.Println(trimmedVersion)

	//İstek kısmi api falan

	api := "https://wpscan.com/api/v3+/wordpresses/"
	api = api + trimmedVersion
	fmt.Println("api", api)

	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		fmt.Println("http isteği başarısız", err)
		return
	}
	req.Header.Set("Authorization", "Token token=rz0NCyIZlJYfEtZTaCUlsPkomaVBdv1db04IuKHmYjU")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP isteği başarısız:", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body okuma hatası:", err)
		return
	}
	fmt.Println(string(body))
}
