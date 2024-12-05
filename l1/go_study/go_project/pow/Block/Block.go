package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
区块-区块-区块
区块保存信息
HashCode:当前节点的hash
PreCode:前一节点的hash
timestamp:时间戳
Diff:网络难以系数，若前导0符合难度系数，则挖矿成功
Data:交易信息
Index:区块高度（第一个块高度为1）
Nonce:随机值
创世区块
*/
//通过代码挖矿
//1.声明区块的结构体
type Block struct {
	//上一个区块的hash
	Prehash string
	//当前区块的hash
	HashCode string
	//时间戳
	TimeStamp string
	//当前网络的难度系数,
	//控制hash值有几个前导0
	Diff int
	//存交易信息
	Data string
	//区块高度
	Index int
	//随机值
	Nonce int
}

// 创建创世区块（链中的第一个区块）
func GenerateFirstBlock(data string) Block {
	//创建一个区块
	var firstblock Block
	//赋值
	firstblock.Prehash = "0"
	firstblock.TimeStamp = time.Now().String()
	firstblock.Diff = 4 //暂且设置为4，也就是前置0的个数
	firstblock.Data = data
	firstblock.Index = 1
	firstblock.Nonce = 0
	//当前块的hash
	//用sha256算一个真正的hash
	firstblock.HashCode = GenerationHashValue(firstblock)
	return firstblock
}

// 生成区块的hash值
func GenerationHashValue(block Block) string {
	var hashdata = strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Diff) + block.TimeStamp
	//hash算法
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hased := sha.Sum(nil)
	//将字节转为字符串
	return hex.EncodeToString(hased)
}

// 产生新的区块
func GenerateNextBlock(data string, oldBlock Block) Block {
	//产生一个新的区块
	var newBlock Block
	newBlock.Prehash = oldBlock.HashCode
	newBlock.TimeStamp = time.Now().String()
	newBlock.Diff = 4    //暂且设置为4，也就是前置0的个数
	newBlock.Data = data //交易信息
	newBlock.Index = 2
	//0是由旷工去调整的
	newBlock.Nonce = 0
	//根据diff写
	//利用pow进行挖矿
	newBlock.HashCode = pow(newBlock.Diff, &newBlock)
	return newBlock
}

// PoW工作量证明算法进行hash碰撞
// 传指针，保证操作的是同一对象
func pow(diff int, block *Block) string {
	//不停的去挖矿
	for {
		hash := GenerationHashValue(*block)
		//每挖一次，打印一次hash值
		fmt.Println(hash)
		//判断hash值前缀是否为4个0
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			fmt.Println("挖矿成功")
			return hash
		} else {
			//没挖到
			//随机值自增
			block.Nonce++
		}
	}
}
