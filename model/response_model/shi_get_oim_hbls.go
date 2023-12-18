package response_model

type ShiGetOimHbls struct {
	Type    any                   `json:"type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rows    []ShiGetOimHblsRows   `json:"rows"`
	Footer  []ShiGetOimHblsFooter `json:"footer"`
	Total   int64                 `json:"total"`
	Success bool                  `json:"success"`
	Err     bool                  `json:"err"`
}

type ShiGetOimHblsFooter struct {
	TableKeyName      any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter          any     `json:"IsFooter"`
	Selector          any     `json:"selector"`
	KeyId             any     `json:"KeyId"`
	MasterId          any     `json:"MasterId"`
	Caller            any     `json:"caller"`
	OihHblId          int64   `json:"oih_hbl_id"`
	OimMasterId       int64   `json:"oim_master_id"`
	EfcCompanyId      int64   `json:"efc_company_id"`
	EfcBranchId       int64   `json:"efc_branch_id"`
	OihHblSeqNo       int64   `json:"oih_hbl_seq_no"`
	OihHblDisplaySeq  any     `json:"oih_hbl_display_seq"`
	OihDocCode        any     `json:"oih_doc_code"`
	StatusDesc        any     `json:"status_desc"`
	Rowversion        any     `json:"rowversion"`
	OihHblId01        int64   `json:"oih_hbl_id_01"`
	EfcCompanyId01    any     `json:"efc_company_id_01"`
	EfcBranchId01     any     `json:"efc_branch_id_01"`
	Rowversion01      any     `json:"rowversion_01"`
	ConsigneeName     any     `json:"consignee_name"`
	BrokerName        any     `json:"broker_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	TotalPackage      int64   `json:"total_package"`
	ShipperName       any     `json:"shipper_name"`
	TotalWeightKg     int64   `json:"total_weight_kg"`
	TotalWeightLb     float64 `json:"total_weight_lb"`
	TotalVolumeCbm    int64   `json:"total_volume_cbm"`
	TotalVolumeCbf    float64 `json:"total_volume_cbf"`
	HouseNo           any     `json:"house_no"`
	CollectInvoiceAmt float64 `json:"collect_invoice_amt"`
	HousePpCcCode     any     `json:"house_pp_cc_code"`
	LastFreeDate      any     `json:"last_free_date"`
	HouseServiceType  any     `json:"house_service_type"`
	ToDoorSw          any     `json:"to_door_sw"`
	AvailableDate     any     `json:"available_date"`
	CreatedDate       any     `json:"created_date"`
	ModifiedDate      any     `json:"modified_date"`
	CreatedByCode     any     `json:"created_by_code"`
	ModifiedByCode    any     `json:"modified_by_code"`
	ConsigneeId       any     `json:"consignee_id"`
	ObjectState       int64   `json:"ObjectState"`
}

type ShiGetOimHblsRows struct {
	TableKeyName      any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter          any     `json:"IsFooter"`
	Selector          any     `json:"selector"`
	KeyId             int64   `json:"KeyId"`
	MasterId          any     `json:"MasterId"`
	Caller            any     `json:"caller"`
	OihHblId          int64   `json:"oih_hbl_id"`
	OimMasterId       int64   `json:"oim_master_id"`
	EfcCompanyId      int64   `json:"efc_company_id"`
	EfcBranchId       int64   `json:"efc_branch_id"`
	OihHblSeqNo       int64   `json:"oih_hbl_seq_no"`
	OihHblDisplaySeq  int64   `json:"oih_hbl_display_seq"`
	OihDocCode        any     `json:"oih_doc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	StatusDesc        any     `json:"status_desc"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rowversion        []int64 `json:"rowversion"`
	OihHblId01        int64   `json:"oih_hbl_id_01"`
	EfcCompanyId01    int64   `json:"efc_company_id_01"`
	EfcBranchId01     int64   `json:"efc_branch_id_01"`
	Rowversion01      []int64 `json:"rowversion_01"`
	ConsigneeName     any     `json:"consignee_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BrokerName        any     `json:"broker_name"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
	TotalPackage      int64   `json:"total_package"`
	ShipperName       any     `json:"shipper_name"`
	TotalWeightKg     int64   `json:"total_weight_kg"`
	TotalWeightLb     float64 `json:"total_weight_lb"`
	TotalVolumeCbm    int64   `json:"total_volume_cbm"`
	TotalVolumeCbf    float64 `json:"total_volume_cbf"`
	HouseNo           any     `json:"house_no"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CollectInvoiceAmt float64 `json:"collect_invoice_amt"`
	HousePpCcCode     any     `json:"house_pp_cc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	LastFreeDate      any     `json:"last_free_date"`
	HouseServiceType  any     `json:"house_service_type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ToDoorSw          any     `json:"to_door_sw"`
	AvailableDate     any     `json:"available_date"`
	CreatedDate       any     `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedDate      any     `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CreatedByCode     any     `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedByCode    any     `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ConsigneeId       int64   `json:"consignee_id"`
	ObjectState       int64   `json:"ObjectState"`
}
