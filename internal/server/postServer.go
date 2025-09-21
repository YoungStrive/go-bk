package server

import (
	"errors"
	"go-bk/internal/model"
	"go-bk/internal/repositorie"
	"strconv"
)

func AddPost(post *model.AddPost) error {

	postCount := repositorie.GetPostCount(post.UserId)

	postEntity := &model.Post{
		Title:     post.Title,
		Content:   post.Content,
		RefUserID: post.UserId,
	}
	return repositorie.AddPost(postEntity, postCount)

}

func ListPost(keyword string, userId uint) ([]model.Post, error) {
	return repositorie.ListPost(keyword, userId)

}

func GetPost(postId string) (map[string]interface{}, error) {
	return repositorie.GetPost(postId)
}

// 修改文章
func UpdatePost(userId uint, post *model.UpdatePost) error {
	postId := post.ID
	postEntity := repositorie.GetPostById(postId)
	if postEntity.ID == 0 {
		return errors.New("文章不存在")
	}
	if postEntity.RefUserID != userId {
		return errors.New("没有权限修改此文章")
	}
	postEntity.Title = post.Title
	postEntity.Content = post.Content
	repositorie.UpdatePost(postEntity)
	return nil
}

func DeletePost(userId uint, postId string) error {

	id, _ := strconv.Atoi(postId)
	postEntity := repositorie.GetPostById(id)
	if postEntity.ID == 0 {
		return errors.New("文章不存在")
	}
	if postEntity.RefUserID != userId {
		return errors.New("没有权限删除此文章")
	}
	repositorie.DeletePost(postEntity)
	return nil
}
