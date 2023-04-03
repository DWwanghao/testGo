package database

import (
	"pero-redis/config"
	"pero-redis/interface/resp"
	"pero-redis/lib/logger"
	"pero-redis/resp/reply"
	"strconv"
	"strings"
)

type Database struct {
	dbSet []*DB
}

func NewDataBase() *Database {
	if config.Properties.Databases == 0 {
		config.Properties.Databases = 16
	}
	database := &Database{}
	database.dbSet = make([]*DB, config.Properties.Databases)
	for i := range database.dbSet {
		db := makeDB()
		db.index = i
		database.dbSet[i] = db
	}
	return database
}

func (database *Database) Exec(client resp.Connection, args [][]byte) resp.Reply {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()
	cmdName := strings.ToLower(string(args[0]))
	if cmdName == "select" {
		if len(args) != 2 {
			return reply.MakeArgNumErrReply("select")
		}
		return execSelect(client, database, args[1:])
	}
	index := client.GetDBIndex()
	db := database.dbSet[index]
	return db.Exec(client, args)
}

func (database *Database) AfterClientClose(c resp.Connection) {

}

func (database *Database) Close() {

}

func execSelect(c resp.Connection, database *Database, args [][]byte) resp.Reply {
	dbIndex, err := strconv.Atoi(string(args[0]))
	if err != nil || dbIndex >= len(database.dbSet) {
		return reply.MakeErrReply("Err invalidate db index")
	}
	c.SelectDB(dbIndex)
	return reply.MakeOkReply()
}
