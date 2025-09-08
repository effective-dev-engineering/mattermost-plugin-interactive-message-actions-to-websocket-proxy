package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type ApiInteractiveMessageActionResponse struct {
	Update *struct {
		Message *string `json:"message"`
	} `json:"update"`
	Props            *any    `json:"props"`
	EphemeralText    *string `json:"ephemeral_text"`
	SkipSlackParsing *bool   `json:"skip_slack_parsing"`
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
// The root URL is currently <siteUrl>/plugins/com.mattermost.plugin-starter-template/api/v1/. Replace com.mattermost.plugin-starter-template with the plugin ID.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()

	// Middleware to require that the user is logged in
	router.Use(p.MattermostAuthorizationRequired)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.HandleFunc("/handler", p.Handler).Methods(http.MethodPost)

	router.ServeHTTP(w, r)
}

func (p *Plugin) MattermostAuthorizationRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("Mattermost-User-ID")
		if userID == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func respondErr(w http.ResponseWriter, code int, err error) (int, error) {
	http.Error(w, err.Error(), code)
	return code, err
}

func (p *Plugin) Handler(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := io.ReadAll(r.Body)
	if bodyReadError != nil {
		p.API.LogError("Error when reading body: ", bodyReadError.Error())
		respondErr(w, http.StatusBadRequest, bodyReadError)
		return
	}

	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		p.API.LogError("Failed to unmarshal request body: ", err.Error())
		respondErr(w, http.StatusBadRequest, err)
		return
	}

	context := data["context"]

	var userId string
	if ctxMap, ok := context.(map[string]any); ok {
		userId = fmt.Sprintf("%v", ctxMap["user_id"])
	} else {
		respondErr(w, http.StatusBadRequest, errors.New("failed to parse context"))
		return
	}

	if len(userId) > 0 {
		p.API.PublishWebSocketEvent(WsEventAction, data, &model.WebsocketBroadcast{
			UserId: userId,
		})
	}
}
