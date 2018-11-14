# 处理Web程序的输入与输出

## 基本要求
1. 支持静态文件服务
   
下图对data.txt中的内容进行了访问

![staticFileServer](/image/staticFileServe2.png)
   
2. 支持简单js访问

下图使用了老师博客中的index.html，使用ajax对id和content进行了请求

![staticFileServer](/image/staticFileServe.png)

3. 提交表单，并输出一个表格

通过对上图用户名和密码进行填写与提交，得到下图

![login](/image/login.png)

4. 对`/unknown`给出开发中的提示，返回码5xx

使用curl进行测试，得到下图

![500](/image/500.png)

## 提高要求

[gzip过滤器源码分析](https://blog.csdn.net/u012837895/article/details/84064709)

