package gen

import (
	"fmt"

	"gopx/model/sql/template"
	"gopx/util"
	"gopx/util/pathx"
)

func genNew(table Table, withCache, postgreSql bool) (string, error) {
	text, err := pathx.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source()) + "`"
	}

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]any{
			"table":                 t,
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func genNewPrefix(table Table, withCache, postgreSql bool, prefix string) (string, error) {
	text, err := pathx.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))

	if len(prefix) > 0 {
		prefixName := fmt.Sprintf("%v%v", prefix, table.Name.Source())
		t = fmt.Sprintf(`"%s"`, wrapWithRawString(prefixName, postgreSql))
	}
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source()) + "`"
	}

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]any{
			"table":                 t,
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
