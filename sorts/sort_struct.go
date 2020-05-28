package sorts

//对自定义类型进行排序的解决方法

type Person struct {
	Name   string
	Age    int
	Height int
}

type Persons []Person

func (p Persons) Len() int      { return len(p) }
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//继承 Person 的所有属性和方法
type SortByName struct{ Persons }

//降序
func (p SortByName) Less(i, j int) bool {
	return len(p.Persons[i].Name) > len(p.Persons[j].Name)
}

type SortByAge struct{ Persons }

//降序
func (p SortByAge) Less(i, j int) bool {
	return p.Persons[i].Age > p.Persons[j].Age
}

type SortByHeight struct{ Persons }

//降序
func (p SortByHeight) Less(i, j int) bool {
	return p.Persons[i].Height > p.Persons[j].Height
}
