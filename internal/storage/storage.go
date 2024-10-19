package storage

import "errors"

var (
	ErrURLNotExists   = errors.New("url doesnt exists")
	ErrUrlsNotSaved   = errors.New("failed to save url to database")
	ErrFailedToGetUrl = errors.New("failed to get url from database")
)
