package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"ByteDance/config"
	"strconv"
	"time"
	"ByteDance/model"
)

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	file, err := c.FormFile("data")
	if err != nil {
		fmt.Println("form failed ....", err)
	}
	local_filename := "./public/" +  file.Filename
	fmt.Println(local_filename)

	c.SaveUploadedFile(file, local_filename)

	title := c.PostForm("title")
	fmt.Println(title)
	objectName := model.ObjectName + file.Filename
	fmt.Println(objectName)

	err = model.Bucket.PutObjectFromFile(objectName, local_filename)
	if err != nil {
		model.HandleError(err)
	}

   err = os.Remove(local_filename)
   if err != nil {
	   fmt.Println("remove local file failed",err)
   }


   video := model.Video_sql{}
   user_id,_ := c.Get("user_id")
   uid := config.GetInterfaceToString(user_id)

   video.Author_id,_= strconv.Atoi(uid)

   play_url := model.Base_url + file.Filename
   coverUrl := model.Static_url
   video.PlayUrl = play_url
   video.CoverUrl = coverUrl
   video.Title = title
   video.CreateTime = time.Now().Unix()

   model.Db_write.Table("videos").Create(&video)

	c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg: "nil",
		})





	//filename := "./public/" + file.Filename
	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//saveFile := filepath.Join("./public/", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, Response{
	//	StatusCode: 0,
	//	StatusMsg:  finalName + " uploaded successfully",
	//})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	user_id := c.Query("user_id")

	//得到用户的信息
	user := model.User{}
	model.Db_read.Table("user").Where("id = ?",user_id).First(&user)


    video_sql :=[]model.Video_sql{}


	model.Db_read.Table("videos").Where("author_id",user_id).Find(&video_sql)
	video := make([]model.Video,len(video_sql))
	for i := 0 ; i < len(video_sql) ; i++ {
		video[i].Id = video_sql[i].Id
		video[i].Title = video_sql[i].Title
		video[i].CommentCount = video_sql[i].CommentCount
		video[i].FavoriteCount = video_sql[i].FavoriteCount
		video[i].PlayUrl = video_sql[i].PlayUrl
		video[i].CoverUrl = video_sql[i].CoverUrl
		video[i].Author = user
		video[i].IsFavorite = false
	}


	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg: "nil",
		},
		VideoList: video,
	})
}
