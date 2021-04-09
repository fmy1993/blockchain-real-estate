package routers

import (
	"encoding/json"
	"fmt"

	"github.com/fmy1993/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/fmy1993/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//新建房地产(管理员)
func CreateCrop(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 2 {
		return shim.Error("参数个数不满足")
	}
	id := args[0] //accountId用于验证是否为管理员
	hashinfo := args[1]

	if id == "" || hashinfo == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	// var formattedTotalArea float64
	// if val, err := strconv.ParseFloat(totalArea, 64); err != nil {
	// 	return shim.Error(fmt.Sprintf("totalArea参数格式转换出错: %s", err))
	// } else {
	// 	formattedTotalArea = val
	// }

	//判断业主是否存在
	// resultsProprietor, err := utils.GetStateByPartialCompositeKeys(stub, lib.AccountKey, []string{proprietor})
	// if err != nil || len(resultsProprietor) != 1 {
	// 	return shim.Error(fmt.Sprintf("业主proprietor信息验证失败%s", err))
	// }
	Crop := &lib.Crop{
		Id:       id,
		HashInfo: hashinfo,
	}
	// 写入账本
	if err := utils.WriteLedger(Crop, stub, lib.RealEstateKey, []string{Crop.Id, Crop.HashInfo}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	CropByte, err := json.Marshal(Crop)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(CropByte)
}
