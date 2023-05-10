// Copyright 2019 vnextcanary Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}

type PageContent struct {
	Title string
	Pics  string
}
