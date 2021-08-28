package main

import (
	"testing"
)

func TestSendMsg(t *testing.T) {
	// 支持换行符\n
	SendMsg("#123\n#456\n#789")
}
