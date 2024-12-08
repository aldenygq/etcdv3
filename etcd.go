package etcdv3
import (
    "context"
    "go.etcd.io/etcd/client/v3"
    "time"
)

type EtcdClient struct {
    Client *clientv3.Client
}

func NewEtcdClient(etcdEndpoints []string) (*EtcdClient, error) {
    client, err := clientv3.New(clientv3.Config{
        Endpoints:   etcdEndpoints,
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        return nil, err
    }
    return &EtcdClient{Client: client}, nil
}

//服务注册
func (ec *EtcdClient) RegisterService(key, value string, ttl int64) (clientv3.LeaseID, error) {
    lease, err := ec.Client.Grant(context.TODO(), ttl)
    if err != nil {
        return 0, err
    }
    _, err = ec.Client.Put(context.TODO(), key, value, clientv3.WithLease(lease.ID))
    if err != nil {
        ec.Client.Revoke(context.TODO(), lease.ID)
        return 0, err
    }
    return lease.ID, nil
}
//服务发现
func (ec *EtcdClient) DiscoverService(key string) (interface{},error) {
    var kvmap map[string]interface{} = make(map[string]interface{},0)
    getResp, err := ec.Client.Get(context.TODO(), key, clientv3.WithPrefix())
    if err != nil {
        return nil,err
    }
    for _, kv := range getResp.Kvs {
        if _,ok := kvmap[string(kv.Key)];!ok {
            kvmap[string(kv.Key)] = string(kv.Value)
        }
    }

    return kvmap,nil
}

//监听服务
func (ec *EtcdClient) WatchService(serviceName string,callback func(*clientv3.Event)) {
    rch := ec.Client.Watch(context.Background(), serviceName, clientv3.WithPrefix())
    for wresp := range rch {
        for _, ev := range wresp.Events {
            callback(ev)
        }
    }
}
//租约续期
func (ec *EtcdClient) LeaseKeepAlive(leaseID clientv3.LeaseID) (int64,error) {
    _, err := ec.Client.KeepAlive(context.TODO(), leaseID)
    if err != nil {
        return 0,err
    }
    ttl, err := ec.Client.TimeToLive(context.Background(), leaseID, clientv3.WithAttachedKeys())
    if err != nil {
        return 0,err
    }
    return ttl.TTL,nil
}

//服务注销
func (ec *EtcdClient) DeregisterService(serviceKey string)  error {
    _, err := ec.Client.Delete(context.Background(), serviceKey)
    if err != nil {
        return err
    }
    return nil
}
