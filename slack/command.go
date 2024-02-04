package slack

import (
	"net/http"
	"strings"
)

const (
	CommandPlan   = "plan"
	CommandDeploy = "deploy"
)

func HandleCommand(w http.ResponseWriter, r *http.Request) {
	userInput := r.FormValue("text")

	command, params, isFound := strings.Cut(userInput, " ")
	if !isFound {

	}

	switch command {
	case CommandPlan:
		//TODO: call github
	case CommandDeploy:
		//TODO: call github
	default:
		//TODO: return 400
	}
}
