/*
 * (C) Copyright 2016 Michael Dowling (http://github.com/michaeldowling) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ethyl


type EthereumNetworkError struct {
	Code    int64 `json:"code"`
	Message string `json:"message"`
}

type EthereumNetworkRequest struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Method         string `json:"method"`
	Params         []interface{} `json:"params"`
	Id             string `json:"id"`
}

type EthereumFilterNetworkRequest struct {
	JsonRpcVersion string `json:"jsonrpc"`
	Method         string `json:"method"`
	Params         []FilterOptions `json:"params"`
	Id             string `json:"id"`
}

type StringResultEthereumNetworkResponse struct {
	Id             string `json:"id"`
	JsonRpcVersion string `json:"jsonrpc"`
	Result         string `json:"result"`
	Error          EthereumNetworkError `json:"error"`
}

type GenericSliceResultEthereumNetworkResponse struct {
	Id             string `json:"is"`
	JsonRpcVersion string `json:"jsonepc"`
	Result         []interface{} `json:"result"`
	Error          EthereumNetworkError `json:"error"`
}

type BooleanResultEthereumNetworkResponse struct {
	Id             string `json:"id"`
	JsonRpcVersion string `json:"jsonrpc"`
	Result         bool   `json:"result"`
	Error          EthereumNetworkError `json:"error"`
}




// for transactions
type TransactionObject struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}

type TransactionObjectResultEthereumNetworkResponse struct {
	Id             string `json:"id"`
	JsonRpcVersion string `json:"jsonrpc"`
	Result         TransactionObject `json:"result"`
}

type TransactionReceiptObject struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress"`
	// Logs []string - TODO
}

type TransactionReceiptObjectResultEthereumNetworkResponse struct {
	Id             string `json:"id"`
	JsonRpcVersion string `json:"jsonrpc"`
	Result         TransactionReceiptObject `json:"result"`
}

type FilterOptions struct {
	FromBlock string `json:"fromBlock,omitempty"`
	ToBlock   string `json:"toBlock,omitempty"`
	Address   string `json:"address,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

type FilterLogObject struct {
	Type             string `json:"type"`
	LogIndex         string `json:"logIndex"`
	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"transactionHash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Address          string `json:"address"`
	Data             string `json:"data"`
	Topics           []string `json:"topics"`
}

type FilterLogObjectResultEthereumNetworkResponse struct {

	Id             string `json:"id"`
	JsonRpcVersion string `json:"jsonrpc"`
	Result         []FilterLogObject `json:"result"`

}