package main

import (
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/recipe/userroles/userrolesmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

var smtpUsername = "..."
var smtpSettings = emaildelivery.SMTPSettings{
	Host: // "smtp.mail.ovh.net",
	From: emaildelivery.SMTPFrom{
		Name:  "",
		Email: "",
	},
	Port:     465,
	Username: &smtpUsername, // this is optional. In case not given, from.email will be used
	Password: """",
	Secure:   true,
	// this is optional. TLS config is used if Secure is set to true, or server supports STARTTLS
	// if not provided, the SDK will use a default config
	// TLSConfig: &tls.Config{
	// ...
	//},
}

var SuperTokensConfig = supertokens.TypeInput{
	Supertokens: &supertokens.ConnectionInfo{
		ConnectionURI: "http://localhost:3567",
	},
	AppInfo: supertokens.AppInfo{
		AppName:       "SuperTokens Demo App",
		APIDomain:     "http://localhost:3000",
		WebsiteDomain: "http://localhost:3000",
	},
	RecipeList: []supertokens.Recipe{
		emailpassword.Init(&epmodels.TypeInput{
			EmailDelivery: &emaildelivery.TypeInput{
				Service: emailpassword.MakeSMTPService(emaildelivery.SMTPServiceConfig{
					Settings: smtpSettings,
				}),
			},
		}),
		emailverification.Init(evmodels.TypeInput{
			Mode: evmodels.ModeRequired, // or evmodels.ModeOptional
			EmailDelivery: &emaildelivery.TypeInput{
				Service: emailverification.MakeSMTPService(emaildelivery.SMTPServiceConfig{
					Settings: smtpSettings,
				}),
			},
		}),
		userroles.Init(&userrolesmodels.TypeInput{
			SkipAddingRolesToAccessToken:       true,
			SkipAddingPermissionsToAccessToken: true,
		}),

		session.Init(nil),
		dashboard.Init(nil),
	},
}
