package driver

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/vimcoders/sql"
)

func WithQueryer(row sql.Row) sql.Builder {
	return &queryer{
		row: row,
	}
}

func WithInserter(row sql.Row) sql.Builder {
	return &inserter{
		row: row,
	}
}

type queryer struct {
	row sql.Row
}

func (q *queryer) Clone() sql.Row {
	t := reflect.TypeOf(q.row).Elem()

	if t == nil {
		return nil
	}

	i := reflect.New(t).Interface()

	if i == nil {
		return nil
	}

	return i.(sql.Row)
}

func (q *queryer) TableName() string {
	if tableName := q.row.TableName(); len(tableName) > 0 {
		return tableName
	}

	if t := reflect.TypeOf(q.row).Elem(); t != nil {
		return strings.ToLower(t.Name())
	}

	return ""
}

func (q *queryer) Prepare() (query string, args []interface{}) {
	t, v := reflect.TypeOf(q.row).Elem(), reflect.ValueOf(q.row).Elem()

	if t == nil {
		return
	}

	var keys []string
	fields := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields[i] = field.Name

		if field.Tag.Get("key") != "true" {
			continue
		}

		args = append(args, v.Field(i).Interface())
		keys = append(keys, fmt.Sprintf("`%v`=?", field.Name))
	}

	return fmt.Sprintf("SELECT %v FROM `%v` WHERE %v", strings.Join(fields, ","), q.TableName(), strings.Join(keys, ",")), args
}

func (q *queryer) Scan(scan func(dest ...interface{}) error) sql.Row {
	if q.row == nil {
		return nil
	}

	if scanner, ok := q.row.(sql.Scanner); ok {
		return scanner.Scan(scan)
	}

	table := q.Clone()

	t, v := reflect.TypeOf(table).Elem(), reflect.ValueOf(table).Elem()

	dest := make([]interface{}, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			dest[i] = new(string)
		default:
			dest[i] = v.Field(i).Addr().Interface()
		}
	}

	if err := scan(dest...); err != nil {
		return nil
	}

	for i := 0; i < len(dest); i++ {
		switch v.Field(i).Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			str, ok := dest[i].(*string)

			if !ok {
				return nil
			}

			convertor, ok := v.Field(i).Interface().(sql.Convertor)

			if ok {
				v.Field(i).Set(reflect.ValueOf(convertor.Convert(*str)))
				continue
			}

			json.Unmarshal([]byte(*str), v.Field(i).Interface())
		}
	}

	return table
}

type inserter struct {
	row sql.Row
}

func (i *inserter) TableName() string {
	if tableName := i.row.TableName(); len(tableName) > 0 {
		return tableName
	}

	if t := reflect.TypeOf(i.row).Elem(); t != nil {
		return strings.ToLower(t.Name())
	}

	return ""
}

func (i *inserter) Prepare() (query string, args []interface{}) {
	t, v := reflect.TypeOf(i.row).Elem(), reflect.ValueOf(i.row).Elem()

	if t == nil {
		return "", nil
	}

	// ---bad
	//var keys, values string

	//for i := 0; i < t.NumField(); i++ {
	//	switch {
	//	case i <= 0:
	//		keys = fmt.Sprintf("%v", t.Field(i).Name)
	//		values = "?"
	//	default:
	//		keys = fmt.Sprintf("%v,`%v`", keys, t.Field(i).Name)
	//		values = fmt.Sprintf("%v,?", values)
	//	}

	//	args = append(args, v.Field(i).Interface())
	//}

	// --just one
	//var keys, values bytes.Buffer

	//for i := 0; i < t.NumField(); i++ {
	//	if i > 0 {
	//		keys.WriteRune(',')
	//		values.WriteRune(',')
	//	}

	//	keys.WriteString(t.Field(i).Name)
	//	values.WriteString("?")
	//	args = append(args, v.Field(i).Interface())
	//}

	// --maybe
	var keys, values []string

	for i := 0; i < t.NumField(); i++ {
		keys = append(keys, t.Field(i).Name)
		values = append(values, "?")

		iface := v.Field(i).Addr().Interface()

		if stringer, ok := iface.(sql.Stringer); ok {
			args = append(args, stringer.ToString())
			continue
		}

		switch v.Field(i).Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			b, err := json.Marshal(iface)

			if err != nil {
				return "", nil
			}

			args = append(args, string(b))
		default:
			args = append(args, v.Field(i).Interface())
		}
	}

	return fmt.Sprintf("INSERT INTO `%v` (%v) VALUES(%v)", i.TableName(), strings.Join(keys, ","), strings.Join(values, ",")), args
}
