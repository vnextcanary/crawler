// Copyright 2019 vnextcanary Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package parser

import (
	"crawler/configs"
	"crawler/engine"
	"regexp"
	"strings"
)

const pageRe = `<a href="(html_XXXX[^"]+)"[^>]*>([^<]+)</a>`
const picsRe = `<img src="([^"]+)" border="0" onclick="if`

func GetPageList(c []byte) engine.ParserResult {
	//configs.Logs("GetPageList",string(c))
	re := regexp.MustCompile(pageRe)
	matches := re.FindAllSubmatch(c, -1)
	result := engine.ParserResult{}
	i := 0
	for _, m := range matches {
		title := string(m[2])
		url := configs.SubPageUrlSplite + string(m[1])
		result.Items = append(result.Items, title)
		//configs.Logs("GetPageListResult: ",url,title)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(b []byte) engine.ParserResult {
				return GetPageContent(b, title, url)
			},
		})
		i++
		if i > 1 {
			//break
		}
	}

	configs.Logs("GetPageListResult: ", len(matches))
	return result

}

func GetPageContent(c []byte, t string, url string) engine.ParserResult {
	//configs.Logs("GetPageContent",string(c))
	//configs.Logs("GetPageContent",t,url)
	//return engine.NilParser(c)
	var itemCount int
	re := regexp.MustCompile(picsRe)
	matches := re.FindAllSubmatch(c, -1)
	i := 0
	urls := ""
	for _, m := range matches {
		err := configs.Db.QueryRow("select count(*) as count  from "+configs.PicsTable+" where url=?", string(m[1])).Scan(&itemCount)
		if err != nil {
			configs.LogErr(err)
		}
		if itemCount == 0 {
			_, _ = configs.Db.Exec("insert into "+configs.PicsTable+"(url) values (?)", string(m[1]))
		}
		//configs.Logs("GetContent: ",string(m[1]))
		//SavePic(string(m[1]))
		path := strings.Split(string(m[1]), "/")
		urls += path[len(path)-2] + "/" + path[len(path)-1] + "^"
		//urls+=string(m[1])+"^"
		i++
		if i > 1 {
			//break
		}
	}
	err := configs.Db.QueryRow("select count(*) as count  from "+configs.MainTable+" where url=?", url).Scan(&itemCount)
	if err != nil {
		configs.LogErr(err)
	}
	if itemCount == 0 {
		_, _ = configs.Db.Exec("insert into "+configs.MainTable+"(page,title,pics,url,err) values (?,?,?,?,?)", configs.MyConfigs.IntStartPage, t, urls, url, "")
	}
	configs.Logs("GetPageContent:", len(matches))
	return engine.NilParser(c)
}
