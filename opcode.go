package opcode


import "fmt"

type Opcode uint16

const (
	OpHeartbeat Opcode = iota
	OpSubstance
	OpThird
	OpReload
	OpOffline
	OpDeleted
	OpUpgrade
)

const (
	OpEvent Opcode = iota + 100
	OpTask
	OpSpdx
)

const (
	OpCPU Opcode = iota + 200
	OpDiskIO
	OpFileSystem
	OpListen
	OpMemory
	OpNetwork
	OpService
	OpSocket
	OpSysInfo
)

const (
	OpAccountDiff Opcode = iota + 300
	OpAccountFull
	OpProcessDiff
	OpProcessFull
)

// opcodes Opcode 对应名称
var opcodes = map[Opcode]string{
	OpHeartbeat: "minion 发出的心跳包",
	OpSubstance: "minion 配置更新",
	OpThird:     "三方文件更新",
	OpReload:    "重新加载指定配置",
	OpOffline:   "节点下线",
	OpDeleted:   "删除节点",
	OpUpgrade:   "节点客户端升级",

	OpEvent: "上报事件",
	OpTask:  "上报 rock-go 内部服务运行信息",
	OpSpdx:  "上报节点 SPDX 清单",

	OpCPU:        "上报 CPU 信息",
	OpDiskIO:     "上报磁盘 I/O",
	OpFileSystem: "上报文件系统",
	OpListen:     "上报端口监听",
	OpMemory:     "上报内存信息",
	OpNetwork:    "上报网络信息",
	OpService:    "上报系统服务信息",
	OpSocket:     "上报 socket 连接信息",
	OpSysInfo:    "上报节点基本信息",

	OpAccountDiff: "上报账户差异信息",
	OpAccountFull: "上报账户全量信息",
	OpProcessDiff: "上报进程差异信息",
	OpProcessFull: "上报进程全量信息",
}

// String implement fmt.Stringer
func (op Opcode) String() string {
	if str, exist := opcodes[op]; exist {
		return str
	}
	return fmt.Sprintf("<unnamed minion opcode: %d>", op)
}
