package controllers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/revel/modules/db/app"
	"github.com/creer/revel-contact-sample/app/models"
)

var (
	Dbm *gorp.DbMap // このデータベースマッパーからSQLを流す
)

func InitDB() {

	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	//テーブル&カラム定義
	Dbm.AddTableWithName(models.Contact{}, "contact").SetKeys(true, "Id")

	//DB初期化
	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	//初期データ挿入
	contacts := []*models.Contact{
		&models.Contact{0, "kato", "mail@go8.jp", "コメント"},
		&models.Contact{0, "kato2", "mail2@go8.jp", "コメント2"},
	}
	for _, contact := range contacts {
		if err := Dbm.Insert(contact); err != nil {
			panic(err)
		}
	}

}


//transaction処理のためのGorpControllerを定義
type GorpController struct {
	*revel.Controller
	Transaction *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	transaction, err := Dbm.Begin() // ここで開始したtransactionをCOMMITする
	if err != nil {
		panic(err)
	}
	c.Transaction = transaction
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	if err := c.Transaction.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Transaction == nil {
		return nil
	}
	if err := c.Transaction.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Transaction = nil
	return nil
}
