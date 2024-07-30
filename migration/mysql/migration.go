package mysql

import (
	"time"

	"gorm.io/gorm"
)

type TblEcomCustomer struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	MemberId         int       `gorm:"type:int"`
	FirstName        string    `gorm:"type:varchar(255)"`
	LastName         string    `gorm:"type:varchar(255)"`
	Email            string    `gorm:"type:varchar(255)"`
	MobileNo         string    `gorm:"type:varchar(255)"`
	Username         string    `gorm:"type:varchar(255)"`
	Password         string    `gorm:"type:varchar(255)"`
	StreetAddress    string    `gorm:"type:varchar(255)"`
	City             string    `gorm:"type:varchar(255)"`
	State            string    `gorm:"type:varchar(255)"`
	Country          string    `gorm:"type:varchar(255)"`
	ZipCode          string    `gorm:"type:varchar(255)"`
	IsActive         int       `gorm:"type:int"`
	CreatedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int"`
	DeletedOn        time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:int;DEFAULT:NULL"`
	Count            int       `gorm:"-:migration;<-:false"`
	ProfileImage     string    `gorm:"type:varchar(255)"`
	ProfileImagePath string    `gorm:"type:varchar(255)"`
	NameString       string    `gorm:"-:migration;<-:false"`
	ShippingAddress  string    `gorm:"-:migration;<-:false"`
	TenantId         int       `gorm:"type:int"`
}

type TblEcomProduct struct {
	Id                 int `gorm:"primaryKey;auto_increment"`
	CategoriesId       string
	ProductName        string
	ProductSlug        string
	ProductDescription string
	ProductImagePath   string
	ProductYoutubePath string
	ProductVimeoPath   string
	Sku                string
	ProductPrice       int
	Tax                int
	Totalcost          int
	Priority           int       `gorm:"-:migration;<-:false"`
	Price              int       `gorm:"-:migration;<-:false"`
	StartDate          time.Time `gorm:"-:migration;<-:false"`
	EndDate            time.Time `gorm:"-:migration;<-:false"`
	Type               string    `gorm:"-:migration;<-:false"`
	Quantity           int       `gorm:"-:migration;<-:false"`
	Order_id           int       `gorm:"-:migration;<-:false"`
	Product_id         int       `gorm:"-:migration;<-:false"`
	Quantityprice      int       `gorm:"-:migration;<-:false"`
	IsActive           int
	Stock              int
	CreatedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy          int       `gorm:"DEFAULT:NULL"`
	ModifiedOn         time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted          int       `gorm:"DEFAULT:0"`
	DeletedOn          time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	Imgpath            []string  `gorm:"-"`
	TenantId           int       `gorm:"type:int"`
}

type TblEcomProductPricing struct {
	Id        int `gorm:"primaryKey;auto_increment"`
	ProductId int
	Priority  int
	Price     int
	StartDate time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	EndDate   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	Type      string
	IsDeleted int       `gorm:"DEFAULT:0"`
	Startdate string    `gorm:"-:migration;<-:false"`
	Enddate   string    `gorm:"-:migration;<-:false"`
	DeletedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:int;DEFAULT:NULL"`
	TenantId  int       `gorm:"type:int"`
}

type TblEcomProductOrder struct {
	Id              int                  `gorm:"primaryKey;auto_increment"`
	Uuid            string               `gorm:"type:varchar(255)"`
	CustomerId      int                  `gorm:"type:int"`
	OrderStatus     int                  `gorm:"type:int"`
	ShippingAddress string               `gorm:"type:varchar(255)"`
	IsDeleted       int                  `gorm:"type:int"`
	Username        string               `gorm:"-:migration;<-:false"`
	Email           string               `gorm:"-:migration;<-:false"`
	MobileNo        string               `gorm:"-:migration;<-:false"`
	StreetAddress   string               `gorm:"-:migration;<-:false"`
	City            string               `gorm:"-:migration;<-:false"`
	State           string               `gorm:"-:migration;<-:false"`
	Country         string               `gorm:"-:migration;<-:false"`
	ZipCode         string               `gorm:"-:migration;<-:false"`
	CreatedOn       time.Time            `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedOn      time.Time            `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedDate    string               `gorm:"-:migration"`
	CreatedDate     string               `gorm:"-:migration"`
	Price           int                  `gorm:"type:int"`
	Tax             int                  `gorm:"type:int"`
	TotalCost       int                  `gorm:"type:int"`
	FirstName       string               `gorm:"-:migration;<-:false"`
	LastName        string               `gorm:"-:migration;<-:false"`
	NameString      string               `gorm:"-:migration;<-:false"`
	Orders          []TblEcomOrderStatus `gorm:"foreignKey:OrderId;references:Id"`
	StatusValue     string               `gorm:"-:migration;<-:false"`
	StatusPriority  int                  `gorm:"-:migration;<-:false"`
	StatusColor     string               `gorm:"-:migration;<-:false"`
	DeletedOn       time.Time            `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy       int                  `gorm:"type:int;DEFAULT:NULL"`
	TenantId        int                  `gorm:"type:int"`
}

type TblEcomProductOrderDetail struct {
	Id         int `gorm:"primaryKey;auto_increment"`
	Order_id   int `gorm:"type:int"`
	Product_id int `gorm:"type:int"`
	Quantity   int `gorm:"type:int"`
	Price      int `gorm:"type:int"`
	Tax        int `gorm:"type:int"`
	TenantId   int `gorm:"type:int"`
}

type TblEcomOrderStatus struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	OrderId     int       `gorm:"type:int"`
	OrderStatus int       `gorm:"type:int"`
	CreatedBy   int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
	TenantId    int       `gorm:"type:int"`
}

type TblEcomCurrency struct {
	Id              int    `gorm:"primaryKey;auto_increment"`
	CurrencyName    string `gorm:"type:varchar(255)"`
	CurrencyType    string `gorm:"type:varchar(255)"`
	CurrencySymbol  string `gorm:"type:varchar(255)"`
	IsActive        int    `gorm:"type:int"`
	CurrencyDefault int    `gorm:"type:int"`
	CreatedOn       time.Time
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted       int       `gorm:"type:int"`
	DeletedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"type:int"`
	DateString      string    `gorm:"-"`
	TenantId        int       `gorm:"type:int"`
}

