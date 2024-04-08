package sendserver

import (
	"gopkg.in/gomail.v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func DownloadRemoteFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func IsLocalPath(path string) bool {
	isNetWork := strings.Contains(path, "://")
	// 判断 Scheme 是否是网络协议（http、https等）
	return !isNetWork
}

type Smtp struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderPassword string
}

func (s *Smtp) Send(toEmail, subject, body string, attch []string) error {
	// 创建邮件对象
	m := gomail.NewMessage()
	m.SetHeader("From", s.SenderEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	if len(attch) > 0 {
		for _, val := range attch {
			if IsLocalPath(val) {
				m.Attach(val)
			} else {
				// 将字节流作为附件添加到邮件中
				fileBytes, err := DownloadRemoteFile(val)
				if err == nil {
					fileSetting := gomail.SetCopyFunc(func(writer io.Writer) error {
						_, writeError := writer.Write(fileBytes)
						return writeError
					})
					m.Attach(val, fileSetting)

				}
			}
		}
	}
	// 使用 SSL 连接发送邮件
	d := gomail.NewDialer(s.Server, s.Port, s.SenderEmail, s.SenderPassword)
	d.SSL = true
	var err error
	err = d.DialAndSend(m)
	return err
}
