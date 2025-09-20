package service

import (
	"fmt"
)

// メール送信の共通インターフェース
type EmailService interface {
	SendInvitationEmail(email, token string) error
}

// メールテンプレートを表す
type EmailTemplate struct {
	Subject string
	Body    string
}

// 招待メールのテンプレートを返します
func GetInvitationEmailTemplate(email, token string) *EmailTemplate {
	return &EmailTemplate{
		Subject: "【Twitter Clone】メールアドレスの確認をお願いします",
		Body: fmt.Sprintf("%s様\n\nこのたびはご登録ありがとうございます。\n\n以下のリンクをクリックして、メールアドレスの確認を完了してください。\n\n▼メールアドレス確認リンク\n\n%s\n\n※このリンクの有効期限は24時間です。\n\n※心当たりがない場合は、このメールを無視してください。\n\n引き続きよろしくお願いいたします。\n\nTwitter Clone サポートチーム",
			email,
			"http://localhost:8080/verify?token="+token,
		),
	}
}
