package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	bc "github.com/fmy1993/blockchain-real-estate/application/blockchain"
	"github.com/fmy1993/blockchain-real-estate/application/pkg/app"
	"github.com/gin-gonic/gin"
)

type Crop struct {
	DataType string `json:"datatype"` // add column here
	Id       string `json:"id"`
	HashInfo string `json:"hashinfo"`
}

type CropIdBody struct {
	CropId string `json:"cropid"`
	//DataType string `json:"datatype"`
}

//分别定义数组和 数组中的 k,v可以直接绑定jsonarray
type CropRequestBody struct {
	Args []CropIdBody `json:"args"`
}

// @Summary 增加上链数据
// @Param Crop body Crop true "crop"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/addCrop [post]
func AddCrop(c *gin.Context) {
	appG := app.Gin{C: c}
	cropBody := new(Crop)
	if err := c.ShouldBind(&cropBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(cropBody.DataType)) //add column here，注意顺序，加到集合最后
	bodyBytes = append(bodyBytes, []byte(cropBody.Id))
	bodyBytes = append(bodyBytes, []byte(cropBody.HashInfo))
	//var resp channel.Response
	//调用智能合约
	//resp, err := bc.ChannelExecute("createCrop", bodyBytes)  这样会带来编译错误，避免错误的方法是 改成_
	_, err := bc.ChannelExecute("createCrop", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	//resp.Payload存放着链码返回的相应信息
	/*  做测试返回插入数据用的，已经没用了
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	*/
	bciResp, err := bc.ChannelQueryBlockinfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	maxHeight := bciResp.BCI.Height - 1
	blockResp, err := bc.QueryBlockinfoByBlockHeight(maxHeight)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	blockInfo := new(BlockInfo) // 一个包里的结构体可以相互使用
	blockInfo.BlockHeight = blockResp.Header.Number
	blockInfo.CurrentBlockHash = blockResp.Header.DataHash
	blockInfo.PreviousBlockHash = blockResp.Header.PreviousHash
	appG.Response(http.StatusOK, "成功", blockInfo)

}

// @Summary 查询上链信息,cropid="datatype"+"-"+"id" eg:"test-711"
// @Param crop body CropRequestBody true "crop"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryCrop [post]
func QueryCrop(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CropRequestBody) //CropRequestBody
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.CropId))
	}
	//调用智能合约,这里必须要修改
	resp, err := bc.ChannelQuery("queryCrop", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 增加上链数据
// @Param Crop body Crop true "crop"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/updateCrop [post]
func UpdateCrop(c *gin.Context) {
	appG := app.Gin{C: c}
	cropBody := new(Crop) //CropRequestBody
	//解析Body参数
	if err := c.ShouldBind(&cropBody); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte

	//bodyBytes = append(bodyBytes, []byte(val.CropId))
	bodyBytes = append(bodyBytes, []byte(cropBody.DataType)) //add column here，注意顺序，加到集合最后
	bodyBytes = append(bodyBytes, []byte(cropBody.Id))
	bodyBytes = append(bodyBytes, []byte(cropBody.HashInfo))

	//调用智能合约,这里必须要修改 resp
	resp, err := bc.ChannelExecute("updateCrop", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	// var data []map[string]interface{}
	// if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
	// 	appG.Response(http.StatusInternalServerError, "失败", err.Error())
	// 	return
	// }
	appG.Response(http.StatusOK, "成功", resp)
}
