package utils

import (
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node  //节点
)

func InitSnowflake(MachineID int64) error {
	snowflake.Epoch = 1704067200000   //设置时间戳起始点为2024-01-01 00:00:00
	var err error
	//创造一个节点
	node, err = snowflake.NewNode(MachineID)
	return err
}

//生成int64类型的id
func GenerateID() int64 {
	if node == nil {
		panic("snowflake node not initialized")
	}
	return node.Generate().Int64()
}


//生成string类型的id
func GenerateIDStr() string {
	if node == nil {
		panic("snowflake node not initialized")
	}
	return node.Generate().String()
}
