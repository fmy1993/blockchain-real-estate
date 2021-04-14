package v1

import (
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

// @Summary test post
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
