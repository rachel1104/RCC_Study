package test02

import "fmt"

// DB的结构体
type DB struct {
	path string
	data map[string][]byte
}

//模拟开启连接

func New(path string) (*DB, error) {
	self := DB{
		path: path,
		data: make(map[string][]byte),
	}
	return &self, nil
}

// 模拟关闭连接
func (self *DB) Close() error {
	return nil
}
func (self *DB) Put(key []byte, value []byte) error {
	self.data[string(key)] = value
	return nil
}
func (self *DB) Get(key []byte) ([]byte, error) {
	if v, ok := self.data[string(key)]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("Not found")
	}
}
func (self *DB) Del(key []byte) error {
	if _, ok := self.data[string(key)]; ok {
		delete(self.data, string(key))
	} else {
		return fmt.Errorf("Not found")
	}
}

//模拟遍历取值
func (self *DB) Iterator
func main() {

}
