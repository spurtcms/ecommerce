package ecommerce

import (
	"time"

	"gorm.io/gorm"
)

type TblEcomCustomers struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId         int       `gorm:"type:integer"`
	FirstName        string    `gorm:"type:character varying"`
	LastName         string    `gorm:"type:character varying"`
	Email            string    `gorm:"type:character varying"`
	MobileNo         string    `gorm:"type:character varying"`
	Username         string    `gorm:"type:character varying"`
	Password         string    `gorm:"type:character varying"`
	StreetAddress    string    `gorm:"type:character varying"`
	City             string    `gorm:"type:character varying"`
	State            string    `gorm:"type:character varying"`
	Country          string    `gorm:"type:character varying"`
	ZipCode          string    `gorm:"type:character varying"`
	IsActive         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	Count            int       `gorm:"-:migration;<-:false"`
	ProfileImage     string    `gorm:"type:character varying"`
	ProfileImagePath string    `gorm:"type:character varying"`
	NameString       string    `gorm:"-:migration;<-:false"`
	ShippingAddress  string    `gorm:"-:migration;<-:false"`
}

type TblEcomProducts struct {
	Id                 int `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId       string
	ProductName        string
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
	CreatedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy          int       `gorm:"DEFAULT:NULL"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted          int       `gorm:"DEFAULT:0"`
	DeletedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	Imgpath            []string  `gorm:"-"`
}

type TblEcomProductPricings struct {
	Id        int `gorm:"primaryKey;auto_increment;type:serial"`
	ProductId int
	Priority  int
	Price     int
	StartDate time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	EndDate   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Type      string
	IsDeleted int       `gorm:"DEFAULT:0"`
	Startdate string    `gorm:"-:migration;<-:false"`
	Enddate   string    `gorm:"-:migration;<-:false"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
}

type TblEcomProductOrders struct {
	Id              int                  `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid            string               `gorm:"type:character varying"`
	CustomerId      int                  `gorm:"type:integer"`
	OrderStatus     int                  `gorm:"type:integer"`
	ShippingAddress string               `gorm:"type:character varying"`
	IsDeleted       int                  `gorm:"type:integer"`
	Username        string               `gorm:"-:migration;<-:false"`
	Email           string               `gorm:"-:migration;<-:false"`
	MobileNo        string               `gorm:"-:migration;<-:false"`
	StreetAddress   string               `gorm:"-:migration;<-:false"`
	City            string               `gorm:"-:migration;<-:false"`
	State           string               `gorm:"-:migration;<-:false"`
	Country         string               `gorm:"-:migration;<-:false"`
	ZipCode         string               `gorm:"-:migration;<-:false"`
	CreatedOn       time.Time            `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedOn      time.Time            `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedDate    string               `gorm:"-:migration"`
	CreatedDate     string               `gorm:"-:migration"`
	Price           int                  `gorm:"type:integer"`
	Tax             int                  `gorm:"type:integer"`
	TotalCost       int                  `gorm:"type:integer"`
	FirstName       string               `gorm:"-:migration;<-:false"`
	LastName        string               `gorm:"-:migration;<-:false"`
	NameString      string               `gorm:"-:migration;<-:false"`
	Orders          []TblEcomOrderStatus `gorm:"foreignKey:OrderId;references:Id"`
	StatusValue     string               `gorm:"-:migration;<-:false"`
	StatusPriority  int                  `gorm:"-:migration;<-:false"`
	StatusColor     string               `gorm:"-:migration;<-:false"`
	DeletedOn       time.Time            `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int                  `gorm:"type:integer;DEFAULT:NULL"`
}

type TblEcomProductOrderDetails struct {
	Id         int `gorm:"primaryKey;auto_increment;type:serial"`
	Order_id   int `gorm:"type:integer"`
	Product_id int `gorm:"type:integer"`
	Quantity   int `gorm:"type:integer"`
	Price      int `gorm:"type:integer"`
}

type TblEcomOrderStatus struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	OrderId     int       `gorm:"type:integer"`
	OrderStatus int       `gorm:"type:integer"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
}

type TblEcomCurrency struct {
	Id              int    `gorm:"primaryKey;auto_increment;type:serial"`
	CurrencyName    string `gorm:"type:character varying"`
	CurrencyType    string `gorm:"type:character varying"`
	CurrencySymbol  string `gorm:"type:character varying"`
	IsActive        int    `gorm:"type:integer"`
	CurrencyDefault int    `gorm:"type:integer"`
	CreatedOn       time.Time
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted       int       `gorm:"type:integer"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"type:integer"`
	DateString      string    `gorm:"-"`
}

type TblEcomStatus struct {
	Id          int    `gorm:"primaryKey;auto_increment;type:serial"`
	Status      string `gorm:"type:character varying"`
	Description string `gorm:"type:character varying"`
	IsActive    int    `gorm:"type:integer"`
	Priority    int    `gorm:"type:integer"`
	ColorCode   string `gorm:"type:character varying"`
	CreatedOn   time.Time
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer"`
}

type TblEcomPayment struct {
	Id           int    `gorm:"primaryKey;auto_increment;type:serial"`
	PaymentName  string `gorm:"type:character varying"`
	Description  string `gorm:"type:character varying"`
	PaymentImage string `gorm:"type:character varying"`
	IsActive     int    `gorm:"type:integer"`
	CreatedOn    time.Time
	CreatedBy    int
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:integer"`
}

type TblEcomSettings struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	StoreName       string    `gorm:"type:character varying"`
	DisplayStock    int       `gorm:"type:integer"`
	StockWarning    int       `gorm:"type:integer"`
	StockCheckout   int       `gorm:"type:integer"`
	CurrencyDefault int       `gorm:"type:integer"`
	PaymentDefault  int       `gorm:"type:integer"`
	StatusDefault   int       `gorm:"type:integer"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
}

// MigrateTable creates this package related tables in your database
func MigrateTables(db *gorm.DB) {

	db.AutoMigrate(&TblEcomCustomers{}, &TblEcomOrderStatus{}, &TblEcomProductOrderDetails{}, &TblEcomProductOrders{}, &TblEcomProductPricings{}, &TblEcomProducts{}, TblEcomCurrency{},
		TblEcomStatus{}, TblEcomPayment{}, TblEcomSettings{})

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
