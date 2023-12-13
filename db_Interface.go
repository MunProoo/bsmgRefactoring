package main

type DBInterface interface {
	ConnectDB() (err error)
}
