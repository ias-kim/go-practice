package main

type OurDB struct {
	Name string
}


type DB interface {
	GetData() string
	WriteData(data string)
	Close() error 
}

func (db *OurDB) GetData() string {
...
}

func (db *OurDB) WriteData(data string) {
...
}

func (db *OurDB) Close() error {
...
}