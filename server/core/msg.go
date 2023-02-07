package core

import (
	"errors"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/th1nk-er/fast-get/server/conf"
	"gopkg.in/yaml.v3"
)

func GetMessage(m GetModel) (code int, msg string, e error) {

	code = 0
	e = nil
	msg = ""
	setting := conf.GetSystemSetting()
	if !setting.Get.On {
		e = errors.New("get messages is disabled")
		code = 1002
		return
	}
	if setting.Get.Key != "" && m.SystemKey != setting.Get.Key {
		e = errors.New("your system key is incorrect")
		code = 1003
		return
	}

	if re, _ := regexp.Compile(`^[a-zA-Z0-9]+$`); !re.MatchString(m.MsgID) {
		e = errors.New("your message id is incorrect")
		code = 1004
		return
	}

	var filePath = "./data/" + m.MsgID + ".yaml"

	if _, err := os.Stat(filePath); err != nil {
		e = errors.New("your message not exists")
		code = 1005
		return
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		e = errors.New("can not read your message")
		code = 500
		return
	}

	var message Message
	if yaml.Unmarshal(bytes, &message) != nil {
		e = errors.New("unmarshal your message failed")
		code = 500
		return
	}
	if message.MsgKey != "" && message.MsgKey != m.MsgKey {
		e = errors.New("your message key is incorrect")
		code = 1006
		return
	}

	msg = message.Msg

	return
}

func SaveMessage(m SaveModel) (code int, msgID string, e error) {

	code = 0
	msgID = ""
	e = nil

	setting := conf.GetSystemSetting()

	if !setting.Save.On {
		code = 1010
		e = errors.New("save messages is disabled")
		return
	}

	if m.EditKey == "" {
		code = 1011
		e = errors.New("your edit key can not be null")
		return
	}

	if setting.Save.Key != "" && m.SystemKey != setting.Save.Key {
		code = 1012
		e = errors.New("your system key is incorrect")
		return
	}

	if len(m.Msg) == 0 || len(m.Msg) > setting.Message.MaxLength {
		code = 1013
		e = errors.New("your message length should be 1 - " + strconv.Itoa(setting.Message.MaxLength))
		return
	}

	path := getRandomStr(8)
	for true {
		if _, err := os.Stat("./data/" + path + ".yaml"); err != nil {
			break
		}
		path = getRandomStr(8)
	}
	msgID = path
	path = "./data/" + path + ".yaml"

	message := Message{Msg: m.Msg, MsgKey: m.MsgKey, EditKey: m.EditKey}

	bytes, err := yaml.Marshal(&message)
	if err != nil {
		msgID = ""
		code = 500
		e = errors.New("marshal your message failed")
		return
	}

	if err = os.WriteFile(path, bytes, 0666); err != nil {
		msgID = ""
		code = 500
		e = errors.New("write file failed")
		return
	}
	return
}

func EditMessage(m EditModel) (code int, e error) {

	code = 0
	e = nil

	setting := conf.GetSystemSetting()

	if !setting.Edit.On {
		code = 1020
		e = errors.New("edit messages is disabled")
		return
	}

	if setting.Edit.Key != "" && m.SystemKey != setting.Edit.Key {
		code = 1021
		e = errors.New("your system key is incorrect")
		return
	}

	if re, _ := regexp.Compile(`^[a-zA-Z0-9]+$`); !re.MatchString(m.MsgID) {
		e = errors.New("your message id is incorrect")
		code = 1022
		return
	}

	var filePath = "./data/" + m.MsgID + ".yaml"

	if _, err := os.Stat(filePath); err != nil {
		e = errors.New("your message not exists")
		code = 1023
		return
	}

	bytes, e := os.ReadFile(filePath)
	if e != nil {
		e = errors.New("can not read your message")
		code = 500
		return
	}

	var message Message
	if yaml.Unmarshal(bytes, &message) != nil {
		e = errors.New("unmarshal your message failed")
		code = 500
		return
	}

	if m.EditKey != message.EditKey {
		code = 1024
		e = errors.New("your edit key is incorrect")
		return
	}

	if len(m.Msg) == 0 || len(m.Msg) > setting.Message.MaxLength {
		code = 1025
		e = errors.New("your message length should be 1 - " + strconv.Itoa(setting.Message.MaxLength))
		return
	}

	path := "./data/" + m.MsgID + ".yaml"

	message.Msg = m.Msg

	bytes, err := yaml.Marshal(&message)
	if err != nil {
		code = 500
		e = errors.New("marshal your message failed")
		return
	}

	if err := os.WriteFile(path, bytes, 0666); err != nil {
		code = 500
		e = errors.New("write file failed")
		return
	}
	return
}

func getRandomStr(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(3)
		if num == 0 {
			s += strconv.Itoa(rand.Intn(10))
		} else if num == 1 {
			s += string('a' + rune(rand.Intn(26)))
		} else {
			s += string('A' + rune(rand.Intn(26)))
		}
	}
	return s
}
