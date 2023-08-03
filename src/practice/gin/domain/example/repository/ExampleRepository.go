package repository

var db = make(map[string]string)

func Add(key string, value string) {
	db[key] = value
}

func Remove(key string) {
	delete(db, key)
}

func Get(key string) (string, bool) {
	value, ok := db[key]
	return value, ok
}
