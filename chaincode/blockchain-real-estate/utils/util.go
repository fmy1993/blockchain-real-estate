/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 1:51 下午
 * @Description: 读写账本工具
 */
package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//写入账本
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	// 序列化对象
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-序列化json数据失败出错: %s", objectType, err))
	}
	//写入区块链账本
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-写入区块链账本出错: %s", objectType, err))
	}
	return nil
}

//删除账本
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//创建复合主键,重新创建复合主键，根据这个复合主键删除账本中的数据
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-创建复合主键出错 %s", objectType, err))
	} else {
		key = val
	}
	//写入区块链账本
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-删除区块链账本出错: %s", objectType, err))
	}
	return nil
}

// objectType是复合键的类型
//根据复合主键查询数据(适合获取全部，多个，单个数据)
//将keys拆分查询
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// 传入的keys长度为0，则查找并返回所有数据
		// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
		// shim.StateQueryIteratorInterface,err
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
		}
		defer resultIterator.Close()

		//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组，有提供一个迭代器遍历集合
		for resultIterator.HasNext() {
			val, err := resultIterator.Next() // 这里取到值  (*queryresult.KV, error)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
			}

			results = append(results, val.GetValue()) //只取 value(结构体)，返回的是字节数组
			// 因此results就是 [][] byte    []byte的一个父集
		}
	} else {
		// 传入的keys长度不为0，查找相应的数据并返回
		// 就是一个for循环，每次查一个
		for _, v := range keys {
			// 创建组合键
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-创建组合键出错: %s", objectType, err))
			}
			// 从账本中获取数据
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-获取数据出错: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes) // 再组合起来
			}
		}
	}

	return results, nil
}

//根据复合主键查询数据(适合获取全部或指定的数据)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// 通过主键从区块链查找相关的数据，相当于对主键的模糊查询
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-获取全部数据出错: %s", objectType, err))
	}
	defer resultIterator.Close()

	//检查返回的数据是否为空，不为空则遍历数据，否则返回空数组
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-返回的数据出错: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
