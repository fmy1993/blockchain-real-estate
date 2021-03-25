/**
 * @Author: fmy1993
 * @Email:
 * @Date: 2021/3/25
 * @Description: 获取当前区块信息接口
 */
package v1

import (
	"net/http"

	bc "github.com/fmy1993/blockchain-real-estate/application/blockchain"
	"github.com/fmy1993/blockchain-real-estate/application/pkg/app"
	"github.com/gin-gonic/gin"
)

type BlockInfo struct {
	CurrentBlockHash  []byte `json:"currentblockinfo"`
	PreviousBlockHash []byte `json:"previousblockinfo"`
	BlockHeight       uint64 `json:"blockheight"`
}

// @Summary 输出blockInfo
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getBlockInfo [get]
func GetBlockInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	bciResp := bc.ChannelQueryBlockinfo()
	blockInfo := new(BlockInfo)
	blockInfo.BlockHeight = bciResp.BCI.Height
	blockInfo.CurrentBlockHash = bciResp.BCI.CurrentBlockHash
	blockInfo.PreviousBlockHash = bciResp.BCI.PreviousBlockHash
	appG.Response(http.StatusOK, "成功", blockInfo)

}
