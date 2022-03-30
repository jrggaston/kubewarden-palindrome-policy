package main

import (
	"encoding/json"
	"testing"
	kubewarden_testing "github.com/kubewarden/policy-sdk-go/testing"
)


func TestIngressObject(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/ingress.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestIngressObjectWithPalindromeLabel(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/ingressPalindromeLabel.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestPodWithNoPalindromeLabels(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/pod.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestPodWithNoLabels(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podNoLabels.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestPodWithNoPalindromeLabelWithPrefix(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podNoPalindromeLabelWithPrefix.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestPodWithOneCharLabel(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podOneCharLabel.json",
        &settings)
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
        t.Error("Unexpected reject")
    }
}
func TestPodWithPalindromeLabel(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podPalindromeLabel.json",
        &settings)
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
        t.Error("Unexpected accept")
    }

    expected_message := "Label level is palindrome"
    if response.Message != expected_message {
        t.Errorf("Got '%s' instead of '%s'", response.Message, expected_message)
    }
}
func TestPodPalindromeLabelCapitalLetters(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podPalindromeLabelCapitalLetters.json",
        &settings)
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
        t.Error("Unexpected accept")
    }

    expected_message := "Label levEL is palindrome"
    if response.Message != expected_message {
        t.Errorf("Got '%s' instead of '%s'", response.Message, expected_message)
    }
}
func TestPodPalindromeLabelEven(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podPalindromeLabelEven.json",
        &settings)
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
        t.Error("Unexpected accept")
    }

    expected_message := "Label abccba is palindrome"
    if response.Message != expected_message {
        t.Errorf("Got '%s' instead of '%s'", response.Message, expected_message)
    }
}
func TestPodPalindromeLabelWithPrefix(t *testing.T) {
    settings := Settings{}

    payload, err := kubewarden_testing.BuildValidationRequest(
        "test_data/podPalindromeLabelWithPrefix.json",
        &settings)
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
        t.Error("Unexpected accept")
    }

    expected_message := "Label abc/cba is palindrome"
    if response.Message != expected_message {
        t.Errorf("Got '%s' instead of '%s'", response.Message, expected_message)
    }
}

