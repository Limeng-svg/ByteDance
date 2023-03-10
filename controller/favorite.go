package controller

import (
	"github.com/Limeng-svg/ByteDance/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetFavoriteList(c *gin.Context) {
	userID, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	targetUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}
	id := c.GetInt("user_id")
	videoList, err := service.GetFavoriteList(id, targetUserID)
	if err != nil {
		c.JSON(200, &videoListResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &videoListResp{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}

func AddOrCancelFavorite(c *gin.Context) {
	vID, ok := c.GetQuery("video_id")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "video_id参数错误",
		})
		return
	}
	uID := c.GetInt("user_id")
	isAdd, ok := c.GetQuery("action_type")
	if !ok {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "action_type参数错误",
		})
		return
	}
	videoID, err := strconv.Atoi(vID)
	if err != nil {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "video_id参数错误",
		})
		return
	}
	actionType, err := strconv.Atoi(isAdd)
	if err != nil || actionType != 1 && actionType != 2 {
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  "actionType参数错误",
		})
		return
	}
	err = service.AddOrCancelFavorite(uID, videoID, actionType)
	if err != nil {
		var errMsg string
		if actionType == 1 {
			errMsg = "点赞失败"
		} else if actionType == 2 {
			errMsg = "取消点赞失败"
		}
		c.JSON(200, &gin.H{
			"status_code": -1,
			"status_msg":  errMsg,
		})
		return
	}
	c.JSON(200, &gin.H{
		"status_code": 0,
		"status_msg":  "success",
	})
}
