package controllers

import (
"reflect"
_"fmt"
"net/http"
"golang.org/x/oauth2"
"os"
"golang.org/x/oauth2/google"
_"golang.org/x/net/context"
_"time"

"github.com/revel/revel"

"log"
  )

type App struct {
	*revel.Controller
}
var (
    googleOauthConfig = &oauth2.Config{
        RedirectURL:    "http://localhost:3000/GoogleCallback",
        ClientID:     os.Getenv("googlekey"), // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
        ClientSecret: os.Getenv("googlesecret"), // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
        Scopes:       []string{},
        Endpoint:     google.Endpoint,
    }
// Some random string, random for each request
    oauthStateString = "random"
)
func (c App) Index() revel.Result {
	//
	//
	// if userIns["token"] == nil {
	// 	// gothic.BeginAuthHandler(res, req)
	// 	 userIns["token"]  = "gothic resp"
	// }

	return c.Render()
}

func (c App) Login() revel.Result {
	var w http.ResponseWriter
  var r *http.Request
  w= c.Controller.Response.In.GetRaw().(*http.Response)
	r= c.Controller.Request.In.GetRaw().(*http.Request)
	log.Printf("Reflection from gothic  : %v",reflect.TypeOf(r) )
	log.Println()
	log.Println()
	log.Println()
	log.Printf("Header response %v " ,w.Header)
	log.Println("WHAT IS Header response %v " ,w)
	log.Println("WHAT IS Header response %v " ,w)
	log.Println("WHAT IS Header response %v " ,w)
	log.Printf("WHAT IS Header response %v " ,w)
	log.Println()
	log.Println()
	r= c.Controller.Request.In.GetRaw().(*http.Request)
  //log.Println("%v" ,r.)
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

return c.Render()
}

func (c App) LoginSub() revel.Result {
	var w http.ResponseWriter
	var r *http.Request
	r= c.Controller.Request.In.GetRaw().(*http.Request)

	log.Printf("Reflection from gothic  : %v",reflect.TypeOf(r) )
	log.Printf("Loginsub  : %v", r )
	log.Printf("Loginsub  : %v", w )
    state := r.FormValue("state")
		log.Printf("Loginsub  : %v", state )
return c.Render()
    // if state != oauthStateString {
    //     fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
    //     http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//
    //     return
    // }
		//
    // code := r.FormValue("code")
    // token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
    // if err != nil {
    //     fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
    //     http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
    //     return
    // }
		//
    // client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
}
