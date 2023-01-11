package logic

import (
	"time"
  "github.com/influxdata/influxdb/client/v2"
)

// write batchs
func Write(c client.Client, db, point string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "m", // 精度, 默认ns
	})
	if err != nil {
		return err
	}
	pt, err := client.NewPoint(point, tags, fields, t)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	err = c.Write(bp)
	if err != nil {
		return err
	}
	return nil
}
