package getcdv3

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/qmarliu/gopkg/log"
	"github.com/qmarliu/gopkg/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type EtcdConn struct {
	rEtcd                                                   *RegEtcd
	schema, serviceName, etcdAddr, host, userName, password string
	port, ttl                                               int

	// targetRpcDefaultPorts
	nameResolver        map[string]*Resolver
	defaultPorts        map[string][]int
	rwNameResolverMutex sync.RWMutex

	conn4UniqueList    []*grpc.ClientConn
	conn4UniqueListMtx sync.RWMutex
	isUpdateStart      bool
	isUpdateStartMtx   sync.RWMutex
}

type RegEtcd struct {
	cli    *clientv3.Client
	ctx    context.Context
	cancel context.CancelFunc
	key    string
}

func NewEtcdConn(schema, etcdAddr, host, userName, password string, port, ttl int) (conn *EtcdConn, err error) {
	if host == "" {
		host, err = utils.GetLocalIP()
		if err != nil {
			return nil, err
		}
	}
	return &EtcdConn{
		schema:   schema,
		etcdAddr: etcdAddr,
		host:     host,
		port:     port,
		ttl:      ttl,
		userName: userName,
		password: password,
	}, nil
}

// "%s:///%s/" ->  "%s:///%s:ip:port"
func (e *EtcdConn) RegisterEtcd4Unique(operationID, serviceName string) error {
	serviceName = serviceName + ":" + net.JoinHostPort(e.host, strconv.Itoa(e.port))
	return e.RegisterEtcd(operationID, serviceName)
}

func (e *EtcdConn) GetTarget(serviceName string) string {
	return GetPrefix4Unique(e.schema, serviceName) + ":" + net.JoinHostPort(e.host, strconv.Itoa(e.port)) + "/"
}

// etcdAddr separated by commas
func (e *EtcdConn) RegisterEtcd(operationID, serviceName string) error {
	e.serviceName = serviceName
	args := e.schema + " " + e.etcdAddr + " " + e.host + " " + serviceName + " " + utils.Int32ToString(int32(e.port))
	ttl := e.ttl * 3
	cli, err := clientv3.New(clientv3.Config{
		Username:  e.userName,
		Password:  e.password,
		Endpoints: strings.Split(e.etcdAddr, ","), DialTimeout: 5 * time.Second})

	log.Info(operationID, "RegisterEtcd args: ", args, ttl)
	if err != nil {
		log.Error(operationID, "clientv3.New failed ", args, ttl, err.Error())
		return fmt.Errorf("create etcd clientv3 client failed, errmsg:%v, etcd addr:%s", err, e.etcdAddr)
	}
	//lease
	ctx, cancel := context.WithCancel(context.Background())
	resp, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		log.Error(operationID, "Grant failed ", err.Error(), ctx, ttl)
		return fmt.Errorf("grant failed")
	}
	log.Info(operationID, "Grant ok, resp ID ", resp.ID)

	//  schema:///serviceName/ip:port ->ip:port
	serviceValue := net.JoinHostPort(e.host, strconv.Itoa(e.port))
	serviceKey := GetPrefix(e.schema, serviceName) + serviceValue

	//set key->value
	if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
		log.Error(operationID, "cli.Put failed ", err.Error(), ctx, args, resp.ID)
		return fmt.Errorf("put failed, errmsg:%v， key:%s, value:%s", err, serviceKey, serviceValue)
	}

	//keepalive
	kresp, err := cli.KeepAlive(ctx, resp.ID)
	if err != nil {
		log.Error(operationID, "KeepAlive failed ", err.Error(), args, resp.ID)
		return fmt.Errorf("keepalive failed, errmsg:%v, lease id:%d", err, resp.ID)
	}
	log.Info(operationID, "RegisterEtcd ok ", args)

	go func() {
		for {
			select {
			case pv, ok := <-kresp:
				if ok == true {
					log.Debug(operationID, "KeepAlive kresp ok", pv, args)
				} else {
					log.Error(operationID, "KeepAlive kresp failed ", pv, args)
					t := time.NewTicker(time.Duration(ttl/2) * time.Second)
					for {
						select {
						case <-t.C:
						}
						ctx, _ := context.WithCancel(context.Background())
						resp, err := cli.Grant(ctx, int64(ttl))
						if err != nil {
							log.Error(operationID, "Grant failed ", err.Error(), args)
							continue
						}

						if _, err := cli.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
							log.Error(operationID, "etcd Put failed ", err.Error(), args, " resp ID: ", resp.ID)
							continue
						} else {
							log.Info(operationID, "etcd Put ok ", args, " resp ID: ", resp.ID)
						}
					}
				}
			}
		}
	}()
	e.rEtcd = &RegEtcd{ctx: ctx,
		cli:    cli,
		cancel: cancel,
		key:    serviceKey}
	return nil
}

func (e *EtcdConn) UnRegisterEtcd() error {
	//delete
	e.rEtcd.cancel()
	_, err := e.rEtcd.cli.Delete(e.rEtcd.ctx, e.rEtcd.key)
	return err
}
