package cluster

import (
	"github.com/hashicorp/memberlist"
	"io/ioutil"
	"stathat.com/c/consistent"
	"time"
)

type Node interface {
	ShouldProcess(key string) (string, bool)
	Members() []string //该函数被consistent实现
	Addr() string
}
type node struct {
	*consistent.Consistent
	addr string
}

func New(addr, cluster string) (Node, error) {
	//创建gossip新节点的config
	config := memberlist.DefaultLANConfig()
	config.Name = addr
	config.BindAddr = addr
	config.LogOutput = ioutil.Discard
	//创建新节点
	mbl, err := memberlist.Create(config)
	if err != nil {
		return nil, err
	}
	if cluster == "" {
		cluster = addr
	}
	existing := []string{cluster}
	//连接到集群
	_, err = mbl.Join(existing)
	if err != nil {
		return nil, err
	}
	//创建一致性哈希的节点实例
	circle := consistent.New()
	//设置虚拟节点数量
	circle.NumberOfReplicas = 256
	go func() {
		for {
			m := mbl.Members()
			nodes := make([]string, len(m))
			for i, n := range m {
				nodes[i] = n.Name
			}
			//每隔1s将集群节点列表m更新到circle中
			circle.Set(nodes)
			time.Sleep(time.Second)
		}
	}()
	return &node{circle, addr}, nil

}

func (n *node) ShouldProcess(key string) (string, bool) {
	addr, _ := n.Get(key)
	return addr, addr == n.addr
}

func (n *node) Addr() string {
	return n.addr
}
