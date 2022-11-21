package test

import (
	"github.com/g1022/gbtag"
	"testing"
)

func TestVar(t *testing.T) {

	gbtag.Base64Data = "eyJ0aW1lIjoiMjAyMjExMjFUMTk1NzUwKzA4MDAiLCJnaXRfY29tbWl0IjoiYzkxNWFiOSIsImJ1aWxkX21hY2hpbmUiOiJzaW1pbi56aGVuZ0Bnbzk1MjciLCJyZW1hcmsiOiIifQ=="

	//mapData := make(map[string]interface{})
	//
	//log.Println(gbtag.ToString())
	//
	//err := json.Unmarshal([]byte(gbtag.DecodeBase64()), &mapData)
	//if err != nil {
	//	t.FailNow()
	//}
	//
	//if mapData["git_commit"] != "c915ab9" {
	//	t.FailNow()
	//}

}
