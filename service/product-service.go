package service

import (
	"github.com/mashingan/smapping"
	"github.com/sonngocminh99/GoDemo/dto"
	"github.com/sonngocminh99/GoDemo/entity"
	"github.com/sonngocminh99/GoDemo/repository"
)

type ProductService interface {
	InsertProduct(b dto.ProductCreateDTO) entity.Product
	UpdateProduct(b dto.ProductUpdateDTO) entity.Product
	DeleteProduct(b entity.Product)
	AllProducts() []entity.Product
	FindProductById(id int) entity.Product
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{
		repository: repository,
	}
}

func (service *productService) InsertProduct(b dto.ProductCreateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		panic(err)
	}
	res := service.repository.InsertProduct(product)
	return res
}

func (service *productService) UpdateProduct(b dto.ProductUpdateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&b))
	if err != nil {
		panic(err)
	}

	res := service.repository.UpdateProduct(product)
	return res
}

func (service *productService) DeleteProduct(b entity.Product) {
	service.repository.DeleteProduct(b)
}

func (service *productService) AllProducts() []entity.Product {
	return service.repository.AllProducts()
}

func (service *productService) FindProductById(id int) entity.Product {
	return service.repository.FindProductById(id)
}
