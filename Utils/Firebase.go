package Utils

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

var fbapp *firebase.App

func initFirebase() (*firebase.App, error) {
	if fbapp != nil {
		return fbapp, nil
	}
	opt := option.WithCredentialsFile("./burgher-adminsdk-firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	fbapp = app
	return app, nil
}

func VerifyToken(idToken string, email string) (*map[string]interface{}, error) {
	app, err := initFirebase()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client, err3 := app.Auth(ctx)
	if err3 != nil {
		return nil, err3
	}
	tok, err2 := client.VerifyIDToken(ctx, idToken)
	if err2 != nil {
		return nil, err2
	}
	if (tok.Claims)["email"] != email {
		return nil, fmt.Errorf("email mismatch")
	}
	return &tok.Claims, nil
}

// func init() {

// 	// fmt.Println(os.Environ())
// 	resp, err := VerifyToken("eyJhbGciOiJSUzI1NiIsImtpZCI6ImMxNTQwYWM3MWJiOTJhYTA2OTNjODI3MTkwYWNhYmU1YjA1NWNiZWMiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiU2hvYmhpdCBNYWhlc2h3YXJpIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0k1XzNHMkFxNDlRd1kxMDg3QkFPbkF6S2ZhRmhnTjNleXhmYWxCNE9pNHFXWjE4VzJ0VlE9czk2LWMiLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vYnVyZ2hlci1mNDViOCIsImF1ZCI6ImJ1cmdoZXItZjQ1YjgiLCJhdXRoX3RpbWUiOjE3MjE1MTkzOTcsInVzZXJfaWQiOiI5T0Jnajh5a1FQV0FwVWY4Y2dicjE4Rk1nem4yIiwic3ViIjoiOU9CZ2o4eWtRUFdBcFVmOGNnYnIxOEZNZ3puMiIsImlhdCI6MTcyMTUxOTM5OCwiZXhwIjoxNzIxNTIyOTk4LCJlbWFpbCI6InNob2JoaXRtYWhlc2h3YXJpMThAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMTU5MDQ4NjQwNzQ0OTczMzA0ODMiXSwiZW1haWwiOlsic2hvYmhpdG1haGVzaHdhcmkxOEBnbWFpbC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJnb29nbGUuY29tIn19.md_mbNfnMh1SybZjBuHbSJEK9mOJ0oczyVSXsT-IOWA7ZfNRTgOZpH9nl1_1RzUcrYEU6Azj9fSuJxma-uINxCaznJ1Qqdp_ZFaH-dGwZbaRCyjsoe5TlpERSoRCKG7HfDLxDNHjv-NQVZK6dB_N2XzGTHF5sHTPcsXW5h1MrfSkcaonAdB6C90RadBWKtvs5sMJcBmH63lBji9fM_XGAobI9qlvs3-fmqna4vfWrygtSVdto-o3u8H45ct2jMLiFmUnV36PM1VTjq6gr_d0kXvQMKs82vEccMU32CVrqyX-WCHzqyOrfrAQ2AVgGp5jk6dkn2qnNCQllYhZwn2CKw")
// 	fmt.Println(resp)
// 	fmt.Println(err)
// }