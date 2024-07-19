func ByteToString(a []byte) string {
	return *(*string)(unsafe.Pointer(&a))
}
func StringToByte(a string) []byte {
	return *(*[]byte)(unsafe.Pointer(&a))
}

func main() {
	// head = {address, 10, 10}
	// body = [1,2,3,4,5,6,7,8,9,10]
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var address = (**[10]int)(unsafe.Pointer(&s))
	var len2 = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	var cap = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(address, *len2, *cap)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", *address)
	var body = **address
	for i := 0; i < len(body); i++ {
		fmt.Printf("%d ", body[i])
	}

}
func main() {
	length := 6
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	// [0 1 2 3 4 5]
	// 取slice的第5个元素：通过计算第1个元素 + 4 个元素的size 得出
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 4*unsafe.Sizeof(arr[0]))

	fmt.Println(*(*int)(end)) // 4
	fmt.Println(arr[4])       // 4

}

func main() {
	a := "Kylin Lab"
	var b []byte
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&a))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Data = strHeader.Data
	sliceHeader.Len = strHeader.Len
	sliceHeader.Cap = strHeader.Len
	fmt.Println(len(a)) //9
	fmt.Println(len(b)) //9
	fmt.Println(cap(b)) //9

}



type Person struct {
	Name string
	age  int
}

func main() {
	p := &Person{
		Name: "张三",
	}

	fmt.Println(p)
	// *Person是不能直接转换为*string的，所以这里先将*Person转为unsafe.Pointer，再将unsafe.Pointer转为*string
	pName := (*string)(unsafe.Pointer(p))
	*pName = "李四"

	fmt.Println(p)
	// 正常手段是不能操作Person.age的这里先通过uintptr(unsafe.Pointer(pName))得到Person.Name的地址
	// 通过unsafe.Sizeof(p.Name)得到Person.Name占用的字节数
	// Person.Name的地址 + Person.Name占用的字节数就得到了Person.age的地址，然后将地址转为int指针。
	pAge := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(pName)) + unsafe.Sizeof(p.Name))))

	//pAge := (*int)(unsafe.Pointer((uintptr(unsafe.Pointer(pName)) + unsafe.Offsetof(p.age))))
	// 将p的age字段修改为12
	*pAge = 12

	fmt.Println(p)
}
