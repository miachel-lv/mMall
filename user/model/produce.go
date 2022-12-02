package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"user/cache"
)

const (
	//每日排名
	RankKey = "rank"
	//ElectricalRank 家电排名
	ElectricalRank = "elecRank"
	//AccessoryRank 配件排名
	AccessoryRank = "acceRank"
)

type Product struct {
	gorm.Model
	Name 			string `gorm:"size:255;index"`
	CategoryID 		uint `grom:"not null"`
	Title 			string
	Info 			string `gorm:"size:1000"`
	ImgPath 		string
	Price 			string
	DiscountPrice 	string
	OnSale          bool `gorm:"default:false"`
	Num           	int
	BossID        	int
	BossName      	string
	BossAvatar    	string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//商品游览
func (product *Product) AddView() {
	//增加视频点击数
	cache.RedisClient.Incr(ProductViewKey(product.ID))
	//增加排行点击数
	cache.RedisClient.ZIncrBy(RankKey, 1, strconv.Itoa(int(product.ID)))
}

//增加配件排行点击数
func (product *Product) AddAcceRank() {
	cache.RedisClient.ZIncrBy(AccessoryRank, 1, strconv.Itoa(int(product.ID)))
}

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}


