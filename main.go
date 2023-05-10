// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package main

import (
	"crawler/configs"
	"crawler/engine"
	"crawler/parser/1024"
	"fmt"
	"strconv"
)

func main() {
	configs.Initconfigs()
	configs.Initdb()

	fmt.Println(*configs.IntConfigPara)
	fmt.Println("DbConnStr=", configs.DbConnStr)
	fmt.Println("MainTable=", configs.MainTable)
	fmt.Println("PicsTable=", configs.PicsTable)
	fmt.Println("StartUrl=", configs.StartUrl)
	fmt.Println("SubPageUrlSplite=", configs.SubPageUrlSplite)
	fmt.Println("PicsPath=", configs.PicsPath)
	fmt.Println("ThreadCount=", configs.ThreadCount)
	fmt.Println("IntStartPage=", configs.IntStartPage)
	fmt.Println("IntEndPage=", configs.IntEndPage)

	for i := configs.IntStartPage; i < configs.IntEndPage; i++ {
		configs.WriteconfigsPage(i)
		engine.Run(engine.Request{
			Url:        configs.StartUrl + strconv.Itoa(i),
			ParserFunc: parser.GetPageList,
		})
	}
	fmt.Println("finished!")
	//tail -f /root/crawler2.out
}
