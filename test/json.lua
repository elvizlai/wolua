--
-- Created by IntelliJ IDEA.
-- User: sdrzlyz
-- Date: 2017/4/14
-- Time: 下午5:57
-- To change this template use File | Settings | File Templates.
--

encode = json.encode
decode = json.decode

local x = {
    name = "elvizlai",
    age = 25,
    addr = {
        addr1 = "hello",
        arrr2 = "world"
    },
    fav = { "a", "b", "c" }
}

print(encode(decode(encode(x))))
