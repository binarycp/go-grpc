package _interface

import "google.golang.org/grpc"

// 上位机和调试代理服务的数据流转换
type Stream interface {
	// 数据流转换
	To(p []byte)
	// 拓展容量
	Expansion(n int) int
}

// 连接池
type Pool interface {
	// 取得一个可用的连接
	Get() (*grpc.ClientConn, error)
	// 关闭连接池
	Close() error
	// 返回连接池的状态：没有空闲连接，连接池连接数max...
	// 最大只能有32或者64种状态
	Status() int
}

// 链表
type Link interface {
	// 获取头节点
	Unshift() (*Link, error)
	// 添加头节点
	Shift(link *Link)
	// 获取链表长度
	Len() uint
}
