package getcdv3

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/qmarliu/gopkg/log"
	"github.com/qmarliu/gopkg/utils"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	cc                 resolver.ClientConn
	serviceName        string
	grpcClientConn     *grpc.ClientConn
	cli                *clientv3.Client
	schema             string
	etcdAddr           string
	watchStartRevision int64
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if r.cli == nil {
		return nil, fmt.Errorf("etcd clientv3 client failed, etcd:%s", target)
	}
	r.cc = cc
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//     "%s:///%s"
	prefix := GetPrefix(r.schema, r.serviceName)
	// get key first
	resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err == nil {
		var addrList []resolver.Address
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: string(resp.Kvs[i].Value)})
		}
		r.cc.UpdateState(resolver.State{Addresses: addrList})
		r.watchStartRevision = resp.Header.Revision + 1
		go r.watch(prefix, addrList)
	} else {
		return nil, fmt.Errorf("etcd get failed, prefix: %s", prefix)
	}
	return r, nil
}

func (r *Resolver) watch(prefix string, addrList []resolver.Address) {
	rch := r.cli.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithPrefix())
	for n := range rch {
		flag := 0
		for _, ev := range n.Events {
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, string(ev.Kv.Value)) {
					flag = 1
					addrList = append(addrList, resolver.Address{Addr: string(ev.Kv.Value)})
					log.Debug("", "after add, new list: ", addrList)
				}
			case mvccpb.DELETE:
				log.Debug("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value))
				i := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if i < 0 {
					return
				}
				t := string(ev.Kv.Key)[i+1:]
				log.Debug("remove addr key: ", string(ev.Kv.Key), "value:", string(ev.Kv.Value), "addr:", t)
				if s, ok := remove(addrList, t); ok {
					flag = 1
					addrList = s
					log.Debug("after remove, new list: ", addrList)
				}
			}
		}

		if flag == 1 {
			r.cc.UpdateState(resolver.State{Addresses: addrList})
			log.Debug("update: ", addrList)
		}
	}
}

func (r *Resolver) Scheme() string { return r.schema }

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {}

func (e *EtcdConn) SetDefaultEtcdConfig(serviceName string, defaultPorts []int) {
	e.defaultPorts[serviceName] = defaultPorts
}

func (e *EtcdConn) GetDefaultEtcdConfig(serviceName string) []int {
	conf, ok := e.defaultPorts[serviceName]
	if !ok {
		return conf
	}
	return nil
}

func (e *EtcdConn) NewResolver(operationID string, serviceName string) (*Resolver, error) {
	var r Resolver
	r.serviceName = serviceName
	r.cli = e.rEtcd.cli
	r.schema = e.schema
	r.etcdAddr = e.etcdAddr
	resolver.Register(&r)
	//
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, GetPrefix(e.schema, serviceName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure(), WithRequestMateData())
	if err == nil {
		r.grpcClientConn = conn
	}
	return &r, utils.Wrap(err, "")
}

func (e *EtcdConn) getConn(operationID, serviceName string) *grpc.ClientConn {
	e.rwNameResolverMutex.RLock()
	defer e.rwNameResolverMutex.RUnlock()
	r, ok := e.nameResolver[e.schema+serviceName]
	if ok {
		return r.grpcClientConn
	}

	r, ok = e.nameResolver[e.schema+serviceName]
	if ok {
		return r.grpcClientConn
	}

	r, err := e.NewResolver(operationID, serviceName)
	if err != nil {
		return nil
	}
	e.nameResolver[e.schema+serviceName] = r
	return r.grpcClientConn
}

func (e *EtcdConn) getConfigConn(serviceName string, operationID string) *grpc.ClientConn {
	defaultPorts := e.GetDefaultEtcdConfig(serviceName)
	if len(defaultPorts) == 0 {
		log.Error(operationID, "len(configPortList) == 0  ")
		return nil
	}
	target := e.host + ":" + utils.Int32ToString(int32(defaultPorts[0]))
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Error(operationID, "grpc.Dail failed ", err.Error())
		return nil
	}
	return conn
}

