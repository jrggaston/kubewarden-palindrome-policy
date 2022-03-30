package main
import (
	kubewarden "github.com/kubewarden/policy-sdk-go"
)


type Settings struct {
    
}

func validateSettings(payload []byte) ([]byte, error) {
    
    return kubewarden.AcceptSettings()
}