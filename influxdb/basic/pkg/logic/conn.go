package logic

import (
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/researchlab/gbp/influxdb/basic/pkg/conf"
)

func Conn() (client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     conf.Data.Addr,
		Username: conf.Data.Usr,
		Password: conf.Data.Pwd,
		Timeout:  5 * time.Second,
	})
	return c, err
}
