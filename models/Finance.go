package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Finance struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Count     uint32    `gorm:"size:255;not null;" json:"count"`
	Sub       string    `gorm:"size:255;not null;" json:"sub"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Finance) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	// p.Count = html.EscapeString(strings.TrimSpace(p.Coun))
	p.Sub = html.EscapeString(strings.TrimSpace(p.Sub))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Finance) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Count == 0 {
		return errors.New("Required Count")
	}
	// if p.Sub == "" {
	// 	return errors.New("Required Sub")
	// }
	return nil
}

func (p *Finance) SaveFinance(db *gorm.DB) (*Finance, error) {

	var err error
	err = db.Debug().Model(&Finance{}).Create(&p).Error
	if err != nil {
		return &Finance{}, err
	}
	// var err error
	// err = db.Debu().Model(&Finance{}).Create(&p).Error
	// if err != nil {
	//	return &Finance{}, err
	//  }
	// if p.ID !=0 {
	//  	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Finance{}, err
	// 	}
	// }
	return p, nil
}

func (p *Finance) FindAllFinances(db *gorm.DB) (*[]Finance, error) {
	var err error
	posts := []Finance{}
	err = db.Debug().Model(&Finance{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Finance{}, err
	}
	//if len(posts) > 0 {
	// 	for i, _ := range posts {
	// 		err := db.Debg().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
	//  		if err != nil {
	// 			return &[]Finance{}, err
	// 		}
	// 	}
	// }
	return &posts, nil
}

func (p *Finance) FindFinanceByID(db *gorm.DB, pid uint64) (*Finance, error) {
	var err error
	err = db.Debug().Model(&Finance{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Finance{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Finance{}, errors.New("Finance Not Found")
	}
	// if p.ID != 0 {
	// 	err = db.ebug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Finance{}, err
	// 	}
	// }
	return p, nil
}

func (p *Finance) UpdateAFinance(db *gorm.DB) (*Finance, error) {

	var err error
	// d = db.Debug().Model(&Finance{}).Where("id = ?", pid).Take(&Finance{}).UpdateColumns(
	// 	map[string]inerface{}{
	// 		"title":      p.Name,
	// 		"content":    .Count,
	// 		"updated_at": time.ow(),
	// 	}
	// )
	// err = db.Debug().Model(&Finance{}).Where("id = ?", pid).Take(&p).Error
	// if err != ni {
	// 	return &Finance{, err
	//  }
	// if p.ID !=  {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Eror
	// 	if err != ni {
	// 		return &Finance{, err
	//	}
	//  }
	err = db.Debug().Model(&Finance{}).Where("id = ?", p.ID).Updates(Finance{Name: p.Name, Count: p.Count, Sub: p.Sub, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Finance{}, err
	}
	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
	// 	if err != nil {
	// 		return &Finance}, err
	// 	}
	// }
	return p, nil
}

func (p *Finance) DeleteAFinance(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Finance{}).Where("id = ? and author_id = ?", pid, uid).Take(&Finance{}).Delete(&Finance{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Finance not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
