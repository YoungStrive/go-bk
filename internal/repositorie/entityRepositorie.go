package repositorie

import (
	"go-bk/internal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func CreateUser(user *model.User) error {
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户名查找数据判断是否重名
func GetCountByUserName(name string) int64 {
	var nameCount int64
	DB.Model(&model.User{}).Where(&model.User{Name: name}).Count(&nameCount)
	return nameCount
}

func ListUserByName(name string) ([]map[string]interface{}, error) {
	var resultList []map[string]interface{}
	user := &model.User{}
	if name == "" {
		DB.Debug().Model(user).Select("id", "name", "sex").Scan(&resultList)
	} else {
		DB.Debug().Model(user).Select("id", "name", "sex").Where("name LIKE ? ", "%"+name+"%").Find(&resultList)
	}
	return resultList, nil
}

// 根据用户名获取用户信息
func GetByUserName(name string) *model.User {
	user := &model.User{}
	DB.Debug().Model(user).Where(&model.User{Name: name}).First(user)
	return user
}

// 根据用户id 获取发表的文章数
func GetPostCount(userId uint) uint64 {
	var postCount int64
	user := &model.User{ID: userId}
	DB.Debug().Model(&model.User{}).Select("post_count").
		Where(user).Find(&postCount)
	return uint64(postCount)
}

func AddPost(post *model.Post, postCount uint64) error {
	if err := DB.InstanceSet("postCount", postCount).Create(post).Error; err != nil {
		return err
	}
	return nil
}

func ListPost(keyword string, userId uint) ([]model.Post, error) {
	var resultList []model.Post
	post := &model.Post{}
	if keyword == "" {
		DB.Debug().Model(post).Select("id", "title", "content").Where("ref_user_id", userId).Find(&resultList)
	} else {
		DB.Debug().Model(post).Select("id", "title", "content").Where("ref_user_id=? ", userId).Where("title LIKE  '%" + keyword + "%'  or content LIKE  '%" + keyword + "%'").Find(&resultList)
	}
	return resultList, nil
}

// 获取文章详情
func GetPost(id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	DB.Debug().Model(&model.Post{}).Select("id", "title", "content").
		Where("id=?", id).First(&result)
	return result, nil
}

// 根据文章id 获取文章
func GetPostById(id int) *model.Post {
	var post model.Post
	DB.Debug().Find(&post, id)
	return &post
}

// 更改
func UpdatePost(post *model.Post) {
	DB.Debug().Save(post)
}

func DeletePost(post *model.Post) {
	DB.Debug().Delete(post)
}
