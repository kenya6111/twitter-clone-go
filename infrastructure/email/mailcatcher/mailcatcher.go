package mailcatcher

import (
	"fmt"
	"net/smtp"
	"strings"
)

type mailCatcherEmailService struct {
	frontEndRootUrl string
}

var (
	hostname = "mailcatcher"
	port     = 1025
	username = "user@example.com"
	password = "password"
)

func NewMainCatcherEmailService(frontEndRootUrl string) *mailCatcherEmailService {
	return &mailCatcherEmailService{
		frontEndRootUrl: frontEndRootUrl,
	}
}

func (s *mailCatcherEmailService) SendInvitationEmail(email, token string) error {

	from := "gopher@example.net"
	recipients := []string{"foo@example.com", "bar@example.com"}
	subject := "件名: 【Twitter Clone】メールアドレスの確認をお願いします"
	body := fmt.Sprintf("%s様\n\nこのたびはご登録ありがとうございます。\n\n以下のリンクをクリックして、メールアドレスの確認を完了してください。\n\n▼メールアドレス確認リンク\n\n%s\n\n※このリンクの有効期限は24時間です。\n\n※心当たりがない場合は、このメールを無視してください。\n\n引き続きよろしくお願いいたします。\n\nTwitter Clone サポートチーム",
		email,
		"http://localhost:8080/verify?token="+token)

	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body), "\n", "\r\n"))
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), nil, from, recipients, msg); err != nil {
		return err
	}
	return nil
}