type TblEcomStatus struct {
	Id          int    `gorm:"primaryKey;auto_increment"`
	Status      string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	IsActive    int    `gorm:"type:int"`
	Priority    int    `gorm:"type:int"`
	ColorCode   string `gorm:"type:varchar(255)"`
	CreatedOn   time.Time
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int"`
	DeletedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:int"`
	TenantId    int       `gorm:"type:int"`
}

type TblEcomPayment struct {
	Id           int    `gorm:"primaryKey;auto_increment"`
	PaymentName  string `gorm:"type:varchar(255)"`
	Description  string `gorm:"type:varchar(255)"`
	PaymentImage string `gorm:"type:varchar(255)"`
	IsActive     int    `gorm:"type:int"`
	CreatedOn    time.Time
	CreatedBy    int
	ModifiedOn   time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:int"`
	DeletedOn    time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:int"`
	TenantId     int       `gorm:"type:int"`
}

type TblEcomSettings struct {
	Id              int       `gorm:"primaryKey;auto_increment"`
	StoreName       string    `gorm:"type:varchar(255)"`
	DisplayStock    int       `gorm:"type:varchar(255)"`
	StockWarning    int       `gorm:"type:varchar(255)"`
	StockCheckout   int       `gorm:"type:varchar(255)"`
	CurrencyDefault int       `gorm:"type:varchar(255)"`
	PaymentDefault  int       `gorm:"type:varchar(255)"`
	StatusDefault   int       `gorm:"type:varchar(255)"`
	CreatedOn       time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	TenantId        int       `gorm:"type:int"`
}

type TblEcomOrderPayment struct {
	Id          int    `gorm:"primaryKey;auto_increment"`
	OrderId     int    `gorm:"type:int"`
	PaymentMode string `gorm:"type:varchar(255)"`
	TenantId    int    `gorm:"type:int"`
}

type TblEcomCart struct {
	Id         int       `gorm:"primaryKey;auto_increment;"`
	ProductId  int       `gorm:"type:int"`
	CustomerId int       `gorm:"type:int"`
	Quantity   int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	ModifiedOn time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int"`
	DeletedOn  time.Time `gorm:"type:datetime;DEFAULT:NULL"`
	TenantId   int       `gorm:"type:int"`
}

// MigrateTable creates this package related tables in your database
func MigrationTables(db *gorm.DB) {

	db.AutoMigrate(
		&TblEcomCustomer{},
		&TblEcomOrderStatus{},
		&TblEcomProductOrderDetail{},
		&TblEcomProductOrder{},
		&TblEcomProductPricing{},
		&TblEcomProduct{},
		&TblEcomCurrency{},
		&TblEcomStatus{},
		&TblEcomPayment{},
		&TblEcomSettings{},
		&TblEcomOrderPayment{},
		&TblEcomCart{},
	)

	db.Exec(`CREATE INDEX IF NOT EXISTS email_unique
    ON public.tbl_ecom_customers USING btree
    (email COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default
    WHERE is_deleted = 0;`)

	db.Exec(`CREATE INDEX IF NOT EXISTS mobile_no_unique
    ON public.tbl_ecom_customers USING btree
    (mobile_no COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default
    WHERE is_deleted = 0;`)

}
