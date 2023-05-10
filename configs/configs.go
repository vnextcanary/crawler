// Copyright 2019 vnextcanary Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var IntConfigPara = flag.Int("c", 0, "IntConfigPara")

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

type MyConfig struct {
	DbConnStr        string // mysql Connect String
	MainTable        string //  mysql main table
	PicsTable        string // mysql pics table
	StartUrl         string
	SubPageUrlSplite string
	PageRe           string
	PicsRe           string
	PicsPath         string // save Pics Path
	ThreadCount      int    // ThreadCount
	IntStartPage     int    // start page
	IntEndPage       int    // end page
}

var (
	MyConfigs          *MyConfig
	StrConfigsFileName string
	DbConnStr          string
	MainTable          string
	PicsTable          string
	StartUrl           string
	SubPageUrlSplite   string
	PageRe             string
	PicsRe             string
	PicsPath           string
	ThreadCount        int
	IntStartPage       int
	IntEndPage         int
)

func Initconfigs() {

	JsonParse := NewJsonStruct()
	flag.Parse()
	switch *IntConfigPara {
	case 0:
		StrConfigsFileName = "crawler_config.json"
	case 9:
		StrConfigsFileName = "/root/crawler_config_test.json"
	default:
		StrConfigsFileName = "crawler_config.json"
	}
	JsonParse.Load(StrConfigsFileName, &MyConfigs)
	DbConnStr = MyConfigs.DbConnStr
	MainTable = MyConfigs.MainTable
	PicsTable = MyConfigs.PicsTable
	StartUrl = MyConfigs.StartUrl
	SubPageUrlSplite = MyConfigs.SubPageUrlSplite
	PageRe = MyConfigs.PageRe
	PicsRe = MyConfigs.PicsRe
	PicsPath = MyConfigs.PicsPath
	ThreadCount = MyConfigs.ThreadCount
	IntStartPage = MyConfigs.IntStartPage
	IntEndPage = MyConfigs.IntEndPage
}
func WriteconfigsPage(i int) {
	MyConfigs.IntStartPage = i
	byteValue, _ := json.Marshal(MyConfigs)
	_ = ioutil.WriteFile(StrConfigsFileName, byteValue, 0777)
}
