package ecommerce

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spurtcms/member"
	"golang.org/x/crypto/bcrypt"
)

type tblecomcustomers struct {
	Id               int `gorm:"primaryKey;auto_increment;type:serial"`
	MemberId         int
	FirstName        string
	LastName         string
	Email            string
	MobileNo         string
	Username         string
	Password         string
	StreetAddress    string
	City             string
	State            string
	Country          string
	ZipCode          string
	IsActive         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer"`
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"DEFAULT:NULL"`
	Count            int       `gorm:"-:migration;<-:false"`
	ProfileImage     string
	ProfileImagePath string
	NameString       string `gorm:"-:migration;<-:false"`
}

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

// EcommerceSetup used initialize Ecommerce configruation
func EcommerceSetup(config Config) *Ecommerce {

	// MigrateTables(config.DB)

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

func (ecommerce *Ecommerce) CustomerList(limit, offset int, filter Filter) (customer []TblEcomCustomers, count int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomCustomers{}, 0, AuthErr
	}

	customerlist, _, _ := Ecommercemodel.CustomersList(limit, offset, filter, ecommerce.DB)

	_, totalcount, _ := Ecommercemodel.CustomersList(0, 0, filter, ecommerce.DB)

	var finalcustomerlist []TblEcomCustomers

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
	})

	if err != nil {
		log.Println(err)
	}

	return cmember, nil

}

// Create Customer

func (ecommerce *Ecommerce) CreateCustomer(Cc CreateCustomerReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var ccustomer TblEcomCustomers

	ccustomer.MemberId = Cc.MemberId

	ccustomer.City = Cc.City

	ccustomer.Country = Cc.Country

	ccustomer.StreetAddress = Cc.StreetAddress

	ccustomer.ZipCode = Cc.ZipCode

	ccustomer.State = Cc.State

	ccustomer.CreatedBy = Cc.CreatedBy

	ccustomer.FirstName = Cc.FirstName

	ccustomer.LastName = Cc.LastName

	ccustomer.Email = Cc.Email

	ccustomer.MobileNo = Cc.MobileNo

	ccustomer.Password = Cc.Password

	ccustomer.ProfileImage = Cc.ProfileImage

	ccustomer.ProfileImagePath = Cc.ProfileImagePath

	ccustomer.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CustomerCreate(ccustomer, ecommerce.DB)

	if err1 != nil {

		return err1
	}

	return nil
}

// Edit Customer

func (ecommerce *Ecommerce) EditCustomer(id int) (customers TblEcomCustomers, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCustomers{}, AuthErr
	}

	customer, err := Ecommercemodel.CustomerEdit(id, ecommerce.DB)

	if err != nil {
		return TblEcomCustomers{}, err
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

func (ecommerce *Ecommerce) UpdateCustomer(Cc CreateCustomerReq, memberid int) error {

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
	}, memberid)
	// cmember.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err != nil {
		log.Println(err)
	}

	var upadtecustomer TblEcomCustomers

	upadtecustomer.MemberId = memberid

	upadtecustomer.City = Cc.City

	upadtecustomer.Country = Cc.Country

	upadtecustomer.StreetAddress = Cc.StreetAddress

	upadtecustomer.ZipCode = Cc.ZipCode

	upadtecustomer.State = Cc.State

	upadtecustomer.CreatedBy = Cc.CreatedBy

	upadtecustomer.FirstName = Cc.FirstName

	upadtecustomer.LastName = Cc.LastName

	upadtecustomer.Email = Cc.Email

	upadtecustomer.MobileNo = Cc.MobileNo

	upadtecustomer.Password = Cc.Password

	upadtecustomer.ProfileImage = Cc.ProfileImage

	upadtecustomer.ProfileImagePath = Cc.ProfileImagePath

	upadtecustomer.IsActive = Cc.IsActive

	upadtecustomer.ModifiedBy = Cc.ModifiedBy

	upadtecustomer.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CustomerUpdate(upadtecustomer, ecommerce.DB)

	if err1 != nil {

		return err1
	}

	return nil

}

// Delete Particular customer id

func (ecommerce *Ecommerce) DeleteCustomer(id int, deletedby int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	db := ecommerce.DBconf()

	err := db.DeleteMember(id, deletedby)

	if err != nil {

		return err
	}

	var customer TblEcomCustomers

	customer.DeletedBy = deletedby

	customer.MemberId = id

	customer.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CustomerDelete(customer, ecommerce.DB)

	if err1 != nil {
		return err1
	}

	return nil

}

// multi customer delete
func (ecommerce *Ecommerce) MultiSelectCustomerDelete(id []int, deletedby int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	db := ecommerce.DBconf()

	flg, err := db.MultiSelectedMemberDelete(id, deletedby)

	if err != nil {
		return err
	}

	if flg {

		var customer TblEcomCustomers

		customer.IsDeleted = 1

		customer.DeletedBy = deletedby

		customer.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		err1 := Ecommercemodel.MultiSelectDeleteCustomers(customer, id, ecommerce.DB)

		if err1 != nil {
			return err1
		}

	}

	return nil
}

// multi customer status change
func (ecommerce *Ecommerce) MultiSelectCustomersStatus(customerid []int, status int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var customer TblEcomCustomers

	customer.IsActive = status

	customer.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.MultiSelectCustomerIsactive(customer, customerid, ecommerce.DB)

	if err != nil {
		return err
	}

	return nil
}

// To Check email , username, mobileno ia already exists
func (ecommerce *Ecommerce) CheckDuplicateValue(memberid int, email string, username string, mobileno string) (bool, error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	db := ecommerce.DBconf()

	flag, err := db.CheckEmailInMember(memberid, email)

	flag1, err1 := db.CheckNameInMember(memberid, email)

	flag2, err2 := db.CheckNumberInMember(memberid, mobileno)

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
func (ecommerce *Ecommerce) CustomerOrderInfo(uuid string) (productorder []TblEcomProducts, order TblEcomProductOrders, address ShippingAddress, laststatus string, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProducts{}, TblEcomProductOrders{}, ShippingAddress{}, "", AuthErr
	}
	cusinfo, err := Ecommercemodel.GetOrderDetailsbyuuid(uuid, ecommerce.DB)

	log.Println("cusinfo", cusinfo)

	if err != nil {
		return []TblEcomProducts{}, TblEcomProductOrders{}, ShippingAddress{}, "", err
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

	var lstatus string

	var Orderstatus []TblEcomOrderStatuses

	for _, val := range cusinfo.Orders {

		val.CreatedDate = val.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		Orderstatus = append(Orderstatus, val)
		// lstatus = val.OrderStatus

	}

	productinfo, err1 := Ecommercemodel.GetProductdetailsByOrderId(orderid, ecommerce.DB)

	if err1 != nil {

		return []TblEcomProducts{}, TblEcomProductOrders{}, ShippingAddress{}, "", err1
	}

	var product_id []int

	for _, val := range productinfo {

		product_id = append(product_id, val.Product_id)

	}

	productdetails, err2 := Ecommercemodel.GetProductdetailsByProductId(product_id, ecommerce.DB)

	if err2 != nil {
		return []TblEcomProducts{}, TblEcomProductOrders{}, ShippingAddress{}, "", err2

	}
	var productList []TblEcomProducts

	for i, val := range productdetails {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

		}
		if i < len(productinfo) {
			quantity := productinfo[i].Quantity
			price := productinfo[i].Price
			quantityPrice := quantity * price
			productList = append(productList, TblEcomProducts{

				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           productinfo[i].Quantity,
				Price:              productinfo[i].Price,
				Quantityprice:      quantityPrice,
			})
		} else {

			productList = append(productList, TblEcomProducts{
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

	return productList, cusinfo, shippingAddress, lstatus, nil

}

// To get customer details
func (ecommerce *Ecommerce) CustomerInfo(limit, offset, customerid int) (customers TblEcomCustomers, products []TblEcomProductOrders, totalcount int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCustomers{}, []TblEcomProductOrders{}, 0, AuthErr
	}

	customer, err := Ecommercemodel.GetCustomerDetails(customerid, ecommerce.DB)

	if err != nil {
		return TblEcomCustomers{}, []TblEcomProductOrders{}, 0, err
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

	orders, _, err1 := Ecommercemodel.GetOrderDetailsbyCustomerId(limit, offset, customerid, ecommerce.DB)

	_, count, err2 := Ecommercemodel.GetOrderDetailsbyCustomerId(0, 0, customerid, ecommerce.DB)

	if err1 != nil {
		return TblEcomCustomers{}, []TblEcomProductOrders{}, 0, err1
	}

	if err2 != nil {
		return TblEcomCustomers{}, []TblEcomProductOrders{}, 0, err2
	}

	var finalorderlist []TblEcomProductOrders

	for _, orders := range orders {

		orders.CreatedDate = orders.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		if !orders.ModifiedOn.IsZero() {

			orders.ModifiedDate = orders.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		} else {

			orders.ModifiedDate = orders.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		}

		finalorderlist = append(finalorderlist, orders)
	}

	return customer, finalorderlist, count, nil
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
func (Ecommerce *Ecommerce) GetCustomer(memberId int) (customer TblEcomCustomers, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return TblEcomCustomers{}, AuthErr
	}

	customer, err = Ecommercemodel.GetCustomer(memberId, Ecommerce.DB)
	if err != nil {
		return TblEcomCustomers{}, err
	}

	return customer, nil
}
