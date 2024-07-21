package User

import (
	Utils "burgher/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var user UserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, accessToken, refreshToken, notcreated :=
		create(
			User{
				Id:       Utils.GenerateId(),
				UserName: user.UserName,
				Name:     user.Name,
				Tag:      user.Tag,
				Email:    user.Email,
				ImageUrl: nil,
			},
			user.FirebaseAuthIdToken,
		)
	if notcreated != nil {
		w.WriteHeader(503)
		fmt.Println(notcreated)
		w.Write([]byte(notcreated.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.
		NewEncoder(w).
		Encode(
			UserResponse{
				Id:           userresp.Id,
				Name:         userresp.Name,
				UserName:     userresp.UserName,
				Tag:          userresp.Tag,
				IsVerified:   userresp.IsVerified,
				Email:        userresp.Email,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		Username string `json:"username"`
		Tag      int    `json:"tag"`
		UserId   string `json:"userId"`
	}
	var readRequest ReadRequest
	err := json.NewDecoder(r.Body).Decode(&readRequest)
	fmt.Println(readRequest)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, notfound := read(readRequest.Username, readRequest.Tag, readRequest.UserId)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.
		NewEncoder(w).
		Encode(
			UserResponse{
				Id:         userresp.Id,
				Name:       userresp.Name,
				UserName:   userresp.UserName,
				Email:      userresp.Email,
				Tag:        userresp.Tag,
				IsVerified: userresp.IsVerified,
				ImageUrl:   userresp.ImageUrl,
			})
}
func readWithEmailHandler(w http.ResponseWriter, r *http.Request) {
	// reqDump, err2 := httputil.DumpRequest(r, true)
	// fmt.Println(string(reqDump))
	// fmt.Println(err2)

	type ReadRequest struct {
		Email               string `json:"email"`
		FirebaseAuthIdToken string `json:"firebaseAuthIdToken"`
	}
	var readRequest ReadRequest
	// io.Copy(os.Stdout, r.Body)
	err := json.NewDecoder(r.Body).Decode(&readRequest)
	fmt.Println(readRequest)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, accessToken, refreshToken, _ := readWithEmail(readRequest.Email, readRequest.FirebaseAuthIdToken)

	// if notfound != nil {
	// 	w.WriteHeader(503)
	// 	fmt.Println(notfound)
	// 	w.Write([]byte(notfound.Error()))
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.
		NewEncoder(w).
		Encode(
			UserResponse{
				Id:           userresp.Id,
				Name:         userresp.Name,
				UserName:     userresp.UserName,
				Tag:          userresp.Tag,
				IsVerified:   userresp.IsVerified,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				ImageUrl:     userresp.ImageUrl,
			},
		)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(Utils.ContextUserKey)
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	// for k, v := range r.Form {
	// if you are certain v[0] is present
	// fmt.Println(k)
	// fmt.Println(v)

	// }
	if err != nil {
		// If there is an error that means form is empty. Return nil for err in order
		// to validate result as required.
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Println("wiritng file")
	updateProfilePicture(userId.(string), &file)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.
		NewEncoder(w).
		Encode(map[string]string{"status": "success"})
}
