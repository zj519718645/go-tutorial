package model

type Tag struct {
	// id
	Id int32 `json:"id"`
	// 创建人
	Name string `json:"name"`
	// 创建时间
	CreatedOn int32 `json:"created_on"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 修改时间
	ModifiedOn int32 `json:"modified_on"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
	// 删除时间
	DeletedOn int32 `json:"deleted_on"`
	// 是否删除0为未删除、1为已删除
	IsDel int8 `json:"is_del"`
	// 状态0为禁用、1为启用
	State int8 `json:"state"`
}

func (model Tag) TableName() string {
	return "blog_tag"
}
