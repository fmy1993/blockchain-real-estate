/**
 * @Author: fmy1993
 * @Email:
 * @Date: 2021/3/25
 * @Description: 获取当前区块信息接口
 */
package v1

import (
	"net/http"
	"strconv"
	"unicode/utf8"

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
	bciResp, err := bc.ChannelQueryBlockinfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	blockInfo := new(BlockInfo)
	blockInfo.BlockHeight = bciResp.BCI.Height
	blockInfo.CurrentBlockHash = bciResp.BCI.CurrentBlockHash
	blockInfo.PreviousBlockHash = bciResp.BCI.PreviousBlockHash
	appG.Response(http.StatusOK, "成功", blockInfo)

}

// @Summary 输出blockInfo
// @Produce  json
// @Param height query string true "区块高度"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getBlockInfoByBlockHeight [get]
func GetBlockInfoByBlockHeight(c *gin.Context) {
	appG := app.Gin{C: c}
	//var height uint64
	str_height := c.Query("height")
	str_height_len := utf8.RuneCountInString(str_height)

	//blockInfo := new(BlockInfo)

	if str_height_len == 0 {
		appG.Response(http.StatusBadRequest, "参数为区块高度，是大于0的整数", nil)
		return
	}

	blockheight, err := strconv.ParseUint(str_height, 10, 64)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	bciResp, err := bc.ChannelQueryBlockinfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	maxHeight := bciResp.BCI.Height - 1
	if blockheight > maxHeight {
		msg := map[string]interface{}{"msg": "系统数据区块高度为：" + strconv.FormatUint(maxHeight, 10) + ",查询参数超过上限"}
		appG.Response(http.StatusInternalServerError, "失败", msg)
		return
	}
	blockResp, err := bc.QueryBlockinfoByBlockHeight(blockheight)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	blockInfo := new(BlockInfo)
	blockInfo.BlockHeight = blockResp.Header.Number
	blockInfo.CurrentBlockHash = blockResp.Header.DataHash
	blockInfo.PreviousBlockHash = blockResp.Header.PreviousHash
	appG.Response(http.StatusOK, "成功", blockInfo)
}

// @Summary 输出最大数据区块高度
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getMaxDataBlockHeight [get]
func GetMaxDataBlockHeight(c *gin.Context) {
	appG := app.Gin{C: c}
	bciResp, err := bc.ChannelQueryBlockinfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}

	maxBlockHeight := bciResp.BCI.Height - 1

	appG.Response(http.StatusOK, "成功", map[string]interface{}{
		"maxdatablockheight": strconv.FormatUint(maxBlockHeight, 10),
	})
}

// @Summary 检查查询的区块高度是否符合要求
// @Produce  json
// @Param datablockheight query string true "区块高度"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/checkDataBlockHeight [get]
func CheckDataBlockHeight(c *gin.Context) {
	appG := app.Gin{C: c}
	queryHeight := c.Query("datablockheight")
	blockheight, _ := strconv.ParseInt(queryHeight, 10, 64)
	if blockheight <= 0 {
		appG.Response(http.StatusOK, "区块高度必须大于0", map[string]interface{}{
			"result": "false",
		})
		return
	}

	//strconv.
	bciResp, err := bc.ChannelQueryBlockinfo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	maxBlockHeight := (int64)(bciResp.BCI.Height - 1)
	if blockheight > maxBlockHeight {
		appG.Response(http.StatusOK, "区块高度必须小于系统最大数据区块高度："+strconv.FormatInt(maxBlockHeight, 10), map[string]interface{}{
			"result": "false",
		})
		return
	}

	appG.Response(http.StatusOK, "成功", map[string]interface{}{
		"result": "true",
	})
}
