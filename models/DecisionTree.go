package models

import "github.com/go-gota/gota/dataframe"

type DecisionTree struct {
	temp int
}

func (dt *DecisionTree) fit(dataframe.DataFrame) error {
	return nil
}

/*
func (dt *DecisionTree) predict(dataframe.DataFrame) dataframe.DataFrame,error{
	return dataframe.DataFrame(),nil
}
*/
