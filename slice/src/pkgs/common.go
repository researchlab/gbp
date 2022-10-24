package pkgs 


type Info struct {
	Name string
	Num  int
}

type infos []*Info


func MakeInfo()*infos {
	var in infos 
	for idx ,v := range []string{"mike","jack","jim","tony","jac", "tom"}{
		in = append(in , &Info{
			Name:v,
			Num:idx,
		})
	}
	return &in
}

func (p *infos)Modify(idx int){
	// 这个是有问题的， 因为 len(*p) 没有变化 
	//*p = append((*p)[:idx], (*p)[idx+1:]...)

}
