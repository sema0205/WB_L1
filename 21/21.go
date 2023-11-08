package main

import "fmt"

type Message struct {
	id   int
	name string
	data []byte
}

// target
type SaveRecord interface {
	SaveRecord(message Message)
}

// adaptees
type PostgresDB struct {
	names   []string
	storage []Message
}

func (p *PostgresDB) SaveRecord(name string, message Message) {
	p.storage = append(p.storage, message)
	p.names = append(p.names, name)
}

type RedisDB struct {
	id      []int
	storage []Message
}

func (r *RedisDB) SaveRecord(id int, message Message) {
	r.storage = append(r.storage, message)
	r.id = append(r.id, id)
}

// adapters
type PostgresDBAdapter struct {
	adaptee *PostgresDB
}

func (p *PostgresDBAdapter) SaveRecord(message Message) {
	p.adaptee.SaveRecord(message.name, message)
}

type RedisDBAdapter struct {
	adaptee *RedisDB
}

func (r *RedisDBAdapter) SaveRecord(message Message) {
	r.adaptee.SaveRecord(message.id, message)
}

func main() {

	someMessage := []Message{{1, "name1", []byte("some data 1")}, {2, "name2", []byte("some data 2")}}

	ps := &PostgresDB{}
	rs := &RedisDB{}

	var psAdapter = &PostgresDBAdapter{ps}
	var rsAdapter = &RedisDBAdapter{rs}

	SendMessage(psAdapter, someMessage[0])
	SendMessage(rsAdapter, someMessage[1])

	fmt.Println(string(ps.storage[0].data))
	fmt.Println(string(rs.storage[0].data))
}

func SendMessage(storage SaveRecord, message Message) {
	storage.SaveRecord(message)
}
