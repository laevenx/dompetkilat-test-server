package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ConventionalInvoice struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Amount    uint32    `gorm:"size:255;not null;" json:"amount"`
	Tenor     uint32    `gorm:"size:255;not null;" json:"tenor"`
	Grade     string    `gorm:"size:255;not null;" json:"grade"`
	Rate      uint32    `gorm:"size:255;not null;" json:"rate"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *ConventionalInvoice) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Grade = html.EscapeString(strings.TrimSpace(p.Grade))

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *ConventionalInvoice) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Grade == "" {
		return errors.New("Required Grade")
	}
	if p.Amount == 0 {
		return errors.New("Required Amount")
	}

	return nil
}

func (p *ConventionalInvoice) SaveConventionalInvoice(db *gorm.DB) (*ConventionalInvoice, error) {
	var err error
	err = db.Debug().Model(&ConventionalInvoice{}).Create(&p).Error
	if err != nil {
		return &ConventionalInvoice{}, err
	}

	return p, nil
}

func (p *ConventionalInvoice) FindAllConventionalInvoices(db *gorm.DB) (*[]ConventionalInvoice, error) {
	var err error
	posts := []ConventionalInvoice{}
	err = db.Debug().Model(&ConventionalInvoice{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]ConventionalInvoice{}, err
	}

	return &posts, nil
}

func (p *ConventionalInvoice) FindConventionalInvoiceByID(db *gorm.DB, pid uint64) (*ConventionalInvoice, error) {
	var err error
	err = db.Debug().Model(&ConventionalInvoice{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &ConventionalInvoice{}, err
	}

	return p, nil
}

func (p *ConventionalInvoice) UpdateAConventionalInvoice(db *gorm.DB) (*ConventionalInvoice, error) {

	var err error

	err = db.Debug().Model(&ConventionalInvoice{}).Where("id = ?", p.ID).Updates(ConventionalInvoice{Name: p.Name, Amount: p.Amount, Tenor: p.Tenor, Grade: p.Grade, Rate: p.Rate, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &ConventionalInvoice{}, err
	}

	return p, nil
}

func (p *ConventionalInvoice) DeleteAConventionalInvoice(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&ConventionalInvoice{}).Where("id = ?", pid).Take(&ConventionalInvoice{}).Delete(&ConventionalInvoice{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ConventionalInvoice not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
