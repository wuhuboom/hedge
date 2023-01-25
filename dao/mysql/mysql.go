/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err)
		panic(err)
		return err
	}

	//设置连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	////////////////////////////////////////////////////////////////////////模型初始化
	//管理员列表
	model.CheckIsExistModelAdminAdmin(DB)
	model.CheckIsExistModelConfig(DB)
	model.CheckIsExistModelMerchant(DB)
	model.CheckIsExistModelPaidOrder(DB)
	model.CheckIsExistModelPayOrder(DB)
	model.CheckIsExistModelLogger(DB)
	modelPay.CheckIsExistModelBankInformation(DB)
	modelPay.CheckIsExistModelBank(DB)
	modelPay.CheckIsExistModelChannel(DB)
	modelPay.CheckIsExistModelChannelBank(DB)
	modelPay.CheckIsExistModelCollection(DB)
	modelPay.CheckIsExistModelAmountChange(DB)
	modelPay.CheckIsExistModelStatistics(DB)
	model.CheckIsExistModelAgencyRunner(DB)
	model.CheckIsExistModelRunner(DB)
	model.CheckIsExistModelAgencyAccountChange(DB)
	////////////////////////////////////////////////////////////////////////模型初始化
	return err
}

func Close() {
	defer DB.Close()
}
