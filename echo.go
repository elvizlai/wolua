/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/14 17:28.
 */

package main

import (
	"bytes"
	"encoding/json"

	"github.com/yuin/gopher-lua"
)

func Echo(l *lua.LState) int {
	buff := bytes.Buffer{}

	for i, j := 1, l.GetTop(); i <= j; i++ {
		switch x := l.Get(i).(type) {
		case *lua.LTable:
			err := l.CallByParam(lua.P{
				Fn:      l.GetField(l.GetGlobal("json"), "encode"),
				NRet:    1,
				Protect: true,
			}, x)
			CheckErr(err)
			buff.WriteString(l.Get(-1).String())
		default:
			buff.WriteString(x.String())
		}
		buff.WriteString(" ")
	}

	Debug(buff.String())

	return 0
}

func Pretty(l *lua.LState) int {
	buff := bytes.Buffer{}

	for i, j := 1, l.GetTop(); i <= j; i++ {
		switch x := l.Get(i).(type) {
		case *lua.LTable:
			err := l.CallByParam(lua.P{
				Fn:      l.GetField(l.GetGlobal("json"), "encode"),
				NRet:    1,
				Protect: true,
			}, x)
			CheckErr(err)

			var y interface{}
			err = json.Unmarshal([]byte(l.Get(-1).String()), &y)
			CheckErr(err)

			data, err := json.MarshalIndent(y, "", "  ")
			CheckErr(err)

			buff.Write(data)
		default:
			buff.WriteString(x.String())
		}
		buff.WriteString("\n")
	}

	Debug(buff.String())

	return 0
}
