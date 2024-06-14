package ecommerce

import (
	"fmt"
	"log"
	"time"
)

// Store list
func (ecommerce *Ecommerce) StoreList() (storlis TblEcomSettings, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomSettings{}, AuthErr
	}
	storelist, err := Ecommercemodel.StoreLists(ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return storelist, nil
}

// Payment list

func (ecommerce *Ecommerce) PaymentList(offset, limit int) (paymentlists []TblEcomPayment, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomPayment{}, AuthErr
	}
	paymentlist, err := Ecommercemodel.PaymentLists(offset, limit, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return paymentlist, nil
}

// Status List
func (ecommerce *Ecommerce) StatusList(offset, limit int) (statuslists []TblEcomStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomStatus{}, AuthErr
	}
	statuslist, err := Ecommercemodel.StatusLists(offset, limit, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return statuslist, nil
}

// Status List
func (ecommerce *Ecommerce) CurrencyList(offset, limit int) (currencylists []TblEcomCurrency, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomCurrency{}, AuthErr
	}
	currencylist, err := Ecommercemodel.CurrencyLists(offset, limit, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return currencylist, nil
}

// Create Setting
func (ecommerce *Ecommerce) CreateSettings(Ss CreateSettingReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var setting TblEcomSettings

	setting.CreatedBy = Ss.CreatedBy

	setting.CurrencyDefault = Ss.CurrencyDefault

	setting.PaymentDefault = Ss.PaymentDefault

	setting.StatusDefault = Ss.StatusDefault

	setting.StockWarning = Ss.StockWarning

	setting.DisplayStock = Ss.DisplayStock

	setting.StockCheckout = Ss.StockCheckout

	setting.StoreName = Ss.StoreName

	setting.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.CreateSetting(setting, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Create Setting
func (ecommerce *Ecommerce) UpdateSettings(Ss CreateSettingReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var setting TblEcomSettings

	setting.Id = Ss.Id

	setting.ModifiedBy = Ss.ModifiedBy

	setting.CurrencyDefault = Ss.CurrencyDefault

	setting.PaymentDefault = Ss.PaymentDefault

	setting.StatusDefault = Ss.StatusDefault

	setting.StockWarning = Ss.StockWarning

	setting.DisplayStock = Ss.DisplayStock

	setting.StockCheckout = Ss.StockCheckout

	setting.StoreName = Ss.StoreName

	setting.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdateSetting(setting, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Create Currency
func (ecommerce *Ecommerce) CreateCurrency(Cc CreateCurrencyReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var currency TblEcomCurrency

	currency.CreatedBy = Cc.CreatedBy

	currency.CurrencyName = Cc.CurrencyName

	currency.CurrencyDefault = Cc.CurrencyDefault

	currency.CurrencySymbol = Cc.CurrencySymbol

	currency.CurrencyType = Cc.CurrencyType

	currency.IsActive = Cc.IsActive

	currency.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if currency.CurrencyDefault != 0 {

		money, err := Ecommercemodel.FindDefault(ecommerce.DB)
		if err != nil {

			log.Println(err)
		}

		if money.CurrencyName != "" {

			err1 := Ecommercemodel.ChangeDefaultValue(currency, ecommerce.DB)

			if err1 != nil {

				log.Println(err1)
			}

			err := Ecommercemodel.CurrencyCreate(currency, ecommerce.DB)

			if err != nil {
				log.Println(err)
			}
		}

	} else {

		err := Ecommercemodel.CurrencyCreate(currency, ecommerce.DB)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

// Update Currency
func (ecommerce *Ecommerce) UpdateCurrency(Cc CreateCurrencyReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var currency TblEcomCurrency

	currency.Id = Cc.Id

	currency.ModifiedBy = Cc.CreatedBy

	currency.CurrencyName = Cc.CurrencyName

	currency.CurrencyDefault = Cc.CurrencyDefault

	currency.CurrencySymbol = Cc.CurrencySymbol

	currency.CurrencyType = Cc.CurrencyType

	currency.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if currency.CurrencyDefault != 0 {

		money, err := Ecommercemodel.FindDefault(ecommerce.DB)
		if err != nil {

			log.Println(err)
		}

		if money.CurrencyName != "" {

			err1 := Ecommercemodel.ChangeDefaultValue(currency, ecommerce.DB)

			if err1 != nil {

				log.Println(err1)
			}

			err := Ecommercemodel.UpdateCurrency(currency, ecommerce.DB)

			if err != nil {
				log.Println(err)
			}
		}

	} else {

		err := Ecommercemodel.UpdateCurrency(currency, ecommerce.DB)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

//Delete Currency

func (ecommerce *Ecommerce) DeleteCurrency(id int, userid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}
	var money TblEcomCurrency

	money.Id = id

	money.IsDeleted = 1

	money.DeletedBy = userid

	money.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.CurrencyDelete(money, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Currency IsActive

func (ecommerce *Ecommerce) CurrencyIsActive(Cc CreateCurrencyReq) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var money TblEcomCurrency

	money.Id = Cc.Id

	money.IsActive = Cc.IsActive

	money.ModifiedBy = Cc.ModifiedBy

	money.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.InActiveCurrency(money, ecommerce.DB)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Create Status
func (ecommerce *Ecommerce) StatusCreate(Cs CreateStatusReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}
	var status TblEcomStatus

	status.Status = Cs.Status

	status.ColorCode = Cs.ColorCode

	status.Description = Cs.Description

	status.CreatedBy = Cs.CreatedBy

	status.IsActive = Cs.IsActive

	status.Priority = Cs.Priority

	status.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.CreateStatus(status, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Update Status
func (ecommerce *Ecommerce) StatusUpdate(Cs CreateStatusReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}
	var status TblEcomStatus

	status.Id = Cs.Id

	status.Status = Cs.Status

	status.ColorCode = Cs.ColorCode

	status.Description = Cs.Description

	status.ModifiedBy = Cs.ModifiedBy

	status.Priority = Cs.Priority

	status.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdateStatus(status, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Delete Status
func (ecommerce *Ecommerce) StatusDelete(id int, userid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var status TblEcomStatus

	status.Id = id

	status.DeletedBy = userid

	status.IsDeleted = 1

	status.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.DeleteStatus(status, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Status IsActive
func (ecommerce *Ecommerce) StatusIsActive(Cs CreateStatusReq) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var status TblEcomStatus

	status.Id = Cs.Id

	status.IsActive = Cs.IsActive

	status.ModifiedBy = Cs.ModifiedBy

	status.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.OrderStatusIsActive(status, ecommerce.DB)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Create Payment
func (ecommerce *Ecommerce) CreatePayment(Cp CreatePaymentReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var pay TblEcomPayment

	pay.CreatedBy = Cp.CreatedBy

	pay.PaymentName = Cp.PaymentName

	pay.Description = Cp.Description

	pay.IsActive = Cp.IsActive

	pay.PaymentImage = Cp.PaymentImage

	pay.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.PaymentCreate(pay, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Update Payment
func (ecommerce *Ecommerce) UpdatePayment(Cp CreatePaymentReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var pay TblEcomPayment

	pay.Id = Cp.Id

	pay.ModifiedBy = Cp.ModifiedBy

	pay.PaymentName = Cp.PaymentName

	pay.Description = Cp.Description

	pay.PaymentImage = Cp.PaymentImage

	pay.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdatePayment(pay, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Delete Payment
func (ecommerce *Ecommerce) DeletePayment(id int, userid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var pay TblEcomPayment

	pay.Id = id

	pay.IsDeleted = 1

	pay.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	pay.DeletedBy = userid

	err := Ecommercemodel.DeletePayment(pay, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Payment isactive
func (ecommerce *Ecommerce) PaymentIsActive(Cp CreatePaymentReq) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var pay TblEcomPayment

	pay.Id = Cp.Id

	pay.IsActive = Cp.IsActive

	pay.ModifiedBy = Cp.ModifiedBy

	pay.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.PaymentIsActive(pay, ecommerce.DB)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Edit Currency

func (ecommerce *Ecommerce) EditCurrency(id int) (currencys TblEcomCurrency, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCurrency{}, AuthErr
	}

	currency, err := Ecommercemodel.CurrencyGet(id, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return currency, nil
}

// Edit Payment
func (ecommerce *Ecommerce) EditPayment(id int) (payments TblEcomPayment, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomPayment{}, AuthErr
	}
	payment, err := Ecommercemodel.PaymentGet(id, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return payment, nil
}

// Edit Status
func (ecommerce *Ecommerce) EditStatus(id int) (statuss TblEcomStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomStatus{}, AuthErr
	}
	status, err := Ecommercemodel.StatusGet(id, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return status, nil
}

// Check Currency Name

func (ecommerce *Ecommerce) CheckCurrencyName(id int, name string) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var currency TblEcomCurrency

	currency.Id = id

	currency.CurrencyName = name

	flg, err := Ecommercemodel.CheckCurrencyName(currency, ecommerce.DB)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// CheckCurrencyType
func (ecommerce *Ecommerce) CheckCurrencyType(id int, currencytype string) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var currency TblEcomCurrency

	currency.Id = id

	currency.CurrencyType = currencytype

	flg, err := Ecommercemodel.CheckCurrencyType(currency, ecommerce.DB)
	if err != nil {
		log.Println(err)
		return flg, err
	}

	return flg, nil
}

// Check Currecy symbol

func (ecommerce *Ecommerce) CheckCurrencySymbol(id int, csymbol string) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var currency TblEcomCurrency

	currency.Id = id

	currency.CurrencySymbol = csymbol

	flg, err := Ecommercemodel.CheckCurrencySymbol(currency, ecommerce.DB)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Check Status Name

func (ecommerce *Ecommerce) CheckStatusName(id int, name string) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var status TblEcomStatus

	status.Id = id

	status.Status = name

	flg, err := Ecommercemodel.CheckStatusName(status, ecommerce.DB)

	if err != nil {

		fmt.Println(err)

		return false, err
	}

	return flg, nil
}

func (ecommerce *Ecommerce) CheckPaymentName(id int, name string) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var payment TblEcomPayment

	payment.Id = id

	payment.PaymentName = name

	flg, err := Ecommercemodel.CheckPaymentName(payment, ecommerce.DB)

	if err != nil {
		log.Println(err)

		return false, err
	}

	return flg, nil
}

func (ecommerce *Ecommerce) CheckStatusPriority(id int, priority int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var status TblEcomStatus

	status.Id = id

	status.Priority = priority

	flg, err := Ecommercemodel.CheckStatusPriority(status, ecommerce.DB)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}
