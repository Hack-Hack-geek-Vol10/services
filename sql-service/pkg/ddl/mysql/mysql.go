package mysql

import (
	"fmt"
	"strconv"
	"strings"

	ddlerror "github.com/schema-creator/services/sql-service/pkg/ddl/error"
	"github.com/schema-creator/services/sql-service/pkg/utils"
)

type Mysql struct {
	lines []string
}

func NewMysql(data string) *Mysql {
	return &Mysql{lines: utils.RmNilString(strings.Split(data, "\n"))}
}

func (m *Mysql) Convert() (result string, err error) {
	// 行毎に読み取り
	for i, line := range m.lines {
		line = strings.TrimSpace(line)
		words := utils.RmNilString(strings.Fields(line))
		// 先頭の単語を判別
		for _, word := range words {
			switch {
			// コメントアウトされた行
			case strings.Contains(word, "//"):
				continue
			// テーブルを指定した行
			case strings.EqualFold(word, "table"):
				data, err := m.toTable(m.lines[i : i+m.end(i)])
				if err != nil {
					return "", err
				}
				result = fmt.Sprintf("%s\n%s", result, strings.TrimSpace(data))
			// 参照を指定した行
			case strings.EqualFold(word, "->"):
				result = fmt.Sprintf("%s\n%s", result, strings.TrimSpace(toReference(line)))
			}
		}
	}
	return
}

func (m *Mysql) end(s int) (e int) {
	for i, l := range m.lines[s:] {
		if strings.Contains(l, "}") {
			e = i + 1
			break
		}
	}
	return
}

