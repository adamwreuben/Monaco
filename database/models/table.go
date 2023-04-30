package models

type TableInterface interface {
	Add(obj interface{}) string
	GetRow(key string) any
	Query(q any)
	GetAll() []any
	GetAllKeys(allKeys []string) []any
	Update(key string)
	Delete(key string)
}
