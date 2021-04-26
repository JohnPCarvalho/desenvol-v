package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Travel struct {
	ID            uint64 `gorm:"primary_key;auto_increment"`
	DriverID      uint64
	TravelledKm   float32   `json:"travelledkm"`
	LiterSpent    float32   `json:"literSpent"`
	PricePerLiter float32   `json:"priceperliter"`
	CheckoutDate  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"checkout"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Travel) Prepare() {
	t.ID = 0
	t.CheckoutDate = time.Now()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func (t *Travel) Validate() error {

	if t.TravelledKm == 0 {
		return errors.New("required travelled km")
	}

	if t.LiterSpent == 0 {
		return errors.New("required liter spent")
	}

	if t.PricePerLiter == 0 {
		return errors.New("required price per liter")
	}

	if t.CheckoutDate.IsZero() {
		return errors.New("empty date")
	}

	return nil
}

func (t *Travel) SaveTravel(db *gorm.DB) (*Travel, error) {
	var err error
	err = db.Debug().Model(&Travel{}).Create(&t).Error
	fmt.Print("Tentou criar \n")
	if err != nil {
		fmt.Print("Returned in the first call \n")
		return &Travel{}, err
	}
	if t.ID != 0 {
		fmt.Print("deu erro no t.id = 0 \n")
		err = db.Debug().Model(&User{}).Where("id = ?", t.ID).Take(&t.ID).Error
		if err != nil {
			return &Travel{}, err
		}
	}
	return t, nil
}

func (t *Travel) FindAllTravels(db *gorm.DB) (*[]Travel, error) {
	var err error
	Travels := []Travel{}
	err = db.Debug().Model(&Travel{}).Limit(100).Find(&Travels).Error
	if err != nil {
		return &[]Travel{}, err
	}
	if len(Travels) > 0 {
		for i, _ := range Travels {
			err := db.Debug().Model(&User{}).Where("id = ?", Travels[i].ID).Take(&Travels[i].ID).Error
			if err != nil {
				return &[]Travel{}, err
			}
		}
	}
	return &Travels, nil
}

func (t *Travel) FindTravelByID(db *gorm.DB, pid uint64) (*Travel, error) {
	var err error
	err = db.Debug().Model(&Travel{}).Where("id = ?", pid).Take(&t).Error
	if err != nil {
		return &Travel{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.ID).Take(&t.ID).Error
		if err != nil {
			return &Travel{}, err
		}
	}
	return t, nil
}

func (t *Travel) UpdateATravel(db *gorm.DB) (*Travel, error) {

	var err error

	err = db.Debug().Model(&Travel{}).Where("id = ?", t.ID).Updates(Travel{TravelledKm: t.TravelledKm, LiterSpent: t.LiterSpent, PricePerLiter: t.PricePerLiter, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Travel{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.ID).Take(&t.ID).Error
		if err != nil {
			return &Travel{}, err
		}
	}
	return t, nil
}

func (p *Travel) DeleteATravel(db *gorm.DB, pid uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Travel{}).Where("id = ? and author_id = ?", pid, uid).Take(&Travel{}).Delete(&Travel{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Travel not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
