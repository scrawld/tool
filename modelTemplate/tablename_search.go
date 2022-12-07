// %s/__PROJECT_NAME__/ProjectName/g | %s/__TABLE_NAME__/TableName/g | %s/__TABLE_NAME_CAMEL__/TableNameCamel/g | %s/__TABLE_COMMENT__/TableComment/g | 1d | nohl
package __TABLE_NAME__

import (
	"__PROJECT_NAME__/models"
	"__PROJECT_NAME__/utils"

	"github.com/jinzhu/gorm"
)

type __TABLE_NAME_CAMEL__Search struct {
	__TABLE_NAME_CAMEL__
	Limit  int
	Offset int
	Sort   []*models.SortParams
	query  *__TABLE_NAME_CAMEL__Query
}

func New__TABLE_NAME_CAMEL__Search() *__TABLE_NAME_CAMEL__Search {
	o := &__TABLE_NAME_CAMEL__Search{}
	o.query = New(nil)
	o.Limit = 100
	return o
}

func (this *__TABLE_NAME_CAMEL__Search) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__Search {
	utils.Assign(inp, this, excludes...)
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) SetLimit(limit int) *__TABLE_NAME_CAMEL__Search {
	this.Limit = limit
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) SetOffset(offset int) *__TABLE_NAME_CAMEL__Search {
	this.Offset = offset
	return this
}

func (this *__TABLE_NAME_CAMEL__Search) buildCond() (r *gorm.DB) {
	r = this.query.Active()

	/*
	   if len(this.Title) > 0 {
	       r = r.Where("title LIKE ?", "%"+this.Title+"%")
	   }
	*/
	return
}

func (this *__TABLE_NAME_CAMEL__Search) Count() (r int64, err error) {
	err = this.buildCond().Model(&__TABLE_NAME_CAMEL__{}).Count(&r).Error
	return
}

func (this *__TABLE_NAME_CAMEL__Search) Search() (r []*__TABLE_NAME_CAMEL__, err error) {
	orm := this.buildCond().Limit(this.Limit)

	if len(this.Sort) > 0 {
		orm = orm.Order(models.ParseSortParams(this.Sort))
	} else {
		orm = orm.Order(this.TableName() + ".id DESC")
	}
	if this.Offset > 0 {
		orm = orm.Offset(this.Offset)
	}
	err = orm.Find(&r).Error
	return
}
