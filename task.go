package main

import (
	"encoding/json"
)

type Message struct {
	Intpu_files []struct {
		Url string `json:"url"`
	} `json:"intpu_files"`

	Flows []struct {
		Fn_name    string `json:"fn_name"`
		Fn_version string `json:"fn_version"`
		Fn_url     string `json:"args"`
		Args       struct {
		} `json:"args"`
	} `json:"flows"`
}

type Task struct {
	localFuncPath string //本地函数路径
	msg           Message
}

func (task *Task) Parse(msg string) error {
	var data Message
	err := json.Unmarshal([]byte(msg), &data)
	if err == nil {
		task.msg = data
	}
	return nil
}

func (task *Task) Do() error {
	if err := PrepareFunc(task); err != nil {
		return err
	}
	if err := CallFunc(task); err != nil {
		return err
	}
	if err := CallbackResult(task); err != nil {
		return err
	}

	return nil
}

// 下载func，缓存有可不用下载
func PrepareFunc(task *Task) error {
	task.localFuncPath = "/usr/func/v1.0/xxx"
	return nil
}

// 调用python
func CallFunc(task *Task) error {
	return nil
}

// 回调结果
func CallbackResult(task *Task) error {

	return nil
}
