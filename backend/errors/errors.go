package errors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function is designed to handle internal server errors within
// the context of a Gin web server at CONTROLLERS LEVEL and send
// an appropriate response to the client.

// Add the following line at the very beginning of the main function to make sure the
// SendInternalError function is activated properly when running on release mode.

// gin.SetMode(gin.ReleaseMode)

// if the above line is not called in main.go the App will always run in debug mode
// and detailed errors will be shown to the user.

func SendInternalError(ctx *gin.Context, err error) {
	if ctx == nil {
		// Handle the case where the context is nil
		fmt.Printf("Internal Server Error: %v\n", err)
		return
	}
	// The gin.Context object represents the HTTP request and response
	// context, while the error object contains information about the
	// error that occurred.
	if gin.Mode() == "release" {
		// If the mode is set to "release," it implies that the application is running
		// in production mode, where detailed err messages should not be exposed to clients.
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Something went wrong"})
		// In that case it sends a generic error response with HTTP status code 500
		// (Internal Server Error) and a JSON object containing a generic error message.
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		// If not, the application is running in a development or debug mode therefore
		// a JSON object containing the specific error message is sent back.
	}
}
