package ethyl

type EthereumNetworkRequest struct {
    JsonRpcVersion string `json:"jsonrpc"`
    Method         string `json:"method"`
    Params         []string `json:"params"`
    Id             string `json:"id"`
}

type EthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         string `json:"result"`
}

