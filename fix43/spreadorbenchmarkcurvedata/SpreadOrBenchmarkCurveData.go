package spreadorbenchmarkcurvedata

//Component is a fix43 SpreadOrBenchmarkCurveData Component
type Component struct {
	//Spread is a non-required field for SpreadOrBenchmarkCurveData.
	Spread *float64 `fix:"218"`
	//BenchmarkCurveCurrency is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveCurrency *string `fix:"220"`
	//BenchmarkCurveName is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurveName *string `fix:"221"`
	//BenchmarkCurvePoint is a non-required field for SpreadOrBenchmarkCurveData.
	BenchmarkCurvePoint *string `fix:"222"`
}

func New() *Component { return new(Component) }
