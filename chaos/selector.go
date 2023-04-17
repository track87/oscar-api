// Package chaos declare something
// MarsDong 2023/4/3
package chaos

type DataSource string

const (
	DataSourceCMDB   DataSource = "cmdb"
	DataSourceCustom DataSource = "custom"
)

// SelectMode represents the mode to run chaos action.
type SelectMode string

// 支持的mode类型
const (
	// OneMode represents that the system will do the chaos action on one object selected randomly.
	OneMode SelectMode = "one"
	// AllMode represents that the system will do the chaos action on all objects
	// regardless of status (not ready or not running pods includes).
	// Use this label carefully.
	AllMode SelectMode = "all"
	// FixedMode represents that the system will do the chaos action on a specific number of running objects.
	FixedMode SelectMode = "fixed"
	// FixedPercentMode to specify a fixed % that can be inject chaos action.
	FixedPercentMode SelectMode = "fixed-percent"
	// RandomMaxPercentMode to specify a maximum % that can be inject chaos action.
	RandomMaxPercentMode SelectMode = "random-max-percent"
)

// DomainStrategy 范围策略
type DomainStrategy string

const (
	// DomainStrategyMultiAz 多az均匀打散策略
	DomainStrategyMultiAz DomainStrategy = "MultiAz"
	// DomainStrategySpecialAz 圈定az范围
	// 尽量但无法完全保证
	DomainStrategySpecialAz DomainStrategy = "SpecialAz"
	// DomainStrategySingleAz 单az
	DomainStrategySingleAz DomainStrategy = "SingleAz"
	// DomainStrategyRandom 完全随机
	// 默认值
	DomainStrategyRandom DomainStrategy = "Random"
)

// DomainSelector defines domain selector
type DomainSelector struct {
	// Strategy 策略名称
	Strategy DomainStrategy `json:"Strategy"`
	Azs      []string       `json:"Azs"`
}

type CommonSelector struct {
	// Source defines data source
	// Supported source: cmdb, custom
	Source DataSource `json:"Source,omitempty"`

	// Mode defines the mode to run chaos action.
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent
	Mode SelectMode `json:"Mode,omitempty"`

	// Value is required when the mode is set to `FixedMode` / `FixedPercentMode` / `RandomMaxPercentMode`.
	// If `FixedMode`, provide an integer of pods to do chaos action.
	// If `FixedPercentMode`, provide a number from 0-100 to specify the percent of pods the server can do chaos action.
	// IF `RandomMaxPercentMode`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	// +optional
	Value string `json:"Value,omitempty"`

	Domain *DomainSelector `json:"Domain,omitempty"`
}

