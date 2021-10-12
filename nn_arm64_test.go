//go:build arm64
// +build arm64

package gorgonia

import "gorgonia.org/tensor"

var bnAllCases = []struct {
	desc string

	Dtype tensor.Dtype

	X      interface{}
	XShape tensor.Shape

	ScaleInit  InitWFn
	ScaleShape tensor.Shape

	BiasInit  InitWFn
	BiasShape tensor.Shape

	ExpectedTrainResult, ExpectedOutputGrad, ExpectedBiasGrad, ExpectedScaleGrad, ExpectedMean, ExpectedVariance interface{}
	ExpectedEvalResult                                                                                           interface{}
}{
	{
		desc:                "Float32 (3,2)",
		Dtype:               tensor.Float32,
		X:                   RangedFromWithStep(0.5, 0.01),
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float32{-0.16072161, -0.32144323, 0.2, 0.4, 0.5607227, 1.1214454},
		ExpectedMean:        []float32{0.4680, 0.4770},
		ExpectedVariance:    []float32{0.10036002, 0.10036002},
		ExpectedOutputGrad:  []float32{0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666},
		ExpectedBiasGrad:    []float32{0.5, 0.5},
		ExpectedScaleGrad:   []float32{0, -4.199363e-08},
		ExpectedEvalResult:  []float32{0.23030189, 0.46249762, 0.24924053, 0.5003749, 0.26817924, 0.5382523},
	},
	{
		desc:                "Float32 (3,2)",
		Dtype:               tensor.Float32,
		X:                   []float32{-0.1607, -0.3214, 0.2000, 0.4000, 0.5607, 1.1214},
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float32{-0.1674023, -0.33483633, 0.2, 0.4, 0.5674023, 1.1348363},
		ExpectedMean:        []float32{0.17999999, 0.35999998},
		ExpectedVariance:    []float32{0.21709406, 0.5683762},
		ExpectedOutputGrad:  []float32{0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666},
		ExpectedBiasGrad:    []float32{0.5, 0.5},
		ExpectedScaleGrad:   []float32{6.0190774e-07, 6.0190774e-07},
		ExpectedEvalResult:  []float32{-0.01936099, -0.14229004, 0.2128771, 0.4318339, 0.44511518, 1.0059578},
	},
	{
		desc:                "Float32 (3,4)",
		Dtype:               tensor.Float32,
		X:                   RangedFromWithStep(0.5, 0.01),
		XShape:              tensor.Shape{3, 4},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 4},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 4},
		ExpectedTrainResult: []float32{-0.16571283, -0.33142576, -0.49713972, -0.6628527, 0.20000063, 0.40000117, 0.5999998, 0.8000001, 0.5657136, 1.131427, 1.6971393, 2.262853},
		ExpectedMean:        []float32{0.486, 0.495, 0.50400007, 0.513},
		ExpectedVariance:    []float32{0.10144002, 0.10144002, 0.10144002, 0.10144002},
		ExpectedOutputGrad:  []float32{0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336, 0.083333336},
		ExpectedBiasGrad:    []float32{0.25, 0.25, 0.25, 0.25},
		ExpectedScaleGrad:   []float32{0, -2.1287331e-08, -2.1287331e-08, -2.1287331e-08},
		ExpectedEvalResult:  []float32{0.21318635, 0.42825645, 0.64521015, 0.8640479, 0.25086156, 0.5036068, 0.7582357, 1.0147486, 0.28853667, 0.5789571, 0.87126124, 1.1654494},
	},
	{
		desc:                "Float64 (3,2)",
		Dtype:               tensor.Float64,
		X:                   RangedFromWithStep(0.5, 0.01),
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float64{},
		ExpectedMean:        []float64{0.4680, 0.4770},
		ExpectedVariance:    []float64{0.10035999999999998, 0.10035999999999998},
		ExpectedOutputGrad:  []float64{0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666},
		ExpectedBiasGrad:    []float64{0.5, 0.5},
		ExpectedScaleGrad:   []float64{-2.433484712520345e-18, -2.433484712520345e-18},
		ExpectedEvalResult:  []float64{0.23030185885707807, 0.46249758389272355, 0.2492405206427519, 0.5003749074640712, 0.2681791824284257, 0.5382522310354189},
	},
	{
		desc:                "Float32 (3,2)",
		Dtype:               tensor.Float32,
		X:                   RangedFromWithStep(0.1, 0.001),
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           Ones(),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            Zeroes(),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float32{-0.5619533, -0.5619525, -9.449284e-07, -1.0085682e-07, 0.5619514, 0.5619523},
		ExpectedMean:        []float32{0.0918, 0.0927},
		ExpectedVariance:    []float32{0.10000362, 0.10000362},
		ExpectedOutputGrad:  []float32{0.16666667, 0.16666667, 0.16666667, 0.16666667, 0.16666667, 0.16666667},
		ExpectedBiasGrad:    []float32{0.5, 0.5},
		ExpectedScaleGrad:   []float32{1.0630707e-07, 9.8129604e-08},
		ExpectedEvalResult:  []float32{0.025928926, 0.02624516, 0.03225304, 0.032569274, 0.038577177, 0.038893387},
	},
	{
		desc:                "Float64 (2,2,2,2)",
		Dtype:               tensor.Float64,
		X:                   RangedFromWithStep(0.1, 2),
		XShape:              tensor.Shape{2, 2, 2, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2, 1, 1},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2, 1, 1},
		ExpectedTrainResult: []float64{-0.19727328639512015, -0.12504177977782557, -0.05281027316053104, 0.019421233456763526, -0.3945465727902401, -0.2500835595556509, -0.1056205463210618, 0.03884246691352733, 0.3805787665432364, 0.45281027316053096, 0.5250417797778255, 0.5972732863951201, 0.761157533086473, 0.9056205463210623, 1.0500835595556512, 1.1945465727902405},
		ExpectedMean:        []float64{9.990000000000002, 17.189999999999998},
		ExpectedVariance:    []float64{71.07142857142858, 71.07142857142858},
		ExpectedOutputGrad:  []float64{0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625},
		ExpectedBiasGrad:    []float64{0.5, 0.5},
		ExpectedScaleGrad:   []float64{-5.3462054499923185e-17, 1.6038616349976955e-16},
		ExpectedEvalResult:  []float64{-0.15194110496796842, -0.08077000183996672, -0.009598898711965044, 0.06157220441603665, -0.24694532743353526, -0.10460312117753189, 0.0377390850784715, 0.1800812913344749, 0.41742772005604517, 0.48859882318404685, 0.5597699263120486, 0.6309410294400503, 0.8917923226144919, 1.0341345288704953, 1.1764767351264986, 1.318818941382502},
	},
	{
		desc:                "Float32 (2,2,2,2)",
		Dtype:               tensor.Float64,
		X:                   RangedFromWithStep(0.1, 2),
		XShape:              tensor.Shape{2, 2, 2, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2, 1, 1},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2, 1, 1},
		ExpectedTrainResult: []float64{-0.19727328639512015, -0.12504177977782557, -0.05281027316053104, 0.019421233456763526, -0.3945465727902401, -0.2500835595556509, -0.1056205463210618, 0.03884246691352733, 0.3805787665432364, 0.45281027316053096, 0.5250417797778255, 0.5972732863951201, 0.761157533086473, 0.9056205463210623, 1.0500835595556512, 1.1945465727902405},
		ExpectedMean:        []float64{9.990000000000002, 17.189999999999998},
		ExpectedVariance:    []float64{71.07142857142858, 71.07142857142858},
		ExpectedOutputGrad:  []float64{0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625, 0.0625},
		ExpectedBiasGrad:    []float64{0.5, 0.5},
		ExpectedScaleGrad:   []float64{-5.3462054499923185e-17, 1.6038616349976955e-16},
		ExpectedEvalResult:  []float64{-0.15194110496796842, -0.08077000183996672, -0.009598898711965044, 0.06157220441603665, -0.24694532743353526, -0.10460312117753189, 0.0377390850784715, 0.1800812913344749, 0.41742772005604517, 0.48859882318404685, 0.5597699263120486, 0.6309410294400503, 0.8917923226144919, 1.0341345288704953, 1.1764767351264986, 1.318818941382502},
	},
	{
		desc:                "Float64 (2,3,2,2)",
		Dtype:               tensor.Float64,
		X:                   RangedFromWithStep(0.1, 2),
		XShape:              tensor.Shape{2, 3, 2, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 3, 1, 1},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 3, 1, 1},
		ExpectedTrainResult: []float64{-0.16865435186261707, -0.11950043828093479, -0.07034652469925252, -0.021192611117570252, -0.3373087037252339, -0.23900087656186939, -0.14069304939850485, -0.042385222235140324, -0.5059630555878506, -0.35850131484280384, -0.211039574097757, -0.0635778333527102, 0.4211926111175702, 0.47034652469925253, 0.5195004382809348, 0.568654351862617, 0.8423852222351406, 0.9406930493985051, 1.0390008765618697, 1.137308703725234, 1.263577833352711, 1.411039574097758, 1.5585013148428046, 1.7059630555878516},
		ExpectedMean:        []float64{13.590000000000002, 20.79, 27.99},
		ExpectedVariance:    []float64{153.35714285714286, 153.3571428571429, 153.35714285714286},
		ExpectedOutputGrad:  []float64{0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664},
		ExpectedBiasGrad:    []float64{0.3333333333333333, 0.3333333333333333, 0.3333333333333333},
		ExpectedScaleGrad:   []float64{0, 3.638120440587923e-17, 1.0914361321763771e-16},
		ExpectedEvalResult:  []float64{-0.12679935135317508, -0.07834874329488374, -0.029898135236592415, 0.01855247282169892, -0.21483821625971697, -0.11793700014313431, -0.021035784026551645, 0.07586543209003101, -0.2641165947196256, -0.11876477054475164, 0.026587053630122356, 0.17193887780499634, 0.454607945346321, 0.5030585534046124, 0.5515091614629036, 0.599959769521195, 0.9479763771392751, 1.0448775932558576, 1.1417788093724404, 1.238680025489023, 1.4801052953788623, 1.6254571195537364, 1.7708089437286103, 1.9161607679034842},
	},
	{
		desc:                "Float32 (2,3,2,2)",
		Dtype:               tensor.Float32,
		X:                   RangedFromWithStep(0.1, 2),
		XShape:              tensor.Shape{2, 3, 2, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 3, 1, 1},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 3, 1, 1},
		ExpectedTrainResult: []float32{-0.16865435186261707, -0.11950043828093479, -0.07034652469925252, -0.021192611117570252, -0.3373087037252339, -0.23900087656186939, -0.14069304939850485, -0.042385222235140324, -0.5059630555878506, -0.35850131484280384, -0.211039574097757, -0.0635778333527102, 0.4211926111175702, 0.47034652469925253, 0.5195004382809348, 0.568654351862617, 0.8423852222351406, 0.9406930493985051, 1.0390008765618697, 1.137308703725234, 1.263577833352711, 1.411039574097758, 1.5585013148428046, 1.7059630555878516},
		ExpectedMean:        []float32{13.590000000000002, 20.79, 27.99},
		ExpectedVariance:    []float32{153.35714285714286, 153.3571428571429, 153.35714285714286},
		ExpectedOutputGrad:  []float32{0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664, 0.041666666666666664},
		ExpectedBiasGrad:    []float32{0.3333333333333333, 0.3333333333333333, 0.3333333333333333},
		ExpectedScaleGrad:   []float32{1.4547279e-08, 3.5706957e-08, -7.009144e-08},
		ExpectedEvalResult:  []float32{-0.12679935135317508, -0.07834874329488374, -0.029898135236592415, 0.01855247282169892, -0.21483821625971697, -0.11793700014313431, -0.021035784026551645, 0.07586543209003101, -0.2641165947196256, -0.11876477054475164, 0.026587053630122356, 0.17193887780499634, 0.454607945346321, 0.5030585534046124, 0.5515091614629036, 0.599959769521195, 0.9479763771392751, 1.0448775932558576, 1.1417788093724404, 1.238680025489023, 1.4801052953788623, 1.6254571195537364, 1.7708089437286103, 1.9161607679034842},
	},
}

