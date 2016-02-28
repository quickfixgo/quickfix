package triggeringinstruction

//Component is a fix50sp1 TriggeringInstruction Component
type Component struct {
	//TriggerType is a non-required field for TriggeringInstruction.
	TriggerType *string `fix:"1100"`
	//TriggerAction is a non-required field for TriggeringInstruction.
	TriggerAction *string `fix:"1101"`
	//TriggerPrice is a non-required field for TriggeringInstruction.
	TriggerPrice *float64 `fix:"1102"`
	//TriggerSymbol is a non-required field for TriggeringInstruction.
	TriggerSymbol *string `fix:"1103"`
	//TriggerSecurityID is a non-required field for TriggeringInstruction.
	TriggerSecurityID *string `fix:"1104"`
	//TriggerSecurityIDSource is a non-required field for TriggeringInstruction.
	TriggerSecurityIDSource *string `fix:"1105"`
	//TriggerSecurityDesc is a non-required field for TriggeringInstruction.
	TriggerSecurityDesc *string `fix:"1106"`
	//TriggerPriceType is a non-required field for TriggeringInstruction.
	TriggerPriceType *string `fix:"1107"`
	//TriggerPriceTypeScope is a non-required field for TriggeringInstruction.
	TriggerPriceTypeScope *string `fix:"1108"`
	//TriggerPriceDirection is a non-required field for TriggeringInstruction.
	TriggerPriceDirection *string `fix:"1109"`
	//TriggerNewPrice is a non-required field for TriggeringInstruction.
	TriggerNewPrice *float64 `fix:"1110"`
	//TriggerOrderType is a non-required field for TriggeringInstruction.
	TriggerOrderType *string `fix:"1111"`
	//TriggerNewQty is a non-required field for TriggeringInstruction.
	TriggerNewQty *float64 `fix:"1112"`
	//TriggerTradingSessionID is a non-required field for TriggeringInstruction.
	TriggerTradingSessionID *string `fix:"1113"`
	//TriggerTradingSessionSubID is a non-required field for TriggeringInstruction.
	TriggerTradingSessionSubID *string `fix:"1114"`
}

func New() *Component { return new(Component) }

func (m *Component) SetTriggerType(v string)                { m.TriggerType = &v }
func (m *Component) SetTriggerAction(v string)              { m.TriggerAction = &v }
func (m *Component) SetTriggerPrice(v float64)              { m.TriggerPrice = &v }
func (m *Component) SetTriggerSymbol(v string)              { m.TriggerSymbol = &v }
func (m *Component) SetTriggerSecurityID(v string)          { m.TriggerSecurityID = &v }
func (m *Component) SetTriggerSecurityIDSource(v string)    { m.TriggerSecurityIDSource = &v }
func (m *Component) SetTriggerSecurityDesc(v string)        { m.TriggerSecurityDesc = &v }
func (m *Component) SetTriggerPriceType(v string)           { m.TriggerPriceType = &v }
func (m *Component) SetTriggerPriceTypeScope(v string)      { m.TriggerPriceTypeScope = &v }
func (m *Component) SetTriggerPriceDirection(v string)      { m.TriggerPriceDirection = &v }
func (m *Component) SetTriggerNewPrice(v float64)           { m.TriggerNewPrice = &v }
func (m *Component) SetTriggerOrderType(v string)           { m.TriggerOrderType = &v }
func (m *Component) SetTriggerNewQty(v float64)             { m.TriggerNewQty = &v }
func (m *Component) SetTriggerTradingSessionID(v string)    { m.TriggerTradingSessionID = &v }
func (m *Component) SetTriggerTradingSessionSubID(v string) { m.TriggerTradingSessionSubID = &v }
