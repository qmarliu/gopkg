package webhook

import (
	"fmt"
	"os"
	"testing"
)

func TestWebHook(t *testing.T) {
	keyword := "对冲"
	webHook := NewWebHook("b93508635ae9ae4c8be0a671baa2c526d9bd7ac570809e80f01a2b61ccc82d58")
	payLoad := &PayLoad{}

	webHook.apiUrl = ""
	err := webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("api request error should be catch!")
	}

	webHook.accessToken = ""
	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error("json unmarshal error should be catch!")
	}

	webHook.resetApiUrl()
	err = webHook.sendPayload(payLoad)
	if nil == err {
		t.Error(err)
	}

	webHook.resetApiUrl()
	webHook.accessToken = "b93508635ae9ae4c8be0a671baa2c526d9bd7ac570809e80f01a2b61ccc82d58"
	payLoad = &PayLoad{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: "test msg",
		},
	}

	// test send text message
	err = webHook.SendTextMsg(keyword+"Test text message", false, "")
	if nil != err {
		t.Error(err)
	}

	// test send link message
	err = webHook.SendLinkMsg(keyword+"A link message", "Click me to baidu search", "", "https://www.baidu.com")
	if nil != err {
		t.Error(err)
	}

	// test send markdown message
	err = webHook.SendMarkdownMsg(keyword+"A markdown message", "## This is title \n > Hello World", false, "13800138000")
	if nil != err {
		t.Error(err)
	}
	msg := fmt.Sprintf("进程ID %v <br> 用户名 %v <br> 交易所 %v <br> 检测报警幅度 %v <br>", os.Getegid(), "admin", "ftx", "0.1")
	title := "策略启动"
	err = webHook.SendMarkdownMsg("A markdown message", "# "+keyword+" "+title+" \n > "+msg, false)
	if nil != err {
		t.Error(err)
	}

	// test send action card message
	err = webHook.SendActionCardMsg(keyword+"A action card message", "This is a action card message", []string{}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles cannot be null error should be catch!")
	}

	err = webHook.SendActionCardMsg(keyword+"A action card message", "This is a action card message", []string{"Title 1"}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles length not equal error should be catch!")
	}

	err = webHook.SendActionCardMsg(keyword+"A action card message", "This is a action card message", []string{"Baidu Search"}, []string{"https://www.baidu.com"}, true, true)
	if nil != err {
		t.Error(err)
	}

	// test send link card message
	err = webHook.SendLinkCardMsg([]LinkMsg{{Title: "Hello Bob", MessageURL: "https://www.google.com", PicURL: ""}})
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	t.Log("All test had pass ..")
}

func TestWebHook2(t *testing.T) {
	keyword := "对冲"
	webHook := NewWebHook("b93508635ae9ae4c8be0a671baa2c526d9bd7ac570809e80f01a2b61ccc82d58")

	msg := fmt.Sprintf("进程ID: %v  \n  用户名: %v  \n  交易所: %v  \n  检测报警幅度: %v  \n  ", os.Getegid(), "admin", "ftx", "0.1")
	title := "策略启动"
	err := webHook.SendMarkdownMsg("A markdown message", "#### "+keyword+" "+title+" \n > "+msg, false)
	if nil != err {
		t.Error(err)
	}
}
