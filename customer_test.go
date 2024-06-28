package ecommerce

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "picco123@",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurt_cms_apr_3",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to database:", err)

	}
	if err != nil {

		return nil, err

	}

	return db, nil
}

// test listCustomerList function
func TestCustomerList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		customer, count, err := ecommerce.CustomerList(10, 0, Filter{})

		if err != nil {

			log.Println(err)
		}

		fmt.Println(customer, count)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Create Customer Function

func TestCreateCustomer(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		err := ecommerce.CreateCustomer(CreateCustomerReq{FirstName: "tester", Username: "Tester", Email: "tester@gmail.com", MobileNo: "9080706050", Password: "Tester@123", StreetAddress: "21/2 Mariyammnkalai street", City: "Thiruvannamali", ZipCode: "606701", Country: "India", State: "Tamil Nadu", CreatedBy: 1})

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Edit Customer Function

func TestEditCustomer(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	id := 5

	if permisison {

		customer, err := ecommerce.EditCustomer(id)

		if err != nil {

			panic(err)
		}

		log.Println("customer", customer)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Update Customer Function

func TestUpdateCustomer(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		err := ecommerce.UpdateCustomer(CreateCustomerReq{FirstName: "Tester", LastName: "New", Username: "Tester", Email: "tester@gmail.com", MobileNo: "9080706050", Password: "Tester@123"})

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Delete Customer Function
func TestDeleteCustomer(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	id := 6
	deletedby := 1

	if permisison {

		err := ecommerce.DeleteCustomer(id, deletedby)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Multi select delete Customer Function
func TestMultiSelectCustomerDelete(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})
	var id = []int{1, 2, 3}

	deletedby := 1

	if permisison {

		flg, err := ecommerce.MultiSelectCustomerDelete(id, deletedby)

		if err != nil {

			panic(err)
		}

		fmt.Println("flg", flg)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test Multi select  status change  Function
func TestMultiSelectCustomersStatus(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	var id = []int{1, 2, 3}

	status := 1

	if permisison {

		err := ecommerce.MultiSelectCustomersStatus(id, status)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test customer name , email , mobile number is Already exists or not  Function
func TestCheckDuplicateValue(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})
	id := 5
	name := "Tester"
	email := "tester@gmail.com"
	mobileno := "8824556600"

	if permisison {

		flg, err := ecommerce.CheckDuplicateValue(id, name, email, mobileno)

		if err != nil {

			panic(err)
		}
		log.Println("flg", flg)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test To Get Customer order details  Function
func TestCustomerOrderInfo(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})
	id := "SP03052024"

	if permisison {

		product, cusinfo, address, status, err := ecommerce.CustomerOrderInfo(id)

		if err != nil {

			panic(err)
		}

		log.Println("order,product,status", product, status, cusinfo, address)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// Test to get customer details  Function
func TestCustomerInfo(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Customer", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		customer, product, ordercount, err := ecommerce.CustomerInfo(10, 0, 1)

		if err != nil {

			panic(err)
		}

		log.Println("customer & product & ordercount", customer, product, ordercount)

	} else {

		log.Println("permissions enabled not initialised")

	}

}
