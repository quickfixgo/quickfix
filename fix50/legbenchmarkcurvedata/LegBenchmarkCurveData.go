package legbenchmarkcurvedata

//Component is a fix50 LegBenchmarkCurveData Component
type Component struct {
	//LegBenchmarkCurveCurrency is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurveCurrency *string `fix:"676"`
	//LegBenchmarkCurveName is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurveName *string `fix:"677"`
	//LegBenchmarkCurvePoint is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkCurvePoint *string `fix:"678"`
	//LegBenchmarkPrice is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkPrice *float64 `fix:"679"`
	//LegBenchmarkPriceType is a non-required field for LegBenchmarkCurveData.
	LegBenchmarkPriceType *int `fix:"680"`
}

func New() *Component { return new(Component) }
