package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ruapi-generate-md/pkg"
	"github.com/ruapi-generate-md/pkg/db"
	"os"
	"strings"
	"sync"
)

func GeneratePageByItemID(outPath string, projectName string) {
	sucessData := []pkg.Response{
		{Name: "errCode", Type: "int", Remark: "错误码,0表示成功"},
		{Name: "errMsg", Type: "string", Remark: "错误简要信息,无错误时为空"},
		{Name: "errDlt", Type: "errDlt", Remark: "错误详细信息,无错误时为空"},
		{Name: "data", Type: "object", Remark: "对象，具体结构见下方"},
	}
	for _, v := range sucessData {
		successMap[v.Name] = v
	}
	var wg sync.WaitGroup
	dbFileName := "/showdoc_data/html/Sqlite" + "/showdoc.db.php"
	//dbFileName := "./showdoc.db.php"
	data := db.NewDataBase(dbFileName)
	data.Init()
	dir := outPath + "/" + projectName
	item, err := data.Item.TakeItem(projectName)
	if err != nil {
		panic("not found this project:" + projectName)
	}
	headerArgs, err := data.RunapiGlobalParam.TakeRunapiGlobalHeaderParam(item.ItemId)
	if err != nil {
		panic("not found this header")
	}
	globalHeader := getGlobalHeader(headerArgs.ContentJsonStr)
	catalogs, _ := data.Catalog.TakeCatalogs(item.ItemId)
	for _, catalog := range catalogs {
		newCataLog := *catalog
		fmt.Println(newCataLog.CatName.String, newCataLog.CatId.Int32)
		wg.Add(1)
		go func() {
			tempPath := dir + "/" + newCataLog.CatName.String
			if _, err := os.Stat(tempPath); os.IsNotExist(err) {
				if err := os.MkdirAll(tempPath, 0755); err != nil {
					fmt.Printf("创建目录 %s 失败：%s\n", dir, err)
					return
				}
			}
			pages, _ := data.Page.TakePages(newCataLog.CatId, item.ItemId)
			for _, page := range pages {
				//fmt.Print(i, page.PageContent.String, globalHeader, page.PageTitle.String)
				generateOnePageMarkDown(page.PageContent.String, globalHeader, page.PageTitle.String, tempPath)
			}

			wg.Done()
		}()
	}
	wg.Wait()
}

func getGlobalHeader(jsonStr string) []pkg.Header {
	//jsonStr := "[{&quot;name&quot;:&quot;operationID&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;testGordon1111&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;operationID用于全局链路追踪&quot;},{&quot;name&quot;:&quot;token&quot;,&quot;type&quot;:&quot;string&quot;,&quot;value&quot;:&quot;eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJvcGVuSU1BZG1pbiIsIlBsYXRmb3JtIjoiIiwiZXhwIjoxNjg2NzEyMzkwLCJuYmYiOjE2Nzg5MzYwOTAsImlhdCI6MTY3ODkzNjM5MH0.Pr4NqWCVF54oSzFePAU0gFWDHrBGzyEnb_K1BymnC88&quot;,&quot;require&quot;:&quot;1&quot;,&quot;remark&quot;:&quot;管理员或者用户token&quot;}]"
	jsonStr = strings.ReplaceAll(jsonStr, "&quot;", "\"")
	data := []pkg.Header{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
	return data

}
