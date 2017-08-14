/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/14 15:00.
 */

package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cjoudrey/gluahttp"
	"github.com/yuin/gluare"
	"github.com/yuin/gopher-lua"
	ljson "layeh.com/gopher-json"
)

var callStackSize = 1024
var registrySize = 1024 * 20

var fileList = []string{}
var testList = []string{}

var wolua *lua.LState

var xxx = "--------------------------------------------------"
var sss = "**************************************************"

func init() {
	flag.BoolVar(&verbose, "v", verbose, "verbose")
	flag.Parse()

	if verbose {
		xxx += xxx
		sss += sss
	}

	if len(flag.Args()) == 0 {
		fis, err := ioutil.ReadDir("./")
		CheckErr(err)

		for _, v := range fis {
			if !v.IsDir() && strings.HasSuffix(v.Name(), "lua") {
				fileList = append(fileList, "./"+v.Name())
			}
		}
	} else {
		for _, v := range flag.Args() {
			if strings.HasSuffix(v, "lua") {
				fileList = append(fileList, v)
			} else {
				testList = append(testList, v)
			}
		}
	}

	wolua = lua.NewState(lua.Options{
		CallStackSize:       callStackSize,
		RegistrySize:        registrySize,
		IncludeGoStackTrace: true,
	})

	// lib load
	wolua.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	wolua.PreloadModule("json", ljson.Loader)
	wolua.PreloadModule("regex", gluare.Loader)
	wolua.PreloadModule("pg", PGLoader)

	var err error

	err = wolua.DoString(`json = require("json")
http = require("http")
regex = require("regex")
	`)
	if err != nil {
		panic(err)
	}

	wolua.SetGlobal("echo", wolua.NewFunction(Echo))
	wolua.SetGlobal("pretty", wolua.NewFunction(Pretty))
	wolua.SetGlobal("get", wolua.NewFunction(Get))
	wolua.SetGlobal("post", wolua.NewFunction(Post))
	wolua.SetGlobal("multipartfile", wolua.NewFunction(MultiPartFile))
}

func main() {
	var start = time.Now()
	var total = 0
	var passed = 0
	var failed = 0

	for i := range fileList {
		fs, err := wolua.LoadFile(fileList[i])
		if err != nil {
			Fatalf(err.Error())
		}

		if len(testList) == 0 {
			for i := range fs.Proto.Constants {
				if strings.HasPrefix(strings.ToLower(fs.Proto.Constants[i].String()), "test") {
					testList = append(testList, fs.Proto.Constants[i].String())
				}
			}
		} else {
			//temp := []string{}
			//for i := range testList {
			//	for j := range fs.Proto.Constants {
			//		if strings.HasPrefix(strings.ToLower(fs.Proto.Constants[j].String()), testList[i]) {
			//			temp = append(temp, fs.Proto.Constants[j].String())
			//		}
			//	}
			//}
			//testList = temp
		}

		wolua.Push(fs)

		err = wolua.PCall(0, lua.MultRet, nil)
		if err != nil {
			Fatalf(err.Error())
		}

		total += len(testList)

		for i := range testList {
			Printf(xxx)
			err = wolua.CallByParam(lua.P{
				Fn:      wolua.GetGlobal(testList[i]),
				Protect: true,
			})
			if err == nil {
				passed += 1
				Infof("%s ✔", testList[i])
			} else {
				failed += 1
				Errorf(err.Error())
				Errorf("%s ✘", testList[i])
			}
		}
	}

	if total != 0 {
		Printf(sss)
		if failed == 0 {
			Infof("all %d test passed, using: %s", passed, time.Now().Sub(start))
		} else {
			Fatalf("%d test passed, %d test failed, using: %s", passed, failed, time.Now().Sub(start))
		}
	}

}
