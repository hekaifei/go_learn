package parser

import (
	"github.com/antchfx/htmlquery"
	"go_learn/src/crawler_parallel/engine"
	"strings"

	"golang.org/x/net/html"
)

const (
	cityXpath    = `//*[@id="app"]/article[2]/dl/*/a[@href]`
	profilexpath = `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div[@class='list-item']`

	profileNickName  = `//table/tbody/tr/th`
	profileIntroduce = `/div/div`
	profileBase      = `//table/tbody/tr/td`
	profileImg       = `/div/a/img/@src`
)

func ParseCity(node *html.Node) engine.ParseResult {
	reqs := make([]engine.Request, 0)
	// 提取城市
	citys := htmlquery.Find(node, cityXpath)
	for _, cityNode := range citys {
		cityHref := htmlquery.SelectAttr(cityNode, "href")
		reqs = append(reqs, engine.Request{
			Url:        cityHref,
			ParserFunc: ParseCityList,
		})
	}
	return engine.ParseResult{Requests: reqs}
}

func ParseCityList(node *html.Node) engine.ParseResult {
	items := make([]interface{}, 0)
	n1 := htmlquery.Find(node, profilexpath)
	for _, n2 := range n1 {
		item := engine.Profile{}
		if find := htmlquery.Find(n2, profileNickName); find != nil {
			nameNode := find[0]
			if n3 := htmlquery.Find(nameNode, "//a/@href"); n3 != nil && len(n3) > 0 {
				item.UserUrl = htmlquery.SelectAttr(n3[0], "href")
				item.UserId = item.UserUrl[strings.LastIndex(item.UserUrl, "/")+1:]
			}
			item.Name = htmlquery.InnerText(nameNode)
		}
		if find := htmlquery.Find(n2, profileIntroduce); find != nil {
			item.Introduce = htmlquery.InnerText(find[0])
		}
		if find := htmlquery.Find(n2, profileImg); find != nil && len(find) > 0 {
			item.ImageUrl = htmlquery.SelectAttr(find[0], "src")
		}

		n2Nodes := htmlquery.Find(n2, profileBase)
		for index, n3 := range n2Nodes {
			n3.RemoveChild(n3.FirstChild)
			switch index {
			case 0:
				item.Gender = htmlquery.InnerText(n3)
			case 1:
				item.Location = htmlquery.InnerText(n3)
			case 2:
				item.Age = htmlquery.InnerText(n3)
			case 3:
				item.Education = htmlquery.InnerText(n3)
			case 4:
				item.Marriage = htmlquery.InnerText(n3)
			case 5:
				item.Height = htmlquery.InnerText(n3)
			default:
			}
		}

		items = append(items, item)
	}

	return engine.ParseResult{Items: items}
}
