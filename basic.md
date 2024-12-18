# web

[03.1. Web 工作方式 | 第三章. Web 基础 |《Go Web 编程》| Go 技术论坛](https://learnku.com/docs/build-web-application-with-golang/031-web-working-mode/3168)

### 一些基础

***\*一、概念\****

- *HTTP*，即超文本传输协议，是 *HyperText Transfer Protocol*的缩写。浏览网页时在浏览器地址栏中输入的*URL*前面都是以*"http://"*开始的。*HTTP*定义了信息如何被格式化、如何被传输，以及在各种命令下服务器和浏览器所采取的响应。
- *HTTPS*（全称：*Hypertext Transfer Protocol over Secure Socket Layer*），是以安全为目标的*HTTP*通道，简单讲是*HTTP*的安全版。即*HTTP*下加入*SSL*层，*HTTPS*的安全基础是*SSL*，因此加密的详细内容就需要*SSL*。

它是一个*URI scheme*（抽象标识符体系），句法类同*http:*体系。用于安全的*HTTP*数据传输。*https:URL*表明它使用了*HTTP*，但*HTTPS*存在不同于*HTTP*的默认端口及一个加密*/*身份验证层（在*HTTP*与*TCP*之间）。这个系统的最初研发由网景公司进行，提供了身份验证与加密通讯方法，现在它被广泛用于万维网上安全敏感的通讯，例如交易支付方面。

- *FTP*则是*File Transfer Protocol*文件传输协议。
- *TCP*是传输协议，*HTTP*是应用协议。

 

***\*二、使用端口号不一样\****

*HTTP*：*80*

*HTTPS*：*443*

*FTP*：*21*

*TCP*：很多端口，作用都不一样。

 

***\*三、作用与描述\****

- 简单说*HTTP*是面向网页的，而*FTP*是面向文件的
- 要从*FTP*上下载东西，你需要的是支持*FTP*协议的客户端，其实*IE*就是一个，但是有更好的工具可供选择，比如*CuteFTP*或者*FlashFXP*都是不错的工具。
- *HTTP*的连接很简单*,*是无状态的。
- *HTTPS*协议是由*SSL+HTTP*协议构建的可进行加密传输、身份认证的网络协议*,*要比*HTTP*协议安全。
- *HTTP*承载在*TCP*之上。打个比喻，网络是路，*TCP*是跑在路上的车，*HTTP*是车上的人。每个网站内容不一样，就像车上的每个人有不同的故事一样。





### web工作原理

- 客户机通过 TCP/IP 协议建立到服务器的 TCP 连接
- 客户端向服务器发送 HTTP 协议请求包，请求服务器里的资源文档
- 服务器向客户机发送 HTTP 协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理 “动态内容”，并将处理得到的数据返回给客户端
- 客户机与服务器断开。由客户端解释 HTML 文档，在客户端屏幕上渲染图形结果





### DNS工作原理

在浏览器输入域名，操作系统优先检查本地hosts文件是否与其有映射关系。如果有则进行调用，完成解析。如果没有，查找DNS解析器缓存。



### HTTP协议

- HTTP 是一种让 Web 服务器与浏览器 (客户端) 通过 Internet 发送与接收数据的协议，它建立在 TCP 协议之上。

- 它是一个请求、响应协议 -- 客户端发出一个请求，服务器响应这个请求。

- HTTP 协议是无状态的，同一个客户端的这次请求和上次请求是没有对应关系，对 HTTP 服务器来说，它并不知道这两个请求是否来自同一个客户端。为了解决这个问题， Web 程序引入了 Cookie 机制来维护连接的可持续状态。

- 请求包

  - HTTP 协议定义了很多与服务器交互的请求方法，最基本的有 4 种，分别是 GET, POST, PUT, DELETE。一个 URL 地址用于描述一个网络上的资源，而 HTTP 中的 GET, POST, PUT, DELETE 就对应着对这个资源的查，增，改，删 4 个操作。我们最常见的就是 GET 和 POST 了。GET 一般用于获取 / 查询资源信息，而 POST 一般用于更新资源信息。

- 响应包

  - HTTP/1.1 协议中定义了 5 类状态码， 状态码由三位数字组成，第一个数字定义了响应的类别

    1XX 提示信息 - 表示请求已被成功接收，继续处理
    2XX 成功 - 表示请求已被成功接收，理解，接受
    3XX 重定向 - 要完成请求必须进行更进一步的处理
    4XX 客户端错误 - 请求有语法错误或请求无法实现
    5XX 服务器端错误 - 服务器未能实现合法的请求

    例如：200表示正常信息，302表示跳转
