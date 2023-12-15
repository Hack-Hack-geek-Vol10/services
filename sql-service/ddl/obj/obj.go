package obj

import (
	"fmt"
	"strings"

	ddlerror "github.com/schema-creator/services/sql-service/pkg/ddl/error"
	"github.com/schema-creator/services/sql-service/pkg/utils"
)

type OBJ struct {
	lines []string
}

func NewOBJ(data string) *OBJ {
	return &OBJ{lines: utils.RmNilString(strings.Split(data, "\n"))}
}

func (o *OBJ) Convert() (*Object, error) {
	// 行毎に読み取り
	var result Object
	result.RawData = strings.Join(o.lines, "\n")
	for i, line := range o.lines {
		line = strings.TrimSpace(line)
		words := utils.RmNilString(strings.Fields(line))
		// 先頭の単語を判別
		for _, word := range words {
			switch {
			// コメントアウトされた行
			case strings.Contains(word, "//"):
				break
			// テーブルを指定した行
			case strings.EqualFold(word, "table"):
				table, err := toTable(o.lines[i : i+o.end(i)])
				if err != nil {
					return nil, err
				}
				result.Tables = append(result.Tables, table)
			// Enumを指定した行
			case strings.EqualFold(word, "enum"):
				result.Enums = append(result.Enums, toEnum(o.lines[i:i+o.end(i)-1]))
			// 参照を指定した行
			case strings.EqualFold(word, "->"):
				result.Relations = append(result.Relations, toReference(line))
			}
		}
	}
	return &result, nil
}

func (o *OBJ) end(s int) (e int) {
	for i, l := range o.lines[s:] {
		if strings.Contains(l, "}") || strings.EqualFold(l, "table") {
			e = i + 1
			break
		}
	}
	return
}

func toTable(lines []string) (*Table, error) {
	const (
		name = iota
		typeName
	)
	var result Table
	for _, line := range lines {
		words := utils.RmNilString(strings.Fields(line))

		switch {
		case strings.Contains(line, "//"):
			continue
		case len(words) < 2:
			continue
		case strings.EqualFold(words[0], "table"):
			result.Name = words[1]
			continue
		}

		column := &Column{
			Name: words[name],
			Type: words[typeName],
		}

		line = strings.Join(words[2:], " ")
		if len(line) > 0 {
			for _, attribute := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "]", ""), "[", ""), ",")) {
				switch {
				case strings.EqualFold(attribute, "pk"), strings.EqualFold(attribute, "primary key"):
					column.Options = append(column.Options, "PK")
				case strings.EqualFold(attribute, "nn"), strings.EqualFold(attribute, "not null"):
					column.Options = append(column.Options, "NN")
				case strings.EqualFold(attribute, "uq"), strings.EqualFold(attribute, "unique"):
					column.Options = append(column.Options, "UQ")
				case strings.EqualFold(strings.TrimSpace(attribute), "df"), strings.EqualFold(strings.TrimSpace(attribute), "default"):
					if len(strings.Split(attribute, ":")) < 1 {
						err := ddlerror.ErrInvalidConvertType
						return nil, err
					}

					if strings.Contains(strings.Split(attribute, ":")[1], `"`) {
						column.Options = append(column.Options, fmt.Sprintf("DEFAULT '%s'", utils.RmNilString(strings.Split(attribute, ":"))[1]))
					} else {
						column.Options = append(column.Options, fmt.Sprintf("DEFAULT %s", strings.ReplaceAll(strings.ReplaceAll(utils.RmNilString(strings.Split(attribute, ":"))[1], "`", utils.NilString), " ", utils.NilString)))
					}
				}
			}
		}
		result.Columns = append(result.Columns, column)
	}
	return &result, nil
}

func toEnum(lines []string) *Enum {
	var result Enum
	for _, line := range lines {
		switch {
		case strings.Contains(line, "//"):
			continue
		case strings.EqualFold(utils.RmNilString(strings.Fields(line))[0], "enum"):
			result.Name = utils.RmNilString(strings.Fields(line))[1]
			continue
		}
		result.Fields = append(result.Fields, strings.TrimSpace(line))
	}
	return &result
}

func toReference(line string) *Relation {
	const (
		right = iota
		left
	)
	words := utils.RmNilString(strings.Split(line, "->"))

	return &Relation{
		FromCol: strings.TrimSpace(words[right]),
		ToCol:   strings.TrimSpace(words[left]),
	}
}
