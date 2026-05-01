package pop

type Collection[T PopInfo] []T

func (c Collection[T]) Len() int {
	return len(c)
}

func (c Collection[T]) IsEmpty() bool {
	return len(c) == 0
}

func (c Collection[T]) IsId(id int) (bool, T) {
	for _, v := range c {
		if v.GetID() == id {
			return true, v
		}
	}
	var zero T
	return false, zero
}

func (c Collection[T]) GetById(id int) T {
	for _, v := range c {
		if v.GetID() == id {
			return v
		}
	}
	var zero T
	return zero
}

func (c Collection[T]) GetIDByAddr(addr string) int {
	for _, v := range c {
		if v.GetIP() == addr && v.GetID() != 0 {
			return v.GetID()
		}
	}

	return 0
}

type PopProvider interface {
	GetCollection() Collection[PopInfo]
}
