package Blockchain

import (
	"fmt"
	"go_project/pow/Block"
)

// 通过链表的形式，维护区块链中的业务
type Node struct {
	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block
}

func main() {

}

// 创建头节点，保存创世区块
func CreateHeaderNode(data *Block.Block) *Node {
	//创建头结点
	headerNode := new(Node)
	//头节点指针域指向nil
	headerNode.NextNode = nil
	//数据域保存传入的data
	headerNode.Data = data
	//返回头节点
	return headerNode
}

// 添加节点
// 当挖矿成功，添加区块
func AddNode(data *Block.Block, node *Node) *Node {
	//创建新节点
	var newNode *Node = new(Node)
	//指针域指向nil
	newNode.NextNode = nil
	//数据域是传入的data
	newNode.Data = data
	newNode.NextNode = newNode
	return newNode
}

// 查看链表中的数据
func ShowNodes(node *Node) {
	n := node
	for {
		//没有下一个节点，则结束循序
		if n.NextNode == nil {
			fmt.Println(n.Data)
			break
		} else {
			fmt.Println(n.Data)
			n = n.NextNode
		}
	}
}
