package controller

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	//"github.com/gin-gonic/gin"
	"crypto/md5"
	"encoding/hex"
	"github.com/melonws/goweb/libs/logHelper"
	"io"
	"os"
	"photo/config"
	"photo/dao/models"
	"strconv"
	"time"
)

//创建一条数据记录
func CreateItem(c *gin.Context) {
	file, header, err := c.Request.FormFile("imgs")
	title := c.PostForm("title")
	//文件不存在，返回错误信息
	if err != nil {
		logHelper.WriteLog("未获取到用户上传文件,错误原因"+err.Error(), "error/service")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "缺少上传文件",
			"data":    []int{},
		})
		return
	}

	fileName := header.Filename
	md := md5.New()
	timestamp := time.Now().Unix() // 获取时间戳
	println(timestamp)
	md.Write([]byte(strconv.FormatInt(timestamp, 10) + fileName)) // 转换为md5格式，时间戳+文件名
	cipherStr := md.Sum(nil)
	fileName = config.Config["imgPath"] + hex.EncodeToString(cipherStr) + ".jpg" // 拼接文件路径名称
	out, err := os.Create(fileName)
	if err != nil {
		println(err)
		logHelper.WriteLog("读取文件失败,错误原因"+err.Error(), "error/service")
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "读取文件失败,请联系管理员",
			"data":    []int{},
		})
		return
	}
	defer out.Close()

	i, err := io.Copy(out, file)
	println(i)
	if err != nil {
		println(err)
		logHelper.WriteLog("存储文件失败,错误原因:"+err.Error(), "error/service")
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadGateway,
			"message": "存储文件失败,请联系管理员",
			"data":    []int{},
		})
		return
	}
	logHelper.WriteLog("接收文件，存储路径为"+fileName, "system/access")
	stat, err := models.AddItem(&models.PhotoItem{title, fileName, time.Now().Unix()})
	//抛出stat为1正常，0失败
	if stat == 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "上传成功，文件路径为" + fileName,
			"data":    []int{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"stataus": 502,
			"message": "存储失败",
			"data":    nil,
		})
	}
}

//获取全部todo
func GetList(c *gin.Context) {
	status, data, err := models.GetList()
	fmt.Println(data)
	checkData(c, status, data, err)
}

/**
 * @Author      ljw
 * @DateTime    2018-07-28
 * @title       [data controll]
 * @description [检查函数执行状态，并返回对应处理的response]
 * @return      {[type]}
 */
func checkData(c *gin.Context, status int, data interface{}, err error) {
	println(c, status, data, err)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusBadGateway,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		if status != 1 {
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  http.StatusBadGateway,
				"message": "failed",
				"data":    [0]int{},
			})
		} else {
			// 返回数据方式 1、c.string，将json转为string；2、c.json返回map；3、c.json返回gin.H对象，应该是一个map转为json的方法
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"data":    data,
			})
		}
	}
}
