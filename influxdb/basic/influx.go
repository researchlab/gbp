package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/researchlab/gbp/influxdb/basic/pkg/logic"
)

const (
	// TimeLayout for influxdb time format
	timeLayout string = "2006-01-02T15:04:05Z"
)

var metrics = []string{
	"cluster_cpu_prediction",
	"cluster_mem_prediction",
	"cluster_storage_prediction",
}

func main() {
	c, _ := logic.Conn()
	// 01 查询已存在的指标 正常返回
	//queryInstance(c, "cpu_used_percent")
	// 02 写入
	// 02 01 正常写入无tags 的指标 写入正常
	// 02 01 正常写入有tags 的指标， 同时写入相同时间的另外一条指标
	// 指标写入成功，但是被覆盖;
	//writeInstance(c)
	//queryInstance(c, "cpu_used_percent")
	//queryInstance(c, "mem_memused_percent")
	//queryInstance(c, "df_bytes_used_percent")
	//queryInstance(c, metrics[0])
	//queryInstance(c, metrics[1])
	//queryInstance(c, metrics[2])
	// 03 查询不存在的指标，不会报错
	//queryInstance(c, "xtmp_cpu_usage")
	//MultiWriteInstance(c)
	t := time.Now().Unix()
	querymock(c, t, "cpu_used_percent", metrics[0])
	querymock(c, t, "mem_memused_percent", metrics[1])
	querymock(c, t, "df_bytes_used_percent", metrics[2])
}
func MultiWriteInstance(c client.Client) {
	fields := map[string]interface{}{
		"idle":   201.2,
		"system": 33.3,
		"user":   "ggg",
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	t := time.Now()

	logic.Write(c, "opsultra", "tmp_cpu_info", tags, fields, t)
	tt := time.Unix(t.Unix()+60, 0)
	logic.Write(c, "opsultra", "tmp_cpu_info", tags, fields, tt)
	ttt := time.Unix(tt.Unix()+60, 0)
	logic.Write(c, "opsultra", "tmp_cpu_info", tags, fields, ttt)
}
func writeInstance(c client.Client) {
	fields := map[string]interface{}{
		"idle":   201.2,
		"system": 33.3,
		"user":   "tmp",
	}
	fields1 := map[string]interface{}{
		"idle":   101.1,
		"system": 13.3,
		"user":   "tmp",
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	t := time.Now()

	logic.Write(c, "opsultra", "tmp_cpu_info",nil, fields, time.Now())
	// 因为t 相同 所以会被fields1 覆盖
	logic.Write(c, "opsultra", "tmp_cpu_info", tags, fields, t)
	logic.Write(c, "opsultra", "tmp_cpu_info", tags, fields1, t)
}
func queryInstance(c client.Client, metric string) {
	//qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", metric, 5)
	//qs := fmt.Sprintf("SELECT * FROM  %s WHERE time > now() - 5m  tz('Asia/Shanghai')", metric)
	qs := fmt.Sprintf("SELECT value FROM  %s order by time desc limit 1", metric)
	//qs := fmt.Sprintf("SELECT * FROM  %s WHERE time > now() - 5m", metric)
	//qs := fmt.Sprintf("SELECT * FROM  %s WHERE 'time > now() - 5m  tz('Asia/Shanghai')'  LIMIT %d", metric, 5)
	res, err := logic.Query(c, "opsultra", qs)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := make(map[string]int)
	for i, v := range res[0].Series[0].Columns {
		m[v] = i
	}
	fmt.Println(m)
	// protection condition
	if res != nil && res[0].Series != nil {
		for _, row := range res[0].Series[0].Values {
			fmt.Println("row:", row)
			//fmt.Println("row[2]:", row[2])
			//fmt.Println("row[last]:", row[len(row)-1])

			val, err := row[m["value"]].(json.Number).Float64()
			if err != nil {
				fmt.Println("vdata type:", reflect.TypeOf(row[m["value"]]), " err:", err)
			}
			fmt.Println("val:", val)
			//vdata, ok := v[m["value"]].(string)
			//if !ok {
			//	fmt.Println("vdata type:", reflect.TypeOf(v[m["value"]]))
			//	continue
			//}
			//val, _ := strconv.ParseFloat(vdata, 64)
			//vlower, err := row[m["lower"]].(json.Number).Float64()
			//if err !=nil {
			//	fmt.Println("vlower type:", reflect.TypeOf(row[m["lower"]]))
			//}
			////lower, _ := strconv.ParseFloat(vlower, 64)
			////fmt.Println("vlower:", vlower, " lower:", lower)
			//vupper, err := row[m["upper"]].(json.Number).Float64()
			//if err !=nil {
			//	fmt.Println("vupper type:", reflect.TypeOf(row[m["upper"]]))
			//}
			////upper, _ := strconv.ParseFloat(vupper, 64)
			////fmt.Println("vupper:", vupper, " upper:", upper)
			//fmt.Println("vlower:",vlower, " vupper:",vupper)
			// item
			//for j, value := range row {
			//	fmt.Printf("j:%d value:%v\n", j,value)
			//}
		}
	}
}
func querymock(c client.Client, t int64, metric1, metric string) {
	qs := fmt.Sprintf("SELECT * FROM  %s where endpoint='3c95f809-b9ac-4f32-8857-2e8a209632ae' order by time desc limit 1", metric1)
	res, err := logic.Query(c, "opsultra", qs)
	if err != nil {
		fmt.Println(err)
		return
	}
	// protection condition
	if res != nil && res[0].Series != nil {
		row := res[0].Series[0].Values[0]
		fmt.Println("row[2]:", row[2])
		fmt.Println("row[last]:", row[len(row)-1])
		ff, _ := (row[len(row)-1]).(json.Number).Float64()
		mocker(c, row[2].(string), t, ff, 10, metric)
	}
}
func mocker(c client.Client, endpoint string, t int64, base float64, base1 float64, metric string) {
	// mock
	rr := func() float64 {
		rand.Seed(time.Now().UnixNano())
		return 10 + rand.Float64()*(50-10)
	}

	for i := 0; i < 12*30; i++ {
		//for i := 0; i < 2; i++ {
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  "opsultra",
			Precision: "m", // 精度, 默认ns
		})
		if err != nil {
			fmt.Println("bp:", err)
			return
		}
		for j := 0; j < 24*12; j++ {
			t = t + 5*60
			u := base + base1 + rr()
			l := base - base1 - rr()
			if l < 0 {
				l = 0
			}
			v := (u + l) / 2
			tags := map[string]string{
				"endpoint": endpoint,
			}
			fields := map[string]interface{}{
				"value": v,
				"upper": u,
				"lower": l,
			}
			tt := time.Unix(t, 0)
			pt, err := client.NewPoint(metric, tags, fields, tt)
			if err != nil {
				fmt.Println("pt:", err)
				return
			}
			bp.AddPoint(pt)

		}
		err = c.Write(bp)
		if err != nil {
			fmt.Println("c:", err)
			return
		}
	}
}
