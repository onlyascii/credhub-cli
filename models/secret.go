package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Secret struct {
	Name       string
	SecretBody SecretBody
}

func NewSecret(name string, secretBody SecretBody) Item {
	return Secret{
		Name:       name,
		SecretBody: secretBody,
	}
}

func (secret Secret) String() string {
	lines := []string{}

	secretBody := secret.SecretBody
	lines = append(lines,
		fmt.Sprintf("Type:		%s", secretBody.ContentType),
		fmt.Sprintf("Name:		%s", secret.Name),
	)

	if secretBody.ContentType == "value" {
		value := secretBody.Credential.(string)
		lines = append(lines, fmt.Sprintf("Credential:	%s", value))
	} else {
		// We are marshaling again here because there isn't a simple way
		// to convert map[string]interface{} to a Certificate struct
		json_cert, _ := json.Marshal(secretBody.Credential)
		cert := Certificate{}
		json.Unmarshal(json_cert, &cert)
		if cert.Ca != "" {
			lines = append(lines, fmt.Sprintf("CA:		%s", cert.Ca))
		}

		if cert.Certificate != "" {
			lines = append(lines, fmt.Sprintf("Certificate:		%s", cert.Certificate))
		}

		if cert.Private != "" {
			lines = append(lines, fmt.Sprintf("Private:	%s", cert.Private))
		}
	}

	lines = append(lines, fmt.Sprintf("Updated:	%s", secretBody.UpdatedAt))

	return strings.Join(lines, "\n")
}