func (e *EtcdConn) GetConn(operationID, serviceName string) *grpc.ClientConn {
	con := e.getConn(operationID, serviceName)
	if con != nil {
		return con
	}
	log.NewWarn(operationID, utils.GetSelfFuncName(), "conn is nil !!!!! use config conn", e.etcdAddr, serviceName, operationID)
	con = e.getConfigConn(serviceName, operationID)
	return con
}

func (e *EtcdConn) GetDefaultGatewayConn4Unique(operationID string, serviceName string) []*grpc.ClientConn {
	e.isUpdateStartMtx.Lock()
	defer e.isUpdateStartMtx.Unlock()
	if e.isUpdateStart == false {
		e.conn4UniqueList = e.getConn4Unique(operationID, serviceName)
		go func() {
			for {
				select {
				case <-time.After(time.Second * time.Duration(30)):
					e.conn4UniqueListMtx.Lock()
					e.conn4UniqueList = e.getConn4Unique(operationID, serviceName)
					e.conn4UniqueListMtx.Unlock()
				}
			}
		}()
	}
	e.isUpdateStart = true
	if len(e.conn4UniqueList) == 0 {
		return e.getDefaultGatewayConn4UniqueFromConfig(operationID, serviceName)
	}
	return e.conn4UniqueList
}

func (e *EtcdConn) getDefaultGatewayConn4UniqueFromConfig(operationID, serviceName string) []*grpc.ClientConn {
	var conns []*grpc.ClientConn
	configPortList := e.GetDefaultEtcdConfig(serviceName)
	for _, port := range configPortList {
		target := e.host + ":" + utils.Int32ToString(int32(port))
		log.Info(operationID, "rpcRegisterIP ", e.host, " port ", configPortList, " grpc target: ", target, " serviceName: ", "msgGateway")
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			log.Error(operationID, "grpc.Dail failed ", err.Error())
			continue
		}
		conns = append(conns, conn)
	}
	return conns
}

func (e *EtcdConn) getConn4Unique(operationID, serviceName string) []*grpc.ClientConn {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//     "%s:///%s"
	prefix := GetPrefix4Unique(e.schema, serviceName)
	resp, err := e.rEtcd.cli.Get(ctx, prefix, clientv3.WithPrefix())
	//  "%s:///%s:ip:port"   -> %s:ip:port
	allService := make([]string, 0)
	if err == nil {
		for i := range resp.Kvs {
			k := string(resp.Kvs[i].Key)

			b := strings.LastIndex(k, "///")
			k1 := k[b+len("///"):]

			e := strings.Index(k1, "/")
			k2 := k1[:e]
			allService = append(allService, k2)
		}
	} else {
		e.rEtcd.cli.Close()
		return nil
	}
	e.rEtcd.cli.Close()

	allConn := make([]*grpc.ClientConn, 0)
	for _, v := range allService {
		r := e.getConn(operationID, v)
		allConn = append(allConn, r)
	}
	return allConn
}

//var (
//	service2pool   = make(map[string]*Pool)
//	service2poolMu sync.Mutex
//)

//func GetconnFactory(schema, etcdaddr, servicename string) (*grpc.ClientConn, error) {
//	c := getConn(schema, etcdaddr, servicename, "0")
//	if c != nil {
//		return c, nil
//	} else {
//		return c, fmt.Errorf("GetConn failed")
//	}
//}

//func GetConnPool(schema, etcdaddr, servicename string) (*ClientConn, error) {
//	//get pool
//	p := NewPool(schema, etcdaddr, servicename)
//	//poo->get
//
//	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1000*time.Millisecond))
//
//	c, err := p.Get(ctx)
//	//log.Info("", "Get ", err)
//	return c, err
//
//}

//func NewPool(schema, etcdaddr, servicename string) *Pool {
//
//	if _, ok := service2pool[schema+servicename]; !ok {
//		//
//		service2poolMu.Lock()
//		if _, ok1 := service2pool[schema+servicename]; !ok1 {
//			p, err := New(GetconnFactory, schema, etcdaddr, servicename, 5, 10, 1)
//			if err == nil {
//				service2pool[schema+servicename] = p
//			}
//		}
//		service2poolMu.Unlock()
//	}
//
//	return service2pool[schema+servicename]
//}
