package model

import (
	"gorm.io/gorm"
	"log"
	"telegram-dice-bot/internal/utils"
)

type QuickThereLotteryRecord struct {
	Id           string `json:"id" gorm:"type:varchar(64);not null;primaryKey"`
	ChatGroupId  string `json:"chat_group_id" gorm:"type:varchar(64);not null"`
	IssueNumber  string `json:"issue_number" gorm:"type:varchar(64);not null"`
	ValueA       int    `json:"value_a" gorm:"type:int(11);not null"`
	ValueB       int    `json:"value_b" gorm:"type:int(11);not null"`
	ValueC       int    `json:"value_c" gorm:"type:int(11);not null"`
	Total        int    `json:"total" gorm:"type:int(11);not null"`
	SingleDouble string `json:"single_double" gorm:"type:varchar(255);not null"`
	BigSmall     string `json:"big_small" gorm:"type:varchar(255);not null"`
	Triplet      int    `json:"triplet" gorm:"type:int(11);not null"`
	CreateTime   string `json:"create_time" gorm:"type:varchar(255);not null"`
}

func (c *QuickThereLotteryRecord) Create(db *gorm.DB) error {
	if c.Id == "" {
		id, err := utils.NextID()
		if err != nil {
			log.Println("SnowFlakeId create error")
			return err
		}
		c.Id = id
	}

	result := db.Create(c)
	if result.Error != nil {
		return result.Error
	}

	return nil
}