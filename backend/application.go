package main

type applicationSubmitRequest struct {
	FullName                      string `json:"fullName"`
	Email                         string `json:"email"`
	PhoneCountryDial              string `json:"phoneCountryDial"`
	PhoneNumber                   string `json:"phoneNumber"`
	CountryOfResidence            string `json:"countryOfResidence"`
	City                          string `json:"city"`
	Age                           string `json:"age"`
	OrganizationName              string `json:"organizationName"`
	ParticipatedBefore            string `json:"participatedBefore"`
	PreferredParticipationType    string `json:"preferredParticipationType"`
	Motivation                    string `json:"motivation"`
	AmbassadorCode                string `json:"ambassadorCode"`
	AgreedToTermsAndPrivacyPolicy bool   `json:"agreedToTermsAndPrivacyPolicy"`
}

type applicationSubmitResponse struct {
	ApplicationID string `json:"applicationId"`
	PaymentID     string `json:"paymentId"`
	RedirectURL   string `json:"redirectUrl"`
	Status        string `json:"status"`
}
