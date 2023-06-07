package view

import (
	"crawler_book/engine"
	"crawler_book/frontend/model"
	model2 "crawler_book/model"
	"os"
	"testing"
)

func TestSearchResult(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "https://book.douban.com/subject/33414479/",
		Type: "douban",
		Id:   "33414479",
		Payload: model2.Book{
			Name:      "深度学习的数学",
			Author:    "[日]涌井良幸",
			Publisher: "人民邮电出版社",
			Pages:     236,
			Price:     "69.00元",
			Score:     8.4,
			Intro:     "《深度学习的数学》基于丰富的图示和具体示例，通俗易懂地介绍了深度学习相关的数学知识。第1章介绍神经网络的概况；第2章介绍理解神经网络所需的数学基础知识；第3章介绍神经网络的最优化；第4章介绍神经网络和误差反向传播法；第5章介绍深度学习和卷积神经网络。书中使用Excel进行理论验证，帮助读者直观地体验深度学习的原理。",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
