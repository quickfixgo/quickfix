package spreadorbenchmarkcurvedata

//Component is a fix50sp2 SpreadOrBenchmarkCurveData Component
type Component struct {
	//Spread is a non-required field for SpreadOrBenchmarkCurveData.
	Spread *float64 `fix:"218"`
	//BenchmarkCurveCurrency is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveCurrency *string `fix:"220"`
	//BenchmarkCurveName is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveName *string `fix:"221"`
	//BenchmarkCurvePoint is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurvePoint *string `fix:"222"`
	//BenchmarkPrice is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkPrice *float64 `fix:"662"`
	//BenchmarkPriceType is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkPriceType *int `fix:"663"`
	//BenchmarkSecurityID is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkSecurityID *string `fix:"699"`
	//BenchmarkSecurityIDSource is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkSecurityIDSource *string `fix:"761"`
}

func New() *Component { return new(Component) }
