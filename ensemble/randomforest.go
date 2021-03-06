package ensemble

import (
	"errors"
	"fmt"
	"github.com/kevinchapron/golearn/base"
	"github.com/kevinchapron/golearn/meta"
	"github.com/kevinchapron/golearn/trees"
)

// RandomForest classifies instances using an ensemble
// of bagged random decision trees.
type RandomForest struct {
	base.BaseClassifier
	ForestSize int
	Features   int
	Model      *meta.BaggedModel
}

// NewRandomForest generates and return a new random forests
// forestSize controls the number of trees that get built
// features controls the number of features used to build each tree.
func NewRandomForest(forestSize int, features int) *RandomForest {
	ret := &RandomForest{
		base.BaseClassifier{},
		forestSize,
		features,
		nil,
	}
	return ret
}

// Fit builds the RandomForest on the specified instances
func (f *RandomForest) Fit(on base.FixedDataGrid) error {
	numNonClassAttributes := len(base.NonClassAttributes(on))
	if numNonClassAttributes < f.Features {
		return errors.New(fmt.Sprintf(
			"Random forest with %d features cannot fit data grid with %d non-class attributes",
			f.Features,
			numNonClassAttributes,
		))
	}

	f.Model = new(meta.BaggedModel)
	f.Model.RandomFeatures = f.Features
	for i := 0; i < f.ForestSize; i++ {
		tree := trees.NewID3DecisionTree(0.00)
		f.Model.AddModel(tree)
	}
	f.Model.Fit(on)
	return nil
}

// Predict generates predictions from a trained RandomForest.
func (f *RandomForest) Predict(with base.FixedDataGrid) (base.FixedDataGrid, error) {
	return f.Model.Predict(with), nil
}
// Predict Ratio generates predictions from a trained RandomForest.
func (f *RandomForest) PredictRatio(with base.FixedDataGrid) (map[int](map[string]int), error) {
	return f.Model.PredictRatio(with), nil
}
// Getting Max-voted prediction from ratio and dataset
func (f *RandomForest) GenerateMaxRatio(with base.FixedDataGrid, votes map[int](map[string]int)) (base.FixedDataGrid, error) {
	return f.Model.GenerateMaxRatio(with,votes), nil
}

// String returns a human-readable representation of this tree.
func (f *RandomForest) String() string {
	return fmt.Sprintf("RandomForest(ForestSize: %d, Features:%d, %s\n)", f.ForestSize, f.Features, f.Model)
}
