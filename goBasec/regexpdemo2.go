package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	url := "https://movie.douban.com/subject/1292052/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	sHtml, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(sHtml))
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(string(sHtml), -1)
	// fmt.Printf("%v", result)
	// for _, v := range result {
	// 	fmt.Println(v[1])
	// }
	fmt.Println(result[0][1])
	reg1 := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
	result1 := reg1.FindAllStringSubmatch(string(sHtml), -1)
	fmt.Println(result1[0][1])
}
