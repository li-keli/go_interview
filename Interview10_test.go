package test

type threadSafeSet struct {
	s []interface{}
}

func (set *threadSafeSet) RLock() {

}

func (set *threadSafeSet) RUnlock() {

}

// 下面的迭代会有什么问题？
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
