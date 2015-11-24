package controllers

import (
	"github.com/revel/revel"
	"github.com/creer/revel-contact-sample/app/models"
)

type Contact struct {
	App
}

/**
	お問い合わせ入力画面
 */
func (c Contact) Index() revel.Result {
	return c.Render()
}

/*
	お問い合わせ入力確認＆DBに保存
 */
func (c Contact) Send(Name string, Email string, Comment string) revel.Result {

	c.Validation.Required(Name).Message("名前を入力して下さい")
	c.Validation.MaxSize([]rune(Name), 10).Message("名前は10文字以内で入力して下さい")
	c.Validation.Required(Email).Message("メールアドレスを入力して下さい")
	c.Validation.Email(Email).Message("メールアドレスを正しく入力して下さい")
	c.Validation.MaxSize([]rune(Name), 100).Message("コメントは100文字以内で入力して下さい")

	//エラーあり
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Contact.Index)
	}

	//DBにデータ追加
	contact := &models.Contact{0, Name, Email, Comment}
	if err := Dbm.Insert(contact); err != nil {
		panic(err)
	}

	c.FlashParams() //入力値を一時保存
	return c.Redirect(Contact.Sended)
}

/**
	お問い合わせ完了画面
 */
func (c Contact) Sended() revel.Result {
	return c.Render()
}
