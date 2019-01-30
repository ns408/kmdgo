 //  Copyright © 2018-2019 Satinderjit Singh.
 //
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
 //  the top-level directory of this distribution for the individual copyright
 //  holder information and the developer policies on copyright and licensing.
 //
 //  Unless otherwise agreed in a custom licensing agreement, no part of the
 //  kmdgo software, including this file may be copied, modified, propagated.
 //  or distributed except according to the terms contained in the LICENSE file
 //
 //  Removal or modification of this copyright notice is prohibited.
package main

import (
	"fmt"
    "log"
    "github.com/satindergrewal/kmdgo"
)

func main() {
    // Define appName type from kmdgo package
    var appName kmdgo.AppType

    // Define appname variable. The name value must be the matching value of it's data directory name.
    // Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
    appName = `komodo`

    // define the variable with GetInfo struct from pacakge kmdgo
	var info kmdgo.GetInfo

    // Get output of the ResultGetinfo function and store it to GetInfo struct variable
    info, err := appName.GetInfo()
    if err != nil {
        fmt.Printf("Code: %v\n", info.Error.Code)
        fmt.Printf("Message: %v\n\n", info.Error.Message)
        log.Fatalln("Err happened", err)
    }

    // Can print and use the struct variable outputs in further code logic. Check GetInfo struct in package file.
    fmt.Println(info)
    fmt.Println(info.Result)
    
    fmt.Println("Version:", info.Result.Version)
    fmt.Println("Balance:", info.Result.Balance)
    fmt.Println("Blocks:", info.Result.Blocks)
    fmt.Println("Name:", info.Result.Name)
    fmt.Println("Connections:", info.Result.Connections)
    fmt.Println("Difficulty:", info.Result.Difficulty)
    fmt.Println("Magic:", info.Result.Magic)
    fmt.Println(info.Error)
    fmt.Println(info.ID)
}