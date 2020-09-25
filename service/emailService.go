package service

import (
	"bytes"
	"email-service/dao"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

var (
	r          *rand.Rand
	fromUser   = "1094124771@qq.com" // 填写自己的邮箱
	password   = "bochmoxxmzbobaef"  // 填写授权码
	host       = "smtp.qq.com"
	activeCode = ""
)

// 使用第三方库发送邮件
func SendEmail(toUser, userName string) error {
	activeCode = RandString(6)
	e := email.NewEmail()
	e.From = fromUser
	e.To = []string{toUser}
	e.Subject = "购票成功"

	t, err := template.ParseFiles("templates/email.html")
	if err != nil {
		log.Fatalf("ParseFiles error: %v", err)
		return err
	}

	body := new(bytes.Buffer)

	//作为变量传递给html模板
	t.Execute(body, struct {
		Name       string
		ActiveCode string
	}{
		Name:       userName,
		ActiveCode: activeCode,
	})

	// html形式的消息
	e.HTML = body.Bytes()
	// 从缓冲中将内容作为附件到邮件中
	//e.Attach(body, "email.html", "text/html")
	// 以路径将文件作为附件添加到邮件中
	//e.AttachFile("$GOPATH/src/email/main.go")
	// 发送邮件(如果使用QQ邮箱发送邮件的话，password不是邮箱密码而是授权码)
	err = e.Send("smtp.qq.com:587", smtp.PlainAuth("", fromUser, password, host))
	if err == nil {
		dao.InsertEmail(fromUser, toUser)
	}
	return err
}

func RandString(len int) string {
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
