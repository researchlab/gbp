package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

type NewsSpider struct {
	TargetUrl    string
	TargetSource string
	MysqlDSN     string
	InsertStmt   string
	QueryStmt    string
	Duration     int
}

type Paper struct {
	Title   string
	ImgAddr string
	Desc    string
	Content string
	Author  string
	Time    string
}

var (
	NewsSpiderList = []*NewsSpider{
		&NewsSpider{
			TargetUrl:    "https://news.jin10.com",
			TargetSource: "金十数据",
			MysqlDSN:     "root:123456@tcp(localhost:3306)/phdb?charset=utf8",
			InsertStmt:   "INSERT INTO consult_info( consult_title, pic_addr, consult_desc, consult_content, consult_source, consult_author, create_date ) VALUES( ?, ?, ?, ?, ?, ?, ? )",
			QueryStmt:    "SELECT IF(COUNT(*),'true','false') FROM consult_info WHERE consult_title = ?",
			Duration:     30,
		},
		&NewsSpider{
			TargetUrl:    "https://news.jin10.com",
			TargetSource: "金十数据",
			MysqlDSN:     "root:123456@tcp(localhost:3306)/phdb?charset=utf8",
			InsertStmt:   "INSERT INTO consult_info( consult_title, pic_addr, consult_desc, consult_content, consult_source, consult_author, create_date ) VALUES( ?, ?, ?, ?, ?, ?, ? )",
			QueryStmt:    "SELECT IF(COUNT(*),'true','false') FROM consult_info WHERE consult_title = ?",
			Duration:     5,
		},
	}
)

func main() {
	var wg sync.WaitGroup
	wg.Add(len(NewsSpiderList))
	for _, ns := range NewsSpiderList {
		go ns.Run(&wg)
	}
	wg.Wait()
}

func (ns *NewsSpider) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Duration(ns.Duration) * time.Minute)
	ns.newsSpider()
	for range ticker.C {
		ns.newsSpider()
	}
}

// 新闻爬虫
func (ns *NewsSpider) newsSpider() {
	// mysql 初始化
	db, err := sql.Open("mysql", ns.MysqlDSN)
	if nil != err {
		time.Sleep(time.Duration(1) * time.Minute)
		if db, err = sql.Open("mysql", ns.MysqlDSN); nil != err {
			panic(err.Error())
		}
	}
	defer db.Close()

	doc, err := goquery.NewDocument(ns.TargetUrl)
	if nil != err {
		return
	}

	paper := Paper{}
	// 遍历类节点
	doc.Find(".jin-newsList__item").Each(func(i int, s *goquery.Selection) {
		// 文章封面
		a_img := s.Find(".J_lazyImg").Eq(0)
		if nil == a_img {
			return
		}
		img, _ := a_img.Attr("data-original")

		// 文章链接
		a_href := s.Find("a").Eq(0)
		if nil == a_href {
			return
		}
		// 抓取详情
		article_id, _ := a_href.Attr("href")
		article_href := fmt.Sprintf("%s%s", ns.TargetUrl, article_id)
		// 获取详情的dom
		dom, err := goquery.NewDocument(article_href)
		if nil != err {
			return
		}
		// 设定来源
		//source := "金十数据"
		// 观看次数
		//		hit := dom.Find(".jin-meta p").Eq(0).Text()
		// 评论数
		//		msg := dom.Find(".jin-meta p").Eq(1).Text()

		//	paper := Paper{
		//		Title:   dom.Find(".jin-news-article_title").Text(), // 新闻标题
		//		ImgAddr: img,
		//		Desc:    dom.Find(".jin-news-article_description").Text(), // 文章描述
		//		Content: dom.Find(".jin-news-article_content").Text(),     // 文章内容
		//		Author:  dom.Find(".jin-meta p").Eq(3).Text(),
		//		Time:    dom.Find(".jin-meta p").Eq(2).Text() + " " + s.Find(".jin-newsList__time").Text(), // 发布日期
		//	}

		paper.Title = dom.Find(".jin-news-article_title").Text() // 新闻标题
		paper.ImgAddr = img
		paper.Desc = dom.Find(".jin-news-article_description").Text() // 文章描述
		paper.Content = dom.Find(".jin-news-article_content").Text()  // 文章内容
		paper.Author = dom.Find(".jin-meta p").Eq(3).Text()
		paper.Time = dom.Find(".jin-meta p").Eq(2).Text() + " " + s.Find(".jin-newsList__time").Text() // 发布日期

		// mysql 查询
		if isExist, err := ns.fetchRow(db, paper.Title); nil == err && !isExist {
			ns.insert(db, paper.Title, paper.ImgAddr, paper.Desc, paper.Content, ns.TargetSource, paper.Author, paper.Time)
		}
	})
}

//插入
func (ns *NewsSpider) insert(db *sql.DB, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(ns.InsertStmt)
	if err != nil {
		return 0, err
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

//取一行数据，
func (ns *NewsSpider) fetchRow(db *sql.DB, args ...interface{}) (isExist bool, err error) {
	stmtOut, err := db.Prepare(ns.QueryStmt)
	if err != nil {
		return
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(args...).Scan(&isExist)
	return
}
