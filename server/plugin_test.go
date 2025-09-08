package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin/plugintest"
	"github.com/stretchr/testify/assert"
)

func TestActionHandler(t *testing.T) {
	api := &plugintest.API{}
	plugin := Plugin{}

	api.On("PublishWebSocketEvent", WsEventAction, map[string]any(nil), &model.WebsocketBroadcast{UserId: "1"}).Return(nil).Once()

	assert := assert.New(t)
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{
		"user_id": "test_user_id",
		"user_name": "test_user_name",
		"channel_id": "test_channel_id",
		"channel_name": "test_channel_name",
		"post_id": "test_post_id",
		"trigger_id": "test_trigger_id",
		"type": "button",
		"team_id": "",
		"team_domain": "",
		"data_source": "",
		"context": {
			"action": "test_action",
			"user_id": "1",
		},
	}`)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/handler", body)
	r.Header.Set("Mattermost-User-ID", "test-user-id")
	r.Header.Set("Content-Type", "application/json")

	plugin.ServeHTTP(nil, w, r)

	result := w.Result()
	assert.NotNil(result)
	defer result.Body.Close()

	api.AssertCalled(t, "PublishWebSocketEvent", WsEventAction, map[string]any(nil), &model.WebsocketBroadcast{UserId: "1"})
	api.AssertNumberOfCalls(t, "PublishWebSocketEvent", 1)

}
