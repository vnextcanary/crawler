// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	//resp,err := http.Get(url) //获取页面返回的response
	//if err != nil{
	//	return nil,err
	//}
	//defer resp.Body.Close() //记得每次都要关闭response的body哦
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	bufioReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(bufioReader, Determineencoding(bufioReader).NewDecoder())
	//将其他编码的reader转换为常用的utf8reader
	return ioutil.ReadAll(utf8Reader)
}

func Determineencoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//fmt.Println("get url error: %v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
