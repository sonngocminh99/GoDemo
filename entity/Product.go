package entity

type ProductType struct {
	ID   int    `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
}

type ProductCategory struct {
	ID            int         `gorm:"primary_key:auto_increment" json:"id"`
	Name          string      `gorm:"type:varchar(300)" json:"name"`
	ProductTypeID int         `gorm:"not null" json:"product_type_id"`
	ProductType   ProductType `gorm:"foreignkey:ProductTypeID;constant:onUpdate:CASCADE,onDelete:CASCADE" json:"product_type"`
}

type Product struct {
	ID            int             `gorm:"primary_key:auto_increment" json:"id"`
	Name          string          `gorm:"type:varchar(300)" json:"name"`
	Description   string          `gorm:"type:longtext" json:"description"`
	Price         int             `gorm:"type:int" json:"price"`
	Sold          int             `gorm:"type:int" json:"sold"`
	Instruction   string          `gorm:"type:longtext" json:"instruction"`
	Origin        string          `gorm:"type:varchar(50)" json:"origin"`
	Is_active     int             `gorm:"type:tinyint(1)" json:"Isactive"`
	Is_flash_sale int             `gorm:"type:tinyint(1)" json:"IsFlashsale"`
	PriceSale     int             `gorm:"type:int" json:"priceSale"`
	CategoryID    int             `gorm:"not null" json:"category_id"`
	Category      ProductCategory `gorm:"foreignkey:CategoryID;constant:onUpdate:CASCADE,onDelete:CASCADE" json:"category"`
	Quantity      int             `gorm:"type:int" json:"quantity"`
	ImagePresent  string          `gorm:"type:varchar(100)" json:"imagepresent"`
	Ingredient    string          `gorm:"type:longtext" json:"Ingredient"`
	Brand         string          `gorm:"type:varchar(500)" json:"brand"`
}

type Tabler interface {
	TableName() string
}

func (Product) TableName() string {
	return "ecommerce_product"
}
