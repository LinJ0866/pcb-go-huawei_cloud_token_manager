// main.go
package main

import (
	"fmt"
	"go-huawei_cloud_token_manager/config"
	"go-huawei_cloud_token_manager/internal/database"
	routes "go-huawei_cloud_token_manager/internal/routes/v1"
	"runtime"
)

func printHello() {
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println("# start go-huawei_cloud_token_manager with", runtime.Version())
	fmt.Println("#")
	fmt.Println("# *****  *****  *        *    *   *  ***** ")
	fmt.Println("# *      *   *  *       * *   **  *  *     ")
	fmt.Println("# *  **  *   *  *      *****  * * *  *  ** ")
	fmt.Println("# *   *  *   *  *      *   *  *  **  *   * ")
	fmt.Println("# *****  *****  *****  *   *  *   *  ***** ")
	fmt.Println("#")
	fmt.Println("#-----------------------------------------------------#")
	fmt.Println()
}

func registerServer() {
	config.Init()
	printHello()
	database.InitMysql()
	routes.InitRouter()
}

func main() {
	registerServer()
}
