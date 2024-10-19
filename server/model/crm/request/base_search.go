package request

type baseSearchReq struct {
	CustomerName              string `json:"customerName" form:"customerName" `                          //客户名
	OrderNumber               string `json:"orderNumber" form:"orderNumber"`                             //订单编号
	BillNumber                string `json:"billNumber" form:"billNumber" `                              //账单编号
	BillPaymentNumber         string `json:"billPaymentNumber" form:"billPaymentNumber"`                 //应付账单编号
	PurchaseOrderNumber       string `json:"purchaseOrderNumber" form:"purchaseOrderNumber"`             //订购单编号
	ProcurementContractNumber string `json:"procurementContractNumber" form:"procurementContractNumber"` //订购合同编号
}
