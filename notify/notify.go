package notify

import "svcheck/types"

// SendNotification : send a notification using one of the available notification options
func SendNotification(req types.Request) {
	if req.IsAlive == false {
		EmitPush(req)
	}
}
