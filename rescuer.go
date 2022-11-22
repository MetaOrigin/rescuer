package main

import (
	"fmt"
	
	"github.com/jeyrce/rescuer/cmd"
)

func init() {

}

func main() {
	if err := cmd.Rescuer.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
