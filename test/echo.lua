--
-- Created by IntelliJ IDEA.
-- User: sdrzlyz
-- Date: 2017/4/14
-- Time: 下午4:59
-- To change this template use File | Settings | File Templates.
--



echo(print)
echo(3.14)
echo(true)
echo(123)
echo('abc')
echo({ 'a', 'b' })
echo({
    x = 'abc',
    y = 123
})
echo("json:", { x = 'abc', y = 123 }, "str:", "abc", "num:", 123)

pretty({
    x = 'abc',
    y = 123
})
