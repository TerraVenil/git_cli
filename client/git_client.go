package client

import (
	"log"
	"net/rpc"

	"./../shared"
)

type GitClient struct {
	Client *rpc.Client
}

func (t *GitClient) GetVersion() shared.Version {
	var reply shared.Version
	err := t.Client.Call("GitAPI.GetVersion", struct{}{}, &reply)
	if err != nil {
		log.Fatal("git api error:", err)
	}
	return reply
}
