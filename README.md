#  [![yotron](logo-yotron.png)](http://www.yotron.de)

[YOTRON](http://www.yotron.de) is a consultancy company which is focused on DevOps, Cloudmanagement and 
Data Management with NOSQL and SQL-Databases. Visit us on [ www.yotron.de ](http://www.yotron.de)

## Menu

- [ goConfigurableLogger](#list_1)

>- [ Example Configuration file](#list_2)

>- [ Setup of the configuration file](#list_3)

>- [ Log level and formatting](#list_4)

>- [ Get the logger within your project](#list_5)

>- [ How to use the debugger](#list_6)

>>- [ own credentials](#list_7)

(this menu is created with [Markdown Menu](https://www.markdownmenu.com))

# goConfigurableLogger{#list_1}
This project is a small and simple logging tool for Go projects which is configurable via a configuration File. 
The log entries are all preformatted with the loglevel and a human readable timestamp. 

With this project you can Setup a logger which is configurable centrally with a log file in your project.
The file is located in your root folder of your go project. You are able to manage

- the goal to the logs
- the level of the log

It is possible to write into a file, give it out in the standard output or discard the logs.

## Example Configuration file{#list_2}
An example logfile is looking like that:
```yaml
---
log_level: info
log_writer: file
log_file: metrics_exporter.log
```

## Setup of the configuration file{#list_3}
The configuration file is a file in the root folder of your go-project. It is called `conf.logging.yml`.
The setup of the logs shall be done in that way.

`log_level`: The level of the logging. Possible values are `debug`, `info`, `warn`, `error`

`log_writer`: The goal of the the logging. Possible values are `file`, `stdout`, `discard`

`log_file`: If you write to a file, you have to setup the logfile. If you only want to write to a logfile in you root
folder of the project, you do not need a folder path for the file.

## Log level and formatting{#list_4}
The following log level are provided. 

- `info`
- `warn`
- `error`
- `debug`

As usual the level are hierarchally organized. 
This means for example, in with the log level `debug` all other messages will be handled as well. 

All log level inherits from the ordinary log package of Golang (https://golang.org/pkg/log). 
All string formatting functions like `Println`, `Printf` ... are supported please have a look at (https://golang.org/pkg/log/#Logger)

## Get the logger within your project{#list_5}
You can get the project with a simple:

`go get github.com/yotron/goConfigurableLogger`  

We recommend to import the project into your GOROOT folder.

## How to use the debugger{#list_6}
You can import the logger in your golang package in that way:

`import . "github.com/yotron/goConfigurableLogger"`

When you import the logging that way, you can use the logging tool in that way within your golang-package.

`Info.Println("That is my message.")`

`Warn.Println(msg)`

`Error.Println("File malformed:", path, "Error:", err)`

`Debug.Printf("Debug: %message - %error", mss, err)`

The log entry i preformatted. The result is looking like that example.
```log
ERROR: 2020/01/17 15:46:31 common.go:140: Could not convert string:  Error: strconv.ParseInt: parsing "": invalid syntax
INFO: 2020/01/17 16:01:04 prometheus.go:33: Parallelization set to: 1
```

### own credentials{#list_7}
You are searching for similar clustered Testenvironments. Don't hesitate to get in touch.

created by [Joern Kleinbub](https://github.com/joernkleinbub), YOTRON, 21.01.2020

Vist me at [LinkedIn](https://www.linkedin.com/in/j%C3%B6rn-kleinbub/) 

Or via EMail <joern.kleinbub@yotron.de>, www.yotron.de
