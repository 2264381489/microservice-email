package controller

import (
	"bytes"
	"email-service/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestEmailHandler(t *testing.T) {
	client := http.DefaultClient
	data := map[string]string{
		"toUser":   "chenhb67@chinaunicom.cn",
		"userName": "无意",
	}
	buffer, err := json.Marshal(data)
	if err != nil {
		t.Logf("msrshal fail, err:%v\n", err)
	}
	body := bytes.NewBuffer(buffer)
	response, err := client.Post("http://127.0.0.1:9000/email", "application/json", body)
	if err != nil {
		t.Logf("post fail, err:%v\n", err)
	}
	t.Logf("response:%v\n", response)
	var res = model.Result{}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		t.Logf("unmarshal fail, err: %v\n", err)
	}
	t.Logf("res: %v\n", res)
	assert.Equal(t, 200, response.StatusCode)
}
