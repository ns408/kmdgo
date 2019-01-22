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
)

type ZListOperationIDs struct {
	Result []string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}

func (appName AppType) ZListOperationIDs(status string) (ZListOperationIDs, error) {
	query := APIQuery {
		Method:	`z_listoperationids`,
		Params:	`["`+status+`"]`,
	}
	//fmt.Println(query)

	var z_listoperationids ZListOperationIDs

	z_listoperationidsJson := appName.APICall(query)
	//fmt.Println(z_listoperationidsJson)

	var result APIResult

	json.Unmarshal([]byte(z_listoperationidsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listoperationidsJson), &z_listoperationids)
		return z_listoperationids, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listoperationidsJson), &z_listoperationids)
	return z_listoperationids, nil
}