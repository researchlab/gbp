package main

type DomainInfo struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	CustId         int    `json:"custId"`
	SrcIp          string `json:"srcIp"`
	LogFmt         int    `json:"logFmt"`
	LogInterval    int    `json:"logInterval"`
	LogWild        int    `json:"logWild"`
	Type           int    `json:"type"`
	HType          int    `json:"hType"`
	LogLevel       int    `json:"logLevel"`
	BitRate        int    `json:"bitRate"`
	CostWithParent int    `json:"costWithParent"`
}

func (p *DomainInfo) Update(data DomainInfo) {
	p.Name = data.Name
	p.SrcIp = data.SrcIp
}
