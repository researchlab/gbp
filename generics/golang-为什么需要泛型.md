# golang  为什么需要泛型

泛型解决的问题

- 泛型函数能自动做类型推导。 关注业务逻辑通用化，解耦数据类型;
- 在编译期间即可发现类型安全问题，而不是现在要等到运行期间;
- 可以去掉接口和反射，进而可能提高程序性能

```golang 
func IndexOf[T comparable](arr []T, value T) int {
    for i, v := range arr {
        if v == value {
            return i
        }
    }
    return -1
}
```
> comparable 是go2提供的内置设施，代表所有可以比较类型.


```golang 
func isEqual[T comparable](a,b []T) bool{
    if len(a) != len(b){
        return false
    }
    for i := range a {
        if a[i] != b[i]{
            return false
        }
    }
    return true
}
```

参考

- 01 https://www.cnblogs.com/apocelipes/p/13832224.html
