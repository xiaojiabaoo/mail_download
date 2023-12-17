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
		TableKeyName          any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		IsFooter              any     `json:"IsFooter"`
		Selector              any     `json:"selector"`
		KeyId                 int64   `json:"KeyId"`
		MasterId              any     `json:"MasterId"`
		Caller                any     `json:"caller"`
		PlmChargeCostId       int64   `json:"plm_charge_cost_id"`
		EfcCompanyId          int     `json:"efc_company_id"`
		EfcBranchId           int     `json:"efc_branch_id"`
		PlmChargeCostSeqNo    int     `json:"plm_charge_cost_seq_no"`
		BranchCode            any     `json:"branch_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PliChargeCostItemId   int64   `json:"pli_charge_cost_item_id"`
		PliChargeCostItemNo   int     `json:"pli_charge_cost_item_no"`
		OpType                any     `json:"op_type"`            // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ChargeCostType        any     `json:"charge_cost_type"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmCodeIndicator      any     `json:"plm_code_indicator"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmCatCode            any     `json:"plm_cat_code"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PlmMasterId           int64   `json:"plm_master_id"`
		PlmHblId              int64   `json:"plm_hbl_id"`
		PlmHblDisplaySequence any     `json:"plm_hbl_display_sequence"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeMaster         any     `json:"doc_code_master"`          // 类型string，因为对应的Java语言String的默认值是null，以免报错
		DocCodeHouse          any     `json:"doc_code_house"`           // 类型string，因为对应的Java语言String的默认值是null，以免报错
		PrepaidCharge         float64 `json:"prepaid_charge"`
		CollectCharge         float64 `json:"collect_charge"`
		ArCharge              float64 `json:"ar_charge"`
		PrepaidCost           float64 `json:"prepaid_cost"`
		CollectCost           float64 `json:"collect_cost"`
		ApCost                float64 `json:"ap_cost"`
		CurrencySw            bool    `json:"currency_sw"`
		PrepaidChargeCurrency float64 `json:"prepaid_charge_currency"`
		CollectChargeCurrency float64 `json:"collect_charge_currency"`
		ArChargeCurrency      float64 `json:"ar_charge_currency"`
		PrepaidCostCurrency   float64 `json:"prepaid_cost_currency"`
		CollectCostCurrency   float64 `json:"collect_cost_currency"`
		ApCostCurrency        float64 `json:"ap_cost_currency"`
		CurrencyCode          any     `json:"currency_code"`
		CurrencyRate          any     `json:"currency_rate"`
		CurrencyXDCode        any     `json:"currency_x_d_code"`
		ArChargeOtherCurrency float64 `json:"ar_charge_other_currency"`
		ApCostOtherCurrency   float64 `json:"ap_cost_other_currency"`
		Reference             any     `json:"reference"`
		ExtraDesc             any     `json:"extra_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		VendorId              int64   `json:"vendor_id"`
		CustomerId            int64   `json:"customer_id"`
		ArBilledSw            bool    `json:"ar_billed_sw"`
		ApBilledSw            any     `json:"ap_billed_sw"`
		AgentCalcSw           any     `json:"agent_calc_sw"`
		ExtraInvSw            bool    `json:"extra_inv_sw"`
		ExtraCostSw           any     `json:"extra_cost_sw"`
		DueAgentSw            any     `json:"due_agent_sw"`
		ArmInvoiceId          int64   `json:"arm_invoice_id"`
		ArmInvoiceDetailId    int64   `json:"arm_invoice_detail_id"`
		ApiInvoiceId          any     `json:"api_invoice_id"`
		ApiInvoiceDetailId    any     `json:"api_invoice_detail_id"`
		RateFactorTotal       any     `json:"rate_factor_total"`
		RateUnitCode          any     `json:"rate_unit_code"`
		RateAmt               float64 `json:"rate_amt"`
		CalculateAmt          float64 `json:"calculate_amt"`
		FrtDistributeCost     float64 `json:"frt_distribute_cost"`
		VendorName            any     `json:"vendor_name"`   // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CustomerName          any     `json:"customer_name"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ItemCharge            float64 `json:"item_charge"`
		ItemCost              float64 `json:"item_cost"`
		VatTaxSw              any     `json:"vat_tax_sw"`
		ApTaxPercent          any     `json:"ap_tax_percent"`
		ApTaxAmt              float64 `json:"ap_tax_amt"`
		ApWithholdTaxPercent  int     `json:"ap_withhold_tax_percent"`
		ApWithholdTaxAmt      float64 `json:"ap_withhold_tax_amt"`
		ArWithholdTaxPercent  int     `json:"ar_withhold_tax_percent"`
		ArWithholdTaxAmt      float64 `json:"ar_withhold_tax_amt"`
		ArTaxPercent          any     `json:"ar_tax_percent"`
		ArTaxAmt              float64 `json:"ar_tax_amt"`
		ApTaxCurrency         float64 `json:"ap_tax_currency"`
		ArTaxCurrency         float64 `json:"ar_tax_currency"`
		PeriodId              int     `json:"period_id"`
		AgentArId             any     `json:"agent_ar_id"`
		AgentApId             any     `json:"agent_ap_id"`
		AgentName             any     `json:"agent_name"`
		OldPlmChargeCostId    any     `json:"old_plm_charge_cost_id"`
		OldPlmBranchCode      any     `json:"old_plm_branch_code"`
		PlmDisplayFloatSeq    any     `json:"plm_display_float_seq"`
		CreatedDate           any     `json:"created_date"`     // 类型string，因为对应的Java语言String的默认值是null，以免报错
		CreatedByCode         any     `json:"created_by_code"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedDate          any     `json:"modified_date"`    // 类型string，因为对应的Java语言String的默认值是null，以免报错
		ModifiedByCode        any     `json:"modified_by_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Rowversion            []int   `json:"rowversion"`
		AdvanceItemSw         bool    `json:"advance_item_sw"`
		ChargeCostDesc        any     `json:"charge_cost_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		SelSw                 int     `json:"sel_sw"`
		X12ChargeCode         any     `json:"x12_charge_code"`
		ObjectState           int     `json:"ObjectState"`
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
		Returnobj interface{} `json:"returnobj"`
		Js        interface{} `json:"js"`
		Err       interface{} `json:"err"`
		MsgObj    struct {
			Message      string      `json:"message"`
			Reason       string      `json:"reason"`
			MsgTypeCode  string      `json:"msg_type_code"`
			MsgAskType   interface{} `json:"msg_ask_type"`
			MsgId        int         `json:"msg_id"`
			FocusField   interface{} `json:"focus_field"`
			ExtraMessage interface{} `json:"extra_message"`
			ExtraJs      interface{} `json:"extra_js"`
			Caller       interface{} `json:"caller"`
			Callf        interface{} `json:"callf"`
			CallAction   interface{} `json:"callAction"`
			CallObjStr   interface{} `json:"callObjStr"`
			CallUrl      interface{} `json:"callUrl"`
			CallSubform  interface{} `json:"callSubform"`
			CallButton   interface{} `json:"callButton"`
			CallAskId    interface{} `json:"callAskId"`
			CallBack     interface{} `json:"callBack"`
			WinName      interface{} `json:"win_name"`
		} `json:"msgObj"`
		Ask   string `json:"ask"`
		Pwait int    `json:"pwait"`
	} `json:"obj"`
	Mode   string `json:"mode"`
	Id     int64  `json:"id"`
	Mid    int64  `json:"mid"`
	Code   int    `json:"code"`
	MsgObj struct {
		Message      string      `json:"message"`
		Reason       string      `json:"reason"`
		MsgTypeCode  string      `json:"msg_type_code"`
		MsgAskType   interface{} `json:"msg_ask_type"`
		MsgId        int         `json:"msg_id"`
		FocusField   interface{} `json:"focus_field"`
		ExtraMessage interface{} `json:"extra_message"`
		ExtraJs      interface{} `json:"extra_js"`
		Caller       interface{} `json:"caller"`
		Callf        interface{} `json:"callf"`
		CallAction   interface{} `json:"callAction"`
		CallObjStr   interface{} `json:"callObjStr"`
		CallUrl      interface{} `json:"callUrl"`
		CallSubform  interface{} `json:"callSubform"`
		CallButton   interface{} `json:"callButton"`
		CallAskId    interface{} `json:"callAskId"`
		CallBack     interface{} `json:"callBack"`
		WinName      interface{} `json:"win_name"`
	} `json:"msgObj"`
}
