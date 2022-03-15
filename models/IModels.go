package models

import (
	"github.com/go-gota/gota/dataframe"
)

type Model interface {
	fit(dataframe.DataFrame)
	predict(dataframe.DataFrame)
}
