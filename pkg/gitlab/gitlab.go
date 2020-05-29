package gitlab

import "net/http"

// HookHandler handles webhooks received from GitLab
func HookHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
