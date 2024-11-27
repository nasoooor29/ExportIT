package ExportIT

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func cliNamedParam(funcAny any) (*cobra.Command, error) {
	fNameAndPkg := getFunctionName(funcAny)
	fnAst, err := getFuncByName(fNameAndPkg)
	split := strings.Split(fNameAndPkg, ".")
	if err != nil {
		return nil, err
	}
	cmd := &cobra.Command{
		Use:   split[len(split)-1],
		Short: fNameAndPkg,
		Long:  fnAst.Doc.Text(),
	}

	_, err = iterParamsByFunc(
		funcAny,
		fnAst,
		func(paramName string, short string, val *any) error {
			err := addParamsToFlagset(cmd.Flags(), paramName, short, val)
			if err != nil {
				return errors.Join(err, fmt.Errorf("type "+reflect.TypeOf(val).Name()))
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	cmd.RunE = func(c *cobra.Command, args []string) error {
		values, err := iterParamsByFunc(
			funcAny,
			fnAst,
			func(paramName string, short string, val *any) error {
				return getParamFromFlag(c.Flags(), paramName, val)
			})
		if err != nil {
			return err
		}

		values = reflect.ValueOf(funcAny).Call(values)
		for _, v := range values {
			if isError(v.Type()) {
				return err
			}
			fmt.Printf("v.Interface(): %v\n", v.Interface())
		}

		return nil
	}

	return cmd, nil

}

func iterParamsByFunc(
	funcAny any,
	fnAst *ast.FuncDecl,
	f func(paramName string, short string, val *any) error,
) ([]reflect.Value, error) {
	fn := reflect.TypeOf(funcAny)
	if fn.Kind() != reflect.Func {
		fmt.Printf("t.Name(): %v\n", fn.Name())
		return nil, fmt.Errorf("param is not a function")
	}
	params := getFuncParams(fnAst)
	values := []reflect.Value{}
	for i, param := range params {
		pType := fn.In(i)
		val := reflect.New(pType).Elem().Interface()
		short := getShorthand(param)
		err := f(param, short, &val)
		if err != nil {
			return nil, err
		}
		values = append(values, reflect.ValueOf(val))
	}
	return values, nil
}

func addParamsToFlagset(set *pflag.FlagSet, pName, short string, val *any) error {
	switch (*val).(type) {
	case string:
		ca := (*val).(string)
		set.StringVarP(&ca, pName, short, "", "")
		*val = ca
	case int:
		ca := (*val).(int)
		set.IntVarP(&ca, pName, short, -1, "")
		*val = ca
	case bool:
		ca := (*val).(bool)
		set.BoolVarP(&ca, pName, short, false, "")
		*val = ca
	default:
		return fmt.Errorf("unspported type")
	}
	return nil
}

func getParamFromFlag(set *pflag.FlagSet, pName string, val *any) error {
	switch (*val).(type) {
	case string:
		v, err := set.GetString(pName)
		*val = v
		if err != nil {
			return err
		}
	case int:
		v, err := set.GetInt(pName)
		if err != nil {
			return err
		}
		*val = v
	case bool:
		v, err := set.GetBool(pName)
		if err != nil {
			return err
		}
		*val = v
	default:
		return fmt.Errorf("unspported type")
	}
	return nil
}

func getShorthand(s string) string {
	return string(s[0])
}
