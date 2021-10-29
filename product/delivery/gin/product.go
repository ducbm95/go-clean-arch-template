package gin

import (
	"github.com/ducbm95/go-clean-arch-template/domain"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService domain.ProductService
}

func NewProductController(productService domain.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (pc *ProductController) Create(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
