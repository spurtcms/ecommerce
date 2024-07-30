package ecommerce

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spurtcms/ecommerce/migration"
	"github.com/spurtcms/member"
	"golang.org/x/crypto/bcrypt"
)

type ShippingAddress struct {
	Name    string `json:"name"`
	Area    string `json:"area"`
	Number  string `json:"number"`
	Email   string `json:"email"`
	HouseNo string `json:"houseno"`
	City    string `json:"city"`
	Country string `json:"country"`
	States  string `json:"states"`
}

type TblEcomCustomer struct {
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
	TenantId         int       `gorm:"type:integer"`
}

type TblEcomProduct struct {
	Id                 int `gorm:"primaryKey;auto_increment;type:serial"`
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
	CreatedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy          int       `gorm:"DEFAULT:NULL"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted          int       `gorm:"DEFAULT:0"`
	DeletedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	Imgpath            []string  `gorm:"-"`
	TenantId           int       `gorm:"type:integer"`
}

type TblEcomProductOrder struct {
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
	TenantId        int                  `gorm:"type:integer"`
}

type TblEcomOrderStatus struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	OrderId     int       `gorm:"type:integer"`
	OrderStatus int       `gorm:"type:integer"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
	TenantId    int       `gorm:"type:integer"`
}

// EcommerceSetup used initialize Ecommerce configruation
func EcommerceSetup(config Config) *Ecommerce {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Ecommerce{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}

// Member package connection

func (ecommerce *Ecommerce) DBconf() *member.Member {
	var memberconfig = member.MemberSetup(member.Config{DB: ecommerce.DB, AuthEnable: ecommerce.AuthEnable, PermissionEnable: ecommerce.PermissionEnable})
	return memberconfig
}

// Customers list

func (ecommerce *Ecommerce) CustomerList(limit, offset int, filter Filter, tenantid int) (customer []TblEcomCustomer, count int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomCustomer{}, 0, AuthErr
	}

	Ecommercemodel.DataAccess = ecommerce.DataAccess

	Ecommercemodel.UserId = ecommerce.UserId

	customerlist, _, _ := Ecommercemodel.CustomersList(offset, limit, filter, ecommerce.DB, tenantid)

	_, totalcount, _ := Ecommercemodel.CustomersList(0, 0, filter, ecommerce.DB, tenantid)

	var finalcustomerlist []TblEcomCustomer

	for _, customer := range customerlist {

		var first = customer.FirstName

		var last = customer.LastName

		var firstn = strings.ToUpper(first[:1])

		var lastn string

		if customer.LastName != "" {

			lastn = strings.ToUpper(last[:1])
		}

		var Name = firstn + lastn

		customer.NameString = Name

		finalcustomerlist = append(finalcustomerlist, customer)

	}

	return finalcustomerlist, totalcount, nil

}

// Create Member

func (ecommerce *Ecommerce) CreateMember(Cc CreateCustomerReq) (ccmember member.Tblmember, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return member.Tblmember{}, AuthErr
	}

	db := ecommerce.DBconf()

	cmember, err := db.CreateMember(member.MemberCreationUpdation{

		FirstName: Cc.FirstName,

		LastName: Cc.LastName,

		Email: Cc.Email,

		MobileNo: Cc.MobileNo,

		Username: Cc.Username,

		ProfileImage: Cc.ProfileImage,

		ProfileImagePath: Cc.ProfileImagePath,

		IsActive: Cc.IsActive,

		Password: Cc.Password,

		CreatedBy: Cc.CreatedBy,

		TenantId: Cc.TenantId,
	})

	if err != nil {
		log.Println(err)
	}

	return cmember, nil

}

func (ecommerce *Ecommerce) UpdateMember(Cc CreateCustomerReq, memberid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	db := ecommerce.DBconf()

	err := db.UpdateMember(member.MemberCreationUpdation{

		FirstName: Cc.FirstName,

		LastName: Cc.LastName,

		Email: Cc.Email,

		MobileNo: Cc.MobileNo,

		Username: Cc.Username,

		ProfileImage: Cc.ProfileImage,

		ProfileImagePath: Cc.ProfileImagePath,

		IsActive: Cc.IsActive,

		Password: Cc.Password,

		ModifiedBy: Cc.CreatedBy,

		TenantId: Cc.TenantId,
	}, memberid)

	if err != nil {
		log.Println(err)
	}

	return nil

}

// Create Customer

