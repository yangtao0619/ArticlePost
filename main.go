package main

import (
	_ "shanghai2/routers"
	"github.com/astaxie/beego"
	_ "shanghai2/models"
)

func main() {
	//添加函数映射
	beego.AddFuncMap("next",getNextPageIndex)
	beego.AddFuncMap("pre",getPrePageIndex)
	//测试数据库操作
	beego.Run()
}

func getPrePageIndex(pageIndex int) int {
	//如果已经是第一页就直接返回
	if pageIndex == 1{
		return pageIndex
	}
	return pageIndex - 1
}

func getNextPageIndex(pageIndex int,pageCount int) int {
	//如果已经是最后一页的话,直接返回
	if pageIndex == pageCount{
		return pageIndex
	}
	return pageIndex + 1
}
