package main

import (
	"go_project/pow/Block"
	"go_project/pow/Blockchain"
)

func main() {
	first := Block.GenerateFirstBlock("创世区块")
	//生成下一个区块
	second := Block.GenerateNextBlock("第二个区块", first)
	//创建链表
	header1 := Blockchain.CreateHeaderNode(&first)
	//将第二个区块加入链表
	Blockchain.AddNode(&second, header1)
	//查看链表中的数据
	Blockchain.ShowNodes(header1)
}
