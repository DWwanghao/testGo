package database

import (
	"pero-redis/interface/resp"
	"pero-redis/resp/reply"
)

func Ping(db *DB, args [][]byte) resp.Reply {
	return reply.MakePongReply()
}

//PING
func init() {
	RegisterCommand("ping", Ping, 1)
}
