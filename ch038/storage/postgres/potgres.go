package postgres

import architecture "github.com/web_archi/ch038"

type Db map[int]architecture.Person


func (pg Db) Save(n int, p architecture.Person) {
	pg[n] = p
}

func (pg Db) Retrieve(n int) architecture.Person {
	return pg[n]
}