package dtos

import "ml-test-quasar/utils"

type Response struct {
	Message  string      `json:"message"`
	Position utils.Point `json:"position"`
}
