# 快速启动

获取并运行 crawler:

1. 源代码模式

	git clone https://github.com/kdhly/crawler.git  
	cd crawler  
	go run main.go
	 
 or cd crawler
 nohup go run main.go > /crawler.out 2>&1 &

	## components:

	go get github.com/Sirupsen/logrus  
	go get github.com/go-sql-driver/mysql  
	etc  


# 结构和功能:
golang  
数据库: mysql
1. 可调整抓取线程数和文件保存线程;  
2. 退出时会在配置文件和mysql中自动保存进度;  
3. 便携式设计;   


#### 使用需知:  
1. 使用前请调整配置文件、抓取网站URL及正则过滤,并导入mysql库; 

## 联系作者
mail: 1269866868@qq.com,vnextcanary@gmail.com  
QQ群：920788836  
或访问: http://www.vnextcanary.com/?page=bbs&category=XWdNEvaL_go  

## tips:
如果您有好的项目或建议，我们可以帮助您实现它.  
 

## Give me a star
如果你喜欢或计划使用这个项目，请Give me a star，谢谢！

## Donation
如果本项目让你感觉不错，你可以通过以下链接捐款，以更好地支持本项目和团队的发展: <br /><br />
![10](/static/img/donation/alipay.jpg)   <br /><br /> <br />

![10](/static/img/donation/weixin.jpg)    <br /><br /> <br />

##### 或者https://paypal.me/vnextcanary

## Screenshots ：<br /><br />
#### mainpage 
>![11](/static/img/screenshots/mainpage.jpg)  <br /><br />