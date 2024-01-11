package response_model

type ShiQueryOimMaster struct {
	Rows    []ShiQueryOimMasterRows `json:"rows"`
	Footer  any                     `json:"footer"`
	Total   int64                   `json:"total"`
	Success bool                    `json:"success"`
	Err     bool                    `json:"err"`
}

type ShiQueryOimMasterRows struct {
	TableKeyName                 any   `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter                     any   `json:"IsFooter"`
	Selector                     any   `json:"selector"`
	KeyId                        int64 `json:"KeyId"`
	MasterId                     any   `json:"MasterId"`
	Caller                       any   `json:"caller"`
	OimMasterId                  int64 `json:"oim_master_id"`
	EfcCompanyId                 int64 `json:"efc_company_id"`
	EfcBranchId                  int64 `json:"efc_branch_id"`
	OimMasterSeqNo               any   `json:"oim_master_seq_no"`
	OimDocCode                   any   `json:"oim_doc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BranchCode                   any   `json:"branch_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	OpType                       any   `json:"op_type"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	StatusId                     any   `json:"status_id"`
	MasterBlNo                   any   `json:"master_bl_no"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterPpCcCode               any   `json:"master_pp_cc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	LoadType                     any   `json:"load_type"`         // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ConsignmentType              any   `json:"consignment_type"`
	MasterServiceType            any   `json:"master_service_type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterBlType                 any   `json:"master_bl_type"`
	MasterBlTypeCode             any   `json:"master_bl_type_code"`
	SubMasterBlNo                any   `json:"sub_master_bl_no"`
	FromCountryCode              any   `json:"from_country_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ToCountryCode                any   `json:"to_country_code"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CarrierCode                  any   `json:"carrier_code"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CarrierId                    any   `json:"carrier_id"`
	SeaWayBlSw                   any   `json:"sea_way_bl_sw"`
	OceanRoutingCode             any   `json:"ocean_routing_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	OceanRoutingId               any   `json:"ocean_routing_id"`
	VesselName                   any   `json:"vessel_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	VoyageNo                     any   `json:"voyage_no"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	VoyageNumNo                  any   `json:"voyage_num_no"`
	AgentId                      any   `json:"agent_id"`
	OimReceiptPort               any   `json:"oim_receipt_port"`
	OimReceiptPortUnCode         any   `json:"oim_receipt_port_un_code"`
	LoadingPort                  any   `json:"loading_port"`           // 类型string，因为对应的Java语言String的默认值是null，以免报错
	LoadingPortUnCode            any   `json:"loading_port_un_code"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DischargePort                any   `json:"discharge_port"`         // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DischargePortUnCode          any   `json:"discharge_port_un_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DischargeMetroCode           any   `json:"discharge_metro_code"`
	OimDeliveryPort              any   `json:"oim_delivery_port"`         // 类型string，因为对应的Java语言String的默认值是null，以免报错
	OimDeliveryPortUnCode        any   `json:"oim_delivery_port_un_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	OimReceiptDate               any   `json:"oim_receipt_date"`
	OnBoardDate                  any   `json:"on_board_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	VesselArriveDate             any   `json:"vessel_arrive_date"`
	MblReceiveDate               any   `json:"mbl_receive_date"`
	EtaDate                      any   `json:"eta_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ImDichargeDate               any   `json:"im_dicharge_date"`
	MasterDeliveryEtaDate        any   `json:"master_delivery_eta_date"`
	PeriodId                     any   `json:"period_id"`
	MblSurrenderReqSw            any   `json:"mbl_surrender_req_sw"`
	MblSurrenderDate             any   `json:"mbl_surrender_date"`
	StorageStartDate             any   `json:"storage_start_date"`
	DevanningInstructionSentDate any   `json:"devanning_instruction_sent_date"`
	DevanningCompleteDate        any   `json:"devanning_complete_date"`
	Teu                          any   `json:"teu"`
	ColoaderId                   any   `json:"coloader_id"`
	ColoadSw                     any   `json:"coload_sw"`
	DgSw                         any   `json:"dg_sw"`
	ReeferSw                     any   `json:"reefer_sw"`
	StopItSw                     any   `json:"stop_it_sw"`
	OimFreightLocId              any   `json:"oim_freight_loc_id"`
	DrayageId                    any   `json:"drayage_id"`
	AgentLotRefNo                any   `json:"agent_lot_ref_no"`
	OimNote                      any   `json:"oim_note"`
	CurrencyRate                 any   `json:"currency_rate"`
	CurrencyRateOverride         any   `json:"currency_rate_override"`
	CurrencyXDCode               any   `json:"currency_x_d_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BillingDate                  any   `json:"billing_date"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CloseDate                    any   `json:"close_date"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MblItNo                      any   `json:"mbl_it_no"`
	MblItDate                    any   `json:"mbl_it_date"`
	MblItEtaDate                 any   `json:"mbl_it_eta_date"`
	MblItCity                    any   `json:"mbl_it_city"`
	MblItCityUnCode              any   `json:"mbl_it_city_un_code"`
	MblItDestination             any   `json:"mbl_it_destination"`
	MblItDestinationUnCode       any   `json:"mbl_it_destination_un_code"`
	MblItIndicator               any   `json:"mbl_it_indicator"`
	InbondCarrierId              any   `json:"inbond_carrier_id"`
	ArmInvoiceId                 any   `json:"arm_invoice_id"`
	ApiVendorInvoiceId           any   `json:"api_vendor_invoice_id"`
	ApiCarrierInvoiceId          any   `json:"api_carrier_invoice_id"`
	Hbls                         any   `json:"hbls"`
	ContractNo                   any   `json:"contract_no"`
	InputTypeId                  any   `json:"input_type_id"`
	BillToAgentId                any   `json:"bill_to_agent_id"`
	AppPaymentId                 any   `json:"app_payment_id"`
	CombinedHblSw                any   `json:"combined_hbl_sw"`
	ArPaidSw                     any   `json:"ar_paid_sw"`
	ApPaidSw                     any   `json:"ap_paid_sw"`
	CurrencyCode                 any   `json:"currency_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterCollectAmt             any   `json:"master_collect_amt"`
	DestinationCostAmt           any   `json:"destination_cost_amt"`
	CreditToAgentAmt             any   `json:"credit_to_agent_amt"`
	AgentFee                     any   `json:"agent_fee"`
	AgentBalance                 any   `json:"agent_balance"`
	AgentFeeCurrency             any   `json:"agent_fee_currency"`
	AgentBalanceCurrency         any   `json:"agent_balance_currency"`
	AgentInvoiceAmt              any   `json:"agent_invoice_amt"`
	ScacCode                     any   `json:"scac_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	MasterDirectSw               any   `json:"master_direct_sw"`
	BookingNo                    any   `json:"booking_no"`
	MasterFreightPaidDate        any   `json:"master_freight_paid_date"`
	ContainerCount               any   `json:"container_count"`
	GateOutCount                 any   `json:"gate_out_count"`
	RdCount                      any   `json:"rd_count"`
	FclContainerDesc             any   `json:"fcl_container_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	PendingExRateSw              any   `json:"pending_ex_rate_sw"`
	AgentName                    any   `json:"agent_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CarrierName                  any   `json:"carrier_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ColoaderName                 any   `json:"coloader_name"`
	OimFreightLocName            any   `json:"oim_freight_loc_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BillToAgentName              any   `json:"bill_to_agent_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	InbondCarrierName            any   `json:"inbond_carrier_name"`
	LastStatusId                 any   `json:"last_status_id"`
	LastEventDate                any   `json:"last_event_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	EdiMasterDocCode             any   `json:"edi_master_doc_code"`
	EdiImportTypeId              any   `json:"edi_import_type_id"`
	OldOimId                     any   `json:"old_oim_id"`
	OldDocCodeMaster             any   `json:"old_doc_code_master"`
	OldOimBranchCode             any   `json:"old_oim_branch_code"`
	DrayageName                  any   `json:"drayage_name"`
	AbiReleaseStatusId           any   `json:"abi_release_status_id"`
	MasterCostByHblDistr         any   `json:"master_cost_by_hbl_distr"`
	MasterCostByHblCount         any   `json:"master_cost_by_hbl_count"`
	MasterOtherYymmSum           any   `json:"master_other_yymm_sum"`
	OimFirmCode                  any   `json:"oim_firm_code"`
	AgentGroupCode               any   `json:"agent_group_code"`
	ContainerLoadingId           any   `json:"container_loading_id"`
	ContainerLoadingName         any   `json:"container_loading_name"`
	CarrierInlandSw              any   `json:"carrier_inland_sw"`
	SocSw                        any   `json:"soc_sw"`
	PrevStatusId                 any   `json:"prev_status_id"`
	CarrierReleaseDate           any   `json:"carrier_release_date"`
	MasterAlertNo                any   `json:"master_alert_no"`
	FetchedFirmCode              any   `json:"fetched_firm_code"`
	SelSw                        any   `json:"sel_sw"`
	SelByCode                    any   `json:"sel_by_code"`
	CreatedDate                  any   `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedDate                 any   `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CreatedByCode                any   `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ModifiedByCode               any   `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	DevanningId                  any   `json:"devanning_id"`
	DevanningName                any   `json:"devanning_name"`
	DevanFirmCode                any   `json:"devan_firm_code"`
	Rowversion                   any   `json:"rowversion"`
	OimMasterId01                any   `json:"oim_master_id_01"`
	EfcCompanyId01               any   `json:"efc_company_id_01"`
	EfcBranchId01                any   `json:"efc_branch_id_01"`
	ExEmptyDispatched            any   `json:"ex_empty_dispatched"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ExDateOfReceipt              any   `json:"ex_date_of_receipt"`
	ExOnRailFeeder               any   `json:"ex_on_rail_feeder"`
	ExRailFeederUnload           any   `json:"ex_rail_feeder_unload"`
	ExTerminalGateIn             any   `json:"ex_terminal_gate_in"`
	ExOnBoardDate                any   `json:"ex_on_board_date"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
	LoadingPortAtdDate           any   `json:"loading_port_atd_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	TranshipDate                 any   `json:"tranship_date"`
	TranshipDate1                any   `json:"tranship_date1"`
	ImTerminalGateOut            any   `json:"im_terminal_gate_out"`
	ImOnRailFeeder               any   `json:"im_on_rail_feeder"`
	ImRailFeederDepart           any   `json:"im_rail_feeder_depart"`
	ImRailFeederArrive           any   `json:"im_rail_feeder_arrive"`
	ImRailFeederUnload           any   `json:"im_rail_feeder_unload"`
	OimAvailableDate             any   `json:"oim_available_date"`
	OimLastFreeDate              any   `json:"oim_last_free_date"`
	ImLastGateOutDate            any   `json:"im_last_gate_out_date"`
	ImDateOfDelivery             any   `json:"im_date_of_delivery"`
	RdDate                       any   `json:"rd_date"`
	SystemCloseDate              any   `json:"system_close_date"`
	CarrierAnRecvSw              any   `json:"carrier_an_recv_sw"`
	CarrierAnRecvDate            any   `json:"carrier_an_recv_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CarrierEtaDate               any   `json:"carrier_eta_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
	BerthingDate                 any   `json:"berthing_date"`
	BookEtdDate                  any   `json:"book_etd_date"`
	BookEtaDate                  any   `json:"book_eta_date"`
	BookDestEtaDate              any   `json:"book_dest_eta_date"`
	RailRefNo                    any   `json:"rail_ref_no"`
	Rowversion01                 any   `json:"rowversion_01"`
	StatusDesc                   any   `json:"status_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ObjectState                  any   `json:"ObjectState"`
}
