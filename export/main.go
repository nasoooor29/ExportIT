package export

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GenerateCustomNamedCommand(fn ExportedFunc) error {
	fname := fn.Func.Name.Name
	toFill := GetCustomCommandTemplate()
	cliParams := ""
	funcGlobVars := ""
	funcParams := ""

	shorts := GenerateShorthands(GetFuncParams(fn.Func))
	for i, p := range fn.Func.Type.Params.List {
		name := p.Names[0].Name
		t := fmt.Sprintf("%v", p.Type)
		sh := shorts[i]
		paramName := fmt.Sprintf("%vCmd_%v", fname, name)
		funcGlobVars += fmt.Sprintf("\t%v %v\n", paramName, t)
		funcParams += "\t\t\t" + paramName + ",\n"

		cliParams += fmt.Sprintf(
			"\t"+`%vCmd.Flags().%vVarP(&%v, "%v", "%v", %v, "")`+"\n",
			fname,
			CapitalizeFirst(t),
			paramName,
			name,
			string(sh),
			GenerateDefaultValues(t),
		)

	}

	path := "cmd/" + fname + ".go"

	toFill = strings.ReplaceAll(toFill, "$FUNC_NAME", fname)
	toFill = strings.ReplaceAll(toFill, "$FUNC_COMMENT", "`"+fn.Func.Doc.Text()+"`")
	toFill = strings.ReplaceAll(toFill, "$PKG_PATH", fn.Pkg.PkgPath)
	toFill = strings.ReplaceAll(toFill, "$PKG_NAME", fn.Pkg.Name)
	toFill = strings.ReplaceAll(toFill, "$CLI_PARAMS", cliParams)
	toFill = strings.ReplaceAll(toFill, "$CLI_VARS", funcGlobVars)
	toFill = strings.ReplaceAll(toFill, "$FUNC_PARAMS", funcParams)
	err := os.WriteFile(path, []byte(toFill), 0777)
	if err != nil {
		return err
	}

	cmdFmt := exec.Command("gofmt", "-w", path)
	cmdFmt.Stdout = os.Stdout
	cmdFmt.Stderr = os.Stderr
	if err := cmdFmt.Run(); err != nil {
		return err
	}
	return nil
}
