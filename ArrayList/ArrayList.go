package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newVal interface{}) error    // 设置修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
	Iterator() Iterator                         // 返回迭代器
}

// 数据结构，字符串，整数，实数
type ArrayList struct {
	dataStore [] interface{} // 数组存储
	theSize   int            // 数组大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      //初始化结构体
	list.dataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.theSize = 0
	return list
}

func (list *ArrayList) checkIsFull() {
	if list.theSize == cap(list.dataStore) { //判断内存使用
		// make 中间的参数，0，中间的参数没有开辟内存
		newDataStore := make([]interface{}, 2*list.theSize, 2*list.theSize) //开辟双倍内存
		copy(newDataStore, list.dataStore)                                  //拷贝
		list.dataStore = newDataStore                                       //赋值
		fmt.Println(list.dataStore)
		fmt.Println(newDataStore)
	}
}

/*
接口实现
*/

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(append(list.dataStore, newVal)) //叠加数据
	list.theSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index > list.theSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newVal
	return nil
}

func (list *ArrayList) Insert(index int, newVal interface{}) error {
	list.checkIsFull()                               //检测内存，如果满了自动追加
	list.dataStore = list.dataStore[:list.theSize+1] //插入数据，内存移动一位
	for i := list.theSize; i > index; i-- { //从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal
	list.theSize++
	return nil
}

func (list *ArrayList) Clear(index int, newVal interface{}) {
	list.dataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.theSize = 0
}

func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.theSize--
	return nil
}
