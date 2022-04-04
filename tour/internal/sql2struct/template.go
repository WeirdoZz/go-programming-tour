package sql2struct

import (
	"fmt"
	"html/template"
	"os"
	"tour/internal/word"
)

/*
将mysql.go中从数据库映射来的数据映射到模板中
*/

const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ Extra close brace or missing open bracelength 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ Extra close brace or missing open bracetypeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

// StructColumn 保存列的相关信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// StructTemplateDB 保存表名和其中的列
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

// NewStructTemplate 新建一个模板对象
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: strcutTpl}
}

// AssemblyColumns 对数据库直接映射过来的结构体进行进一步分解转换（比如数据库类型到go类型的转换）
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		//设置结构体的tag，比如其json名
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

// Generate 对模板中的函数做定义并且渲染模板
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
