/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/fmy1993)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/3 11:24 下午
 * @Description: 测试api
 */
package v1

import (
	"net/http"

	bc "github.com/fmy1993/blockchain-real-estate/application/blockchain"
	"github.com/fmy1993/blockchain-real-estate/application/pkg/app"
	"github.com/gin-gonic/gin"
)

// @Summary 测试输出Hello
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/hello [get]
func Hello(c *gin.Context) {
	appG := app.Gin{C: c}
	BciResp := bc.ChannelQueryBlockinfo()
	//appG.Response(http.StatusOK, "成功", bci.BlockchainInfo.GetCurrentBlockHash())
	appG.Response(http.StatusOK, "成功", map[string]interface{}{
		"当前区块hash值": BciResp.BCI.CurrentBlockHash,
		"前区块hash值":  BciResp.BCI.PreviousBlockHash,
		"区块高度":      BciResp.BCI.Height,
	})
	// appG.Response(http.StatusOK, "成功", map[string]interface{}{
	// 	"msg": "test",
	// })
}
