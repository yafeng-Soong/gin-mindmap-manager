package router

import (
	"net/http"

	"github.com/yafeng-Soong/gin-mindmap-manager/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func (t *TestRouter) InitTestRouter(r *gin.Engine) {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/sayHello", func(c *gin.Context) {
		if currentUser, err := utils.GetCurrentUser(c); err != nil {
			c.String(http.StatusForbidden, err.Error())
		} else {
			c.String(http.StatusOK, "Hello "+currentUser.Username)
		}
	})

	r.GET("/currentUser", func(c *gin.Context) {
		session := sessions.Default(c)
		currentUser := session.Get("currentUser")
		c.JSON(http.StatusOK, gin.H{"currentUser": currentUser})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
}