// テーブルの形に変換
func (m *Mysql) toTable(lines []string) (result string, err error) {
	const (
		column = iota
		typeName
	)

	for _, line := range lines {
		var row string
		words := utils.RmNilString(strings.Fields(line))

		switch {
		case strings.Contains(line, "//"):
			continue
		case len(words) < 2:
			continue
		case strings.EqualFold(words[0], "table"):
			result = fmt.Sprintf("%s CREATE TABLE `%s` (", result, words[1])
			continue
		}

		row = fmt.Sprintf("`%s` %s", words[column], m.checkType(words[typeName]))
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
						row = fmt.Sprintf("%s DEFAULT '%s'", row, utils.RmNilString(strings.Split(attribute, ":"))[1])
					} else {
						row = fmt.Sprintf("%s DEFAULT %s", row, strings.ReplaceAll(strings.ReplaceAll(utils.RmNilString(strings.Split(attribute, ":"))[1], "`", utils.NilString), " ", utils.NilString))
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
	result = fmt.Sprintf("REFERENCES `%s (`%s`);", refs[table], refs[column])

	refs = strings.Split(strings.TrimSpace(words[left]), ".")
	result = fmt.Sprintf("ALTER TABLE `%s` ADD FOREIGN KEY (`%s`) %s", refs[table], refs[column], result)
	return
}

// int => TINYINT[M,U,Z] | SMALLINT[M,U,Z] | MEDIUMINT[M,U,Z] | INT[M,U,Z] | BIGINT[M,U,Z] "M"=>1~255, "U"=> unsigned, "Z"=> zerofill
// float => FLOAT[M,D,U,Z] | DOUBLE[M,D,U,Z] | DECIMAL[M,D,U,Z] "M"=>1~255, "D"=>0~30, "U"=> unsigned, "Z"=> zerofill
// bit => BIT[M] "M"=>1~64
// date => DATE | DATETIME | TIMESTAMP | TIME | YEAR
// varchar => VARCHAR[M,cs:"charset",col:"collate"] "M"=>1~65535 | CHAR[M] "M"=>0~65535 (charset,collateはそのままぶち込む)
// binary => BINARY[M] "M"=>1~255 | VARBINARY[M] "M"=>1~65535
// blob => TINYBLOB | BLOB[M] | MEDIUMBLOB | LONGBLOB "M"=>1~65535
// text => TINYTEXT[charset,collate] | TEXT[M,charset,collate] | MEDIUMTEXT[charset,collate] | LONGTEXT[charset,collate]
// enum => default に
// set => 非対応
// word -> type
func (m *Mysql) checkType(word string) (result string) {
	types := strings.TrimSpace(strings.Split(word, "(")[0])
	switch {
	case strings.EqualFold(types, "TINYINT"), strings.EqualFold(types, "SMALLINT"), strings.EqualFold(types, "MEDIUMINT"), strings.EqualFold(types, "INT"), strings.EqualFold(types, "BIGINT"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		result = types
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)

			if _, err := strconv.Atoi(v); err == nil {
				result = fmt.Sprintf("%s(%s)", result, v)
				continue
			}

			switch {
			case strings.EqualFold(v, "U"):
				result = fmt.Sprintf("%s UNSIGNED", result)
			case strings.EqualFold(v, "Z"):
				result = fmt.Sprintf("%s ZEROFILL", result)
			default:
				continue
			}
		}
	case strings.EqualFold(types, "FLOAT"), strings.EqualFold(types, "DOUBLE"), strings.EqualFold(types, "DECIMAL"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		result = types
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			// if reflect.TypeOf(v).Kind() == reflect.Int {
			// 	result = fmt.Sprintf("%s(%d)", result, v)
			// 	continue
			// }
			switch {
			case strings.EqualFold(v, "U"):
				result = fmt.Sprintf("%s UNSIGNED", result)
			case strings.EqualFold(v, "Z"):
				result = fmt.Sprintf("%s ZEROFILL", result)
			default:
				continue
			}
		}
	case strings.EqualFold(types, "BIT"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			if _, err := strconv.Atoi(v); err == nil {
				result = fmt.Sprintf("%s(%s)", result, v)
				continue
			}
		}
	case strings.EqualFold(types, "DATE"), strings.EqualFold(types, "DATETIME"), strings.EqualFold(types, "TIMESTAMP"), strings.EqualFold(types, "TIME"), strings.EqualFold(types, "YEAR"):
		return types
	case strings.EqualFold(types, "VARCHAR"), strings.EqualFold(types, "CHAR"):
		if len(strings.Split(word, "(")) == 1 {
			return word + "(255)"
		}
		result = types
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			if _, err := strconv.Atoi(v); err == nil {
				result = fmt.Sprintf("%s(%s)", result, v)
				continue
			}

			cscol := utils.RmNilString(strings.Split(v, ":"))
			switch {
			case strings.EqualFold(cscol[0], "charset"):
				result = fmt.Sprintf("%s CHARACTER SET %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			case strings.EqualFold(cscol[0], "collate"):
				result = fmt.Sprintf("%s COLLATE %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			default:
				continue
			}
		}
	case strings.EqualFold(types, "BINARY"), strings.EqualFold(types, "VARBINARY"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			if _, err := strconv.Atoi(v); err == nil {
				result = fmt.Sprintf("%s(%s)", result, v)
				continue
			}
		}
	case strings.EqualFold(types, "TINYBLOB"), strings.EqualFold(types, "BLOB"), strings.EqualFold(types, "MEDIUMBLOB"), strings.EqualFold(types, "LONGBLOB"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		result = types
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			if strings.EqualFold(types, "BLOB") {
				if _, err := strconv.Atoi(v); err == nil {
					result = fmt.Sprintf("%s(%s)", result, v)
					continue
				}
			}

			cscol := utils.RmNilString(strings.Split(v, ":"))
			switch {
			case strings.EqualFold(cscol[0], "charset"):
				result = fmt.Sprintf("%s CHARACTER SET %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			case strings.EqualFold(cscol[0], "collate"):
				result = fmt.Sprintf("%s COLLATE %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			default:
				continue
			}
		}
	case strings.EqualFold(types, "TINYTEXT"), strings.EqualFold(types, "TEXT"), strings.EqualFold(types, "MEDIUMTEXT"), strings.EqualFold(types, "LONGTEXT"):
		if len(strings.Split(word, "(")) == 1 {
			return word
		}
		result = types
		for _, v := range utils.RmNilString(strings.Split(strings.ReplaceAll(strings.Split(word, "(")[1], ")", ""), ",")) {
			v = strings.TrimSpace(v)
			if strings.EqualFold(types, "TEXT") {
				if _, err := strconv.Atoi(v); err == nil {
					result = fmt.Sprintf("%s(%s)", result, v)
					continue
				}
			}

			cscol := utils.RmNilString(strings.Split(v, ":"))
			switch {
			case strings.EqualFold(cscol[0], "charset"):
				result = fmt.Sprintf("%s CHARACTER SET %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			case strings.EqualFold(cscol[0], "collate"):
				result = fmt.Sprintf("%s COLLATE %s", result, strings.ReplaceAll(cscol[1], `"`, ""))
			default:
				continue
			}
		}
	default:
		return m.toEnum(strings.TrimSpace(types))
	}
	return
}

func (m *Mysql) toEnum(name string) (result string) {
	var (
		readFlag bool = false
	)
	for _, line := range m.lines {
		if readFlag {
			if strings.Contains(line, "}") {
				break
			}
			result = fmt.Sprintf("%s'%s', ", result, strings.TrimSpace(line))
		}
		v := utils.RmNilString(strings.Fields(line))
		if len(v) == 0 {
			continue
		}
		switch {
		case strings.Contains(line, "//"):
			continue
		case strings.EqualFold(v[0], "enum"):
			if strings.Fields(line)[1] == name {
				readFlag = true
				result = fmt.Sprint("ENUM (")
			}
		}
	}
	if len(result) == 0 {
		return
	}
	result = fmt.Sprintf("%s)", result[:len(result)-2])
	return
}
