package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KubeOperator/kobe/api"
	"github.com/KubeOperator/kobe/pkg/constant"
	"google.golang.org/grpc"
	"os"
)

type Result map[string]map[string]interface{}

type kobeInventoryProvider struct {
	host string
	port int
}

func NewKobeInventoryProvider(host string, port int) *kobeInventoryProvider {
	return &kobeInventoryProvider{
		host: host,
		port: port,
	}
}
func (r Result) String() string {
	b, err := json.Marshal(&r)
	if err != nil {
		return ""
	}
	return string(b)
}

func (kip kobeInventoryProvider) getInventory(id string) (*api.Inventory, error) {

	conn, err := kip.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := &api.GetInventoryRequest{
		Id: id,
	}
	resp, err := client.GetInventory(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp.Item, nil
}

func (kip kobeInventoryProvider) ListHandler() (Result, error) {
	id, err := kip.getInventoryId()
	if err != nil {
		return nil, err
	}
	inventory, _ := kip.getInventory(id)
	groups := make(map[string]map[string]interface{})
	all := api.Group{
		Name:     "all",
		Hosts:    []string{},
		Children: []string{},
		Vars:     inventory.Vars,
	}
	for _, group := range inventory.Groups {
		m := parseGroupToMap(*group)
		groups[group.Name] = m
		all.Children = append(all.Children, group.Name)
	}
	meta := map[string]interface{}{}
	hostVars := map[string]interface{}{}
	for _, host := range inventory.Hosts {
		vars := make(map[string]interface{})
		all.Hosts = append(all.Hosts, host.Name)
		hostVars[host.Name] = map[string]interface{}{
			"ansible_ssh_host": host.Ip,
			"ansible_ssh_port": host.Port,
			"ansible_ssh_user": host.User,
			"ansible_ssh_pass": host.Password,
		}
		if host.Vars != nil {
			for k, v := range host.Vars {
				vars[k] = v
			}
			hostVars["vars"] = vars
		}
	}
	groups[all.Name] = parseGroupToMap(all)
	meta["hostvars"] = hostVars
	groups["_meta"] = meta
	return groups, nil
}

func parseGroupToMap(group api.Group) map[string]interface{} {
	m := map[string]interface{}{}
	if group.Hosts != nil {
		m["hosts"] = group.Hosts
	}
	if group.Children != nil {
		m["children"] = group.Children
	}
	if group.Vars != nil {
		m["vars"] = group.Vars
	}
	return m
}

func (kip kobeInventoryProvider) getInventoryId() (string, error) {
	id := os.Getenv(constant.TaskEnvKey)
	if id == "" {
		return "", errors.New(fmt.Sprintf("invalid id: %s", id))
	}
	return id, nil
}

func (kip kobeInventoryProvider) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", kip.host, kip.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
