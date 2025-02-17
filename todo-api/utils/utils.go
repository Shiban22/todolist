package utils

import "github.com/bwmarrin/snowflake"

func GenSnowflakeId() int64 {
	node, err := snowflake.NewNode(1) // 节点 ID 可以根据你的实际需求设置
	if err != nil {
		// 处理错误，这里简单返回 0
		return 0
	}
	return node.Generate().Int64()
}
