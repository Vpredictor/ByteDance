package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ByteDance/model"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	last_time := c.Query("last_time")
	video_sql := []model.Video_sql{}

	model.Db_read.Table("videos").Limit(5).Order("create_time desc").Where("create_time < ?",last_time).Find(&video_sql)

	next_time := video_sql[len(video_sql)-1].CreateTime

	video := make([]model.Video,len(video_sql))
	for i := 0 ; i < len(video_sql) ; i++ {
		video[i].Id = video_sql[i].Id
		user := model.User{}
		model.Db_read.Table("user").Where("id = ?",video[i].Id).First(&user)
		video[i].Title = video_sql[i].Title
		video[i].CommentCount = video_sql[i].CommentCount
		video[i].FavoriteCount = video_sql[i].FavoriteCount
		video[i].PlayUrl = video_sql[i].PlayUrl
		video[i].CoverUrl = video_sql[i].CoverUrl
		video[i].Author = user
		video[i].IsFavorite = false
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: video,
		NextTime:  next_time,
	})





	////得到用户的信息
	//user := User{}
	//model.Db_read.Table("user").Where("id = ?",user_id).First(&user)
	//
	//
	//video_sql :=[]Video_sql{}
	//
	//
	//model.Db_read.Table("videos").Where("author_id",user_id).Find(&video_sql)
	//video := make([]Video,len(video_sql))
	//for i := 0 ; i < len(video_sql) ; i++ {
	//	video[i].Id = video_sql[i].Id
	//	video[i].Title = video_sql[i].Title
	//	video[i].CommentCount = video_sql[i].CommentCount
	//	video[i].FavoriteCount = video_sql[i].FavoriteCount
	//	video[i].PlayUrl = video_sql[i].PlayUrl
	//	video[i].CoverUrl = video_sql[i].CoverUrl
	//	video[i].Author = user
	//	video[i].IsFavorite = false
	//}
	//
	//
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//		StatusMsg: "nil",
	//	},
	//	VideoList: video,
	//})
	//
	//
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
}
