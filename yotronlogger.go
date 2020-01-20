/*
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 */

package goConfigurableLogger

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type ConfLog struct {
	Log_level  string `yaml:"log_level"`
	Log_writer string `yaml:"log_writer"`
	Log_file   string `yaml:"log_file"`
}

type writer struct {
	Debug io.Writer
	Warn  io.Writer
	Info  io.Writer
	Error io.Writer
}

var (
	Debug *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger

	conf      ConfLog
	logWriter writer
)

func init() {
	conf.readMainConfig()
	setWriter()
	handleLogger()
}

func setWriter() {
	logWriter = writer{ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard}
	if conf.Log_writer != "discard" {
		lw := getLogWriter()
		switch conf.Log_level {
		case "error":
			logWriter = writer{ioutil.Discard, ioutil.Discard, ioutil.Discard, lw}
		case "warn":
			logWriter = writer{ioutil.Discard, ioutil.Discard, lw, lw}
		case "info":
			logWriter = writer{ioutil.Discard, lw, lw, lw}
		case "debug":
			logWriter = writer{lw, lw, lw, lw}
		default:
			panic("Failed to set LogLevel. Value wrong:" + conf.Log_level)
		}
	}
}

func getLogFile() io.Writer {
	logFile, err := os.OpenFile(conf.Log_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic("Could not Create or Write to file: " + conf.Log_file)
	}
	return logFile
}

func getLogWriter() io.Writer {
	switch conf.Log_writer {
	case "file":
		return getLogFile()
	case "stdout":
		return os.Stdout
	case "discard":
		return ioutil.Discard
	default:
		panic("Failed to set LogWriter. Value wrong:" + conf.Log_level)
	}
}

func handleLogger() {
	Debug = log.New(logWriter.Debug,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logWriter.Error,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(logWriter.Warn,
		"WARN: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logWriter.Info,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (conf *ConfLog) readMainConfig() {
	err := yaml.Unmarshal(readYAMLFile("conf.logging.yml", conf), &conf)
	if err != nil {
		panic("Error during unmarshalling logfile")
	}
}

func readYAMLFile(path string, type_struct interface{}) []byte {
	err := yaml.Unmarshal(readFile(path), &type_struct)
	if err != nil {
		panic("Logfile malformed")
	}
	byte, err := yaml.Marshal(type_struct)
	return byte
}

func readFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Could not read file: " + path)
	}
	return file
}
