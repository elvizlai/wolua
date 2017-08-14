/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/16 22:08.
 */

package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/yuin/gopher-lua"
)

type luaPG struct {
	db *sql.DB
}

// PGLoader is the module loader function.
func PGLoader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"new": pgConn,
	})

	registerPGType(mod, L)

	L.Push(mod)
	return 1
}

const luaPGRespTypeName = "wopg"

func registerPGType(module *lua.LTable, L *lua.LState) {
	mt := L.NewTypeMetatable(luaPGRespTypeName)

	L.SetField(mt, "__index", L.SetFuncs(module, pgMethods))
}

var pgMethods = map[string]lua.LGFunction{
	"exec": pgExec,
}

func checkPGResp(l *lua.LState) *luaPG {
	ud := l.CheckUserData(1)
	if v, ok := ud.Value.(*luaPG); ok {
		return v
	}
	l.ArgError(1, "lua pg expected")
	return nil
}

func pgExec(l *lua.LState) int {
	conn := checkPGResp(l)

	cmd := l.CheckString(2)

	aff, err := conn.db.Exec(cmd)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	affNum, err := aff.RowsAffected()
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LNumber(affNum))

	return 1
}

// should put conn to x
func pgConn(l *lua.LState) int {
	dsn := l.CheckString(1)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}

	ud := l.NewUserData()
	ud.Value = &luaPG{
		db: db,
	}

	l.SetMetatable(ud, l.GetTypeMetatable(luaPGRespTypeName))

	l.Push(ud)

	return 1
}
