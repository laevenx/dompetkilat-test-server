package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Sbn struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Amount    uint32    `gorm:"size:255;not null;" json:"amount"`
	Tenor     uint32    `gorm:"size:255;not null;" json:"tenor"`
	Rate      uint32    `gorm:"size:255;nonull;" json:"rate"`
	Type      string    `gorm:"size:255;not null;" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Sbn) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Type = html.EscapeString(strings.TrimSpace(p.Type))

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

}

func (p *Sbn) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Amount == 0 {
		return errors.New("Required Amount")
	}
	if p.Rate == 0 {
		return errors.New("Required Rate")
	}
	if p.Type == "" {
		return errors.New("Required Type")
	}

	return nil
}

func (p *Sbn) SaveSbn(db *gorm.DB) (*Sbn, error) {
	var err error
	err = db.Debug().Model(&Sbn{}).Create(&p).Error
	if err != nil {
		return &Sbn{}, err
	}

	return p, nil

}

func (p *Sbn) FindAllSbn(db *gorm.DB) (*[]Sbn, error) {
	var err error
	result := []Sbn{}
	err = db.Debug().Model(&Sbn{}).Limit(100).Find(&result).Error
	if err != nil {
		return &[]Sbn{}, err
	}

	return &result, nil

}

func (p *Sbn) FindSbnByID(db *gorm.DB, pid uint64) (*Sbn, error) {
	var err error
	err = db.Debug().Model(&Sbn{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Sbn{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Sbn{}, errors.New("Sbn Not Found")
	}

	return p, nil
}

func (p *Sbn) UpdateASbn(db *gorm.DB) (*Sbn, error) {

	var err error

	err = db.Debug().Model(&Sbn{}).Where("id = ?", p.ID).Updates(Sbn{Name: p.Name, Amount: p.Amount, Tenor: p.Tenor, Rate: p.Rate, Type: p.Type, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Sbn{}, err
	}

	return p, nil
}

func (p *Sbn) DeleteASbn(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Sbn{}).Where("id = ?", pid).Take(&Sbn{}).Delete(&Sbn{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Sbn not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