func (ecommerce *Ecommerce) CreateCustomer(Cc CreateCustomerReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var ccustomer TblEcomCustomer

	ccustomer.MemberId = Cc.MemberId

	ccustomer.City = Cc.City

	ccustomer.Country = Cc.Country

	ccustomer.StreetAddress = Cc.StreetAddress

	ccustomer.ZipCode = Cc.ZipCode

	ccustomer.State = Cc.State

	ccustomer.CreatedBy = Cc.CreatedBy

	ccustomer.FirstName = Cc.FirstName

	ccustomer.LastName = Cc.LastName

	ccustomer.Username = Cc.Username

	ccustomer.Email = Cc.Email

	ccustomer.MobileNo = Cc.MobileNo

	if Cc.Password != "" {

		password := HashingPassword(Cc.Password)

		ccustomer.Password = password
	}

	ccustomer.ProfileImage = Cc.ProfileImage

	ccustomer.ProfileImagePath = Cc.ProfileImagePath

	ccustomer.IsActive = Cc.IsActive

	ccustomer.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	ccustomer.TenantId = Cc.TenantId

	err1 := Ecommercemodel.CustomerCreate(ccustomer, ecommerce.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

// Edit Customer

func (ecommerce *Ecommerce) EditCustomer(id int, tenantid int) (customers TblEcomCustomer, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCustomer{}, AuthErr
	}

	customer, err := Ecommercemodel.CustomerEdit(id, ecommerce.DB, tenantid)

	if err != nil {
		return TblEcomCustomer{}, err
	}
	var first = customer.FirstName

	var last = customer.LastName

	var firstn = strings.ToUpper(first[:1])

	var lastn string

	if customer.LastName != "" {

		lastn = strings.ToUpper(last[:1])
	}

	var Name = firstn + lastn

	customer.NameString = Name

	return customer, nil

}

// Update Customer

func (ecommerce *Ecommerce) UpdateCustomer(Cc CreateCustomerReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var updatecustomer TblEcomCustomer

	updatecustomer.MemberId = Cc.MemberId

	updatecustomer.City = Cc.City

	updatecustomer.Country = Cc.Country

	updatecustomer.StreetAddress = Cc.StreetAddress

	updatecustomer.ZipCode = Cc.ZipCode

	updatecustomer.State = Cc.State

	updatecustomer.CreatedBy = Cc.CreatedBy

	updatecustomer.FirstName = Cc.FirstName

	updatecustomer.LastName = Cc.LastName

	updatecustomer.Email = Cc.Email

	updatecustomer.MobileNo = Cc.MobileNo

	if Cc.Password != "" {
		password := HashingPassword(Cc.Password)
		updatecustomer.Password = password
	}

	updatecustomer.ProfileImage = Cc.ProfileImage

	updatecustomer.ProfileImagePath = Cc.ProfileImagePath

	updatecustomer.IsActive = Cc.IsActive

	updatecustomer.ModifiedBy = Cc.ModifiedBy

	updatecustomer.Username = Cc.Username

	updatecustomer.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CustomerUpdate(updatecustomer, ecommerce.DB)

	if err1 != nil {

		return err1
	}

	return nil

}

// Delete Particular customer id

func (ecommerce *Ecommerce) DeleteCustomer(id int, deletedby int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var customer TblEcomCustomer

	customer.DeletedBy = deletedby

	customer.MemberId = id

	customer.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CustomerDelete(customer, ecommerce.DB, tenantid)

	if err1 != nil {
		return err1
	}

	return nil

}

// multi customer delete
func (ecommerce *Ecommerce) MultiSelectCustomerDelete(id []int, deletedby int, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var customer TblEcomCustomer

	customer.IsDeleted = 1

	customer.DeletedBy = deletedby

	customer.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err1 := Ecommercemodel.MultiSelectDeleteCustomers(customer, id, ecommerce.DB, tenantid)

	if err1 != nil {
		return false, err1
	}

	return flg, nil
}

// multi customer status change
func (ecommerce *Ecommerce) MultiSelectCustomersStatus(customerid []int, status int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var customer TblEcomCustomer

	customer.IsActive = status

	customer.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.MultiSelectCustomerIsactive(customer, customerid, ecommerce.DB, tenantid)

	if err != nil {
		return err
	}

	return nil
}

// To Check email , username, mobileno ia already exists
func (ecommerce *Ecommerce) CheckDuplicateValue(memberid int, email string, username string, mobileno string, tenantid int) (bool, error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	db := ecommerce.DBconf()

	flag, err := db.CheckEmailInMember(memberid, email, tenantid)

	flag1, err1 := db.CheckNameInMember(memberid, email, tenantid)

	flag2, err2 := db.CheckNumberInMember(memberid, mobileno, tenantid)

	// flags := map[string]bool{
	// 	"emailflag":  flag,
	// 	"nameflag":   flag1,
	// 	"numberflag": flag2,
	// }

	if err != nil {

		return flag, err
	}
	if err1 != nil {

		return flag1, err
	}
	if err2 != nil {

		return flag2, err
	}

	return true, nil
}

// To Get Customer order info details
func (ecommerce *Ecommerce) CustomerOrderInfo(uuid string, tenantid int) (productorder []TblEcomProduct, order TblEcomProductOrder, address ShippingAddress, statusdetails []TblEcomOrderStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProduct{}, TblEcomProductOrder{}, ShippingAddress{}, []TblEcomOrderStatus{}, AuthErr
	}
	cusinfo, err := Ecommercemodel.GetOrderDetailsbyuuid(uuid, ecommerce.DB, tenantid)

	log.Println("cusinfo", cusinfo)

	if err != nil {
		return []TblEcomProduct{}, TblEcomProductOrder{}, ShippingAddress{}, []TblEcomOrderStatus{}, err
	}

	var first = cusinfo.FirstName

	var last = cusinfo.LastName

	var firstn = strings.ToUpper(first[:1])

	var lastn string

	if cusinfo.LastName != "" {

		lastn = strings.ToUpper(last[:1])
	}

	var Name = firstn + lastn

	cusinfo.NameString = Name

	cusinfo.CreatedDate = cusinfo.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

	orderid := cusinfo.Id
	// To get order stauts is particular id
	statusdetails, err6 := Ecommercemodel.OrderStatusDetails(uuid, ecommerce.DB, tenantid)

	if err6 != nil {
		log.Println(err6)
	}

	productinfo, err1 := Ecommercemodel.GetProductdetailsByOrderId(orderid, ecommerce.DB, tenantid)

	if err1 != nil {

		return []TblEcomProduct{}, TblEcomProductOrder{}, ShippingAddress{}, []TblEcomOrderStatus{}, err1
	}

	var product_id []int

	for _, val := range productinfo {

		product_id = append(product_id, val.Product_id)

	}

	productdetails, err2 := Ecommercemodel.GetProductdetailsByProductId(product_id, ecommerce.DB, tenantid)

	if err2 != nil {
		return []TblEcomProduct{}, TblEcomProductOrder{}, ShippingAddress{}, []TblEcomOrderStatus{}, err2

	}
	var productList []TblEcomProduct

	for i, val := range productdetails {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

		}
		if i < len(productinfo) {
			quantity := productinfo[i].Quantity
			price := productinfo[i].Price
			quantityPrice := quantity * price
			productList = append(productList, TblEcomProduct{

				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           productinfo[i].Quantity,
				Price:              productinfo[i].Price,
				Quantityprice:      quantityPrice,
			})
		} else {

			productList = append(productList, TblEcomProduct{
				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           0,
				Price:              0,
				Quantityprice:      0,
			})
		}
	}

	var shippingAddress ShippingAddress

	err4 := json.Unmarshal([]byte(cusinfo.ShippingAddress), &shippingAddress)
	if err4 != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	return productList, cusinfo, shippingAddress, statusdetails, nil

}

