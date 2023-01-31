package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
)

// GetReceiveCollectionOrder 获取已经领取的订单
func GetReceiveCollectionOrder(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	col := make([]modelPay.Collection, 0)
	status := c.PostForm("status")
	fmt.Println(whoMap)
	mysql.DB.Where("runner_id=?  and  status =?", whoMap.ID, status).Find(&col)
	tools.ReturnSuccess2000DataCode(c, col, "ok")
	return

}
