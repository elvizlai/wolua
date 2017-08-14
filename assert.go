/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/4/14 15:59.
 */

package main

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
