package etcdv3

import (
    "testing"
    "fmt"
    "encoding/json"
)

func Test_NewEtcdClient(t *testing.T) {
    var endpoints []string = []string{"127.0.0.1:2379"}
    client,err := NewEtcdClient(endpoints)
    if err != nil {
        fmt.Printf("new etcd v3 client failed:%v\n",err)
        return
    }
    if client == nil {
        fmt.Printf("new etcd v3 client is null")
        return
    }
    fmt.Printf("new etcd v3 success\n")
}

func Test_RegisterService(t *testing.T) {
    var endpoints []string = []string{"127.0.0.1:2379"}
    client,err := NewEtcdClient(endpoints)
    if err != nil {
        fmt.Printf("new etcd v3 client failed:%v\n",err)
        return
    }
    if client == nil {
        fmt.Printf("new etcd v3 client is null")
        return
    }
    fmt.Printf("new etcd v3 success\n")

    id,err := client.RegisterService("sds/sds.lta.svc","172.16.80.23:8080,172.16.80.23:8080",600)
    if err != nil {
        fmt.Printf("register service failed:%v\n",err)
        return
    }

    fmt.Printf("id:%v\n",id)
}

func Test_DiscoverService(t *testing.T) {
     var endpoints []string = []string{"127.0.0.1:2379"}
     client,err := NewEtcdClient(endpoints)
     if err != nil {
         fmt.Printf("new etcd v3 client failed:%v\n",err)
         return
     }
     if client == nil {
         fmt.Printf("new etcd v3 client is null")
         return
     }
     fmt.Printf("new etcd v3 success\n")
    key := "sds/sds.lta.svc"
    result,err := client.DiscoverService(key)
    if err != nil {
        fmt.Printf("discover service failed:%v\n",err)
        return
    }
    re,_:= json.Marshal(result)
    fmt.Printf("result:%v\n",string(re))
}
func Test_DeregisterService(t *testing.T) {
   var endpoints []string = []string{"127.0.0.1:2379"}
   client,err := NewEtcdClient(endpoints)
   if err != nil {
       fmt.Printf("new etcd v3 client failed:%v\n",err)
       return
   }
   if client == nil {
       fmt.Printf("new etcd v3 client is null")
       return
   }
   fmt.Printf("new etcd v3 success\n")
  key := "sds/sds.lta.svc"
    err = client.DeregisterService(key)
    if err != nil {
        fmt.Printf("delete service failed:%v\n",err)
        return
    }
    fmt.Printf("delete service success\n")
}
