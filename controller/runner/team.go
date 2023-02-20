package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// GetSubordinate 团队中心
func GetSubordinate(c *gin.Context) {
	action := c.Query("action")
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	//获取 所有下级代理
	if action == "list" {
		runner := make([]model.Runner, 0)
		db := mysql.DB.Where("superior=?", whoMap.ID)
		if username, ise := c.GetPostForm("username"); ise == true {
			db = db.Where("username=?", username)
		}
		db.Find(&runner)
		tools.ReturnSuccess2000DataCode(c, runner, "OK")
		return
	}
	//修改其盈利点
	if action == "update" {
		id := c.PostForm("runner_id")
		ru2 := model.Runner{}
		err := mysql.DB.Where("id=? and superior=?", id, whoMap.ID).First(&ru2).Error
		if err != nil {
			tools.ReturnErr101Code(c, "the user is not exist")
			return
		}

		ups := map[string]interface{}{}
		if rate, isE := c.GetPostForm("collection_point"); isE == true {
			float, _ := strconv.ParseFloat(rate, 64)
			if float > whoMap.CollectionPoint {
				tools.ReturnErr101Code(c, "Can not be more than your collection rate")
				return
			}
			ups["CollectionPoint"] = float
		}

		if rate, isE := c.GetPostForm("pay_point"); isE == true {
			float, _ := strconv.ParseFloat(rate, 64)
			if float > whoMap.PayPoint {
				tools.ReturnErr101Code(c, "Can not be more than your pay rate")
				return
			}
			ups["PayPoint"] = float
		}

		err = mysql.DB.Model(&model.Runner{}).Where("id=?", id).Update(ups).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Failed to modify. Try again later")
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return

	}
	//团队接单明细
	if action == "order" {
		db := mysql.DB
		ls := make([]modelPay.Collection, 0)
		if username, ise := c.GetPostForm("username"); ise == true {
			runner := model.Runner{Username: username}
			id := runner.GetRunnerId(mysql.DB)
			err := mysql.DB.Where("id=? and superior=?", id, whoMap.ID).First(&model.Runner{}).Error
			if err != nil {
				tools.ReturnErr101Code(c, "the user is not exist")
				return
			}
			//类型
			if kinds, isE := c.GetPostForm("kinds"); isE == true {
				db = db.Where("kinds=?", kinds)
			}
			db = db.Where("runner_id=?", id).Find(&ls)
			for i, l := range ls {
				rr := model.Runner{ID: l.RunnerId}
				ls[i].RunnerName = rr.GetRunnerUsername(mysql.DB)
			}
			tools.ReturnSuccess2000DataCode(c, ls, "OK")
			return
		}
		//查询下级的
		runnerArray := make([]model.Runner, 0)
		db.Where("superior=?", whoMap.ID).Find(&runnerArray)
		var IdArray []int
		for _, runner := range runnerArray {
			IdArray = append(IdArray, runner.ID)
		}
		db = mysql.DB.Where("runner_id  in  (?)", IdArray)
		if kinds, isE := c.GetPostForm("kinds"); isE == true {
			db = db.Where("kinds=?", kinds)
		}
		db.Find(&ls)
		for i, l := range ls {
			rr := model.Runner{ID: l.RunnerId}
			ls[i].RunnerName = rr.GetRunnerUsername(mysql.DB)
		}

		tools.ReturnSuccess2000DataCode(c, ls, "OK")
		return
	}
	//团队充值明细
	//团队提现明细
	if action == "withdraw" {
		db := mysql.DB
		ls := make([]model.Record, 0)
		if username, ise := c.GetPostForm("username"); ise == true {
			runner := model.Runner{Username: username}
			id := runner.GetRunnerId(mysql.DB)
			err := mysql.DB.Where("id=? and superior=?", id, whoMap.ID).First(&model.Runner{}).Error
			if err != nil {
				tools.ReturnErr101Code(c, "the user is not exist")
				return
			}
			//类型
			if kinds, isE := c.GetPostForm("kinds"); isE == true {
				db = db.Where("kinds=?", kinds)
			}
			db = db.Where("runner_id=?", id).Find(&ls)
			tools.ReturnSuccess2000DataCode(c, ls, "OK")
			return
		}
		//查询下级的
		runnerArray := make([]model.Runner, 0)
		db.Where("superior=?", whoMap.ID).Find(&runnerArray)
		var IdArray []int
		for _, runner := range runnerArray {
			IdArray = append(IdArray, runner.ID)
		}

		db = mysql.DB.Where("runner_id  in  (?)", IdArray)
		if kinds, isE := c.GetPostForm("kinds"); isE == true {
			db = db.Where("kinds=?", kinds)
		}
		db.Find(&ls)
		tools.ReturnSuccess2000DataCode(c, ls, "OK")
		return
	}
	//团队账变
	if action == "amountChange" {
		db := mysql.DB
		//判断群贤
		ls := make([]model.RunnerAmountChange, 0)
		if username, ise := c.GetPostForm("username"); ise == true {
			runner := model.Runner{Username: username}
			id := runner.GetRunnerId(mysql.DB)
			err := mysql.DB.Where("id=? and superior=?", id, whoMap.ID).First(&model.Runner{}).Error
			if err != nil {
				tools.ReturnErr101Code(c, "the user is not exist")
				return
			}
			//类型
			if kinds, isE := c.GetPostForm("kinds"); isE == true {
				db = db.Where("kinds=?", kinds)
			}

			db = db.Where("runner_id=?", id).Find(&ls)
			tools.ReturnSuccess2000DataCode(c, ls, "OK")
			return
		}

		//查询下级的
		runnerArray := make([]model.Runner, 0)
		db.Where("superior=?", whoMap.ID).Find(&runnerArray)
		var IdArray []int
		for _, runner := range runnerArray {
			IdArray = append(IdArray, runner.ID)
		}
		arc := make([]model.RunnerAmountChange, 0)
		db = db.Where(IdArray)
		if kinds, isE := c.GetPostForm("kinds"); isE == true {
			db = db.Where("kinds=?", kinds)
		}
		db.Find(&arc)

		tools.ReturnSuccess2000DataCode(c, arc, "OK")
		return
	}

}
