package main

type db_login struct {
	Adress		string		`json:"address"`
	Username	string		`json:"username"`
	Password	string		`json:"password"`
	Database	string		`json:"database"`
}
