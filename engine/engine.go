// Copyright 2019 vnextcanary Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package engine

import (
	"bytes"
	"crawler/configs"
	"crawler/fetcher"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// var MAX_THREADS = configs.ThreadCount
// var picsPath =configs.PicsPath
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		var wg sync.WaitGroup
		//fetch 2 线程
		for i := 0; i < 2; i++ {
			if len(requests) == 0 {
				break
			}
			r := requests[0]
			requests = requests[1:]
			wg.Add(1) // wg 中的计数器加1
			go func() {
				FetchUrl(&r, &requests)
				defer wg.Done() // wg 中的计数器减1
			}()
		}
		wg.Add(1) // wg 中的计数器加1
		go func() {
			GoSavePic()
			defer wg.Done() // wg 中的计数器减1
		}()
		//wg.Wait() // 持续阻塞等待，直到 wg 中的计数器为0
		var ch = make(chan bool)
		go func() {
			wg.Wait()
			ch <- false
		}()

		select {
		case <-ch:
			fmt.Println("All ok !")
		case <-time.After(time.Duration(120) * time.Second):
			fmt.Println("Timed out waiting for main group")
		}
	}
}
func FetchUrl(r *Request, rs *[]Request) {
	var itemCount int
	err := configs.Db.QueryRow("select count(*) as count  from "+configs.MainTable+" where url=?", r.Url).Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
	}
	if itemCount == 0 {
		fmt.Println("fetcher url :", r.Url)
		body, err := fetcher.Fetch(r.Url)
		//time.Sleep(time.Second*2)
		if err != nil {
			fmt.Println("get content error :", r.Url, err)
		}
		parserResult := r.ParserFunc(body)
		*rs = append(*rs, parserResult.Requests...)
		//for _,item:=range parserResult.Items{
		//log.Printf("get item :  %v",item)
		//}

	} else {
		fmt.Println("skip: " + r.Url)
	}
}
func GoSavePic() {
	var itemCount int
	err := configs.Db.QueryRow("select count(*) FROM " + configs.PicsTable + " where isok=0").Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
	}
	if itemCount < 10 {
		return
	}
	rows, err := configs.Db.Query("select url FROM " + configs.PicsTable + " where isok=0")
	if err != nil {
		configs.LogErr(err)
	}
	defer rows.Close()
	var wg sync.WaitGroup
	for rows.Next() {
		var url sql.NullString
		if err = rows.Scan(&url); err != nil {
			configs.LogErr(err)
		}
		if url.Valid {
			wg.Add(1) // wg 中的计数器加1
			go func() {
				SavePic(url.String)
				defer wg.Done() // wg 中的计数器减1
			}()
		}
	}
	var ch = make(chan bool)
	go func() {
		wg.Wait()
		ch <- false
	}()

	select {
	case <-ch:
		fmt.Println("All save ok !")
		return
	case <-time.After(time.Duration(90) * time.Second):
		fmt.Println("Timed out waiting for wait group")
		//wg.Done()
		return
	}
	//wg.Wait() // 持续阻塞等待，直到 wg 中的计数器为0
}

/*wait
var wg sync.WaitGroup
    for i:=0;i<10;i++{
        wg.Add(1) // wg 中的计数器加1
        go func(){
            defer wg.Done() // wg 中的计数器减1
        }()
    }
    wg.Wait() // 持续阻塞等待，直到 wg 中的计数器为0
*/
/*控制线程数
func GoSavePic(itemCount int) {
	rows, err := configs.Db.Query("select url FROM "+configs.PicsTable+" where isok=0")
	if err != nil {
		configs.LogErr(err)
	}
	defer rows.Close()
	jobs := make(chan string, itemCount)

	var wg sync.WaitGroup
	// 设置需要多少个线程阻塞
	wg.Add(configs.ThreadCount)

	// 根据线程数控制启动多少个消费者线程
	for n := 0; n < configs.ThreadCount; n++ {
		go worker(jobs, &wg)
	}
	// 生产者
	//for i := 0; i < 10; i++ {
	//	jobs <- i
	//}

	for rows.Next() {
		var url sql.NullString
		if err = rows.Scan(&url); err != nil {
			configs.LogErr(err)
		}
		if url.Valid {
			jobs<-url.String
		}
	}
	close(jobs)
	// 等待所有线程执行完毕的阻塞方法
	wg.Wait()

}
func worker(jobs <-chan string, wg *sync.WaitGroup) {
	for job := range jobs {
		SavePic(job)
		//time.Sleep(100* time.Millisecond)
	}
	// 消费完毕则调用 Done，减少需要等待的线程
	wg.Done()
}
*/
func SavePic(url string) {
	_, _ = configs.Db.Exec("update "+configs.PicsTable+" set isok=1 where url=?", url)
	path := strings.Split(url, "/")
	fmt.Println("save pic:" + path[len(path)-2] + "/" + path[len(path)-1])
	fullPath := configs.PicsPath + path[len(path)-2]
	exist, err := configs.PathExists(fullPath)
	if err != nil {
		configs.LogErr(err)
	}
	if !exist {
		_ = os.Mkdir(fullPath, 0777)
	}
	out, err := os.Create(fullPath + "/" + path[len(path)-1])
	if err != nil {
		configs.LogErr(err)
		return
	}
	defer out.Close()
	resp, err := http.Get(url)
	if resp == nil {
		configs.LogErr(err)
		return
	}
	if err != nil {
		configs.LogErr(err)
		return
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		configs.LogErr(err)
		return
	}
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		configs.LogErr(err)
		return
	}
	//defer io.Close()
	_, err = configs.Db.Exec("update "+configs.PicsTable+" set isdownloadok=1 where url=?", url)
	if err != nil {
		configs.LogErr(err)
		return
	}
}
