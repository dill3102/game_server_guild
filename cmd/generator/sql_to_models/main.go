package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	os.MkdirAll("cmd/generator/sql_to_models/models", os.ModePerm)

	file, err := os.Open("cmd/generator/sql_to_models/schema.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tableName string
	var insideTable bool
	var fields []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// コメント行を無視
		if strings.HasPrefix(line, "--") || line == "" {
			continue
		}

		if strings.HasPrefix(strings.ToUpper(line), "CREATE TABLE") {
			insideTable = true
			tableName = extractTableName(line)
			fields = nil
			continue
		}

		if strings.HasPrefix(strings.ToUpper(strings.Split(line, " ")[0]), "UNIQUE") {
			insideTable = true
			fields = nil
			continue
		}

		if insideTable {
			if strings.HasPrefix(line, ");") || strings.HasPrefix(line, ")") {
				insideTable = false
				generateModelFile(tableName, fields)
				continue
			}

			if !strings.HasPrefix(strings.ToUpper(line), "PRIMARY") {
				fields = append(fields, line)
			}
		}
	}

	// gofmtで整形
	exec.Command("gofmt", "-w", "./models").Run()
}

func extractTableName(line string) string {
	re := regexp.MustCompile("`?(\\w+)`?")
	parts := strings.Split(line, " ")
	for _, part := range parts {
		if strings.Contains(part, "(") {
			name := strings.TrimSuffix(part, "(")
			match := re.FindStringSubmatch(name)
			if len(match) > 1 {
				return match[1]
			}
		}
	}
	return ""
}

func generateModelFile(table string, fields []string) {
	structName := toCamelCase(table)
	fileName := fmt.Sprintf("cmd/generator/sql_to_models/models/%s.go", table)

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package models")
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, `  "gorm.io/gorm"`)
	fmt.Fprintln(f, `  "github.com/google/uuid"`)
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f, "")

	fmt.Fprintf(f, "type %s struct {\n", structName)
	fmt.Fprintln(f, "\tgorm.Model")
	fmt.Fprintln(f, "\tId string `gorm:\"column:id;primaryKey\" json:\"id\"`")

	var constructorFields []Column

	for _, field := range fields {
		col := parseFieldLine(field)

		if col.Name == "id" || col.Name == "created_at" || col.Name == "updated_at" || col.Name == "index" || col.Name == "unique" {
			continue
		}

		fmt.Fprintf(f, "\t%s %s `gorm:\"%s\" json:\"%s\"`\n", toCamelCase(col.Name), col.Type, col.Tag, col.Name)
		constructorFields = append(constructorFields, col)
	}

	fmt.Fprintln(f, "}")
	fmt.Fprintln(f, "")

	// NewXxx constructor
	fmt.Fprintf(f, "func New%s(\n", structName)
	for _, col := range constructorFields {
		fmt.Fprintf(f, "    %s %s,\n", col.Name, col.Type)
	}
	fmt.Fprintf(f, ") *%s {\n", structName)
	fmt.Fprintf(f, "    return &%s{\n", structName)
	fmt.Fprintf(f, "        Id: uuid.NewString(),\n")
	for _, col := range constructorFields {
		fmt.Fprintf(f, "        %s: %s,\n", toCamelCase(col.Name), col.Name)
	}
	fmt.Fprintln(f, "    }")
	fmt.Fprintln(f, "}")
}

type Column struct {
	Name        string
	DefaultName string
	Type        string
	Tag         string
}

func parseFieldLine(line string) Column {
	parts := strings.Fields(line)
	name := strings.Trim(parts[0], "`")
	sqlType := parts[1]
	tag := buildTag(line)

	goType := sqlToGoType(sqlType)
	return Column{
		Name:        name,
		DefaultName: name,
		Type:        goType,
		Tag:         tag,
	}
}

func sqlToGoType(sqlType string) string {
	upper := strings.ToUpper(sqlType)
	switch {
	case strings.HasPrefix(upper, "VARCHAR"):
		return "string"
	case upper == "DATETIME":
		return "time.Time"
	case strings.Contains(upper, "INT"):
		return "int"
	case upper == "BOOLEAN" || upper == "BOOL":
		return "bool"
	case upper == "TEXT":
		return "string"
	default:
		return "time.Time"
	}
}

func buildTag(line string) string {
	tag := []string{}
	if strings.Contains(line, "NOT NULL") {
		tag = append(tag, "not null")
	}
	if strings.Contains(line, "UNIQUE") {
		tag = append(tag, "unique")
	}
	if strings.Contains(line, "DEFAULT") {
		defaultVal := regexp.MustCompile(`DEFAULT\s+['"]?([^'"]+)['"]?`).FindStringSubmatch(line)
		if len(defaultVal) > 1 {
			tag = append(tag, fmt.Sprintf("default:%s", defaultVal[1]))
		}
	}
	return strings.Join(tag, ";")
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
