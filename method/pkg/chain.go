package pkg 

import (
	"fmt"
)
type lower struct {
	cmd string
}

func (p *lower)Query()string{
		return p.cmd
}

func New(table string)*lower{
	return  &lower{
		cmd: fmt.Sprintf("select * from %s", table),
	}
}
