package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

func RunnerOperation(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Runner, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		//  状态
		if kinds, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", kinds)
		}
		//用户名  runner_id
		if username, isE := c.GetPostForm("username"); isE == true {
			db = db.Where("username=?", username)
		}
		db.Model(model.Runner{}).Count(&total)
		db = db.Model(&model.Runner{}).Offset((page - 1) * limit).Limit(limit).Order("register_time desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}
	//添加会员
	if action == "add" {
		var ro RunnerOperationAdd
		//检查参数
		if err := c.ShouldBind(&ro); err != nil {
			tools.ReturnVerifyErrCode(c, err)
			return
		}
		add := model.Runner{}
		add.Username = ro.Username
		add.Password = tools.MD5(ro.Password)
		add.AgencyRunnerId = whoMap.ID
		add.CollectionPoint, _ = strconv.ParseFloat(ro.CollectionPoint, 64)
		add.PayPoint, _ = strconv.ParseFloat(ro.PayPoint, 64)
		//对代付 代收的 盈利点  大小判断
		if add.CollectionPoint > whoMap.CollectionPoint || add.PayPoint > whoMap.PayPoint {
			tools.ReturnErr101Code(c, "Payment collection profit point can not be greater than itself")
			return
		}
		err := add.Add(mysql.DB, redis.Rdb)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "Add successfully")
		return
	}
	if action == "update" {
		id, _ := strconv.Atoi(c.PostForm("id"))
		runner := model.Runner{ID: id, AgencyRunnerId: whoMap.ID}
		boolOne, _ := runner.IsExist(mysql.DB)
		if boolOne == false {
			tools.ReturnErr101Code(c, "Sorry, the user does not exist")
			return
		}
		//修改密码
		if password, IsE := c.GetPostForm("password"); IsE == true {
			runner.Password = password
			err := runner.ChangePassword(mysql.DB)
			if err != nil {
				tools.ReturnErr101Code(c, err)
				return
			}
			tools.ReturnErr101Code(c, "Password changed successfully")
			return
		}
		//修改支付密码
		if password, IsE := c.GetPostForm("pay_password"); IsE == true {
			runner.PayPassword = password
			err := runner.ChangePayPassword(mysql.DB)
			if err != nil {
				tools.ReturnErr101Code(c, err)
				return
			}
			tools.ReturnErr101Code(c, "pay_password changed successfully")
			return
		}
		//修改玩家押金
		if cashPledge, isE := c.GetPostForm("cash_pledge"); isE == true {
			result, _ := redis.Rdb.Get("CashPledge_" + runner.Username).Result()
			if result != "" {
				tools.ReturnErr101Code(c, "Don't do the deposit operation at the specified time")
				return
			}
			remark := c.PostForm("remark")
			if remark == "" {
				remark = ""
			}
			m := model.Runner{ID: id, AgencyRunnerId: runner.AgencyRunnerId, Remark: remark}
			m.CashPledge, _ = strconv.ParseFloat(cashPledge, 64)
			err := m.ChangeCashPledge(mysql.DB)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			redis.Rdb.Set("CashPledge_"+runner.Username, "123", 5*time.Second)
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}
		//修改代收金额
		if CollectionLimit, isE := c.GetPostForm("collection_limit"); isE == true {
			result, _ := redis.Rdb.Get("CollectionLimit" + runner.Username).Result()
			if result != "" {
				tools.ReturnErr101Code(c, "Don't do the deposit operation at the specified time")
				return
			}
			runner.CollectionLimit, _ = strconv.ParseFloat(CollectionLimit, 64)
			//修改代收额度
			err := tools.ForFunc(runner.ChangeCollectionLimit, mysql.DB, true)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}

		//修改佣金  +  -

		//

	}

}
