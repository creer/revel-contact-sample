package models

type Contact struct {
	Id      int      `db:", primarykey, autoincrement"`
	Name    string   `db:",size:64"`
	Email   string   `db:",size:128"`
	Comment string   `db:",size:1024"`
}

