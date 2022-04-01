package main

import (
	"fmt"
    "strings"

	//onelog "github.com/francoispqt/onelog"
	"github.com/kubewarden/gjson"
	kubewarden "github.com/kubewarden/policy-sdk-go"
)

// isPalindrome: function to test if a input string is palindrome
// @input: text string: string to evaluate
// @output: true if the input string is palindrome, false otherwise
func isPalindrome(text string) bool {

    //if the lenght of the input string is 1, it is a character. Do not consider a single char palindrome
    len := len(text)
	if len == 1 {
		return false
	}

    //convert to lowercase
    textLower := strings.ToLower(text)

    //loop through the input string
	for idx:= 0; idx<len/2; idx++ {
		if textLower[idx] != textLower[len - 1 - idx] {
			return false
		}
	}
	return true
}

func validate(payload []byte) ([]byte, error) {

    var err error

    // step1: test the json format of the input
    if !gjson.ValidBytes(payload) {
        return kubewarden.RejectRequest(
            kubewarden.Message("Not a valid JSON document"),
            kubewarden.Code(400))
    }

    // step2: Check if the object is a Pod. If not, accect the object
    object := gjson.GetBytes(
        payload,
        "request.object.kind")
    
    if strings.Contains(strings.ToLower(object.String()), "pod") == false {
        return kubewarden.AcceptRequest()
    }

    // step3: get the labels from the input
    data := gjson.GetBytes(
        payload,
        "request.object.metadata.labels")

    // step3: loop through the labels
    data.ForEach(func(key, value gjson.Result) bool {
        label := key.String()

        // step4: if the label is palindrome, reject the object
        if isPalindrome(label) == true {
            err = fmt.Errorf("Label %s is palindrome", label)
            // stop iterating over labels
            return false
        }

        return true
    })

    // NOTE 7
    if err != nil {
        return kubewarden.RejectRequest(
            kubewarden.Message(err.Error()),
            kubewarden.NoCode)
    }

    return kubewarden.AcceptRequest()
}

