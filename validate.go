package main

import (
	"fmt"

	onelog "github.com/francoispqt/onelog"
	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
)

type Settings struct {
}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func validate(payload []byte) ([]byte, error) {
	name := gjson.GetBytes(payload, "request.object.metadata.name").String()
	namespace := gjson.GetBytes(payload, "request.object.metadata.namespace").String()

	logger.DebugWithFields("validating pod object", func(e onelog.Entry) {
		e.String("name", name)
		e.String("namespace", namespace)
	})

	data := gjson.GetBytes(
		payload,
		"request.object.metadata.labels")

	if !data.Exists() {
		logger.Warn("cannot read labels from metadata: accepting request")
		return kubewarden.AcceptRequest()
	}

	labels := data.Map()
	for k := range labels {
		if isPalindrome(k) {
			logger.InfoWithFields("rejecting pod object", func(e onelog.Entry) {
				e.String("name", name)
				e.String("namespace", namespace)
				e.String("label", k)
			})
			return kubewarden.RejectRequest(
				kubewarden.Message(fmt.Sprintf("The '%s' label is a palidrome", k)),
				kubewarden.NoCode,
			)
		}
	}

	return kubewarden.AcceptRequest()
}
