package handler

import (
	"mime/multipart"
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

func (service *ProductService) Create(id uint, uId uint, files []*multipart.FileHeader) serializer.Response {
	//var boss model.User
	//code := e.SUCCESS
	//
	//dao.DB.Model(&model.User{}).Where("id = ?",id).First(&boss)
	//tmp, _ := files[0].Open()
	//path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	//if err != nil {
	//	code = e.ErrorUploadFile
	//	return serializer.Response {
	//		Status: code,
	//		Data: e.GetMsg(code),
	//		Error: path,
	//	}
	//}
	//product := &model.Product{
	//	Name:          service.Name,
	//	CategoryID:    uint(service.CategoryID),
	//	Title:         service.Title,
	//	Info:          service.Info,
	//	ImgPath:       path,
	//	Price:         service.Price,
	//	DiscountPrice: service.DiscountPrice,
	//	Num:           service.Num,
	//	OnSale:        true,
	//	//BossID:        uId,
	//	BossName:      boss.UserName,
	//	//BossAvatar:    boss.Avatar,
	//}
	//code := e.SUCCESS
	//err = dao.DB.Create(&product).Error
	//if err != nil {
	//	//logging.Info(err)
	//	//code = e.ErrorDatabase
	//	return serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//		Error:  err.Error(),
	//	}
	//}
	//
	//for _,file := range files {
	//	tmp, _ := file.Open()
	//	status, info := UploadToQiNiu(tmp, file.Size)
	//	if status != 200 {
	//		return serializer.Response{
	//			Status: status,
	//			Data:   e.GetMsg(status),
	//			Error:  info,
	//		}
	//	}
	//	productImg := model.ProductImg{
	//		ProductID: product.ID,
	//		ImgPath:   info,
	//	}
	//	err = model.DB.Create(&productImg).Error
	//	if err != nil {
	//		code = e.ERROR
	//		return serializer.Response{
	//			Status: code,
	//			Msg:    e.GetMsg(code),
	//			Error:  err.Error(),
	//		}
	//	}
	//}
	//return serializer.Response{
	//	Status: code,
	//	Data:   serializer.BuildProduct(product),
	//	Msg:    e.GetMsg(code),
	//}
}


