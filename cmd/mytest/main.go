package main

import (
	"fmt"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40/newordersingle"
	"github.com/quickfixgo/quickfix/tag"
)

type NoOrders struct {
	quickfix.Group
}

func (m *NoOrders) SetClOrdID(f field.ClOrdIDField) {
	m.Set(f)
}

func (m NoOrders) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

type MyRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

func NewMyRepeatingGroup() MyRepeatingGroup {
	return MyRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.Side,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Price)})}
}

type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{quickfix.NewRepeatingGroup(tag.NoOrders,
		quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), NewMyRepeatingGroup()})}
}

func (m *NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}

func main() {
	fmt.Println("hello world")

	msg := quickfix.NewMessage()

	group := NewNoOrdersRepeatingGroup()
	order1 := group.Add()
	order1.SetClOrdID(field.ClOrdIDField{FIXString: "1"})

	order2 := group.Add()
	order2.SetClOrdID(field.ClOrdIDField{FIXString: "2"})

	msg.Header.Set(field.NewBeginString(enum.BeginStringFIX42))
	msg.Header.Set(field.NewMsgType(enum.MsgType_NEW_ORDER_CROSS))
	msg.Body.SetGroup(group)

	b, _ := msg.Build()
	fmt.Println(string(b))

	m, err := quickfix.ParseMessage(b)
	if err != nil {
		panic(err)
	}

	g := NewNoOrdersRepeatingGroup()
	m.Body.GetGroup(g)

	fmt.Println(g.Len())

	clordid1, _ := g.Get(0).GetClOrdID()
	clordid2, _ := g.Get(1).GetClOrdID()

	fmt.Println(clordid1)
	fmt.Println(clordid2)

	m, err = quickfix.ParseMessage([]byte("8=FIX.4.29=10835=D34=249=TW52=20160421-14:43:5056=ISLD11=ID21=440=154=138=002000.0055=INTC60=20160421-14:43:5010=004"))
	//	m, err = quickfix.ParseMessage([]byte("8=FIX.4.29=10835=D34=249=TW52=20160421-14:43:5056=ISLD11=ID21=438=002000.0040=154=255=INTC60=20160421-14:43:5010=005"))
	//	m, err = quickfix.ParseMessage([]byte("8=FIX.4.29=11835=D34=249=TW52=20160421-14:43:5056=ISLD11=ID21=438=002000.0040=154=255=INTC58=oh hai60=20160421-14:43:5010=218"))

	if err != nil {
		panic(err)
	}

	nos := newordersingle.FromMessage(m)
	if side, err := nos.GetSide(); err != nil {
		panic(err)
	} else {
		fmt.Printf("Side is %v\n", side)
	}
	nos.SetSide(enum.Side_SELL)
	nos.Set(field.NewText("oh hai"))

	m = nos.ToMessage()
	b, err = m.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	newNos := newordersingle.New(
		field.NewClOrdID("1234"),
		field.NewHandlInst("1"),
		field.NewSymbol("TSLA"),
		field.NewSide(enum.Side_BUY),
		field.NewOrderQty(100.0),
		field.NewOrdType(enum.OrdType_MARKET),
	)

	//newNos.Set(field.NewTransactTime(time.Now()))
	newNos.Set(field.NewTransactTimeNoMillis(time.Now()))

	m = newNos.ToMessage()
	b, err = m.Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
