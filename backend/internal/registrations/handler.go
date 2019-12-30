package registrations

import (
	"encoding/json"
	"net/http"
	"strings"
	"vineguard/internal/cache"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	cacheSvc *cache.Svc
}

func NewHander() *Handler {
	cacheSvc := cache.NewService()
	return &Handler{cacheSvc: cacheSvc}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	bearer := strings.Split(auth, "Bearer ")
	if len(bearer) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		msg := map[string]string{"Message": "Bearer must be present in header"}
		msgJson, _ := json.Marshal(msg)
		_, _ = w.Write(msgJson)
		return
	}
	token := bearer[1]
	user := h.getUser(token)

	switch r.Method {
	case http.MethodGet:
		regs, _ := h.cacheSvc.GetUserRegistrations(user)
		regsJson, err := json.Marshal(regs)
		if err != nil {
			logrus.Errorf("Problem formatting. %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(regsJson)
		if err != nil {
			logrus.Errorf("Problem writing message. %s", err)
		}
	case http.MethodPost:
		newReg := map[string]string{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newReg)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			msg := map[string]string{"Message": "Could not format POST body"}
			msgJson, _ := json.Marshal(msg)
			_, _ = w.Write(msgJson)
			return
		}
		_ = h.cacheSvc.CreateNewRegistration(user, newReg)
	default:
		logrus.Errorf("%s. request type not allowed.", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getUser(googleIdToken string) string {
	// hit to verify google id token https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=
	return "nicholas.b.masson@gmail.com"
}

// /registrations POST
// /registrations?userEmail= GET
// /data?userEmail=&deviceId= GET
