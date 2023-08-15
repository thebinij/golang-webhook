package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func handleWebhookEvent(w http.ResponseWriter, r *http.Request) {

	data, err := decodeJSONData(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	fmt.Printf("Webhook event data: %+v\n", data)

	if cmd, ok := data["command"].(string); ok {
		runScript(cmd)
	} else {
		fmt.Println("No script found in the webhook data.")
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Webhook event received"})
}

func runScript(cmd string) {
	fmt.Printf("Running script: %s\n", cmd)

	// Run the shell command using the "sh" shell
	command := exec.Command("sh", "-c", cmd)

	// Set the standard output and error output to the current process's
	// standard output and error
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Start the command and wait for it to finish
	err := command.Run()
	if err != nil {
		fmt.Printf("Error running script: %s\n", err)
	}
}
