// Copyright 2019 golangbbs Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func LogErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}
func Logs(args ...interface{}) {
	logrus.Info("----------------------------------")
	for _, arg := range args {
		logrus.Info(arg)
	}
	logrus.Info("----------------------------------")
}
func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
