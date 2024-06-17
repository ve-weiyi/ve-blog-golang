package cobrax

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

func ParseFlag(cmd *cobra.Command, res any) error {
	val := reflect.ValueOf(res).Elem()
	reType := reflect.TypeOf(res)

	if reType.Kind() != reflect.Ptr || reType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("res must be a pointer to a struct")
	}

	t := reType.Elem()
	for i := 0; i < t.NumField(); i++ {
		if !val.Field(i).CanSet() {
			return fmt.Errorf("field %s is unexported", t.Field(i).Name)
		}

		structField := t.Field(i)
		// 取tag
		name := structField.Tag.Get("name")
		shorthand := structField.Tag.Get("shorthand")
		value := structField.Tag.Get("value")
		usage := structField.Tag.Get("usage")

		cmd.PersistentFlags().StringVarP(val.Field(i).Addr().Interface().(*string), name, shorthand, value, usage)
	}

	return nil
}
