package dtos

import "ml-mutant-test/utils"

type Response struct {
	Message  string      `json:"message"`
	Position utils.Point `json:"position"`
}
