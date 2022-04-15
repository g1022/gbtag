package gbtag

import (
	"encoding/json"
	"fmt"
)

// Build 编译信息--使用这种聚合方式不行 20211204
//var Build = struct {
//	Commit  string
//	RespTime    string
//	Machine string
//	Env     string
//}{
//	"0000000",
//	"19700101-000000",
//	"no",
//	"qa",
//}

var Commit = "0000000"       //git rev-parse --short HEAD
var Time = "19700101-000000" //编译时间
var Machine = "unknown"      //编译机器
var Env = "env"              //运行环境 test/qa/live
var Author = "name@domain.com"

func ToString() string {
	return fmt.Sprintf("BT=%s&BM=%s&C=%s", Time, Machine, Commit)
}

func ToJson() string {
	var info = struct {
		Commit  string `json:"commit,omitempty"`
		Time    string `json:"time,omitempty"`
		Machine string `json:"machine,omitempty"`
		Env     string `json:"env,omitempty"`
		Author  string `json:"author,omitempty"`
	}{
		Commit:  Commit,
		Time:    Time,
		Machine: Machine,
		Env:     Env,
		Author:  Author,
	}

	resByte, _ := json.Marshal(info)
	return string(resByte)
}
