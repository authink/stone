package db

import (
	"reflect"
	"strings"

	"github.com/authink/inkstone/orm/model"
)

// todo support embed struct
func Bind(tb Table, m model.Identifier) {
	vTable := reflect.ValueOf(tb).Elem()

	t := reflect.TypeOf(m).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		dbTag := f.Tag.Get("db")

		var col string
		if dbTag != "" {
			col = strings.Split(dbTag, ",")[0]
		}

		if col == "" {
			col = strings.ToLower(f.Name)
		}

		v := vTable.FieldByName(f.Name)
		if v.IsValid() && v.CanSet() {
			v.SetString(col)
		}
	}
}
