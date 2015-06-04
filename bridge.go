package bridge

import (
	"net"
	"sync"

	"github.com/docker/libnetwork/ipallocator"
	"github.com/docker/libnetwork/portmapper"
	"github.com/docker/libnetwork/sandbox"
	"github.com/docker/libnetwork/types"
)

const (
	networkType             = "bridge"
	vethPrefix              = "veth"
	vethLen                 = 7
	containerVethPrefix     = "eth"
	maxAllocatePortAttempts = 10
	ifaceID                 = 1
)

var (
	ipAllocator *ipallocator.IPAllocator
	portMapper  *portmapper.PortMapper
)

type configuration struct {
	EnableIPForwarding bool
}

type networkConfiguration struct {
	BridgeName          string
	AddressIPv4         net.IPNet
	FixedCIDR           net.IPNet
	EnableIPTables      bool
	EnableIPMasquerade  bool
	EnableICC           bool
	Mtu                 int
	DefaultGatewayIPv4  net.IP
	DefaultBindingIP    net.IP
	AllowDefaultBridge  bool
	EnableUserLandProxy bool
}

// 为什么端口会在endpoint上，端口的资源应属于sandbox（理应在sandbox.info中）
type endpointConfiguration struct {
	MacAddress   net.HardwareAddr
	PortBindings []types.PortBinding
	ExposePorts  []types.ExposePort
}

// link的父子关系用endpoint来表示？服务访问点
type containerConfiguration struct {
	ParentEndpoints   []string
	ClildrenEndpoints []string
}

// 代表bridge网桥吗? 同样为什么有portbinding\ exposeports  和 portmapping
type bridgeEndpoint struct {
	id              types.UUID
	intf            sandbox.Interface
	macAddress      net.HardwareAddr
	config          *endpointConfiguration
	containerConfig *containerConfiguration
	portMappering   []types.PortBinding
}

// 一个网络中一个网桥（bridge）、网桥中有多个endpoints(bridgeendpoints到底具体指什么呢？ vethpair)
type bridgeNetwork struct {
	id        types.UUID
	bridge    *bridgeInterface
	config    *networkConfiguration
	endpoints map[types.UUID]*bridgeEndpoint
	sync.Mutex
}

type dirver struct {
	config   *configuration
	network  *bridgeNetwork
	networks *map[types.UUID]bridgeNetworks
	sync.Mutex
}

func init() {
	ipAllocator = ipallocator.New()
	portMapper = portmapper.New()
}
