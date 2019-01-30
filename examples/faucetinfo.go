/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `TXSCLCC`

	var fctifo kmdgo.FaucetInfo

	fctifo, err := appName.FaucetInfo()
	if err != nil {
		fmt.Printf("Code: %v\n", fctifo.Error.Code)
		fmt.Printf("Message: %v\n\n", fctifo.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("fctifo value", fctifo)
	fmt.Println("-------")
	fmt.Println(fctifo.Result)

	fmt.Println("Result: ", fctifo.Result.Result)
	fmt.Println("Name: ", fctifo.Result.Name)
	fmt.Println("Funding: ", fctifo.Result.Funding)
}