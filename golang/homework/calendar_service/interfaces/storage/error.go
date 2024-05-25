package storage

import "errors"

var ErrEventNotFound = errors.New("event not found")

var ErrEventAlreadyExists = errors.New("event already exists")

var ErrStorageOperation = errors.New("storage operation failed")

var ErrDateBusy = errors.New("That date already used")
