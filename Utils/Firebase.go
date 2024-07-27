package Utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

var fbapp *firebase.App

func initFirebase() (*firebase.App, error) {
	if fbapp != nil {
		return fbapp, nil
	}

	var firebaseconfig map[string]interface{}
	jsonFile, err := os.ReadFile("./burgher-adminsdk-template.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened sdkfile.json")

	if err := json.Unmarshal(jsonFile, &firebaseconfig); err != nil {
		log.Fatal(err)
	}
	var re = regexp.MustCompile(`/\\n/g`)
	pvtKey, _ := base64.StdEncoding.DecodeString(os.Getenv("FIREBASE_PVT_KEY_B64"))
	pvtKeyStr := string(pvtKey)
	firebasepvtkey := re.ReplaceAllString(pvtKeyStr, `\n`)
	firebasepvtkey = strings.Replace(firebasepvtkey, `\n`, "\n", -1)
	firebaseconfig["private_key_id"] = os.Getenv("FIREBASE_PVT_KEY_ID")
	firebaseconfig["private_key"] = firebasepvtkey
	firebaseconfig["client_email"] = os.Getenv("FIREBASE_CLIENT_EMAIL")
	firebaseconfig["client_id"] = os.Getenv("FIREBASE_CLIENT_ID")
	firebaseconfig["client_x509_cert_url"] = os.Getenv("FIREBASE_CLIENT_X509CERT")
	// jsonFile[]
	// fmt.Println(firebaseconfig)
	// fmt.Println(string(jsonFile))
	bytefirebaseconfig, _ := json.Marshal(firebaseconfig)
	_ = os.WriteFile("./burgher-adminsdk-firebase.json", bytefirebaseconfig, 0644)

	opt := option.WithCredentialsFile("./burgher-adminsdk-firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	fbapp = app
	return app, nil
}

//	func init() {
//		initFirebase()
//	}
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
