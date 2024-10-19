package crm

import "gorm.io/gorm"

func UniversalSearchCustomerName(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_customers.customer_name = ?", value)
	}
}

func UniversalSearchOrderNumber(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_order.order_number = ?", value)
	}
}

func UniversalSearchBillNumber(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_bill.bill_number = ?", value)
	}
}

func UniversalSearchBillPaymentNumber(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_bill_payment.bill_payment_number = ?", value)
	}
}

func UniversalSearchPurchaseOrderNumber(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_purchase_order.purchase_order_number = ?", value)
	}
}

func UniversalSearchProcurementContractNumber(db *gorm.DB, value string) {
	if value != "" {
		db = db.Where("crm_procurement_contract.procurement_contract_number = ?", value)
	}
}
