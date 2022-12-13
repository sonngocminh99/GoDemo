package dto

type ProductUpdateDTO struct {
	Id            int    `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
	Price         int    `json:"price" form:"price" binding:"required"`
	Sold          int    `json:"sold" form:"sold" binding:"required"`
	Instruction   string `json:"instruction" form:"instruction" binding:"required"`
	Origin        string `json:"origin" form:"origin" binding:"required"`
	Is_active     int    `json:"Isactive" form:"Isactive" binding:"required"`
	Is_flash_sale int    `json:"IsFlashsale" form:"IsFlashsale" binding:"required"`
	PriceSale     int    `json:"priceSale" form:"priceSale" binding:"required"`
	CategoryID    int    `json:"category_id,omitempty" form:"category_id,omitempty" binding:"required"`
	Quantity      int    `json:"quantity" form:"quantity" binding:"required"`
	ImagePresent  string `json:"imagepresent" form:"imagepresent"`
	Ingredient    string `json:"ingredient" form:"ingredient"`
	Brand         string `json:"brand" form:"brand" binding:"required"`
}

type ProductCreateDTO struct {
	Id            int    `json:"id" form:"id" binding:"required"`
	Name          string `json:"name" form:"name" binding:"required"`
	Description   string `json:"description" form:"description" binding:"required"`
	Price         int    `json:"price" form:"price" binding:"required"`
	Sold          int    `json:"sold" form:"sold" binding:"required"`
	Instruction   string `json:"instruction" form:"instruction" binding:"required"`
	Origin        string `json:"origin" form:"origin" binding:"required"`
	Is_active     int    `json:"Isactive" form:"Isactive" binding:"required"`
	Is_flash_sale int    `json:"IsFlashsale" form:"IsFlashsale" binding:"required"`
	PriceSale     int    `json:"priceSale" form:"priceSale" binding:"required"`
	CategoryID    int    `json:"category_id,omitempty" form:"category_id,omitempty" binding:"required"`
	Quantity      int    `json:"quantity" form:"quantity" binding:"required"`
	ImagePresent  string `json:"imagepresent" form:"imagepresent"`
	Ingredient    string `json:"ingredient" form:"ingredient"`
	Brand         string `json:"brand" form:"brand" binding:"required"`
}
