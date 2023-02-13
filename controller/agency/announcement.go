package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// AnnouncementOperation 公告操作
func AnnouncementOperation(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Announcement, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		//  状态
		if kinds, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", kinds)
		}
		db.Model(model.Announcement{}).Count(&total)
		db = db.Model(&model.Announcement{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "update" {
		id, _ := strconv.Atoi(c.PostForm("id"))
		Announcement := model.Announcement{}
		err := mysql.DB.Where("id=?", id).First(&Announcement).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		//状态单独修改
		if status, isExist := c.GetPostForm("status"); isExist == true {
			sta, _ := strconv.Atoi(status)
			if sta != 2 && sta != 1 {
				tools.ReturnErr101Code(c, "Please enter the correct status")
				return
			}
			err := mysql.DB.Model(&model.Announcement{}).Where("id=?", id).Update(&model.Announcement{Status: sta}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			tools.ReturnSuccess2000Code(c, "Modified successfully")
			return
		}

	}

	if action == "del" {
		id, _ := strconv.Atoi(c.PostForm("id"))
		Announcement := model.Announcement{}
		err := mysql.DB.Where("id=?", id).First(&Announcement).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		err = mysql.DB.Model(&model.Announcement{}).Delete(&model.Announcement{ID: id}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

	if action == "add" {
		txt := c.PostForm("content")
		if txt == "" {
			tools.ReturnErr101Code(c, "content is not null")
			return
		}
		title := c.PostForm("title")
		if title == "" {
			tools.ReturnErr101Code(c, "content is not null")
			return
		}
		announcement := model.Announcement{Content: txt, AgencyRunnerId: whoMap.ID, Title: title}
		err := announcement.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
