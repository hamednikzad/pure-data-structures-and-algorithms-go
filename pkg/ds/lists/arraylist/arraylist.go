package arraylist

type ArrayList struct {
	items    []interface{}
	count    int
	capacity int
}

func New(items ...interface{}) *ArrayList {
	aList := &ArrayList{}
	if len(items) > 0 {
		//Add elements
	}
	return aList
}

func (list *ArrayList) Add(value interface{}){
	if list.count == len(list.items) {
		list.checkCapacity(list.count + 1)
	}

	list.items[list.count] = value
	list.count++
}

func (list *ArrayList) checkCapacity(size int){
	if len(list.items) >= size{
		return
	}

	var newCapacity int
	if len(list.items) == 0{
		newCapacity = 4
	}else{
		newCapacity = list.count * 2
	}

	if newCapacity < size{
		newCapacity = size
	}

	list.changeCapacity(newCapacity)
}

func (list *ArrayList) changeCapacity(newCapacity int) {
	if newCapacity < list.capacity{
		return
	}

	if newCapacity == list.capacity{
		return
	}

	if newCapacity > 0{
		var newItems = new object[newCapacity];

		if (_size > 0)
		{
			Array.Copy(_items, 0, newItems, 0, _size);
		}

		_items = newItems;
	}
	else
	{
		_items = Array.Empty<object>();
	}

	_capacity = newCapacity;
}