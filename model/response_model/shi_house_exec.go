package response_model

type OihHouseExec struct {
	Action int64 `json:"action"`
	Op     any   `json:"op"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Obj    struct {
		Returnobj any `json:"returnobj"`
		Js        any `json:"js"`
		Err       any `json:"err"`
		MsgObj    struct {
			Message      any   `json:"message"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
			Reason       any   `json:"reason"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
			MsgTypeCode  any   `json:"msg_type_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
			MsgAskType   any   `json:"msg_ask_type"`
			MsgId        int64 `json:"msg_id"`
			FocusField   any   `json:"focus_field"`
			ExtraMessage any   `json:"extra_message"`
			ExtraJs      any   `json:"extra_js"`
			Caller       any   `json:"caller"`
			Callf        any   `json:"callf"`
			CallAction   any   `json:"callAction"`
			CallObjStr   any   `json:"callObjStr"`
			CallUrl      any   `json:"callUrl"`
			CallSubform  any   `json:"callSubform"`
			CallButton   any   `json:"callButton"`
			CallAskId    any   `json:"callAskId"`
			CallBack     any   `json:"callBack"`
			WinName      any   `json:"win_name"`
		} `json:"msgObj"`
		Ask   any `json:"ask"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Pwait int `json:"pwait"`
	} `json:"obj"`
	Mode   any   `json:"mode"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
	Id     int64 `json:"id"`
	Mid    int64 `json:"mid"`
	Code   int64 `json:"code"`
	MsgObj struct {
		Message      any   `json:"message"`       // 类型string，因为对应的Java语言String的默认值是null，以免报错
		Reason       any   `json:"reason"`        // 类型string，因为对应的Java语言String的默认值是null，以免报错
		MsgTypeCode  any   `json:"msg_type_code"` // 类型string，因为对应的Java语言String的默认值是null，以免报错
		MsgAskType   any   `json:"msg_ask_type"`
		MsgId        int64 `json:"msg_id"`
		FocusField   any   `json:"focus_field"`
		ExtraMessage any   `json:"extra_message"`
		ExtraJs      any   `json:"extra_js"`
		Caller       any   `json:"caller"`
		Callf        any   `json:"callf"`
		CallAction   any   `json:"callAction"`
		CallObjStr   any   `json:"callObjStr"`
		CallUrl      any   `json:"callUrl"`
		CallSubform  any   `json:"callSubform"`
		CallButton   any   `json:"callButton"`
		CallAskId    any   `json:"callAskId"`
		CallBack     any   `json:"callBack"`
		WinName      any   `json:"win_name"`
	} `json:"msgObj"`
}
