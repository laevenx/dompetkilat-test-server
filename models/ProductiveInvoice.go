package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductiveInvoice struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Amount    uint32    `gorm:"size:255;not null;" json:"amount"`
	Grade     string    `gorm:"size:255;not null;" json:"grade"`
	Rate      uint32    `gorm:"size:255;not null;" json:"rate"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *ProductiveInvoice) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Grade = html.EscapeString(strings.TrimSpace(p.Grade))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *ProductiveInvoice) Validate() error {

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

func (p *ProductiveInvoice) SaveProductiveInvoice(db *gorm.DB) (*ProductiveInvoice, error) {
	var err error
	err = db.Debug().Model(&ProductiveInvoice{}).Create(&p).Error
	if err != nil {
		return &ProductiveInvoice{}, err
	}

	return p, nil
}

func (p *ProductiveInvoice) FindAllProductiveInvoices(db *gorm.DB) (*[]ProductiveInvoice, error) {
	var err error
	posts := []ProductiveInvoice{}
	err = db.Debug().Model(&ProductiveInvoice{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]ProductiveInvoice{}, err
	}

	return &posts, nil
}

func (p *ProductiveInvoice) FindProductiveInvoiceByID(db *gorm.DB, pid uint64) (*ProductiveInvoice, error) {
	var err error
	err = db.Debug().Model(&ProductiveInvoice{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductiveInvoice{}, err
	}

	return p, nil
}

func (p *ProductiveInvoice) UpdateAProductiveInvoice(db *gorm.DB) (*ProductiveInvoice, error) {

	var err error

	err = db.Debug().Model(&ProductiveInvoice{}).Where("id = ?", p.ID).Updates(ProductiveInvoice{Name: p.Name, Amount: p.Amount, Grade: p.Grade, Rate: p.Rate, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &ProductiveInvoice{}, err
	}

	return p, nil
}

func (p *ProductiveInvoice) DeleteAProductiveInvoice(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&ProductiveInvoice{}).Where("id = ?", pid).Take(&ProductiveInvoice{}).Delete(&ProductiveInvoice{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductiveInvoice not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
