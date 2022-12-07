// %s/__PROJECT_NAME__/ProjectName/g | %s/__TABLE_NAME__/TableName/g | %s/__TABLE_NAME_CAMEL__/TableNameCamel/g | %s/__TABLE_COMMENT__/TableComment/g | 1d | nohl
package __TABLE_NAME__

import (
	"errors"
	"time"

	"__PROJECT_NAME__/models"
	"__PROJECT_NAME__/utils"

	"github.com/jinzhu/gorm"
)

// `gorm:"column:;type:;not null;default:'';comment:'';"`

// __TABLE_NAME_CAMEL__ __TABLE_COMMENT__
type __TABLE_NAME_CAMEL__ struct{}

func (this *__TABLE_NAME_CAMEL__) TableName() string {
	return "__TABLE_NAME__"
}

type __TABLE_NAME_CAMEL__Query struct {
	__TABLE_NAME_CAMEL__
	orm        *gorm.DB
	mustColumn []string
}

func New(orm *gorm.DB) *__TABLE_NAME_CAMEL__Query {
	o := &__TABLE_NAME_CAMEL__Query{}
	o.orm = orm
	return o
}

func (this *__TABLE_NAME_CAMEL__Query) Orm() *gorm.DB {
	if this.orm == nil {
		this.orm = models.Orm
	}
	return this.orm
}

func (this *__TABLE_NAME_CAMEL__Query) CreateTable() {
	if !this.Orm().HasTable(&__TABLE_NAME_CAMEL__{}) {
		this.Orm().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='__TABLE_COMMENT__'").CreateTable(&__TABLE_NAME_CAMEL__{})
	}
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Load(inp interface{}, excludes ...string) *__TABLE_NAME_CAMEL__Query {
	utils.Assign(inp, this, excludes...)
	return this
}

func (this *__TABLE_NAME_CAMEL__Query) Active() *gorm.DB {
	return this.Orm().Table(this.TableName()).Where(this.TableName() + ".delete_time=0")
}

// MustCols update use only: must update columns
func (this *__TABLE_NAME_CAMEL__Query) MustCols(cols ...string) *__TABLE_NAME_CAMEL__Query {
	this.mustColumn = append(this.mustColumn, cols...)
	return this
}

func (this *__TABLE_NAME_CAMEL__Query) GetById(id int64) (r *__TABLE_NAME_CAMEL__, err error) {
	r = &__TABLE_NAME_CAMEL__{}
	err = this.Active().Where("id=?", id).Take(r).Error
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Find() (r []*__TABLE_NAME_CAMEL__, r2 map[int64]*__TABLE_NAME_CAMEL__, err error) {
	err = this.Active().
		Find(&r).Error
	r2 = map[int64]*__TABLE_NAME_CAMEL__{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *__TABLE_NAME_CAMEL__Query) Create(inp interface{}) error {
	this.Load(inp)
	this.CreateTime = time.Now().Unix()
	this.UpdateTime = time.Now().Unix()
	return this.Orm().Create(&this.__TABLE_NAME_CAMEL__).Error
}

func (this *__TABLE_NAME_CAMEL__Query) Update(inp interface{}) (err error) {
	if inp != nil {
		this.Load(inp)
	}
	if this.Id == 0 {
		return errors.New("id is not set")
	}
	this.UpdateTime = time.Now().Unix()
	err = this.Orm().Model(&this.__TABLE_NAME_CAMEL__).
		Where("id=?", this.Id).
		Updates(models.ConvertStructToMap(&this.__TABLE_NAME_CAMEL__, this.mustColumn)).Error
	return
}

func (this *__TABLE_NAME_CAMEL__Query) UpdateMapById(id int64, inp map[string]interface{}) (err error) {
	if id == 0 {
		return errors.New("id is not set")
	}
	inp["update_time"] = time.Now().Unix()
	err = this.Orm().Model(&this.__TABLE_NAME_CAMEL__).
		Where("id=?", id).Updates(inp).Error
	return
}
