package main

type Task struct {
	localFuncPath string //本地函数路径
}

func (task *Task) Parse(msg string) error {
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
