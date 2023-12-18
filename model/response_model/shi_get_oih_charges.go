package response_model

type ShiGetOihCharges struct {
	Type    any                      `json:"type"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Rows    []ShiGetOihChargesRows   `json:"rows"`
	Footer  []ShiGetOihChargesFooter `json:"footer"`
	Total   int64                    `json:"total"`
	Success bool                     `json:"success"`
	Err     bool                     `json:"err"`
}

type ShiGetOihChargesRows struct {
	TableKeyName          any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter              any     `json:"IsFooter"`
	Selector              any     `json:"selector"`
	KeyId                 int64   `json:"KeyId"`
	MasterId              any     `json:"MasterId"`
	Caller                any     `json:"caller"`
	PlmChargeCostId       int64   `json:"plm_charge_cost_id"`
	EfcCompanyId          int64   `json:"efc_company_id"`
	EfcBranchId           int64   `json:"efc_branch_id"`
	PlmChargeCostSeqNo    int64   `json:"plm_charge_cost_seq_no"`
	BranchCode            any     `json:"branch_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	PliChargeCostItemId   int64   `json:"pli_charge_cost_item_id"`
	PliChargeCostItemNo   int64   `json:"pli_charge_cost_item_no"`
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
	CurrencyCode          any     `json:"currency_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	CurrencyRate          int64   `json:"currency_rate"`
	CurrencyXDCode        any     `json:"currency_x_d_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ArChargeOtherCurrency float64 `json:"ar_charge_other_currency"`
	ApCostOtherCurrency   float64 `json:"ap_cost_other_currency"`
	Reference             any     `json:"reference"`  // 类型string，因为对应的Java语言String的默认值是null，以免报错
	ExtraDesc             any     `json:"extra_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	VendorId              int64   `json:"vendor_id"`
	CustomerId            int64   `json:"customer_id"`
	ArBilledSw            bool    `json:"ar_billed_sw"`
	ApBilledSw            bool    `json:"ap_billed_sw"`
	AgentCalcSw           any     `json:"agent_calc_sw"`
	ExtraInvSw            bool    `json:"extra_inv_sw"`
	ExtraCostSw           any     `json:"extra_cost_sw"`
	DueAgentSw            any     `json:"due_agent_sw"`
	ArmInvoiceId          int64   `json:"arm_invoice_id"`
	ArmInvoiceDetailId    int64   `json:"arm_invoice_detail_id"`
	ApiInvoiceId          int64   `json:"api_invoice_id"`
	ApiInvoiceDetailId    int64   `json:"api_invoice_detail_id"`
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
	ApWithholdTaxPercent  int64   `json:"ap_withhold_tax_percent"`
	ApWithholdTaxAmt      float64 `json:"ap_withhold_tax_amt"`
	ArWithholdTaxPercent  int64   `json:"ar_withhold_tax_percent"`
	ArWithholdTaxAmt      float64 `json:"ar_withhold_tax_amt"`
	ArTaxPercent          any     `json:"ar_tax_percent"`
	ArTaxAmt              float64 `json:"ar_tax_amt"`
	ApTaxCurrency         float64 `json:"ap_tax_currency"`
	ArTaxCurrency         float64 `json:"ar_tax_currency"`
	PeriodId              int64   `json:"period_id"`
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
	Rowversion            []int64 `json:"rowversion"`
	AdvanceItemSw         bool    `json:"advance_item_sw"`
	ChargeCostDesc        any     `json:"charge_cost_desc"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	SelSw                 int64   `json:"sel_sw"`
	X12ChargeCode         any     `json:"x12_charge_code"`
	ObjectState           int64   `json:"ObjectState"`
}

