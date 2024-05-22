package helpers

import (
	"net/http"
	"time"

	"github.com/hiboedi/zenshop/app/models"
)

var userSession = "user-session"

func SetUserCookie(w http.ResponseWriter, r *http.Request, user models.UserLoginResponse) {
	cookie := http.Cookie{
		Name:     userSession,
		Value:    user.ID,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(userSession)
	if err != nil {
		if err == http.ErrNoCookie {
			w.Write([]byte("no cookie found"))
			return nil, err
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return cookie, nil
}

func DeleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     userSession,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Write([]byte("cookie has been deleted"))
}
