package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// GetReceiveCollectionOrder 获取已经领取的订单
func GetReceiveCollectionOrder(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	col := make([]modelPay.Collection, 0)
	status := c.PostForm("status")
	mysql.DB.Where("runner_id=?  and  status =?", whoMap.ID, status).Find(&col)

	for i, _ := range col {
		col[i].NoticeUrl = ""
		col[i].CallbackContent = ""
	}

	//更新时间
	if whoMap.Working == 2 {
		mysql.DB.Model(&model.Runner{}).Where("id=?", whoMap.ID).Update(&model.Runner{LastGetOrderTime: time.Now().Unix() + 5*60})
	}
	tools.ReturnSuccess2000DataCode(c, col, "ok")
	return
}

// ImWorking 我正在工作
func ImWorking(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	work, _ := strconv.Atoi(c.PostForm("work"))
	if work != 1 && work != 2 {
		tools.ReturnErr101Code(c, "illegal")
		return
	}
	if work == whoMap.Working {
		tools.ReturnErr101Code(c, "Don't repeat requests")
		return
	}
	//判断是否  已经 绑定了收款地址
	err := mysql.DB.Where("runner_id=? and kind=1", whoMap.ID).First(&model.RunnerUpi{}).Error
	if err != nil {
		tools.JsonWrite(c, tools.NoBindingBackCard, nil, "Bind the working upi first")
		return
	}

	//判断 是要有资格接单
	if whoMap.CollectionLimit < 500 && whoMap.Working == 1 {
		tools.ReturnErr101Code(c, "Your collection balance is not enough,The balance needs to be at least 500")
		return
	}
	uup := &model.Runner{Working: work}
	if work == 2 {
		uup.LastGetOrderTime = time.Now().Unix() + 300
	}

	err = mysql.DB.Model(&model.Runner{}).Where("id=?", whoMap.ID).Update(uup).Error
	if err != nil {
		tools.ReturnErr101Code(c, err)
		return
	}

	tools.ReturnSuccess2000Code(c, c.PostForm("work"))
	return
}

// ConfirmTheOrder 确认订单
func ConfirmTheOrder(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	collection := modelPay.Collection{}
	//实际收到多少金额
	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	err := mysql.DB.Where("id=? and  runner_id=?", c.PostForm("id"), whoMap.ID).First(&collection).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Order does not exist")
		return
	}
	//判断 订单状态
	if collection.Status != 1 {
		tools.ReturnErr101Code(c, "You ordered it too fast")
		return
	}
	if amount > collection.Amount {
		tools.ReturnErr101Code(c, "The amount received cannot be greater than the amount of the order")
		return
	}
	if amount < collection.Amount {
		runner := model.Runner{}
		runner.Remark = "Collect less than the order amount"
		runner.ID = whoMap.ID
		runner.CollectionLimit = collection.Amount - amount
		runner.FreezeCollectionLimit = amount - collection.Amount
		runner.Col.ID = collection.ID
		runner.Col.ActualAmount = amount
		err := runner.ChangeCollectionLimit(mysql.DB, false, 4)
		if err != nil {
			tools.ReturnErr101Code(c, err)
			return
		}
		tools.ReturnSuccess2000Code(c, "Confirm success, wait for administrator review")
		return
	}
	err = mysql.DB.Model(&modelPay.Collection{}).Where("id=?", collection.ID).Update(&modelPay.Collection{Status: 5, ActualAmount: amount}).Error
	if err != nil {
		tools.ReturnErr101Code(c, err)
		return
	}
	tools.ReturnSuccess2000Code(c, "Confirm success, wait for administrator review")
	return
}
