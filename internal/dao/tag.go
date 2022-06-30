package dao

import (
	"go-depot/internal/model"
	"go-depot/pkg/app"
)

/**
DAO层
数据访问对象的封装
*/
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.db)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.db, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}

	return tag.Create(d.db)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		ID:         id,
		Name:       name,
		State:      state,
		ModifiedBy: modifiedBy,
	}
	return tag.Update(d.db)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{ID: id}
	return tag.Delete(d.db)
}
