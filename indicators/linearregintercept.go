package indicators

import (
	"github.com/thetruetrade/gotrade"
)

type LinearRegIntercept struct {
	*LinearRegWithoutStorage

	// public variables
	Data []float64
}

func NewLinearRegIntercept(timePeriod int, selectData gotrade.DataSelectionFunc) (indicator *LinearRegIntercept, err error) {
	newInd := LinearRegIntercept{}
	newInd.LinearRegWithoutStorage, err = NewLinearRegWithoutStorage(timePeriod, selectData,
		func(dataItem float64, slope float64, intercept float64, streamBarIndex int) {
			result := intercept

			if result > newInd.LinearRegWithoutStorage.maxValue {
				newInd.LinearRegWithoutStorage.maxValue = result
			}

			if result < newInd.LinearRegWithoutStorage.minValue {
				newInd.LinearRegWithoutStorage.minValue = result
			}

			newInd.Data = append(newInd.Data, result)
		})

	return &newInd, err
}

func NewLinearRegInterceptForStream(priceStream *gotrade.DOHLCVStream, timePeriod int, selectData gotrade.DataSelectionFunc) (indicator *LinearRegIntercept, err error) {
	newInd, err := NewLinearRegIntercept(timePeriod, selectData)
	priceStream.AddTickSubscription(newInd)
	return newInd, err
}