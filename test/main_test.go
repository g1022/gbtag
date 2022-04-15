package test

import (
	"github.com/g1022/gbtag"
	"log"
	"testing"
)

func TestVar(t *testing.T) {
	shortInfo := gbtag.ToString()
	defaultInfo := "BT=19700101-000000&BM=unknown&C=0000000"
	if shortInfo != defaultInfo {
		t.FailNow()
	}

	log.Println(gbtag.ToJson())
	log.Println(gbtag.ToString())
	log.Println(gbtag.Env)
}
