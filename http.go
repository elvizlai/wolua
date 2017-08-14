/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/19 18:19.
 */

package main

import (
	"bytes"
	"github.com/yuin/gopher-lua"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"encoding/base64"
)

func Post(l *lua.LState) int {
	Debugf("[post] %s", l.Get(1).String())

	Debugf("")

	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("pretty"),
		Protect: true,
	}, lua.LString("[req]"), l.Get(2))
	CheckErr(err)

	err = l.CallByParam(lua.P{
		Fn:      l.GetField(l.GetGlobal("http"), "post"),
		NRet:    2,
		Protect: true,
	}, l.Get(1), l.Get(2))
	CheckErr(err)

	r, e := l.Get(-2), l.Get(-1)

	l.Push(r)
	l.Push(e)

	return 2
}

func Get(l *lua.LState) int {
	Debugf("[get] %s", l.Get(1).String())

	Debugf("")

	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("echo"),
		Protect: true,
	}, lua.LString("[req]"), l.Get(2))
	CheckErr(err)

	Debugf("")

	err = l.CallByParam(lua.P{
		Fn:      l.GetField(l.GetGlobal("http"), "get"),
		NRet:    2,
		Protect: true,
	}, l.Get(1), l.Get(2))
	CheckErr(err)

	r, e := l.Get(-2), l.Get(-1)

	l.Push(r)
	l.Push(e)

	return 2
}

// MultiPartFile tries to build an multi-part for http request
// param1 is field name
// param2 is file path
// param3 is file name(is empty, using file path base)
func MultiPartFile(l *lua.LState) int {
	var fieldName string
	var filePath string
	var fileName string

	fieldName = l.ToString(1)
	filePath = l.ToString(2)

	if l.Get(3).Type() == lua.LTNil {
		fileName = filepath.Base(filePath)
	} else {
		fileName = l.ToString(3)
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	buff := &bytes.Buffer{}
	writer := multipart.NewWriter(buff)

	fw, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	_, err = io.Copy(fw, file)

	l.Push(lua.LString(base64.StdEncoding.EncodeToString(buff.Bytes())))
	l.Push(lua.LNil)

	return 2
}
