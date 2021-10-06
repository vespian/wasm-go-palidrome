package main

import (
	"encoding/json"
	"testing"

	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
)

func TestNoLabelsLeadsToApproval(t *testing.T) {

	payload, err := kubewarden_testing.BuildValidationRequest(
		"test_data/pod-no-labels.json",
		Settings{},
	)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_testing.ValidationResponse
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != true {
		t.Error("Unexpected rejection")
	}
}

func TestApproval(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequest(
		"test_data/pod-with-labels-no-palidrome.json",
		Settings{},
	)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_testing.ValidationResponse
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != true {
		t.Error("Unexpected rejection")
	}
}

func TestRejection(t *testing.T) {
	payload, err := kubewarden_testing.BuildValidationRequest(
		"test_data/pod-with-labels-palidrome.json",
		Settings{},
	)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	responsePayload, err := validate(payload)
	if err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	var response kubewarden_testing.ValidationResponse
	if err := json.Unmarshal(responsePayload, &response); err != nil {
		t.Errorf("Unexpected error: %+v", err)
	}

	if response.Accepted != false {
		t.Error("Unexpected approval")
	}
}
