package controller

import (
	"github.com/Limeng-svg/ByteDance/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type loginOrRegisterReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type userInfoResp struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	User interface{} `json:"user"`
}

func GetUserInfo(c *gin.Context) {
	userID := c.Query("user_id")
	targetUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(200, &userInfoResp{
			Code: -1,
			Msg:  "user_id is not valid",
			User: nil,
		})
		return
	}
	id := c.GetInt("user_id")
	userInfo, err := service.GetUserInfo(id, targetUserID)
	if err != nil {
		c.JSON(200, &userInfoResp{
			Code: -1,
			Msg:  err.Error(),
			User: nil,
		})
		return
	}
	c.JSON(200, &userInfoResp{
		Code: 0,
		Msg:  "success",
		User: userInfo,
	})
}

type registerAndLoginResp struct {
	Code   int    `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	userInfo := loginOrRegisterReq{}
	_ = c.ShouldBind(&userInfo)
	err := service.Register(userInfo.Username, userInfo.Password)
	if err != nil {
		c.JSON(200, &registerAndLoginResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	data, err := service.Login(userInfo.Username, userInfo.Password)
	if err != nil {
		c.JSON(200, &registerAndLoginResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &registerAndLoginResp{
		Code:   0,
		Msg:    "success",
		UserID: data.ID,
		Token:  data.Token,
	})
}

func Login(c *gin.Context) {
	userInfo := loginOrRegisterReq{}
	_ = c.ShouldBind(&userInfo)
	data, err := service.Login(userInfo.Username, userInfo.Password)
	if err != nil {
		c.JSON(200, &registerAndLoginResp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, &registerAndLoginResp{
		Code:   0,
		Msg:    "success",
		UserID: data.ID,
		Token:  data.Token,
	})
}
