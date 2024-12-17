package main

import (
	"ExportIT/cmd"
)

func main() {
	cmd.Execute()
	// funcs, err := export.GetExportedFuncs()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	//
	// err = os.MkdirAll("cmd", 0777)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return
	// }
	// for _, fn := range funcs {
	// 	err := export.GenerateCustomNamedCommand(fn)
	// 	if err != nil {
	// 		fmt.Printf("err: %v\n", err)
	// 	}
	// }
	// err := ExportIT.ExecCli("mom",
	// 	"",
	// 	"",
	// 	ExportIT.CliNamedParam(BatchFiles),
	// )
	// fmt.Printf("err: %v\n", err)
}
