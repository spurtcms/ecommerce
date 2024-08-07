package ecommerce

import (
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

func (ecommerce *Ecommerce) PaymentList(offset, limit int, tenantid int) (paymentlists []TblEcomPayment, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomPayment{}, AuthErr
	}

	Ecommercemodel.DataAccess = ecommerce.DataAccess

	Ecommercemodel.UserId = ecommerce.UserId

	paymentlist, err := Ecommercemodel.PaymentLists(offset, limit, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return paymentlist, nil
}

// Status List
func (ecommerce *Ecommerce) StatusList(offset, limit int, tenantid int) (statuslists []TblEcomStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomStatus{}, AuthErr
	}

	Ecommercemodel.DataAccess = ecommerce.DataAccess

	Ecommercemodel.UserId = ecommerce.UserId

	statuslist, err := Ecommercemodel.StatusLists(offset, limit, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return statuslist, nil
}

// Status List
func (ecommerce *Ecommerce) CurrencyList(offset, limit int, tenantid int) (currencylists []TblEcomCurrency, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomCurrency{}, AuthErr
	}

	Ecommercemodel.DataAccess = ecommerce.DataAccess

	Ecommercemodel.UserId = ecommerce.UserId

	currencylist, err := Ecommercemodel.CurrencyLists(offset, limit, ecommerce.DB, tenantid)

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

	setting.TenantId = Ss.TenantId

	err := Ecommercemodel.CreateSetting(setting, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Create Setting
func (ecommerce *Ecommerce) UpdateSettings(Ss CreateSettingReq, tenantid int) error {

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

	err := Ecommercemodel.UpdateSetting(setting, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Create Currency
func (ecommerce *Ecommerce) CreateCurrency(Cc CreateCurrencyReq, tenantid int) error {

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

	currency.TenantId = Cc.TenantId

	currency.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if currency.CurrencyDefault != 0 {

		money, err := Ecommercemodel.FindDefault(ecommerce.DB)

		if err != nil {

			log.Println(err)
		}

		if money.CurrencyName == "" {

			err := Ecommercemodel.CurrencyCreate(currency, ecommerce.DB)
			if err != nil {
				log.Println(err)
			}
		} else if money.CurrencyName != "" {

			err1 := Ecommercemodel.ChangeDefaultValue(money.Id, ecommerce.DB, tenantid)

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
func (ecommerce *Ecommerce) UpdateCurrency(Cc CreateCurrencyReq, tenantid int) error {

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

			err1 := Ecommercemodel.ChangeDefaultValue(money.Id, ecommerce.DB, tenantid)

			if err1 != nil {

				log.Println(err1)
			}

			err := Ecommercemodel.UpdateCurrency(currency, ecommerce.DB, tenantid)

			if err != nil {
				log.Println(err)
			}
		}

	} else {

		err := Ecommercemodel.UpdateCurrency(currency, ecommerce.DB, tenantid)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

//Delete Currency

func (ecommerce *Ecommerce) DeleteCurrency(id int, userid int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}
	var money TblEcomCurrency

	money.Id = id

	money.IsDeleted = 1

	money.DeletedBy = userid

	money.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.CurrencyDelete(money, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Currency IsActive

func (ecommerce *Ecommerce) CurrencyIsActive(Cc CreateCurrencyReq, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var money TblEcomCurrency

	money.Id = Cc.Id

	money.IsActive = Cc.IsActive

	money.ModifiedBy = Cc.ModifiedBy

	money.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.InActiveCurrency(money, ecommerce.DB, tenantid)

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

	status.TenantId = Cs.TenantId

	err := Ecommercemodel.CreateStatus(status, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Update Status
func (ecommerce *Ecommerce) StatusUpdate(Cs CreateStatusReq, tenantid int) error {

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

	err := Ecommercemodel.UpdateStatus(status, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Delete Status
func (ecommerce *Ecommerce) StatusDelete(id int, userid int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var status TblEcomStatus

	status.Id = id

	status.DeletedBy = userid

	status.IsDeleted = 1

	status.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.DeleteStatus(status, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Status IsActive
func (ecommerce *Ecommerce) StatusIsActive(Cs CreateStatusReq, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var status TblEcomStatus

	status.Id = Cs.Id

	status.IsActive = Cs.IsActive

	status.ModifiedBy = Cs.ModifiedBy

	status.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.OrderStatusIsActive(status, ecommerce.DB, tenantid)

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

	pay.TenantId = Cp.TenantId

	err := Ecommercemodel.PaymentCreate(pay, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Update Payment
func (ecommerce *Ecommerce) UpdatePayment(Cp CreatePaymentReq, tenantid int) error {

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

	err := Ecommercemodel.UpdatePayment(pay, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Delete Payment
func (ecommerce *Ecommerce) DeletePayment(id int, userid int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var pay TblEcomPayment

	pay.Id = id

	pay.IsDeleted = 1

	pay.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	pay.DeletedBy = userid

	err := Ecommercemodel.DeletePayment(pay, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return nil
}

// Payment isactive
func (ecommerce *Ecommerce) PaymentIsActive(Cp CreatePaymentReq, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var pay TblEcomPayment

	pay.Id = Cp.Id

	pay.IsActive = Cp.IsActive

	pay.ModifiedBy = Cp.ModifiedBy

	pay.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	flg, err := Ecommercemodel.PaymentIsActive(pay, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Edit Currency

func (ecommerce *Ecommerce) EditCurrency(id int, tenantid int) (currencys TblEcomCurrency, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomCurrency{}, AuthErr
	}

	currency, err := Ecommercemodel.CurrencyGet(id, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return currency, nil
}

// Edit Payment
func (ecommerce *Ecommerce) EditPayment(id int, tenantid int) (payments TblEcomPayment, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomPayment{}, AuthErr
	}
	payment, err := Ecommercemodel.PaymentGet(id, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return payment, nil
}

// Edit Status
func (ecommerce *Ecommerce) EditStatus(id int, tenantid int) (statuss TblEcomStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomStatus{}, AuthErr
	}
	status, err := Ecommercemodel.StatusGet(id, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	return status, nil
}

// Check Currency Name

func (ecommerce *Ecommerce) CheckCurrencyName(id int, name string, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var currency TblEcomCurrency

	flg, err := Ecommercemodel.CheckCurrencyName(id, name, currency, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// CheckCurrencyType
func (ecommerce *Ecommerce) CheckCurrencyType(id int, currencytype string, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var currency TblEcomCurrency

	flg, err := Ecommercemodel.CheckCurrencyType(id, currencytype, currency, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
		return flg, err
	}

	return flg, nil
}

// Check Currecy symbol

func (ecommerce *Ecommerce) CheckCurrencySymbol(id int, csymbol string, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var currency TblEcomCurrency

	flg, err := Ecommercemodel.CheckCurrencySymbol(id, csymbol, currency, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}

// Check Status Name

func (ecommerce *Ecommerce) CheckStatusName(id int, name string, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var status TblEcomStatus

	flg, err := Ecommercemodel.CheckStatusName(id, name, status, ecommerce.DB, tenantid)

	if err != nil {

		return false, err
	}

	return flg, nil
}

func (ecommerce *Ecommerce) CheckPaymentName(id int, name string, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var payment TblEcomPayment

	flg, err := Ecommercemodel.CheckPaymentName(id, name, payment, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)

		return false, err
	}

	return flg, nil
}

func (ecommerce *Ecommerce) CheckStatusPriority(id int, priority int, tenantid int) (flgs bool, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}
	var status TblEcomStatus

	flg, err := Ecommercemodel.CheckStatusPriority(id, priority, status, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return flg, nil
}
