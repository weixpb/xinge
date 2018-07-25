package xinge

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	accessId       int64    = 2100294174
	secretKey      string   = "c60de9af49e9163dd0245d5f08cf0a97"
	accessIdIOS    int64    = 2200294252
	secretKeyIOS   string   = "32323692558b98234691bf5e01e20878"
	account        string   = "9"
	account_list   []string = []string{"I2NF56YKW41N", "I2NF56YKW41N", "I2NF56YKW41N"}
	clientT        *Client  = NewClient(accessId, secretKey)
	messageAndroid          = EasyMessageAndroid("test xinge push msg", "test xinge push Android API")
	messageIOS              = EasyMessageIOS("title", "test xinge push iOS API", IOSENV_DEV)
)

func TestPushSingleAccountIOS(t *testing.T) {
	msg := messageIOS
	custom := map[string]interface{}{}
	// custom["business"] = 1
	// custom["time"] = "2017-05-30 10:19:00"
	msg.SetCustom(custom)
	resultString(t, "PushAccountIOS", PushAccountIOS(accessIdIOS, secretKeyIOS, "title11", "测试信鸽推送diOS API", account, 2))
	// resultString(t, "PushTokenIOS", PushTokenIOS(accessId, secretKey, "测试信鸽推送duo iOS API", deviceToken, 2))
	// resultString(t, "PushAllIOS", PushAllIOS(accessId, secretKey, "测试推送 iOS API", 2))
	// resultString(t, "QueryPushStatus", clientT.QueryPushStatus([]string{"3317610443"}))
}

func resultString(t *testing.T, funcName string, resp XgResponse) {
	byt, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("测试 %s 发生错误：%v\n", funcName, err)
	} else {
		fmt.Printf("测试 %s 返回值为：%s\n", funcName, string(byt))
	}
}
