# Mattermost plugin for forwarding message interactions to WebSocket proxy events
(for bots directly connected to sockets)

## How to use

1. Install this plugin

2. Send a post with the following props (use [this article](https://developers.mattermost.com/integrate/plugins/interactive-messages/) to customize your attachment and action):
```json
props: {
    "attachments": [
        {
            "text": "Text example",
            "color": "#3AA3E3",
            "actions": [
                {
                    // you can use `button` or `select` values
                    "type": "button",
                    // replace `name` to your button name
                    "name": "Button text",
                    "integration": {
                        // static plugin url, don't replace it
                        "url": "/plugins/mattermost-plugin-interactive-message-actions-to-websocket-proxy/api/v1/handler",
                        "context": {
                            // replace `action` to your action name
                            "action": "your_action_name",
                            // replace `user_id` to your bot id
                            "user_id": "your_bot_user_id"

                            // also you can add any new fields to `context` to use it in your websocket event handler
                        }
                    }
                }
            ]
        }
    ]
}
```

1. Press your button` (or choice some option in your select if you use `select` type) under your post in Mattermost to send a PublishWebSocketEvent

2. Handle your websocket event:
```json
{
    // static event name
    "event": "custom_mattermost-plugin-interactive-message-actions-to-websocket-proxy_action",
    "data": {
        // your action type
        "type": "button",
        // user id of the user who pressed the button
        "user_id": "initiator_user_id",
        // user name of the user who pressed the button
        "user_name": "initiator_user_name", 
        // context with you sent data
        "context": {
            "action": "your_action_name",
            "user_id": "your_bot_user_id"
        },
        "post_id": "post_id_with_your_actions",
        "channel_id": "channel_id_with_post_with_your_actions",
        "channel_name": "channel_name_with_post_with_your_actions",
        // you can use `trigger_id` to open an interactive dialog for user:
        // https://developers.mattermost.com/integrate/plugins/interactive-dialogs/
        "trigger_id": "your_trigger_id" 

        // other proxied fields
    },
}
```
