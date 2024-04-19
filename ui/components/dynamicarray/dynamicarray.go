package dynamicarray

type dynamicArray struct {
	size    int
	maxSize int
	dArray  []int
}

var DynArray dynamicArray

func SetupArray(arraySize int) dynamicArray {
	DynArray.maxSize = arraySize
	DynArray.size = 0
	DynArray.dArray = make([]int, arraySize)

	return DynArray
}

func (d *dynamicArray) Get(index int) int {
	return d.dArray[index]
}

func (d *dynamicArray) Set(index int, value int) {
	d.dArray[index] = value
}

func (d *dynamicArray) Pushback(value int) {
	if d.size == d.maxSize {
		d.Resize()
	}
	d.dArray[d.size] = value
	d.size++
}

func (d *dynamicArray) Popback() int {
	if d.size > 0 {
		d.size--
	}
	return d.dArray[d.size]
}

func (d *dynamicArray) Resize() {
	d.maxSize *= 2
	newArray := make([]int, d.maxSize)
	for i := range d.dArray {
		newArray = append(newArray, d.dArray[i])
	}
	d.dArray = make([]int, d.maxSize)
	d.dArray = newArray
}

func (d *dynamicArray) GetSize() int {
	return d.size
}

func (d *dynamicArray) GetCapacity() int {
	return d.maxSize
}
