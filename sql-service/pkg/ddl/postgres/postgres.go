package postgres

import (
	"fmt"
	"strings"

	ddlerror "github.com/schema-creator/services/sql-service/pkg/ddl/error"
	"github.com/schema-creator/services/sql-service/pkg/utils"
)

type Postgres struct {
	lines []string
}

func NewPostgres(data string) *Postgres {
	return &Postgres{lines: utils.RmNilString(strings.Split(data, "\n"))}
}

func (p *Postgres) Convert() (result string, err error) {
	// 行毎に読み取り
	for i, line := range p.lines {
		line = strings.TrimSpace(line)
		words := utils.RmNilString(strings.Fields(line))
		// 先頭の単語を判別
		for _, word := range words {
			switch {
			// コメントアウトされた行
			case strings.Contains(word, "//"):
				result = commentOut(result, line)
			// テーブルを指定した行
			case strings.EqualFold(word, "table"):
				data, err := toTable(p.lines[i : i+p.end(i)])
				if err != nil {
					return "", err
				}
				result = fmt.Sprintf("%s\n%s", result, strings.TrimSpace(data))
			// Enumを指定した行
			case strings.EqualFold(word, "enum"):
				result = fmt.Sprintf("%s\n%s", strings.TrimSpace(toEnum(p.lines[i:i+p.end(i)-1])), result)
			// 参照を指定した行
			case strings.EqualFold(word, "->"):
				result = fmt.Sprintf("%s\n%s", result, strings.TrimSpace(toReference(line)))
			}
		}
	}
	return
}

func (p *Postgres) end(s int) (e int) {
	for i, l := range p.lines[s:] {
		if strings.Contains(l, "}") {
			e = i + 1
			break
		}
	}
	return
}

func commentOut(result, line string) string {
	return fmt.Sprintf("%s \n %s", result, strings.ReplaceAll(line, "//", "--"))
}

// テーブルの形に変換
func toTable(lines []string) (result string, err error) {
	const (
		column = iota
		typeName
	)

	for _, line := range lines {
		var row string
		words := utils.RmNilString(strings.Fields(line))

		switch {
		case strings.Contains(line, "//"):
			commentOut(result, line)
		case len(words) < 2:
			continue
		case strings.EqualFold(words[0], "table"):
			result = fmt.Sprintf(`%s CREATE TABLE "%s" (`, result, words[1])
			continue
		}

		row = fmt.Sprintf(`"%s" %s`, words[column], words[typeName])
		line = strings.Join(words[2:], " ")

		if len(line) > 0 {
			// [pk, nn, uq, df]のような形式なので、カンマで分割し、[]を削除後にそれぞれの要素を判別
			for _, attribute := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "]", ""), "[", ""), ",")) {
				switch {
				case strings.EqualFold(attribute, "pk"), strings.EqualFold(attribute, "primary key"):
					row = fmt.Sprintf("%s PRIMARY KEY", row)
				case strings.EqualFold(attribute, "nn"), strings.EqualFold(attribute, "not null"):
					row = fmt.Sprintf("%s NOT NULL", row)
				case strings.EqualFold(attribute, "uq"), strings.EqualFold(attribute, "unique"):
					row = fmt.Sprintf("%s UNIQUE", row)
				case strings.EqualFold(strings.TrimSpace(attribute), "df"), strings.EqualFold(strings.TrimSpace(attribute), "default"):
					if len(strings.Split(attribute, ":")) < 1 {
						err = ddlerror.ErrInvalidConvertType
						return
					}

					if strings.Contains(strings.Split(attribute, ":")[1], `"`) {
						row = fmt.Sprintf(`%s DEFAULT "%s"`, row, utils.RmNilString(strings.Split(attribute, ":"))[1])
					} else {
						row = fmt.Sprintf(`%s DEFAULT "%s"`, row, strings.ReplaceAll(strings.ReplaceAll(utils.RmNilString(strings.Split(attribute, ":"))[1], "`", utils.NilString), " ", utils.NilString))
					}
				}
			}
		}
		result = fmt.Sprintf("%-4s\n%s,", result, row)
	}
	result = fmt.Sprintf("%s\n);\n", result[:len(result)-1])
	return
}

// AlterTalbe ~ Add Foreign Key ~ References ~
func toReference(line string) (result string) {
	const (
		table = iota
		column
	)
	const (
		right = iota
		left
	)
	words := utils.RmNilString(strings.Split(line, "->"))

	refs := strings.Split(strings.TrimSpace(words[right]), ".")
	result = fmt.Sprintf(`REFERENCES "%s" ("%s");`, refs[table], refs[column])

	refs = strings.Split(strings.TrimSpace(words[left]), ".")
	result = fmt.Sprintf(`ALTER TABLE "%s" ADD FOREIGN KEY ("%s") %s`, refs[table], refs[column], result)
	return
}

// Create Type ~ As Enum ~
func toEnum(lines []string) (result string) {
	for _, line := range lines {
		switch {
		case strings.Contains(line, "//"):
			commentOut(result, line)
		case strings.EqualFold(utils.RmNilString(strings.Fields(line))[0], "enum"):
			result = fmt.Sprintf(`CREATE TYPE "%s" AS ENUM (`, utils.RmNilString(strings.Fields(line))[1])
			continue
		}
		result = fmt.Sprintf("%-4s\n'%s',", result, strings.TrimSpace(line))
	}
	result = fmt.Sprintf("%s\n);\n", result[:len(result)-1])
	return
}
