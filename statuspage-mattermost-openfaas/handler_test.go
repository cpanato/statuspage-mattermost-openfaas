package function

import "testing"

func TestHandleReturnsCorrectResponse(t *testing.T) {
	// TODO: Improve the test, add a webserver to receive the webhook
	resp := Handle([]byte(`{
    "meta": {
        "unsubscribe": "http://status.xoxoxoxo.com/?unsubscribe=ssss",
        "documentation": "https://help.statuspage.io/knowledge_base/topics/webhook-notifications",
        "generated_at": "2019-02-19T16:19:51.098Z"
    },
    "page": {
        "id": "6w4r0ttlx5ft",
        "status_indicator": "none",
        "status_description": "All Systems Operational"
    },
    "component": {
        "created_at": "2019-01-21T21:45:43.473Z",
        "id": "rms1xkdxxqx1",
        "name": "XOXO Notifications",
        "page_id": "6w4r0ttlx5ft",
        "position": 36,
        "status": "operational",
        "updated_at": "2019-02-19T16:19:50.579Z",
        "showcase": false
    },
    "component_update": {
        "old_status": "degraded_performance",
        "new_status": "operational",
        "created_at": "2019-02-19T16:19:50.521Z",
        "component_type": "Component",
        "state": "sn_created",
        "id": "jyvmqcnby11y",
        "component_id": "rms1xkdxxqx1"
    }
}`))

	if resp != true {
		t.Fatalf("Expected: %v, Got: %v", true, resp)
	}
}