var bnstackedCases = []struct {
	desc string

	Epochs int

	Dtype tensor.Dtype

	XInit  InitWFn
	XShape tensor.Shape

	ScaleInit  InitWFn
	ScaleShape tensor.Shape

	BiasInit  InitWFn
	BiasShape tensor.Shape

	ExpectedTrainResult, ExpectedOutputGrad, ExpectedBiasGrad, ExpectedScaleGrad, ExpectedMean, ExpectedVariance interface{}
	ExpectedEvalResult                                                                                           interface{}
}{
	{
		desc:                "Example (1d Float32)",
		Dtype:               tensor.Float32,
		Epochs:              1,
		XInit:               RangedFromWithStep(float32(0.5), float32(0.01)),
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           RangedFromWithStep(float32(0.3), float32(0.3)),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            RangedFromWithStep(float32(0.2), float32(0.2)),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float32{-0.16740213, -0.33483604, 0.19999963, 0.39999926, 0.5674025, 1.1348367},
		ExpectedMean:        []float32{0.18000033, 0.36000067},
		ExpectedVariance:    []float32{0.21710846, 0.56843376},
		ExpectedOutputGrad:  []float32{0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666},
		ExpectedBiasGrad:    []float32{0.5, 0.5},
		ExpectedScaleGrad:   []float32{1.6863456e-08, -8.432093e-09},
		ExpectedEvalResult:  []float32{0.18238541, 0.43004364, 0.19457874, 0.46041483, 0.20677204, 0.490786},
	},
	{
		desc:                "Example (1d Float64)",
		Dtype:               tensor.Float64,
		Epochs:              1,
		XInit:               RangedFromWithStep(0.5, 0.01),
		XShape:              tensor.Shape{3, 2},
		ScaleInit:           RangedFromWithStep(0.3, 0.3),
		ScaleShape:          tensor.Shape{1, 2},
		BiasInit:            RangedFromWithStep(0.2, 0.2),
		BiasShape:           tensor.Shape{1, 2},
		ExpectedTrainResult: []float64{-0.1674022853682372, -0.33483633412378944, 0.19999999999999998, 0.4, 0.5674022853682372, 1.1348363341237895},
		ExpectedMean:        []float64{0.17999999999999985, 0.36000000000000015},
		ExpectedVariance:    []float64{0.21710843373493974, 0.568433734939759},
		ExpectedOutputGrad:  []float64{0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666, 0.16666666666666666},
		ExpectedBiasGrad:    []float64{0.5, 0.5},
		ExpectedScaleGrad:   []float64{1.6863456e-08, -8.432093e-09},
		ExpectedEvalResult:  []float64{0.18238597332767545, 0.4315682494085479, 0.19457929958738146, 0.461711271524584, 0.20677262584708744, 0.4918542936406202},
	},
}
