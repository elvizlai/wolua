--
-- Created by IntelliJ IDEA.
-- User: sdrzlyz
-- Date: 2017/4/14
-- Time: 下午6:02
-- To change this template use File | Settings | File Templates.
--

function http1()
    local resp = http.get("https://www.baidu.com")

    assert(err == nil)

    echo(resp.body)
end

function http2()
    local resp, err = http.get("https://www.baidu.com", { query = "?xyz=123" })

    assert(err == nil)

    print(resp.body)
end

function http_post()
    local resp, err = post("https://www.baidu.com", {
        body = "xyz123123"
    })
    print(resp)
    print(err)
end

function http_post_withmultipartfile()
    local data, err = multipartfile("media", "/Users/sdrzlyz/Desktop/132.jpeg")
    print(data, err)

    --    data, err = multipartfile("/Users/sdrzlyz/Desktop/123.jpeg", "456.jpeg")
    --    print(data, err)
end