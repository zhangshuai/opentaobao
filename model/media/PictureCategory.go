package media

import (
	"sync"
)

// PictureCategory 结构体
type PictureCategory struct {
	// 图片分类名
	PictureCategoryName string `json:"picture_category_name,omitempty" xml:"picture_category_name,omitempty"`
	// 图片分类型别，sys-fixture代表店铺装修分类(系统分类)，sys-auction代表宝贝图片分类(系统分类)，user-define代表用户自定义分类
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 图片类目的创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 图片分类的修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 图片分类ID
	PictureCategoryId int64 `json:"picture_category_id,omitempty" xml:"picture_category_id,omitempty"`
	// 图片分类排序
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 分类下的图片数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 一级分类的parent_id为0&lt;br/&gt;二级分类的则为其父分类的picture_category_id
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

var poolPictureCategory = sync.Pool{
	New: func() any {
		return new(PictureCategory)
	},
}

// GetPictureCategory() 从对象池中获取PictureCategory
func GetPictureCategory() *PictureCategory {
	return poolPictureCategory.Get().(*PictureCategory)
}

// ReleasePictureCategory 释放PictureCategory
func ReleasePictureCategory(v *PictureCategory) {
	v.PictureCategoryName = ""
	v.Type = ""
	v.Created = ""
	v.Modified = ""
	v.PictureCategoryId = 0
	v.Position = 0
	v.Total = 0
	v.ParentId = 0
	poolPictureCategory.Put(v)
}