// To get customer details
func (ecommerce *Ecommerce) CustomerInfo(limit, offset, customerid int, tenantid int) (customers TblEcomCustomer, products []TblEcomProductOrder, totalcount int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCustomer{}, []TblEcomProductOrder{}, 0, AuthErr
	}

	customer, err := Ecommercemodel.GetCustomerDetails(customerid, ecommerce.DB, tenantid)

	if err != nil {
		return TblEcomCustomer{}, []TblEcomProductOrder{}, 0, err
	}

	var first = customer.FirstName

	var last = customer.LastName

	var firstn = strings.ToUpper(first[:1])

	var lastn string

	if customer.LastName != "" {

		lastn = strings.ToUpper(last[:1])
	}

	var Name = firstn + lastn

	customer.NameString = Name

	orders, _, err1 := Ecommercemodel.GetOrderDetailsbyCustomerId(limit, offset, customerid, ecommerce.DB, tenantid)

	_, count, err2 := Ecommercemodel.GetOrderDetailsbyCustomerId(0, 0, customerid, ecommerce.DB, tenantid)

	if err1 != nil {
		return TblEcomCustomer{}, []TblEcomProductOrder{}, 0, err1
	}

	if err2 != nil {
		return TblEcomCustomer{}, []TblEcomProductOrder{}, 0, err2
	}

	return customer, orders, count, nil
}

// Password Hasing
func (ecommerce *Ecommerce) HashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}

// Get Customer details
func (Ecommerce *Ecommerce) GetCustomer(memberId int, tenantid int) (customer TblEcomCustomer, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return TblEcomCustomer{}, AuthErr
	}

	customer, err = Ecommercemodel.GetCustomer(memberId, Ecommerce.DB, tenantid)
	if err != nil {
		return TblEcomCustomer{}, err
	}

	return customer, nil
}

// Get Customer Details by Id
func (Ecommerce *Ecommerce) GetCustomerDetailsById(memberId int, tenantid int) (customer CustomerDetails, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return CustomerDetails{}, AuthErr
	}

	customer, err = EcommerceModel.GetCustomerDetailsById(EcommerceModel{}, memberId, Ecommerce.DB, tenantid)
	if err != nil {

		return CustomerDetails{}, err
	}

	return customer, err
}

// Update member and customer details
func (Ecommerce *Ecommerce) UpdateCustomerAndMemberDetails(memberId int, memberDetails map[string]interface{}, customerDetails map[string]interface{}, tenantid int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.UpdateMemberDetails(EcommerceModel{}, memberId, memberDetails, Ecommerce.DB, tenantid)
	if err != nil {

		return err
	}

	err = EcommerceModel.UpdateCustomerDetails(EcommerceModel{}, memberId, customerDetails, Ecommerce.DB, tenantid)
	if err != nil {

		return err
	}

	return nil
}
