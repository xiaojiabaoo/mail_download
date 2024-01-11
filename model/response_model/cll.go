package response_model

type LoginResp struct {
	Lang        string `json:"lang"`
	Ver         string `json:"ver"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	Url         string `json:"url"`
	LoginCookie string `json:"login_cookie"`
	MainCookie  string `json:"main_cookie"`
}

type ReadPdf struct {
	InvoiceNo    string
	MasterMawbNo []string
	Amount       string
}

type AirQueryInfo struct {
	Ver         int64 `json:"ver"`
	Vid         any   `json:"vid"`         //类型string，因为对应的Java语言String的默认值是null，以免报错
	Query       any   `json:"query"`       //类型string，因为对应的Java语言String的默认值是null，以免报错
	View        any   `json:"view"`        //类型string，因为对应的Java语言String的默认值是null，以免报错
	DefaultView any   `json:"defaultView"` //类型string，因为对应的Java语言String的默认值是null，以免报错
	Syscode     any   `json:"syscode"`     //类型string，因为对应的Java语言String的默认值是null，以免报错
}

type AirSaveAihCharge struct {
	BlockUI string `json:"blockUI"`
	Row     struct {
		TableKeyName          any `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		IsFooter              any `json:"IsFooter"`
		Selector              any `json:"selector"`
		KeyId                 any `json:"KeyId"`
		MasterId              any `json:"MasterId"`
		Caller                any `json:"caller"`
		PlmChargeCostId       any `json:"plm_charge_cost_id"`
		EfcCompanyId          any `json:"efc_company_id"`
		EfcBranchId           any `json:"efc_branch_id"`
		PlmChargeCostSeqNo    any `json:"plm_charge_cost_seq_no"`
		BranchCode            any `json:"branch_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PliChargeCostItemId   any `json:"pli_charge_cost_item_id"`
		PliChargeCostItemNo   any `json:"pli_charge_cost_item_no"`
		OpType                any `json:"op_type"`            // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ChargeCostType        any `json:"charge_cost_type"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmCodeIndicator      any `json:"plm_code_indicator"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmCatCode            any `json:"plm_cat_code"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmMasterId           any `json:"plm_master_id"`
		PlmHblId              any `json:"plm_hbl_id"`
		PlmHblDisplaySequence any `json:"plm_hbl_display_sequence"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeMaster         any `json:"doc_code_master"`          // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeHouse          any `json:"doc_code_house"`           // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PrepaidCharge         any `json:"prepaid_charge"`
		CollectCharge         any `json:"collect_charge"`
		ArCharge              any `json:"ar_charge"`
		PrepaidCost           any `json:"prepaid_cost"`
		CollectCost           any `json:"collect_cost"`
		ApCost                any `json:"ap_cost"`
		CurrencySw            any `json:"currency_sw"`
		PrepaidChargeCurrency any `json:"prepaid_charge_currency"`
		CollectChargeCurrency any `json:"collect_charge_currency"`
		ArChargeCurrency      any `json:"ar_charge_currency"`
		PrepaidCostCurrency   any `json:"prepaid_cost_currency"`
		CollectCostCurrency   any `json:"collect_cost_currency"`
		ApCostCurrency        any `json:"ap_cost_currency"`
		CurrencyCode          any `json:"currency_code"`
		CurrencyRate          any `json:"currency_rate"`
		CurrencyXDCode        any `json:"currency_x_d_code"`
		ArChargeOtherCurrency any `json:"ar_charge_other_currency"`
		ApCostOtherCurrency   any `json:"ap_cost_other_currency"`
		Reference             any `json:"reference"`
		ExtraDesc             any `json:"extra_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		VendorId              any `json:"vendor_id"`
		CustomerId            any `json:"customer_id"`
		ArBilledSw            any `json:"ar_billed_sw"`
		ApBilledSw            any `json:"ap_billed_sw"`
		AgentCalcSw           any `json:"agent_calc_sw"`
		ExtraInvSw            any `json:"extra_inv_sw"`
		ExtraCostSw           any `json:"extra_cost_sw"`
		DueAgentSw            any `json:"due_agent_sw"`
		ArmInvoiceId          any `json:"arm_invoice_id"`
		ArmInvoiceDetailId    any `json:"arm_invoice_detail_id"`
		ApiInvoiceId          any `json:"api_invoice_id"`
		ApiInvoiceDetailId    any `json:"api_invoice_detail_id"`
		RateFactorTotal       any `json:"rate_factor_total"`
		RateUnitCode          any `json:"rate_unit_code"`
		RateAmt               any `json:"rate_amt"`
		CalculateAmt          any `json:"calculate_amt"`
		FrtDistributeCost     any `json:"frt_distribute_cost"`
		VendorName            any `json:"vendor_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CustomerName          any `json:"customer_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ItemCharge            any `json:"item_charge"`
		ItemCost              any `json:"item_cost"`
		VatTaxSw              any `json:"vat_tax_sw"`
		ApTaxPercent          any `json:"ap_tax_percent"`
		ApTaxAmt              any `json:"ap_tax_amt"`
		ApWithholdTaxPercent  any `json:"ap_withhold_tax_percent"`
		ApWithholdTaxAmt      any `json:"ap_withhold_tax_amt"`
		ArWithholdTaxPercent  any `json:"ar_withhold_tax_percent"`
		ArWithholdTaxAmt      any `json:"ar_withhold_tax_amt"`
		ArTaxPercent          any `json:"ar_tax_percent"`
		ArTaxAmt              any `json:"ar_tax_amt"`
		ApTaxCurrency         any `json:"ap_tax_currency"`
		ArTaxCurrency         any `json:"ar_tax_currency"`
		PeriodId              any `json:"period_id"`
		AgentArId             any `json:"agent_ar_id"`
		AgentApId             any `json:"agent_ap_id"`
		AgentName             any `json:"agent_name"`
		OldPlmChargeCostId    any `json:"old_plm_charge_cost_id"`
		OldPlmBranchCode      any `json:"old_plm_branch_code"`
		PlmDisplayFloatSeq    any `json:"plm_display_float_seq"`
		CreatedDate           any `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CreatedByCode         any `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedDate          any `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedByCode        any `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Rowversion            any `json:"rowversion"`
		AdvanceItemSw         any `json:"advance_item_sw"`
		ChargeCostDesc        any `json:"charge_cost_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		SelSw                 any `json:"sel_sw"`
		X12ChargeCode         any `json:"x12_charge_code"`
		ObjectState           any `json:"ObjectState"`
	} `json:"row"`
	Action any `json:"action"`
	Frm    any `json:"frm"`
	Op     any `json:"op"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Mode   any `json:"mode"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Mid    any `json:"mid"`
	Obj    struct {
		Mdate  any `json:"mdate"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Optype any `json:"optype"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	} `json:"obj"`
	Id     int64 `json:"id"`
	Code   int   `json:"code"`
	MsgObj struct {
		Message      any `json:"message"`
		Reason       any `json:"reason"`
		MsgTypeCode  any `json:"msg_type_code"`
		MsgAskType   any `json:"msg_ask_type"`
		MsgId        any `json:"msg_id"`
		FocusField   any `json:"focus_field"`
		ExtraMessage any `json:"extra_message"`
		ExtraJs      any `json:"extra_js"`
		Caller       any `json:"caller"`
		Callf        any `json:"callf"`
		CallAction   any `json:"callAction"`
		CallObjStr   any `json:"callObjStr"`
		CallUrl      any `json:"callUrl"`
		CallSubform  any `json:"callSubform"`
		CallButton   any `json:"callButton"`
		CallAskId    any `json:"callAskId"`
		CallBack     any `json:"callBack"`
		WinName      any `json:"win_name"`
	} `json:"msgObj"`
}

type AirAihHouseExec struct {
	Action int    `json:"action"`
	Op     string `json:"op"`
	Obj    struct {
		Returnobj any `json:"returnobj"`
		Js        any `json:"js"`
		Err       any `json:"err"`
		MsgObj    struct {
			Message      string `json:"message"`
			Reason       string `json:"reason"`
			MsgTypeCode  string `json:"msg_type_code"`
			MsgAskType   any    `json:"msg_ask_type"`
			MsgId        int    `json:"msg_id"`
			FocusField   any    `json:"focus_field"`
			ExtraMessage any    `json:"extra_message"`
			ExtraJs      any    `json:"extra_js"`
			Caller       any    `json:"caller"`
			Callf        any    `json:"callf"`
			CallAction   any    `json:"callAction"`
			CallObjStr   any    `json:"callObjStr"`
			CallUrl      any    `json:"callUrl"`
			CallSubform  any    `json:"callSubform"`
			CallButton   any    `json:"callButton"`
			CallAskId    any    `json:"callAskId"`
			CallBack     any    `json:"callBack"`
			WinName      any    `json:"win_name"`
		} `json:"msgObj"`
		Ask   string `json:"ask"`
		Pwait int    `json:"pwait"`
	} `json:"obj"`
	Mode   string `json:"mode"`
	Id     int64  `json:"id"`
	Mid    int64  `json:"mid"`
	Code   int    `json:"code"`
	MsgObj struct {
		Message      string `json:"message"`
		Reason       string `json:"reason"`
		MsgTypeCode  string `json:"msg_type_code"`
		MsgAskType   any    `json:"msg_ask_type"`
		MsgId        int    `json:"msg_id"`
		FocusField   any    `json:"focus_field"`
		ExtraMessage any    `json:"extra_message"`
		ExtraJs      any    `json:"extra_js"`
		Caller       any    `json:"caller"`
		Callf        any    `json:"callf"`
		CallAction   any    `json:"callAction"`
		CallObjStr   any    `json:"callObjStr"`
		CallUrl      any    `json:"callUrl"`
		CallSubform  any    `json:"callSubform"`
		CallButton   any    `json:"callButton"`
		CallAskId    any    `json:"callAskId"`
		CallBack     any    `json:"callBack"`
		WinName      any    `json:"win_name"`
	} `json:"msgObj"`
}
