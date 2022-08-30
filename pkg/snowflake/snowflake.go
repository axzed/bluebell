// 第三方库: 雪花算法生成器
package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

// 全局node变量
var node *sf.Node

// 初始化上面全局的node
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// 通过节点生成雪花ID
func GenID() int64 {
	return node.Generate().Int64()
}
