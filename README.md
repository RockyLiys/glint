# glint
glint 是一款golang开发的web漏洞主动(被动)扫描器，是目前为止跟上主流技术的开源工具,如有一下功能:


1.xss AST语义检测

2.SQL 注入检测 （正在重构）

3.xray poc 脚本检测（这个偷懒主要参照 https://github.com/jweny/pocassist 
)

4.基于浏览器的爬虫主动扫描

5.被动扫描

6.csrf 检测

7.ssrf 检测 （正在重构）

8.jsonp ast语义检测

9.Xxe 实体注入检测 支持回显和反链平台 （正在重构）

10.CRLF 检测

11.CORS 跨域共享检测

### 目前情况
提交频繁，几乎每天都在改动，此项目目前全程一个人开发，研究者比较难以使用
除了以下推荐命令可以使用，其他的设计还得自己花费时间研究

### 粗略的使用说明
因为启动模式设计得很多，比较混乱，我个人推荐研究人员使用被动扫描,记住装上chrome

```shell
glint.exe  --passiveproxy  --cert server.pem --key server.key
```
然后访问  http://martian.proxy/authority.cer 下载证书浏览器导入就行

浏览器设置代理 (你的局域网ip 如192.168.166.8):8080 ，记住是局域网不是127,当然你在agent.go configure 函数中修改

## 待开发
一般逻辑漏洞的ai检测,极具挑战性的研究功能
OOB反链平台的重构


此项目还在开发阶段,距离发行版放出要我测试直到满意为止。

