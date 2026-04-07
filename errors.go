package er

import "errors"

var (
	// ErrInvalidDatasetName is returned when a dataset name contains
	// path separators or is empty.
	ErrInvalidDatasetName = errors.New("er: invalid dataset name")

	// ErrDatasetNameMustNotBeEmpty is returned when saving with empty name.
	ErrDatasetNameMustNotBeEmpty = errors.New("er: dataset name must not be empty")
)
