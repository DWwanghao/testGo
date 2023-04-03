package dict

import "sync"

type SyncDict struct {
	m sync.Map
}

func (dict *SyncDict) Get(key string) (val interface{}, exists bool) {
	val, ok := dict.m.Load(key)
	return val, ok
}

func (dict *SyncDict) Len() int {
	length := 0
	dict.m.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}

func (dict *SyncDict) Put(key string, val interface{}) (result int) {
	_, exist := dict.m.Load(key)
	dict.m.Store(key, val)
	if exist {
		return 0
	}
	return 1
}

func (dict *SyncDict) PutIfAbsent(key string, val interface{}) (result int) {
	_, exist := dict.m.Load(key)
	if exist {
		return 0
	}
	dict.m.Store(key, val)
	return 1
}

func (dict *SyncDict) PutIfExist(key string, val interface{}) (result int) {
	_, exist := dict.m.Load(key)
	if exist {
		dict.m.Store(key, val)
		return 1
	}
	return 0
}

func (dict *SyncDict) Remove(key string) (result int) {
	_, exist := dict.m.Load(key)
	dict.m.Delete(key)
	if exist {
		return 1
	}
	return 0
}

func (dict *SyncDict) ForEach(consumer Consumer) {
	dict.m.Range(func(key, value interface{}) bool {
		consumer(key.(string), value)
		return true
	})
}

func (dict *SyncDict) Keys() []string {
	result := make([]string, dict.Len())
	i := 0
	dict.m.Range(func(key, value interface{}) bool {
		result[i] = key.(string)
		i++
		return true
	})
	return result
}

func (dict *SyncDict) RandomKeys(limit int) []string {
	result := make([]string, dict.Len())
	for i := 0; i < limit; i++ {
		dict.m.Range(func(key, value interface{}) bool {
			result[i] = key.(string)
			return false
		})
	}
	return result
}

func (dict *SyncDict) RandomDistinctKeys(limit int) []string {
	result := make([]string, dict.Len())
	i := 0
	dict.m.Range(func(key, value interface{}) bool {
		result[i] = key.(string)
		i++
		if i == limit {
			return false
		}
		return true
	})
	return result
}

func (dict *SyncDict) clear() {
	*dict = *MakeSyncDict()
}

func MakeSyncDict() *SyncDict {
	return &SyncDict{}
}