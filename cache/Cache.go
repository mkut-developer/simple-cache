package cache

type Cache interface {
	Set(key string, item interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Print()
}
