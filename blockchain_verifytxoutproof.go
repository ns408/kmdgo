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
/*
Verifies that a proof points to a transaction in a block, returning the transaction it commits to
and throwing an RPC error if the block is not in our best chain

Arguments:
1. "proof"    (string, required) The hex-encoded proof generated by gettxoutproof
*/

package kmdgo

import (
	//"fmt"
	"encoding/json"
	"errors"
)

type VerifyTxOutProof struct {
	Result	[]string `json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
	
}

func (appName AppType) VerifyTxOutProof(pf string) (VerifyTxOutProof, error) {
	query := APIQuery {
		Method:	`verifytxoutproof`,
		Params:	`["`+pf+`"]`,
	}
	//fmt.Println(query)

	var verifytxoutproof VerifyTxOutProof

	verifytxoutproofJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(verifytxoutproofJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifytxoutproofJson), &verifytxoutproof)
		return verifytxoutproof, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(verifytxoutproofJson), &verifytxoutproof)
	return verifytxoutproof, nil
}