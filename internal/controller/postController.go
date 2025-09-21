package controller

import (
	"github.com/gin-gonic/gin"
	"go-bk/internal/model"
	"go-bk/internal/server"
	"go-bk/pkg/response"
	"net/http"
)

// 添加文章
func AddPost(c *gin.Context) {
	//获取用户id 需要登录
	userId, ok := c.Get("userId")
	if !ok {
		return
	}
	post := &model.AddPost{}
	//参数不对
	if err := c.ShouldBind(post); err != nil {
		response.Error(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	post.UserId = userId.(uint)
	err := server.AddPost(post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, post)
}

func ListPost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}
	uId := userId.(uint)
	keyword := c.Query("keyword")
	allPost, err := server.ListPost(keyword, uId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, allPost)
}

func GetPost(c *gin.Context) {
	_, ok := c.Get("userId")
	if !ok {
		return
	}
	id := c.Query("id")
	post, err := server.GetPost(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, post)
}

// 更新评论
func UpdatePost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}
	post := &model.UpdatePost{}
	//参数不对
	if err := c.ShouldBind(post); err != nil {
		response.Error(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	err := server.UpdatePost(userId.(uint), post)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, "更新成功")
}

func DeletePost(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}
	postId := c.Query("postId")
	err := server.DeletePost(userId.(uint), postId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, "删除成功")

}
