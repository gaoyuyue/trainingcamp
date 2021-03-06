package cg

import (
	"encoding/json"
	"errors"
	"github.com/trainingcamp/day3/sample-ipc/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	bytes, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addPlayer", string(bytes))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removePlayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listPlayer", params)
	if resp.Code == "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content:message}
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp,_ := client.Call("broadcast", string(bytes))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}