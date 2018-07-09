package controller

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	//"github.com/gin-gonic/gin"
	"os"
	"io"
)

//创建一条数据记录
func CreateItem(c *gin.Context) {
	file,header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	//if err := c.SaveUploadedFile(file, file.Filename); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusOK,
	//		"message": errors.New("文件讀取失敗"),
	//		"data":    []int{},
	//	})
	//}
	fileName := header.Filename
	out, err := os.Create("./" + fileName)
	if err != nil {
		println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "讀取文件失敗",
			"data":    []int{},
		})
	}
	defer out.Close()

	i,err:=io.Copy(out,file)
	println(i)
	if err != nil {
		println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadGateway,
			"message": "存儲文件失敗",
			"data":    []int{},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "上傳成功，文件路徑為"+fileName,
		"data":    []int{},
	})
}