// GenericSelectorSpec defines some selectors to select objects.
// +k8s:deepcopy-gen=true
type GenericSelectorSpec struct {
	// Namespaces is a set of namespace to which objects belong.
	// +optional
	Namespaces []string `json:"Namespaces,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on fields.
	// +optional
	FieldSelectors map[string]string `json:"FieldSelectors,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on labels.
	// +optional
	LabelSelectors map[string]string `json:"LabelSelectors,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on annotations.
	// +optional
	AnnotationSelectors map[string]string `json:"AnnotationSelectors,omitempty"`
}

// PodSelectorSpec defines the some selectors to select objects.
// If the all selectors are empty, all objects will be used in chaos experiment.
// +k8s:deepcopy-gen=true
type PodSelectorSpec struct {
	GenericSelectorSpec

	// CmdbNodes is a set of cmdb node uuid.
	// +optional
	CmdbNodes []string `json:"CmdbNodes,omitempty"`

	// Nodes is a set of node name and objects must belong to these nodes.
	// +optional
	Nodes []string `json:"Nodes,omitempty"`

	// Pods is a map of string keys and a set values that used to select pods.
	// The key defines the namespace which pods belong,
	// and the each values is a set of pod names.
	// +optional
	Pods map[string][]string `json:"Pods,omitempty"`

	// Map of string keys and values that can be used to select nodes.
	// Selector which must match a node's labels,
	// and objects must belong to these selected nodes.
	// +optional
	NodeSelectors map[string]string `json:"NodeSelectors,omitempty"`

	// PodPhaseSelectors is a set of condition of a pod at the current time.
	// supported value: Pending / Running / Succeeded / Failed / Unknown
	// +optional
	PodPhaseSelectors []string `json:"PodPhaseSelectors,omitempty"`
}

// PodSelector defines selector conditions for pod
// +k8s:deepcopy-gen=true
type PodSelector struct {
	CommonSelector
	// Selector is used to select pods that are used to inject chaos action.
	Selector PodSelectorSpec `json:"Selector,omitempty"`
}

// HostSelectorSpec defines the some selectors to select objects.
// If the all selectors are empty, all objects will be used in chaos experiment.
// +k8s:deepcopy-gen=true
type HostSelectorSpec struct {
	// CmdbNodes is a set of cmdb node uuid.
	// +optional
	CmdbNodes []string `json:"CmdbNodes,omitempty"`

	// Hosts is a set of hosts.
	// +optional
	Hosts []string `json:"Hosts,omitempty"`

	// Map of string keys and values that can be used to select hosts.
	// Selector which must match a node's labels,
	// and objects must belong to these selected nodes.
	// 预留字段
	// +optional
	ClusterSelectors map[string]string `json:"ClusterSelectors,omitempty"`

	// HostPhaseSelectors is a set of condition of a pod at the current time.
	// supported value: Pending / Running / Succeeded / Failed / Unknown todo 收集主机的状态列表
	// +optional
	// default phase selector ['running']
	HostPhaseSelectors []string `json:"HostPhaseSelectors,omitempty"`
}

// HostSelector defines selector conditions for host
// +k8s:deepcopy-gen=true
type HostSelector struct {
	CommonSelector
	// Selector is used to select hosts that are used to inject chaos action.
	Selector HostSelectorSpec `json:"Selector,omitempty"`
}

// MiddlewareResourceType 中间件资源类型
type MiddlewareResourceType string

// MiddlewareResourceTypePod 目前支持的资源类型
const (
	MiddlewareResourceTypePod MiddlewareResourceType = "cloud-native"
)

// MiddlewareType 中间件类型
type MiddlewareType string

// 支持的中间件类型
const (
	// multi master mode or mix mode

	MiddlewareTypeKafka      MiddlewareType = "kafka"
	MiddlewareTypeClickhouse MiddlewareType = "clickhouse"
	MiddlewareTypeRabbitMQ   MiddlewareType = "rabbitmq"

	// master slave mode

	MiddlewareTypeRedis     MiddlewareType = "redis"
	MiddlewareTypeMariadb   MiddlewareType = "mariadb"
	MiddlewareTypeEtcd      MiddlewareType = "etcd"
	MiddlewareTypeConsul    MiddlewareType = "consul"
	MiddlewareTypeZooKeeper MiddlewareType = "zookeeper"
	MiddlewareTypeMongo     MiddlewareType = "mongo"

	MiddlewareTypeOther MiddlewareType = "other"
)

// MasterMode 选主类型， 一般用于中间件实例的目标选择
type MasterMode string

// MasterMode只针对主从模式的中间件生效
const (
	// MasterModeOnly 只从master节点中选择目标
	MasterModeOnly MasterMode = "only"
	// MasterModeExclude 选择目标时排除master节点
	MasterModeExclude MasterMode = "exclude"
	// MasterModeMix 选择所有目标
	MasterModeMix MasterMode = "mix"
)

// MiddlewareSelectorSpec defines the some selectors to select objects.
// If the all selectors are empty, all objects will be used in chaos experiment.
// +k8s:deepcopy-gen=true
type MiddlewareSelectorSpec struct {
	// MasterMode is the master node, used to select master nodes
	MasterMode MasterMode `json:"MasterMode,omitempty"`

	// Namespaces is a set of namespace to which objects belong.
	Namespaces []string `json:"Namespaces,omitempty"`

	// Map of string keys and values that can be used to select objects.
	// A selector based on labels.
	// +optional
	LabelSelectors map[string]string `json:"LabelSelectors,omitempty"`

	// Type defines middleware type, like kafka, redis, mariadb
	Type MiddlewareType `json:"Type,omitempty"`

	// ResourceType defines the type of instance resource
	ResourceType MiddlewareResourceType `json:"ResourceType,omitempty"`

	// InstanceNames is a set of middleware instance names.
	InstanceNames []string `json:"InstanceNames,omitempty"`
}

// MiddlewareSelector defines selector conditions for middleware
// +k8s:deepcopy-gen=true
type MiddlewareSelector struct {
	CommonSelector
	Selector MiddlewareSelectorSpec `json:"Selector"`
}

type Selector struct {
	TargetType string `json:"TargetType"`
	CommonSelector
	HostSelector       *HostSelectorSpec       `json:"HostSelector"`
	PodSelector        *PodSelectorSpec        `json:"PodSelector"`
	MiddlewareSelector *MiddlewareSelectorSpec `json:"MiddlewareSelector"`
}
