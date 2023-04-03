package database

import (
	"pero-redis/interface/resp"
	"pero-redis/lib/wildcard"
	"pero-redis/resp/reply"
)

func execDel(db *DB, args [][]byte) resp.Reply {
	keys := make([]string, len(args))
	for i, v := range args {
		keys[i] = string(v)
	}
	deleted := db.Removes(keys...)
	return reply.MakeIntReply(int64(deleted))
}

func execExists(db *DB, args [][]byte) resp.Reply {
	result := int64(0)
	for _, arg := range args {
		key := string(arg)
		_, exist := db.GetEntity(key)
		if exist {
			result++
		}
	}
	return reply.MakeIntReply(result)
}

func execFlushDB(db *DB, args [][]byte) resp.Reply {
	db.Flush()
	return reply.MakeOkReply()
}

func execType(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		reply.MakeStatusReply("none")
	}
	switch entity.Data.(type) {
	case []byte:
		return reply.MakeStatusReply("string")
	}
	return &reply.UnknownErrReply{}

}

func execRename(db *DB, args [][]byte) resp.Reply {
	src := string(args[0])
	dst := string(args[1])
	entity, exist := db.GetEntity(src)
	if !exist {
		return reply.MakeErrReply("no such key")
	}
	db.PutEntity(dst, entity)
	db.Remove(src)
	return reply.MakeOkReply()
}

func execRenamenx(db *DB, args [][]byte) resp.Reply {
	src := string(args[0])
	dst := string(args[1])
	_, ok := db.GetEntity(dst)
	if ok {
		return reply.MakeIntReply(0)
	}
	entity, exist := db.GetEntity(src)
	if !exist {
		return reply.MakeErrReply("no such key")
	}
	db.PutEntity(dst, entity)
	db.Remove(src)
	return reply.MakeIntReply(1)
}

func execKeys(db *DB, args [][]byte) resp.Reply {
	pattern := wildcard.CompilePattern(string(args[0]))
	result := make([][]byte, 0)
	db.data.ForEach(func(key string, val interface{}) bool {
		if pattern.IsMatch(key) {
			result = append(result, []byte(key))
		}
		return true
	})

	return reply.MakeMultiBulkReply(result)
}

func init() {
	RegisterCommand("del", execDel, -2)
	RegisterCommand("exists", execExists, -2)
	RegisterCommand("flushdb", execFlushDB, -1)
	RegisterCommand("type", execType, 2)
	RegisterCommand("rename", execRename, 3)
	RegisterCommand("renamenx", execRenamenx, 3)
	RegisterCommand("keys", execKeys, 2)
}
