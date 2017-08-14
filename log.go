/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/14 15:28.
 */

package main

import (
	"log"
	"os"
)

var std = log.New(os.Stdout, "", 0)
var verbose = false

func Printf(format string, para ...interface{}) {
	std.Printf(format, para...)
}

func Debug(str string) {
	if verbose {
		std.Println("\033[33m" + str + "\033[0m")
	}
}

func Debugf(format string, para ...interface{}) {
	if verbose {
		std.Printf("\033[33m"+format+"\033[0m", para...)
	}
}

func Infof(format string, para ...interface{}) {
	std.Printf("\033[32m"+format+"\033[0m", para...)
}

func Warnf(format string, para ...interface{}) {
	std.Printf("\033[35m"+format+"\033[0m", para...)
}

func Errorf(format string, para ...interface{}) {
	std.Printf("\033[31m"+format+"\033[0m", para...)
}

func Fatalf(format string, para ...interface{}) {
	std.Fatalf("\033[31m"+format+"\033[0m", para...)
}
