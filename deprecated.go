package gbtag

import "encoding/json"

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

var Commit = "0000000"         //Deprecated: 20221121195828git rev-parse --short HEAD
var Time = "19700101-000000"   //Deprecated: 20221121195828编译时间
var Machine = "unknown"        //Deprecated: 20221121195828编译机器
var Env = "env"                //Deprecated: 20221121195828运行环境 test/qa/live
var Author = "name@domain.com" //Deprecated: 20221121195828

//Deprecated: 20221121195828
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
