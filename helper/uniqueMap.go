package helper

type UniqueMap map[interface{}]struct{}

func NewUniqueMap(initSize int) UniqueMap {
	return UniqueMap(make(map[interface{}]struct{}, initSize))
}

func (u UniqueMap) MAdd(keys ...interface{}) {
	for _, key := range keys {
		u[key] = struct{}{}
	}
}

func (u UniqueMap) MustUniqueAdd(keys ...interface{}) {
	if u.ExistsOne(keys...) {
		panic("dup key")
	}
	u.MAdd(keys...)
}

func (u UniqueMap) Exist(key interface{}) bool {
	_, ok := u[key]
	return ok
}

func (u UniqueMap) ExistsOne(keys ...interface{}) bool {
	for _, key := range keys {
		if u.Exist(key) {
			return true
		}
	}
	return false
}

func (u UniqueMap) ExistsAll(keys ...interface{}) bool {
	for _, key := range keys {
		if !u.Exist(key) {
			return false
		}
	}
	return true
}
