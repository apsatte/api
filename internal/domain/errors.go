package domain

import "errors"

// common
var (
	ErrDatabase     = errors.New("DATABASE_ERROR")
	ErrNotFound     = errors.New("NOT_FOUND")
	ErrUnauthorized = errors.New("UNAUTHORIZED")
)

// modules
var (
	ErrModuleOptionAlreadyExists = errors.New("MODULE_OPTION_ALREADY_EXISTS")
	ErrOptionParamAlreadyExists  = errors.New("MODULE_OPTION_PARAM_ALREADY_EXISTS")
)

// customer
var (
	ErrIncorrectPassword = errors.New("INCORRECT_PASSWORD")
	ErrNameTooLong       = errors.New("NAME_TOO_LONG")
)

// project
var (
	ErrProjectModuleAlreadyExists = errors.New("PROJECT_MODULE_ALREADY_EXISTS")
)

// dish
var (
	ErrDishOptionItemAlreadyExists = errors.New("DISH_OPTION_ITEM_ALREADY_EXISTS")
	ErrDishOptionAlreadyExists     = errors.New("DISH_OPTION_ALREADY_EXISTS")
)
