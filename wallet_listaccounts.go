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
package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)


type ListAccounts struct {
	Result map[string]float64	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ListAccounts(minconf int, incwch bool) (ListAccounts, error) {
	query := APIQuery {
		Method:	`listaccounts`,
		Params:	`[`+strconv.Itoa(minconf)+`, `+strconv.FormatBool(incwch)+`]`,
	}
	//fmt.Println(query)

	var listaccounts ListAccounts

	listaccountsJson := appName.APICall(query)
	//fmt.Println(listaccountsJson)

	var result APIResult

	json.Unmarshal([]byte(listaccountsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listaccountsJson), &listaccounts)
		return listaccounts, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listaccountsJson), &listaccounts)
	return listaccounts, nil
}