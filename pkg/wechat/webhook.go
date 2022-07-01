package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-depot/global"
	"io/ioutil"
	"net/http"
)

type NoticeMessageRequest struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type NoticeMessageReply struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func SendHookTextMsg(text string) (rep *NoticeMessageReply, err error) {
	rep, err = sendHookMsg(&NoticeMessageRequest{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: text,
		},
	})
	return
}

func sendHookMsg(msg *NoticeMessageRequest) (*NoticeMessageReply, error) {
	data, _ := json.Marshal(msg)
	return request(
		http.MethodPost,
		fmt.Sprintf("%s?key=%s", global.WechatSetting.Work.WebHook.EndPoint, global.WechatSetting.Work.WebHook.Key),
		data,
	)
}

func request(method string, url string, data []byte) (*NoticeMessageReply, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	res := &NoticeMessageReply{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
