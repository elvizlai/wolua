--
-- Copyright 2015-2017, Wothing Co., Ltd.
-- All rights reserved.
--
-- Created by elvizlai on 2017/4/18 09:53.
--

local pg = require("pg")

local conn = pg.new("postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable")

print(conn:exec("DROP DATABASE IF EXISTS butler;"))
--print(conn:exec("CREATE DATABASE butler;"))



