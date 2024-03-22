# loging日志库

## 日志输出内容
日期：年月日时分秒，可自定义日期格式。  
内容：msg的值和自定义kv。  
级别：debug、info、warn、error、fatal。  
文件名、行号、函数名：可选。  
 

## 日志输出格式
json格式:{"level":"debug","msg":"read username","time":"2024-03-19T16:06:06+08:00","username":"lgyong"}。  
text格式:time="2024-03-19T16:06:44+08:00" level=debug msg="read username" username=lgyong。  

## 日志等级控制
可以设置日志输出等级,ALL输出所有日志，TRACE输出TRACE-FATAL之间的日志，DEBUG输出DEBUG-FATAL之间的日志。  
FATAL：严重错误级别，表示系统无法继续运行。           
ERROR：错误级别，用于记录错误信息。                   
WARN：警告级别，表示潜在的问题，但不影响程序的运行。  
INFO：信息级别，用于记录程序的正常运行信息。         
DEBUG：调试级别，用于详细记录调试信息。               
TRACE：追踪级别，提供比DEBUG更详细的信息。             
ALL：最低级别，用于启用所有日志记录。     


## 日志输出目标
满足io.writer接口的一个或多个目标。  