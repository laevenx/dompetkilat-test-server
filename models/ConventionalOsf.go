package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ConventionalOsf struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;uniquee" json:"name"`
	Amount    uint32    `gorm:"size:255;not null;" json:"amount"`
	Tenor     uint32    `gorm:"size:255;not null;" json:"tenor"`
	Grade     string    `gorm:"size:255;not null;" json:"grade"`
	Rate      uint32    `gorm:"size:255;not null;" json:"rate"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *ConventionalOsf) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Grade = html.EscapeString(strings.TrimSpace(p.Grade))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *ConventionalOsf) Validate() error {

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

func (p *ConventionalOsf) SaveConventionalOsf(db *gorm.DB) (*ConventionalOsf, error) {
	var err error
	err = db.Debug().Model(&ConventionalOsf{}).Create(&p).Error
	if err != nil {
		return &ConventionalOsf{}, err
	}

	return p, nil
}

func (p *ConventionalOsf) FindAllConventionalOsfs(db *gorm.DB) (*[]ConventionalOsf, error) {
	var err error
	posts := []ConventionalOsf{}
	err = db.Debug().Model(&ConventionalOsf{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]ConventionalOsf{}, err
	}

	return &posts, nil
}

func (p *ConventionalOsf) FindConventionalOsfByID(db *gorm.DB, pid uint64) (*ConventionalOsf, error) {
	var err error
	err = db.Debug().Model(&ConventionalOsf{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &ConventionalOsf{}, err
	}

	return p, nil
}

func (p *ConventionalOsf) UpdateAConventionalOsf(db *gorm.DB) (*ConventionalOsf, error) {

	var err error

	err = db.Debug().Model(&ConventionalOsf{}).Where("id = ?", p.ID).Updates(ConventionalOsf{Name: p.Name, Amount: p.Amount, Tenor: p.Tenor, Grade: p.Grade, Rate: p.Rate, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &ConventionalOsf{}, err
	}

	return p, nil
}

func (p *ConventionalOsf) DeleteAConventionalOsf(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&ConventionalOsf{}).Where("id = ?", pid).Take(&ConventionalOsf{}).Delete(&ConventionalOsf{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ConventionalOsf not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
