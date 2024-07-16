package handlers

import (
	// "net/http"

	// "github.com/gin-contrib/sessions"
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
	var isError = false
	var fromProtected bool = false
const (
	auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
	username_key      string = "username"
	tzone_key         string = "time_zone"
)

const (
	session_name              string = "fmessages"
	session_flashmessages_key string = "flashmessages-key"
)

func getCookieStore() cookie.Store {
	store := cookie.NewStore([]byte(session_flashmessages_key))
	return store
	// return sessions.New([]byte(session_flashmessages_key))
}

// Set adds a new message to the cookie store
func setFlashmessages(c *gin.Context, kind, value string) {
	
	session, _ := getCookieStore().Get(c.Request, session_name)

	session.AddFlash(value, kind)

	session.Save(c.Request, c.Writer)
}

// Get receives flash messages from cookie store
func getFlashmessages(c *gin.Context, kind string) []string {
	session, _ := getCookieStore().Get(c.Request, session_name)

	fm := session.Flashes(kind)

	// if there are some messages…
	if len(fm) > 0 {
		session.Save(c.Request, c.Writer)

		// we start an empty strings slice that we
		// then return with messages
		var flashes []string
		for _, fl := range fm {
			// we add the messages to the slice
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}

	return nil
}

func renderView(c *gin.Context, cmp templ.Component) error {
	c.Header("Content-Type", gin.MIMEHTML)
	// c.Writer.Header().Set(, echo.MIMETextHTML)

	return cmp.Render(c.Request.Context(), c.Writer)
}