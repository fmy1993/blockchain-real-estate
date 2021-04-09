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
	Id       string `json:"id"`
	HashInfo string `json:"hashinfo"`
}

// @Summary test post
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
	bodyBytes = append(bodyBytes, []byte(cropBody.Id))
	bodyBytes = append(bodyBytes, []byte(cropBody.HashInfo))

	//调用智能合约
	resp, err := bc.ChannelExecute("createCrop", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	//appG.Response(http.StatusOK, "成功", data)
	GetBlockInfo(c)
	//appG.Response(http.StatusOK, "成功", cropBody)

}
