package cache

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	// TODO: ваш код
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	// TODO: ваш код
	panic("implement me")
}

func (c Cache) Set(k, v string) {
	// TODO implement me
	panic("implement me")
}

func (c Cache) Get(k string) (v string, ok bool) {
	// TODO implement me
	panic("implement me")
}
