package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Version string `json:"version"`
	Wallet  string `json:"wallet"`
	Account string `json:"account"`
	Node    struct {
		Version                    string `json:"version"`
		PeeringPort                string `json:"peering_port"`
		BootstrapFractionNumerator string `json:"bootstrap_fraction_numerator"`
		ReceiveMinimum             string `json:"receive_minimum"`
		Logging                    struct {
			Version             string `json:"version"`
			Ledger              string `json:"ledger"`
			LedgerDuplicate     string `json:"ledger_duplicate"`
			Vote                string `json:"vote"`
			Network             string `json:"network"`
			NetworkMessage      string `json:"network_message"`
			NetworkPublish      string `json:"network_publish"`
			NetworkPacket       string `json:"network_packet"`
			NetworkKeepalive    string `json:"network_keepalive"`
			NodeLifetimeTracing string `json:"node_lifetime_tracing"`
			InsufficientWork    string `json:"insufficient_work"`
			LogRPC              string `json:"log_rpc"`
			BulkPull            string `json:"bulk_pull"`
			WorkGenerationTime  string `json:"work_generation_time"`
			LogToCerr           string `json:"log_to_cerr"`
			MaxSize             string `json:"max_size"`
		} `json:"logging"`
		WorkPeers                    string   `json:"work_peers"`
		PreconfiguredPeers           []string `json:"preconfigured_peers"`
		PreconfiguredRepresentatives []string `json:"preconfigured_representatives"`
		InactiveSupply               string   `json:"inactive_supply"`
		PasswordFanout               string   `json:"password_fanout"`
		IoThreads                    string   `json:"io_threads"`
		WorkThreads                  string   `json:"work_threads"`
		EnableVoting                 string   `json:"enable_voting"`
		BootstrapConnections         string   `json:"bootstrap_connections"`
		CallbackAddress              string   `json:"callback_address"`
		CallbackPort                 string   `json:"callback_port"`
		CallbackTarget               string   `json:"callback_target"`
	} `json:"node"`
	RPC struct {
		Address              string `json:"address"`
		Port                 string `json:"port"`
		EnableControl        string `json:"enable_control"`
		FrontierRequestLimit string `json:"frontier_request_limit"`
		ChainRequestLimit    string `json:"chain_request_limit"`
	} `json:"rpc"`
	RPCEnable    string `json:"rpc_enable"`
	OpenclEnable string `json:"opencl_enable"`
	Opencl       struct {
		Platform string `json:"platform"`
		Device   string `json:"device"`
		Threads  string `json:"threads"`
	} `json:"opencl"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
