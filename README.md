# glint
glint 是一款golang开发的web漏洞主动(被动)扫描器，如有一下功能:

1.xss AST语义检测 

2.SQL 注入检测 

3.xray poc 脚本检测（这个偷懒主要参照 https://github.com/jweny/pocassist 
)

4.基于浏览器的爬虫,获取更多的Url.

5.csrf 检测

6.ssrf 检测

7.jsonp ast语义检测

本项目以实战为主，有很多改进的处理:
1.发包手段和主动爬虫基于基于浏览器chromedp开发，有许多意外惊喜（指js）
2.payload 注重过waf处理，许多构造手段在网上搜罗。


此项目还在开发阶段,距离发行版放出要我测试直到满意为止。

