package response_model

type AirGetAimReviewInvoiceApis struct {
	Type any `json:"type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rows []struct {
		TableKeyName           any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		IsFooter               any     `json:"IsFooter"`
		Selector               any     `json:"selector"`
		KeyId                  int64   `json:"KeyId"`
		MasterId               any     `json:"MasterId"`
		Caller                 any     `json:"caller"`
		ApiInvoiceId           int64   `json:"api_invoice_id"`
		EfcCompanyId           int     `json:"efc_company_id"`
		EfcBranchId            int     `json:"efc_branch_id"`
		ApiInvoiceSeqNo        int     `json:"api_invoice_seq_no"`
		BranchCode             any     `json:"branch_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		OpType                 any     `json:"op_type"`      // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ApiDocCode             any     `json:"api_doc_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		StatusId               int     `json:"status_id"`
		InvoiceTypeId          int     `json:"invoice_type_id"`
		VendorId               int64   `json:"vendor_id"`
		PayeeId                int64   `json:"payee_id"`
		VendorInvoiceNo        any     `json:"vendor_invoice_no"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PeriodId               int     `json:"period_id"`
		InvoiceAmt             float64 `json:"invoice_amt"`
		ClearAmt               float64 `json:"clear_amt"`
		ClearAmtCurrency       float64 `json:"clear_amt_currency"`
		ItemAmt                float64 `json:"item_amt"`
		VatTaxAmt              float64 `json:"vat_tax_amt"`
		DiscountAmt            float64 `json:"discount_amt"`
		PaidAmt                float64 `json:"paid_amt"`
		InvoiceDate            any     `json:"invoice_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		InvoiceNoAssignDate    any     `json:"invoice_no_assign_date"`
		DueDate                any     `json:"due_date"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PaidDate               any     `json:"paid_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PaidSw                 bool    `json:"paid_sw"`
		DirectEntrySw          bool    `json:"direct_entry_sw"`
		RefNo                  any     `json:"ref_no"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ApiNote                any     `json:"api_note"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ApiActId               int     `json:"api_act_id"`
		PaidByCode             any     `json:"paid_by_code"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeMaster          any     `json:"doc_code_master"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeHouse           any     `json:"doc_code_house"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		MasterId1              int64   `json:"master_id"`
		HouseId                int64   `json:"house_id"`
		InvoicedTransId        int64   `json:"invoiced_trans_id"`
		PaymentTransId         int64   `json:"payment_trans_id"`
		AppPaymentId           int64   `json:"app_payment_id"`
		ApxClearId             any     `json:"apx_clear_id"`
		PaymentBatchId         int64   `json:"payment_batch_id"`
		ApxClearSw             any     `json:"apx_clear_sw"`
		AgentInvSw             bool    `json:"agent_inv_sw"`
		ExtraInvSw             bool    `json:"extra_inv_sw"`
		NonTradeSw             bool    `json:"non_trade_sw"`
		Eta                    any     `json:"eta"`
		Etd                    any     `json:"etd"`
		SelSw                  bool    `json:"sel_sw"`
		SelByCode              any     `json:"sel_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		SelByDate              any     `json:"sel_by_date"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CurrencySw             bool    `json:"currency_sw"`
		CurrencyCode           any     `json:"currency_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CurrencyRate           any     `json:"currency_rate"`
		CurrencyRateOverride   any     `json:"currency_rate_override"`
		CurrencyXDCode         any     `json:"currency_x_d_code"`
		InvoiceAmtCurrency     float64 `json:"invoice_amt_currency"`
		PaidAmtCurrency        float64 `json:"paid_amt_currency"`
		DiscountAmtCurrency    float64 `json:"discount_amt_currency"`
		AdditionalTaxAmt       float64 `json:"additional_tax_amt"`
		OrigInvoiceAmt         float64 `json:"orig_invoice_amt"`
		OrigInvoiceAmtCurrency float64 `json:"orig_invoice_amt_currency"`
		OrigId                 any     `json:"orig_id"`
		ExchRateDiff           float64 `json:"exch_rate_diff"`
		MasterBlNo             any     `json:"master_bl_no"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
		HouseNo                any     `json:"house_no"`            // 类型string，因为对应的Java语言String的默认值是null，以免报错
		EtaDate                any     `json:"eta_date"`            // 类型string，因为对应的Java语言String的默认值是null，以免报错
		EtdDate                any     `json:"etd_date"`            // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CarrierCode            any     `json:"carrier_code"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
		LoadingPortName        any     `json:"loading_port_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DischargePortName      any     `json:"discharge_port_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		VendorName             any     `json:"vendor_name"`         // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PayeeName              any     `json:"payee_name"`          // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CvmTypeCode            any     `json:"cvm_type_code"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
		VatTaxSw               bool    `json:"vat_tax_sw"`
		ApWithholdTaxAmt       float64 `json:"ap_withhold_tax_amt"`
		FclContainerDesc       any     `json:"fcl_container_desc"`
		TaxedById              any     `json:"taxed_by_id"`
		StatementSelSw         bool    `json:"statement_sel_sw"`
		PaymentRefNo           any     `json:"payment_ref_no"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PrepaidSw              bool    `json:"prepaid_sw"`
		ArcClearId             any     `json:"arc_clear_id"`
		ArpPaymentId           any     `json:"arp_payment_id"`
		AgentCcPaidSw          bool    `json:"agent_cc_paid_sw"`
		RequestPayDate         any     `json:"request_pay_date"`
		RequestPayBy           any     `json:"request_pay_by"`
		ReqApprovedDate        any     `json:"req_approved_date"`
		ReqApprovedBy          any     `json:"req_approved_by"`
		PlaceOfDelivery        any     `json:"place_of_delivery"`
		IssuedGovInvoiceSw     bool    `json:"issued_gov_invoice_sw"`
		SelInvSw               bool    `json:"sel_inv_sw"`
		SelInvByCode           any     `json:"sel_inv_by_code"`
		ApproveToPayBy         any     `json:"approve_to_pay_by"`
		ApproveToPayDae        any     `json:"approve_to_pay_dae"`
		CreatedDate            any     `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedDate           any     `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CreatedByCode          any     `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedByCode         any     `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Rowversion             []int   `json:"rowversion"`
		StatusDesc             any     `json:"status_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ObjectState            int     `json:"ObjectState"`
	} `json:"rows"`
	Footer []struct {
		TableKeyName           any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		IsFooter               any     `json:"IsFooter"`
		Selector               any     `json:"selector"`
		KeyId                  any     `json:"KeyId"`
		MasterId               any     `json:"MasterId"`
		Caller                 any     `json:"caller"`
		ApiInvoiceId           int     `json:"api_invoice_id"`
		EfcCompanyId           any     `json:"efc_company_id"`
		EfcBranchId            any     `json:"efc_branch_id"`
		ApiInvoiceSeqNo        any     `json:"api_invoice_seq_no"`
		BranchCode             any     `json:"branch_code"`
		OpType                 any     `json:"op_type"`
		ApiDocCode             any     `json:"api_doc_code"`
		StatusId               any     `json:"status_id"`
		InvoiceTypeId          any     `json:"invoice_type_id"`
		VendorId               int     `json:"vendor_id"`
		PayeeId                any     `json:"payee_id"`
		VendorInvoiceNo        any     `json:"vendor_invoice_no"`
		PeriodId               any     `json:"period_id"`
		InvoiceAmt             any     `json:"invoice_amt"`
		ClearAmt               any     `json:"clear_amt"`
		ClearAmtCurrency       any     `json:"clear_amt_currency"`
		ItemAmt                any     `json:"item_amt"`
		VatTaxAmt              any     `json:"vat_tax_amt"`
		DiscountAmt            any     `json:"discount_amt"`
		PaidAmt                float64 `json:"paid_amt"`
		InvoiceDate            any     `json:"invoice_date"`
		InvoiceNoAssignDate    any     `json:"invoice_no_assign_date"`
		DueDate                any     `json:"due_date"`
		PaidDate               any     `json:"paid_date"`
		PaidSw                 any     `json:"paid_sw"`
		DirectEntrySw          any     `json:"direct_entry_sw"`
		RefNo                  any     `json:"ref_no"`
		ApiNote                any     `json:"api_note"`
		ApiActId               any     `json:"api_act_id"`
		PaidByCode             any     `json:"paid_by_code"`
		DocCodeMaster          any     `json:"doc_code_master"`
		DocCodeHouse           any     `json:"doc_code_house"`
		MasterId1              any     `json:"master_id"`
		HouseId                any     `json:"house_id"`
		InvoicedTransId        any     `json:"invoiced_trans_id"`
		PaymentTransId         any     `json:"payment_trans_id"`
		AppPaymentId           any     `json:"app_payment_id"`
		ApxClearId             any     `json:"apx_clear_id"`
		PaymentBatchId         any     `json:"payment_batch_id"`
		ApxClearSw             any     `json:"apx_clear_sw"`
		AgentInvSw             any     `json:"agent_inv_sw"`
		ExtraInvSw             any     `json:"extra_inv_sw"`
		NonTradeSw             any     `json:"non_trade_sw"`
		Eta                    any     `json:"eta"`
		Etd                    any     `json:"etd"`
		SelSw                  any     `json:"sel_sw"`
		SelByCode              any     `json:"sel_by_code"`
		SelByDate              any     `json:"sel_by_date"`
		CurrencySw             any     `json:"currency_sw"`
		CurrencyCode           any     `json:"currency_code"`
		CurrencyRate           any     `json:"currency_rate"`
		CurrencyRateOverride   any     `json:"currency_rate_override"`
		CurrencyXDCode         any     `json:"currency_x_d_code"`
		InvoiceAmtCurrency     any     `json:"invoice_amt_currency"`
		PaidAmtCurrency        float64 `json:"paid_amt_currency"`
		DiscountAmtCurrency    any     `json:"discount_amt_currency"`
		AdditionalTaxAmt       any     `json:"additional_tax_amt"`
		OrigInvoiceAmt         float64 `json:"orig_invoice_amt"`
		OrigInvoiceAmtCurrency float64 `json:"orig_invoice_amt_currency"`
		OrigId                 any     `json:"orig_id"`
		ExchRateDiff           any     `json:"exch_rate_diff"`
		MasterBlNo             any     `json:"master_bl_no"`
		HouseNo                any     `json:"house_no"`
		EtaDate                any     `json:"eta_date"`
		EtdDate                any     `json:"etd_date"`
		CarrierCode            any     `json:"carrier_code"`
		LoadingPortName        any     `json:"loading_port_name"`
		DischargePortName      any     `json:"discharge_port_name"`
		VendorName             any     `json:"vendor_name"`
		PayeeName              any     `json:"payee_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CvmTypeCode            any     `json:"cvm_type_code"`
		VatTaxSw               any     `json:"vat_tax_sw"`
		ApWithholdTaxAmt       any     `json:"ap_withhold_tax_amt"`
		FclContainerDesc       any     `json:"fcl_container_desc"`
		TaxedById              any     `json:"taxed_by_id"`
		StatementSelSw         any     `json:"statement_sel_sw"`
		PaymentRefNo           any     `json:"payment_ref_no"`
		PrepaidSw              any     `json:"prepaid_sw"`
		ArcClearId             any     `json:"arc_clear_id"`
		ArpPaymentId           any     `json:"arp_payment_id"`
		AgentCcPaidSw          any     `json:"agent_cc_paid_sw"`
		RequestPayDate         any     `json:"request_pay_date"`
		RequestPayBy           any     `json:"request_pay_by"`
		ReqApprovedDate        any     `json:"req_approved_date"`
		ReqApprovedBy          any     `json:"req_approved_by"`
		PlaceOfDelivery        any     `json:"place_of_delivery"`
		IssuedGovInvoiceSw     any     `json:"issued_gov_invoice_sw"`
		SelInvSw               any     `json:"sel_inv_sw"`
		SelInvByCode           any     `json:"sel_inv_by_code"`
		ApproveToPayBy         any     `json:"approve_to_pay_by"`
		ApproveToPayDae        any     `json:"approve_to_pay_dae"`
		CreatedDate            any     `json:"created_date"`
		ModifiedDate           any     `json:"modified_date"`
		CreatedByCode          any     `json:"created_by_code"`
		ModifiedByCode         any     `json:"modified_by_code"`
		Rowversion             any     `json:"rowversion"`
		StatusDesc             any     `json:"status_desc"`
		ObjectState            int     `json:"ObjectState"`
	} `json:"footer"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
	Err     bool `json:"err"`
}
