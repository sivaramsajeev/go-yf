package yf

type Stock struct {
	Meta struct {
		Currency             string
		Symbol               string
		ExchangeName         string
		InstrumentType       string
		FirstTradeDate       int64
		GmtOffset            int64 `json:"gmtoffset"`
		Timezone             string
		RegularMarketPrice   float64
		FiftyTwoWeekHigh     float64
		FiftyTwoWeekLow      float64
		RegularMarketDayHigh float64
		RegularMarketDayLow  float64
		RegularMarketVolume  float64
		LongName             string
		ShortName            string
		ExchangeTimezoneName string
		CurrentTradingPeriod struct {
			Pre     TradingPeriod
			Regular TradingPeriod
			Post    TradingPeriod
		}
		DataGranularity string
		ValidRanges     []string
	}
	Timestamp  []int64
	Indicators struct {
		Quote []struct {
			Low    []float64
			Volume []float64
			High   []float64
			Open   []float64
			Close  []float64
		}
		Unadjclose []struct {
			Unadjclose []float64
		}
		Adjclose []struct {
			Adjclose []float64
		}
	}
}

type TradingPeriod struct {
	Timezone  string
	End       int64
	Start     int64
	GmtOffset int64 `json:"gmtoffset"`
}
