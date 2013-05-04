package field

type SeqNumValue struct{ IntValue }
type LengthValue struct{ IntValue }
type NumInGroupValue struct{ IntValue }
type DayOfMonthValue struct{ IntValue }

type CharValue struct{ StringValue }
type MultipleStringValue struct{ StringValue }
type MultipleCharValue struct{ StringValue }
type CurrencyValue struct{ StringValue }
type MonthYearValue struct{ StringValue }
type DataValue struct{ StringValue }
type LocalMktDateValue struct{ StringValue }
type ExchangeValue struct{ StringValue }
type LanguageValue struct{ StringValue }
type CountryValue struct{ StringValue }
type XMLDataValue struct{ StringValue }
type UTCTimeOnlyValue struct{ StringValue }
type UTCDateOnlyValue struct{ StringValue }
type TZTimeOnlyValue struct{ StringValue }
type TZTimestampValue struct{ StringValue }

type FloatValue struct{ StringValue }
type QtyValue struct{ FloatValue }
type PriceValue struct{ FloatValue }
type PriceOffsetValue struct{ FloatValue }
type AmtValue struct{ FloatValue }
type PercentageValue struct{ FloatValue }

