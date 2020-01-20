#  [![yotron](logo-yotron.png)](http://www.yotron.de)

[YOTRON](http://www.yotron.de) is a consultancy company which is focused on DevOps, Cloudmanagement and 
Data Management with NOSQL and SQL-Databases. Visit us on [ www.yotron.de ](http://www.yotron.de)

# goConfigurableLogger
This project is a small and simple logging tool for Go projects which is configurable via a configuration File.

With this project you can Setup a logger which is configurable centrally with a log file in your project.
The file is located in your root folder of your go project. You are able to manage

- the goal to the logs
- the level of the log

It is possible to write into a file, give it out in the standard output or discard the logs.

## Example Configuration file
An example logfile is looking like that:
```
---
log_level: info
log_writer: file
log_file: metrics_exporter.log
```

## Setup of the configuration file
The configuration file is a file in the root folder of your go-project. It is called `conf.logging.yml`.
The setup of the logs shall be done in that way.

`log_level`: The level of the logging. Possible values are `debug`, `info`, `warn`, `error`

`log_writer`: The goal of the the logging. Possible values are `file`, `stdout`, `discard`

`log_file`: If you write to a file, you have to setup the logfile. If you only want to write to a logfile in you root
folder of the project, you do not need a folder path for the file.

## How to use the debugger
You can import the logger in your package in that way:

. "### TBD ##"

Th following log level are provided. 

- `info`
- `warn`
- `error`
- `debug`

As usual the level are hierarchally organized. 
This means for example, in with the log level `debug` all other messages will be handled as well. 

All log level inherits from the ordinary log package of Golang (https://golang.org/pkg/log). 
All string formatting functions like `Println`, `Printf` ... are supported please have a look at (https://golang.org/pkg/log/#Logger)

Here are some example how to use the logger within your project:

`Info.Println("That is my message.")`

`Warn.Println(msg)`

`Error.Println("File malformed:", path, "Error:", err)`

`Debug.Printf("Debug: %message - %error", mss, err)`
