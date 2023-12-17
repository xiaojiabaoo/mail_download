package response_model

type AirGetAimHbls struct {
	Type    any                   `json:"type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rows    []AirGetAimHblsRows   `json:"rows"`
	Footer  []AirGetAimHblsFooter `json:"footer"`
	Total   int64                 `json:"total"`
	Success bool                  `json:"success"`
	Err     bool                  `json:"err"`
}

type AirGetAimHblsRows struct {
	TableKeyName      any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter          any     `json:"IsFooter"`
	Selector          any     `json:"selector"`
	KeyId             int64   `json:"KeyId"`
	MasterId          any     `json:"MasterId"`
	Caller            any     `json:"caller"`
	AihHblId          int64   `json:"aih_hbl_id"`
	AimMasterId       int64   `json:"aim_master_id"`
	EfcCompanyId      int     `json:"efc_company_id"`
	EfcBranchId       int     `json:"efc_branch_id"`
	AihHblSeqNo       int     `json:"aih_hbl_seq_no"`
	AihHblDisplaySeq  int     `json:"aih_hbl_display_seq"`
	AihDocCode        any     `json:"aih_doc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	StatusId          int     `json:"status_id"`
	HouseNo           any     `json:"house_no"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ShipperName       any     `json:"shipper_name"`
	ChargeableUnitKg  float64 `json:"chargeable_unit_kg"`
	TotalPackage      int     `json:"total_package"`
	CollectInvoiceAmt float64 `json:"collect_invoice_amt"`
	ConsigneeName     any     `json:"consignee_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BrokerName        any     `json:"broker_name"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CreatedDate       any     `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CreatedByCode     any     `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedDate      any     `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedByCode    any     `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rowversion        []int   `json:"rowversion"`
	AihHblId01        int64   `json:"aih_hbl_id_01"`
	EfcCompanyId01    int     `json:"efc_company_id_01"`
	EfcBranchId01     int     `json:"efc_branch_id_01"`
	BankReleaseDate   any     `json:"bank_release_date"`
	CheckReceiveDate  any     `json:"check_receive_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CargoReleaseDate  any     `json:"cargo_release_date"`
	Rowversion01      []int   `json:"rowversion_01"`
	StatusDesc        any     `json:"status_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ConsigneeId       int64   `json:"consignee_id"`
	ObjectState       int     `json:"ObjectState"`
}

type AirGetAimHblsFooter struct {
	TableKeyName      any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter          any     `json:"IsFooter"`
	Selector          any     `json:"selector"`
	KeyId             any     `json:"KeyId"`
	MasterId          any     `json:"MasterId"`
	Caller            any     `json:"caller"`
	AihHblId          int     `json:"aih_hbl_id"`
	AimMasterId       int     `json:"aim_master_id"`
	EfcCompanyId      any     `json:"efc_company_id"`
	EfcBranchId       any     `json:"efc_branch_id"`
	AihHblSeqNo       any     `json:"aih_hbl_seq_no"`
	AihHblDisplaySeq  any     `json:"aih_hbl_display_seq"`
	AihDocCode        any     `json:"aih_doc_code"`
	StatusId          any     `json:"status_id"`
	HouseNo           any     `json:"house_no"`
	ShipperName       any     `json:"shipper_name"`
	ChargeableUnitKg  float64 `json:"chargeable_unit_kg"`
	TotalPackage      int     `json:"total_package"`
	CollectInvoiceAmt any     `json:"collect_invoice_amt"`
	ConsigneeName     any     `json:"consignee_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BrokerName        any     `json:"broker_name"`
	CreatedDate       any     `json:"created_date"`
	CreatedByCode     any     `json:"created_by_code"`
	ModifiedDate      any     `json:"modified_date"`
	ModifiedByCode    any     `json:"modified_by_code"`
	Rowversion        any     `json:"rowversion"`
	AihHblId01        int     `json:"aih_hbl_id_01"`
	EfcCompanyId01    any     `json:"efc_company_id_01"`
	EfcBranchId01     any     `json:"efc_branch_id_01"`
	BankReleaseDate   any     `json:"bank_release_date"`
	CheckReceiveDate  any     `json:"check_receive_date"`
	CargoReleaseDate  any     `json:"cargo_release_date"`
	Rowversion01      any     `json:"rowversion_01"`
	StatusDesc        any     `json:"status_desc"`
	ConsigneeId       any     `json:"consignee_id"`
	ObjectState       int     `json:"ObjectState"`
}
