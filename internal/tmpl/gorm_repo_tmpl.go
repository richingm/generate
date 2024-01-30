package tmpl

const RepoTemplate = NotEditMark + `
package {{.PkgName}}

import (
	"gorm.io/gorm"
)

type {{.StructName}}Wrapper struct {
	{{.StructFieldName}}    *{{.StructName}}
	where   map[string]interface{}
	orderBy string
}

type {{.StructName}}Repo struct {
	db *gorm.DB
}

func New{{.StructName}}Repo(db *gorm.DB) *{{.StructName}}Repo {
	return &{{.StructName}}Repo{db: db}
}

func (Repo {{.StructName}}Repo) QueryWrapper({{.StructFieldName}} *{{.StructName}}) *{{.StructName}}Wrapper {
	return &{{.StructName}}Wrapper{
		{{.StructFieldName}}:  {{.StructFieldName}},
		where: make(map[string]interface{}),
	}
}

func (wrap *{{.StructName}}Wrapper) Where(query string, arg interface{}) *{{.StructName}}Wrapper {
	wrap.where[query] = arg
	return wrap
}

func (wrap *{{.StructName}}Wrapper) OrderBy(query string) *{{.StructName}}Wrapper {
	wrap.orderBy = query
	return wrap
}

func (Repo {{.StructName}}Repo) Insert({{.StructFieldName}} *{{.StructName}}) error {
	return Repo.db.Create({{.StructFieldName}}).Error
}

func (Repo {{.StructName}}Repo) DeleteByWrapper(wrap *{{.StructName}}Wrapper) error {
	orm := Repo.db.Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	return orm.Delete({{.StructName}}{}).Error
}

func (Repo {{.StructName}}Repo) DeleteById(id interface{}) error {
	return Repo.db.Where("id = ?", id).Delete({{.StructName}}{}).Error
}

func (Repo {{.StructName}}Repo) DeleteBatchIds(ids ...interface{}) error {
	return Repo.db.Where("id in ?", ids).Delete({{.StructName}}{}).Error
}

func (Repo {{.StructName}}Repo) UpdateById(id interface{}, {{.StructFieldName}} *{{.StructName}}) error {
	return Repo.db.Where("id = ?", id).Updates({{.StructFieldName}}).Error
}

func (Repo {{.StructName}}Repo) SelectById(id interface{}) (*{{.StructName}}, error) {
	{{.StructFieldName}} := &{{.StructName}}{}
	err := Repo.db.Table(TableName{{.StructName}}).Where("id = ?", id).First({{.StructFieldName}}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return {{.StructFieldName}}, err
}

func (Repo {{.StructName}}Repo) SelectOne(wrap *{{.StructName}}Wrapper) (*{{.StructName}}, error) {
	{{.StructFieldName}} := &{{.StructName}}{}
	orm := Repo.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	err := orm.Order(wrap.orderBy).First({{.StructFieldName}}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return {{.StructFieldName}}, err
}

func (Repo {{.StructName}}Repo) SelectList(wrap *{{.StructName}}Wrapper) ({{.StructFieldName}}List []{{.StructName}}, err error) {
	orm := Repo.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	return {{.StructFieldName}}List, orm.Order(wrap.orderBy).Find(&{{.StructFieldName}}List).Error
}

func (Repo {{.StructName}}Repo) SelectMaps(wrap *{{.StructName}}Wrapper) (list []map[string]interface{}, err error) {
	{{.StructFieldName}}List := make([]{{.StructName}}, 0)
	orm := Repo.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	err = orm.Find(&userList).Error
	for _, {{.StructFieldName}} := range {{.StructFieldName}}List {
		m := make(map[string]interface{})
		obj := {{.StructFieldName}}
		{{range .Fields}} m["{{.Name}}"] = obj.{{.Name}}
		{{end}}
		list = append(list, m)
	}

	return
}

func (Repo {{.StructName}}Repo) SelectCount(wrap *{{.StructName}}Wrapper) (count int64, err error) {
	orm := Repo.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	return count, orm.Count(&count).Error
}

func (Repo {{.StructName}}Repo) SelectPage(limit, offset int, wrap *{{.StructName}}Wrapper) ({{.StructFieldName}}List []{{.StructName}}, err error) {
	orm := Repo.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	err = orm.Limit(limit).Offset(offset).Order(wrap.orderBy).Find(&{{.StructFieldName}}List).Error
	return {{.StructFieldName}}List, err
}

`
