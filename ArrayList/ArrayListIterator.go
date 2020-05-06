package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

type Iterable interface {
	Iterator() Iterator //构造初始化接口
}

// 构造一个指针访问数组
type ArrayListIterator struct {
	list         *ArrayList
	currentIndex int //当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.theSize
}
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := it.list.Get(it.currentIndex) //抓取当前数据
	it.currentIndex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}
func (it *ArrayListIterator) GetIndex() int {
	return 1
}
