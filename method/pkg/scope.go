package pkg 

// undefined Resp
// 可以看到如果是在函数内部定义的Resp 就算在函数返回形参列表中也是未定义的，不能使用, 就算是大写的定义作用域也只限于函数内部，所以大小和小写就没什么区别
func Scope()Resp{
	type Resp struct {
		Name string 
	}
	var r Resp 
	r.Name = "scope01"
	return r 
}
// undefined Resp
func Scope2()Resp{
	// undefined Resp
		var r Resp
		r.Name = "scope02"
		return r
}
