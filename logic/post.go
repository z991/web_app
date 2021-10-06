package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatPost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	//err = redis.CreatPost(p.ID, p.CommunityID)

	return
}
