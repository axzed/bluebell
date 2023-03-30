package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetail 根据id查询社区分类详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
