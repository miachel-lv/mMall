package handler

import (
	"mime/multipart"
	"strconv"
	"sync"
	"user/dao"
	"user/model"
	"user/pkg/e"
	"user/serializer"
)

type ProductService struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	Num           int    `form:"num" json:"num"`
	PageNum       int    `form:"pageNum"`
	PageSize      int    `form:"pageSize"`
}

type ListProductImgService struct {

}

//func (service *ProductService) Show(id string) serializer.Response {
//}

func (service *ProductService) Create(uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss model.User
	code := e.SUCCESS

	dao.DB.Model(&model.User{}).Where("id = ?",uId).First(&boss)
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response {
			Status: code,
			Data: e.GetMsg(code),
			Error: path,
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryID),
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		Num:           service.Num,
		OnSale:        true,
		//BossID:        uId,
		BossName:      boss.UserName,
		//BossAvatar:    boss.Avatar,
	}
	//productDao := dao.NewProductDao(ctx)
	productDao := dao.DB
	err = productDao.Model(&model.Product{}).Create(&product).Error
	if err != nil {
		//logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		//productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		productImgDao := productDao
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorUploadFile
			return serializer.Response{
				Status: code,
				Data:   e.GetMsg(code),
				Error:  path,
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		//err = productImgDao.CreateProductImg(productImg)
		err = productImgDao.Model(&model.ProductImg{}).Create(&productImg).Error
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}

	wg.Wait()

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}


