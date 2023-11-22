package Map

func Map() {
	// 初始化的方式如下
	m1 := map[string]string{
		"key01": "value01",
	}
	m2 := make(map[string]string, 4)
	println(m1, m2)

	// map 中, 元素的读取
	value, status := m1["key01"]
	println(value, status)
	// value: 读取到的元素的值
	// status: 读取的目标元素是否存在

	// map中的遍历
	m1["key02"] = "value02"
	for key, val := range m1 {
		println(key, val)
	}
	// 对于 map 中的遍历依旧是随机的, 也就是说, 遍历多次结果都会不一样

	// map 中元素的删除, 可以通过 delete 方法进行操作
	delete(m1, "key01")
	// 弊端: delete 方法不具备判断删除的值是否存在, 如果想要判断值是否存在,需要进行访问判断
}
