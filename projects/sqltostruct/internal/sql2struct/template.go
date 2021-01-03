package sql2struct

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const structTpl = `type {{.TableName | UnderscoreToUpperCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | UnderscoreToUpperCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | UnderscoreToUpperCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`json:\"%s\"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl, err := template.New("sql2struct").Funcs(template.FuncMap{"UnderscoreToUpperCamelCase": UnderscoreToUpperCamelCase}).Parse(t.structTpl)
	if err != nil {
		return err
	}
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err = tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
