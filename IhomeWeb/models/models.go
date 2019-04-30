package models
import (
//	beego的orm模块
	"github.com/astaxie/beego/orm"
	"study/students_autoerrBooks/IhomeWeb/utils"
	"time"
)
/* 用户 table_name = user */
type User struct {
	Id           int    `json:"user_id"`                       //用户编号
	Name         string `orm:"size(32)" json:"name"`           // 用户名
	PasswordHash string `orm:"size(128)" json:"password_hash"` //用户密码加密处理
	Mobile       string `orm:"size(11);unique"  json:"mobile"` //手机号
	RealName     string `orm:"size(32)" json:"real_name"`      //真实姓名  实名认证
	IdCard       string `orm:"size(20)" json:"id_card"`        //身份证号  实名认证
	AvatarUrl    string `orm:"size(256)" json:"avatar_url"`    //用户头像路径       通过fastdfs进行图片存储
	Houses       []*House      `orm:"reverse(many)" json:"houses"`    //用户发布的房屋信息  一个人多套房
	Orders       []*OrderHouse `orm:"reverse(many)" json:"orders"`    //用户下的订单       一个人多次订单
}
/* 房屋信息 table_name = house */
type House struct {
	Id              int           `json:"house_id"`                                          //房屋编号
	User      *User     `orm:"rel(fk)" json:"user_id"`                                       //房屋主人的用户编号  与用户进行关联
	Area      *Area     `orm:"rel(fk)" json:"area_id"`                                       //归属地的区域编号   和地区表进行关联
	Title     string    `orm:"size(64)" json:"title"`                                        //房屋标题
	Price     int       `orm:"default(0)" json:"price"`                                      //单价,单位:分   每次的价格要乘以100
	Address   string    `orm:"size(512)" orm:"default("")" json:"address"`                   //地址
	RoomCount int       `orm:"default(1)" json:"room_count"`                                 //房间数目
	Acreage   int       `orm:"default(0)" json:"acreage"`                                    //房屋总面积
	Unit      string    `orm:"size(32)" orm:"default("")" json:"unit"`                       //房屋单元,如 几室几厅
	Capacity  int       `orm:"default(1)" json:"capacity"`                                   //房屋容纳的总人数
	Beds      string    `orm:"size(64)" orm:"default("")" json:"beds"`                       //房屋床铺的配置
	Deposit   int       `orm:"default(0)" json:"deposit"`                                    //押金
	Min_days        int           `orm:"default(1)" json:"min_days"`                         //最少入住的天数
	Max_days        int           `orm:"default(0)" json:"max_days"`                         //最多入住的天数 0表示不限制
	Order_count     int           `orm:"default(0)" json:"order_count"`                      //预定完成的该房屋的订单数
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"` //房屋主图片路径
	Facilities      []*Facility   `orm:"reverse(many)" json:"facilities"`                    //房屋设施   与设施表进行关联
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`                      //房屋的图片   除主要图片之外的其他图片地址
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`                        //房屋的订单    与房屋表进行管理
	Ctime     time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

//首页最高展示的房屋数量
var HOME_PAGE_MAX_HOUSES int = 5

//房屋列表页面每页显示条目数
var HOUSE_LIST_PAGE_CAPACITY int = 2
//处理房子信息
func (this *House) ToHouseInfo() interface{} {
	houseInfo := map[string]interface{}{
		"house_id":    this.Id,
		"title":       this.Title,
		"price":       this.Price,
		"area_name":   this.Area.Name,
		"img_url":     utils.AddDomain2Url(this.Index_image_url),
		"room_count":  this.RoomCount,
		"order_count": this.Order_count,
		"address":     this.Address,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
		"ctime":       this.Ctime.Format("2006-01-02 15:04:05"),
	}

	return houseInfo
}
//处理1个房子的全部信息
func(this *House)ToOneHouseDesc()interface{}  {
	houseDesc:=map[string]interface{}{
		"hid":this.Id,
		"user_id":     this.User.Id,
		"user_name":   this.User.Name,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
		"title":       this.Title,
		"price":       this.Price,
		"address":     this.Address,
		"room_count":  this.Room_count,
		"acreage":     this.Acreage,
		"unit":        this.Unit,
		"capacity":    this.Capacity,
		"beds":        this.Beds,
		"deposit":     this.Deposit,
		"min_days":    this.Min_days,
		"max_days":    this.Max_days,
	}
	/* 区域信息 table_name = area */  //区域信息是需要我们手动添加到数据库中的
	type Area struct {
		Id     int      `json:"aid"`                        //区域编号     1    2
		Name   string   `orm:"size(32)" json:"aname"`       //区域名字     昌平 海淀
		Houses []*House `orm:"reverse(many)" json:"houses"` //区域所有的房屋   与房屋表进行关联
	}

	/* 设施信息 table_name = "facility"*/     //设施信息 需要我们提前手动添加的
	type Facility struct {
		Id     int      `json:"fid"`     //设施编号
		Name   string   `orm:"size(32)"` //设施名字
		Houses []*House `orm:"rel(m2m)"` //都有哪些房屋有此设施  与房屋表进行关联的
	}

	/* 房屋图片 table_name = "house_image"*/
	type HouseImage struct {
		Id    int    `json:"house_image_id"`         //图片id
		Url   string `orm:"size(256)" json:"url"`    //图片url     存放我们房屋的图片
		House *House `orm:"rel(fk)" json:"house_id"` //图片所属房屋编号
	}

	const (
		ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
		ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
		ORDER_STATUS_PAID         = "PAID"         //已支付
		ORDER_STATUS_WAIT_COMMENT = "WAIT_COMMENT" //待评价
		ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
		ORDER_STATUS_CANCELED     = "CONCELED"     //已取消
		ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
	)
	/* 订单 table_name = order */
	type orderHouse struct {
		Id        int `json:"order_id"`                                  //订单编号
		User      *User `orm:"rel(fk)" json:"user_id"`                   //下单的用户编号   //与用户表进行关联
		House     *House    `orm:"rel(fk)" json:"house_id"`              //预定的房间编号   //与房屋信息进行关联
		BeginDate time.Time `orm:"type(datetime)"`                       //预定的起始时间
		EndDate   time.Time `orm:"type(datetime)"`                       //预定的结束时间
		Days      int                                                    //预定总天数
HousePrice        int                                                    //房屋的单价
Amount            int                                                    //订单总金额
Status            string    `orm:"default(WAIT_ACCEPT)"`                 //订单状态
Comment           string    `orm:"size(512)"`                            //订单评论
Ctime             time.Time `orm:"auto_now;type(datetime)" json:"ctime"` //每次更新此表，都会更新这个字段
Credit            bool                                                   //表示个人征信情况 true表示良好
	}
	
	//房屋图片
	var imgUrls []string
	for _,imgUrl := range this.Images {
		imgUrls =append(imgUrls,utils.AddDomain2Url(imgUrl.))
	}


}