package ecommerce

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

// To Test Product list
func TestProductsList(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		product, count, err := ecommerce.ProductsList(10, 0, Filter{}, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(product, count)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test Create Product
func TestCreateProduct(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		// layout := "2006-01-02T15:04"
		// time.Parse(layout, offer.Startdate)

		create, err := ecommerce.CreateProduct(CreateProductReq{CategoriesId: "1", ProductDescription: "nice product", ProductName: "Mobile", Sku: "welcome", ProductPrice: 12000, Tax: 1000, Totalcost: 13000, CreatedBy: 1, Type: "discount", Price: 2000, Priority: 1})

		if err != nil {

			panic(err)
		}

		fmt.Println(create)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test Edit Product
func TestEditProduct(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		product, price, discount, err := ecommerce.EditProduct(1, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(product, price, discount)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test update Product Function
func TestUpdateProduct(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	var offerid = []int{1, 2}

	if permisison {

		err := ecommerce.UpdateProduct(CreateProductReq{CategoriesId: "1", ProductDescription: "Bad product", ProductName: "Mobile", Sku: "welcome", ProductPrice: 12000, Tax: 1000, Totalcost: 13000, CreatedBy: 1, Type: "discount", Price: 2000, Priority: 1, ModifiedBy: 1}, offerid, 1)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test Edit Product
func TestDeleteProduct(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	var productid = 1

	var id = 1

	if permisison {

		err := ecommerce.DeleteProduct(productid, id, TenantId)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test Edit Product
func TestCheckSkuName(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		flg, err := ecommerce.CheckSkuName("welcome", 10, TenantId)

		if err != nil {

			panic(err)
		}

		fmt.Println(flg)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test Edit Product
func TestSelectProductsChangeStatus(t *testing.T) {

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

	permisison, _ := Auth.IsGranted("Catalogue", auth.CRUD, TenantId)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	var productid = []int{4, 5}

	if permisison {

		err := ecommerce.SelectProductsChangeStatus(1, productid, TenantId)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}
