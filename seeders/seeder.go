package seeders

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/laevenx/dompetkilat-test-server/models"
)

var users = []models.User{
	models.User{
		Username: "Admin",
		Email:    "steven@gmail.com",
		Password: "admin",
	},
	models.User{
		Username: "User",
		Email:    "luther@gmail.com",
		Password: "user",
	},
}

var invoices = []models.ProductiveInvoice{
	models.ProductiveInvoice{
		Name:   "PT YJK",
		Amount: 1000000,
		Grade:  "B",
		Rate:   16,
	},
	models.ProductiveInvoice{
		Name:   "PT KKY",
		Amount: 4000000,
		Grade:  "B+",
		Rate:   14,
	},
}

var conventionalinvoices = []models.ConventionalInvoice{
	models.ConventionalInvoice{
		Name:   "PT YJK",
		Amount: 1000000,
		Tenor:  120,
		Grade:  "B",
		Rate:   16,
	},
	models.ConventionalInvoice{
		Name:   "PT KKY",
		Amount: 4000000,
		Tenor:  120,
		Grade:  "B+",
		Rate:   14,
	},
}

var conventionalosf = []models.ConventionalOsf{
	models.ConventionalOsf{
		Name:   "PT YJK",
		Amount: 1000000,
		Tenor:  120,
		Grade:  "B",
		Rate:   16,
	},
	models.ConventionalOsf{
		Name:   "PT KKY",
		Amount: 4000000,
		Tenor:  120,
		Grade:  "B+",
		Rate:   14,
	},
}

var reksadana = []models.Reksadana{
	models.Reksadana{
		Name:   "INB",
		Amount: 20000000,
		Return: -1,
	},
	models.Reksadana{
		Name:   "PT TELMOM",
		Amount: 10000000,
		Return: 1,
	},
}

var sbn = []models.Sbn{
	models.Sbn{
		Name:   "SBR XXX",
		Amount: 1000000,
		Tenor:  120,
		Rate:   7,
		Type:   "SBR",
	},
	models.Sbn{
		Name:   "SBR YYY",
		Amount: 2000000,
		Tenor:  120,
		Rate:   8,
		Type:   "SBR",
	},
}

var finances = []models.Finance{
	// models.Finance{
	// 	Name:  "Invoice Financing",
	// 	Count: 35,
	// 	Sub:   "null",
	// },
	// models.Finance{
	// 	Name:  "OSF Financing",
	// 	Count: 25,
	// 	Sub:   "null",
	// },
	models.Finance{
		Name:  "SBN",
		Count: 10,
		Sub:   "null",
	},
	models.Finance{
		Name:  "Reksadana",
		Count: 20,
		Sub:   "null",
	},
	models.Finance{
		Name:  "Conventional Invoice",
		Count: 20,
		Sub:   "invoice",
	},
	models.Finance{
		Name:  "Productive Invoice",
		Count: 15,
		Sub:   "invoice",
	},
	models.Finance{
		Name:  "Conventional OSF",
		Count: 15,
		Sub:   "osf",
	},
	// models.Finance{
	// 	Name:  "Productive OSF",
	// 	Count: 10,
	// 	Sub:   "osf",
	// },
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Finance{}, &models.User{}, &models.ConventionalInvoice{}, &models.ConventionalOsf{}, &models.ProductiveInvoice{}, &models.Reksadana{}, &models.Sbn{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Finance{}, &models.User{}, &models.ConventionalInvoice{}, &models.ConventionalOsf{}, &models.ProductiveInvoice{}, &models.Reksadana{}, &models.Sbn{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range finances {
		err = db.Debug().Model(&models.Finance{}).Create(&finances[i]).Error
		if err != nil {
			log.Fatalf("cannot seed finance table: %v", err)
		}

	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}

	for i, _ := range sbn {
		err = db.Debug().Model(&models.Sbn{}).Create(&sbn[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}

	for i, _ := range reksadana {
		err = db.Debug().Model(&models.Reksadana{}).Create(&reksadana[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}

	for i, _ := range conventionalinvoices {
		err = db.Debug().Model(&models.ConventionalInvoice{}).Create(&conventionalinvoices[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}

	for i, _ := range conventionalosf {
		err = db.Debug().Model(&models.ConventionalOsf{}).Create(&conventionalosf[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}

	for i, _ := range invoices {
		err = db.Debug().Model(&models.ProductiveInvoice{}).Create(&invoices[i]).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}

	}
}
