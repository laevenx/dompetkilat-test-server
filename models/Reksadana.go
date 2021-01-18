package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Reksadana struct{
	ID        	uint64    		`gorm:"primary_key;auto_increment" json:"id"`
	Name    	string      	`gorm:"size:255;not null;" json:"name"`
	Amount   	uint32     		`gorm:"size:255;not null;" json:"amount"`
	Return    	int     		`gorm:"size:255;not null;" json:"return"`
	CreatedAt	time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Reksadana) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

}

func (p *Reksadana) Validate() error {

	if p.Name == "" {
		return errors.New("Required Title")
	}
	if p.Amount == 0 {
		return errors.New("Required Title")
	}
	return nil
}

func (p *Reksadana) SaveReksadana(db *gorm.DB) (*Reksadana, error) {
	var err error
	err = db.Debug().Model(&Reksadana{}).Create(&p).Error
	if err != nil {
		return &Reksadana{}, err
	}
	
	return p, nil
}

func (p *Reksadana)FindAllReksadana(db *gorm.DB) (*[]Reksadana, error) {
	var err error
	posts := []Reksadana{}
	err = db.Debug().Model(&Reksadana{}).Limit(100).Find(&posts).Error
	if err != nil{
		return &[]Reksadana{}, err
	}
	
		
	return &posts, nil
}

func (p *Reksadana)FindReksadanaByID(db *gorm.DB, pid uint64) (*Reksadana, error) {
	var err error
	err = db.Debug().Model(&Reksadana{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Reksadana{},err
	}
	
	return p, nil
}

func (p *Reksadana) UpdateAReksadana(db *gorm.DB) (*Reksadana, error) {

var err error
	
	err = db.Debug().Model(&Reksadana{}).Where("id = ?", p.ID).Updates(Reksadana{Name: p.Name, Amount: p.Amount,Return: p.Return, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Reksadana{}, err
	}
	
	return p, nil
}

func (p *Reksadana)DeleteAReksadana(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

db= db.Debug().Model(&Reksadana{}).Where("id = ?", pid).Take(&Reksadana{}).Delete(&Reksadana{})

if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
		return 0, errors.New("Reksadana not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
