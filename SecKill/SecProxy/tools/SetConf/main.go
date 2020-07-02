package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

const (
	EtcdKey = "/yanglige/backend/secskill/product"
)
type SecInfoConf struct {
	ProductId  int
	StartTime  int
	EndTime    int
	Status     int
	TotalCount int
	LeftCount  int
}
func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connec succ")
	defer cli.Close()
	var SecInfoConfArr []SecInfoConf
	SecInfoConfArr = append(SecInfoConfArr,
		SecInfoConf{
			ProductId: 1031,
			StartTime: 1593536880,
			EndTime: 1593537000,
			Status: 0,
			TotalCount: 1000,
			LeftCount: 1000,
		},
		)
	SecInfoConfArr = append(SecInfoConfArr,
		SecInfoConf{
			ProductId: 1029,
			StartTime: 1593536880,
			EndTime: 1593537000,
			Status: 0,
			TotalCount: 2000,
			LeftCount: 1000,
		},
		)

    data, err := json.Marshal(SecInfoConfArr)
    if err != nil {
    	fmt.Println("json failed", err)
		return
	}
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)


    _, err = cli.Put(ctx, EtcdKey, string(data))
    cancel()
    if err != nil {
    	fmt.Println("put failed, err:", err)
		return
	}

    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    resp, err := cli.Get(ctx, EtcdKey)
    cancel()
    if err != nil {
    	fmt.Println("get failed, err:", err)
		return
	}
    for _, ev := range resp.Kvs {
    	fmt.Printf("%s: %s\n", ev.Key, ev.Value)

	}

}
func main() {
	SetLogConfToEtcd()
}