type ShiGetOihChargesFooter struct {
	TableKeyName          any     `json:"TableKeyName"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	IsFooter              bool    `json:"IsFooter"`
	Selector              any     `json:"selector"`
	KeyId                 any     `json:"KeyId"`
	MasterId              any     `json:"MasterId"`
	Caller                any     `json:"caller"`
	PlmChargeCostId       int64   `json:"plm_charge_cost_id"`
	EfcCompanyId          int64   `json:"efc_company_id"`
	EfcBranchId           int64   `json:"efc_branch_id"`
	PlmChargeCostSeqNo    int64   `json:"plm_charge_cost_seq_no"`
	BranchCode            any     `json:"branch_code"`
	PliChargeCostItemId   any     `json:"pli_charge_cost_item_id"`
	PliChargeCostItemNo   any     `json:"pli_charge_cost_item_no"`
	OpType                any     `json:"op_type"`
	ChargeCostType        any     `json:"charge_cost_type"`
	PlmCodeIndicator      any     `json:"plm_code_indicator"`
	PlmCatCode            any     `json:"plm_cat_code"`
	PlmMasterId           any     `json:"plm_master_id"`
	PlmHblId              any     `json:"plm_hbl_id"`
	PlmHblDisplaySequence any     `json:"plm_hbl_display_sequence"`
	DocCodeMaster         any     `json:"doc_code_master"`
	DocCodeHouse          any     `json:"doc_code_house"`
	PrepaidCharge         any     `json:"prepaid_charge"`
	CollectCharge         float64 `json:"collect_charge"`
	ArCharge              float64 `json:"ar_charge"`
	PrepaidCost           float64 `json:"prepaid_cost"`
	CollectCost           any     `json:"collect_cost"`
	ApCost                float64 `json:"ap_cost"`
	CurrencySw            any     `json:"currency_sw"`
	PrepaidChargeCurrency any     `json:"prepaid_charge_currency"`
	CollectChargeCurrency any     `json:"collect_charge_currency"`
	ArChargeCurrency      any     `json:"ar_charge_currency"`
	PrepaidCostCurrency   any     `json:"prepaid_cost_currency"`
	CollectCostCurrency   any     `json:"collect_cost_currency"`
	ApCostCurrency        any     `json:"ap_cost_currency"`
	CurrencyCode          any     `json:"currency_code"`
	CurrencyRate          any     `json:"currency_rate"`
	CurrencyXDCode        any     `json:"currency_x_d_code"`
	ArChargeOtherCurrency any     `json:"ar_charge_other_currency"`
	ApCostOtherCurrency   any     `json:"ap_cost_other_currency"`
	Reference             any     `json:"reference"`
	ExtraDesc             any     `json:"extra_desc"`
	VendorId              any     `json:"vendor_id"`
	CustomerId            any     `json:"customer_id"`
	ArBilledSw            any     `json:"ar_billed_sw"`
	ApBilledSw            any     `json:"ap_billed_sw"`
	AgentCalcSw           any     `json:"agent_calc_sw"`
	ExtraInvSw            any     `json:"extra_inv_sw"`
	ExtraCostSw           any     `json:"extra_cost_sw"`
	DueAgentSw            any     `json:"due_agent_sw"`
	ArmInvoiceId          any     `json:"arm_invoice_id"`
	ArmInvoiceDetailId    any     `json:"arm_invoice_detail_id"`
	ApiInvoiceId          any     `json:"api_invoice_id"`
	ApiInvoiceDetailId    any     `json:"api_invoice_detail_id"`
	RateFactorTotal       any     `json:"rate_factor_total"`
	RateUnitCode          any     `json:"rate_unit_code"`
	RateAmt               any     `json:"rate_amt"`
	CalculateAmt          any     `json:"calculate_amt"`
	FrtDistributeCost     any     `json:"frt_distribute_cost"`
	VendorName            any     `json:"vendor_name"`
	CustomerName          any     `json:"customer_name"`
	ItemCharge            any     `json:"item_charge"`
	ItemCost              any     `json:"item_cost"`
	VatTaxSw              any     `json:"vat_tax_sw"`
	ApTaxPercent          any     `json:"ap_tax_percent"`
	ApTaxAmt              any     `json:"ap_tax_amt"`
	ApWithholdTaxPercent  any     `json:"ap_withhold_tax_percent"`
	ApWithholdTaxAmt      any     `json:"ap_withhold_tax_amt"`
	ArWithholdTaxPercent  any     `json:"ar_withhold_tax_percent"`
	ArWithholdTaxAmt      any     `json:"ar_withhold_tax_amt"`
	ArTaxPercent          any     `json:"ar_tax_percent"`
	ArTaxAmt              any     `json:"ar_tax_amt"`
	ApTaxCurrency         any     `json:"ap_tax_currency"`
	ArTaxCurrency         any     `json:"ar_tax_currency"`
	PeriodId              any     `json:"period_id"`
	AgentArId             any     `json:"agent_ar_id"`
	AgentApId             any     `json:"agent_ap_id"`
	AgentName             any     `json:"agent_name"`
	OldPlmChargeCostId    any     `json:"old_plm_charge_cost_id"`
	OldPlmBranchCode      any     `json:"old_plm_branch_code"`
	PlmDisplayFloatSeq    any     `json:"plm_display_float_seq"`
	CreatedDate           any     `json:"created_date"`
	CreatedByCode         any     `json:"created_by_code"`
	ModifiedDate          any     `json:"modified_date"`
	ModifiedByCode        any     `json:"modified_by_code"`
	Rowversion            any     `json:"rowversion"`
	AdvanceItemSw         any     `json:"advance_item_sw"`
	ChargeCostDesc        any     `json:"charge_cost_desc"`
	SelSw                 int64   `json:"sel_sw"`
	X12ChargeCode         any     `json:"x12_charge_code"`
	ObjectState           int64   `json:"ObjectState"`
}
