package response_model

type AirQueryAimMaster struct {
	Rows    []AirQueryAimMasterRows `json:"rows"`
	Footer  any                     `json:"footer"`
	Total   int64                   `json:"total"`
	Success bool                    `json:"success"`
	Err     bool                    `json:"err"`
}

type AirQueryAimMasterRows struct {
	TableKeyName           any   `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter               any   `json:"IsFooter"`
	Selector               any   `json:"selector"`
	KeyId                  int64 `json:"KeyId"`
	MasterId               any   `json:"MasterId"`
	Caller                 any   `json:"caller"`
	AimMasterId            int64 `json:"aim_master_id"`
	EfcCompanyId           int   `json:"efc_company_id"`
	EfcBranchId            int   `json:"efc_branch_id"`
	AimMasterSeqNo         int   `json:"aim_master_seq_no"`
	AimDocCode             any   `json:"aim_doc_code"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BranchCode             any   `json:"branch_code"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterPpCcCode         any   `json:"master_pp_cc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	OpType                 any   `json:"op_type"`           // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MawbServiceType        any   `json:"mawb_service_type"`
	StatusId               any   `json:"status_id"`
	CarrierId              any   `json:"carrier_id"`
	AgentId                any   `json:"agent_id"`
	DepartureAirportCode   any   `json:"departure_airport_code"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DestinationAirportCode any   `json:"destination_airport_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DepartureAirport       any   `json:"departure_airport"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DestinationAirport     any   `json:"destination_airport"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DepartureMetroCode     any   `json:"departure_metro_code"`
	DestinationMetroCode   any   `json:"destination_metro_code"`
	MFlightDate            any   `json:"m_flight_date"`
	FlightNoN              any   `json:"flight_no_n"`
	EtaDate                any   `json:"eta_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterBlNo             any   `json:"master_bl_no"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	SubMasterBlNo          any   `json:"sub_master_bl_no"`
	BillingDate            any   `json:"billing_date"`
	CloseDate              any   `json:"close_date"`
	ColoaderId             any   `json:"coloader_id"`
	DevanningId            any   `json:"devanning_id"`
	AgentLotRefNo          any   `json:"agent_lot_ref_no"`
	MawbPiece              any   `json:"mawb_piece"`
	ConsolePallets         any   `json:"console_pallets"`
	AimNote                any   `json:"aim_note"`
	PeriodId               any   `json:"period_id"`
	CurrencyCode           any   `json:"currency_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CurrencyRate           any   `json:"currency_rate"`
	CurrencyRateOverride   any   `json:"currency_rate_override"`
	ArmInvoiceId           any   `json:"arm_invoice_id"`
	ApiVendorInvoiceId     any   `json:"api_vendor_invoice_id"`
	ApiCarrierInvoiceId    any   `json:"api_carrier_invoice_id"`
	Hbls                   any   `json:"hbls"`
	LastFtpProcess         any   `json:"last_ftp_process"`
	FtpSw                  any   `json:"ftp_sw"`
	BillToAgentId          any   `json:"bill_to_agent_id"`
	MawbPrefix             any   `json:"mawb_prefix"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MawbSerial             any   `json:"mawb_serial"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	AirlineCode            any   `json:"airline_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MawbDirectSw           any   `json:"mawb_direct_sw"`
	CombinedHblSw          any   `json:"combined_hbl_sw"`
	OtherCurrencyCode      any   `json:"other_currency_code"`
	MasterCollectAmt       any   `json:"master_collect_amt"`
	DestinationCostAmt     any   `json:"destination_cost_amt"`
	CreditToAgentAmt       any   `json:"credit_to_agent_amt"`
	AgentFee               any   `json:"agent_fee"`
	AgentBalance           any   `json:"agent_balance"`
	AgentFeeCurrency       any   `json:"agent_fee_currency"`
	AgentBalanceCurrency   any   `json:"agent_balance_currency"`
	AgentInvoiceAmt        any   `json:"agent_invoice_amt"`
	CurrencyXDCode         any   `json:"currency_x_d_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ArPaidSw               any   `json:"ar_paid_sw"`
	ApPaidSw               any   `json:"ap_paid_sw"`
	OriginalCcn            any   `json:"original_ccn"`
	SplitSw                any   `json:"split_sw"`
	DeliveryDate           any   `json:"delivery_date"`
	AimFlightNo            any   `json:"aim_flight_no"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	AimFlightDate          any   `json:"aim_flight_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	RcsDate                any   `json:"rcs_date"`
	AtdDate                any   `json:"atd_date"`
	AimArrivalDate         any   `json:"aim_arrival_date"`
	RcfDate                any   `json:"rcf_date"`
	AwdDate                any   `json:"awd_date"`
	NfdDate                any   `json:"nfd_date"`
	PickupDate             any   `json:"pickup_date"`
	LastStatusId           any   `json:"last_status_id"`
	LastEventDate          any   `json:"last_event_date"`
	PendingExRateSw        any   `json:"pending_ex_rate_sw"`
	AgentName              any   `json:"agent_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CarrierName            any   `json:"carrier_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ColoaderName           any   `json:"coloader_name"`
	DevanningName          any   `json:"devanning_name"`
	BillToAgentName        any   `json:"bill_to_agent_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	EdiMasterDocCode       any   `json:"edi_master_doc_code"`
	EdiImportTypeId        any   `json:"edi_import_type_id"`
	OldAimId               any   `json:"old_aim_id"`
	OldDocCodeMaster       any   `json:"old_doc_code_master"`
	OldAimBranchCode       any   `json:"old_aim_branch_code"`
	AbiReleaseStatusId     any   `json:"abi_release_status_id"`
	MasterCostByHblDistr   any   `json:"master_cost_by_hbl_distr"`
	MasterCostByHblCount   any   `json:"master_cost_by_hbl_count"`
	MasterOtherYymmSum     any   `json:"master_other_yymm_sum"`
	AimFirmCode            any   `json:"aim_firm_code"`
	AgentGroupCode         any   `json:"agent_group_code"`
	LastEventDesc          any   `json:"last_event_desc"`
	MasterAlertCode        any   `json:"master_alert_code"`
	SelSw                  any   `json:"sel_sw"`
	SelByCode              any   `json:"sel_by_code"`
	CreatedDate            any   `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedDate           any   `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CreatedByCode          any   `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedByCode         any   `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rowversion             any   `json:"rowversion"`
	StatusDesc             any   `json:"status_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ObjectState            any   `json:"ObjectState"`
}
