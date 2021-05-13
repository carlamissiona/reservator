package controllers

import (
	_ "encoding/json"
	_"fmt"
	"net/http"
	_ "net/url"
	_ "strconv" 
	"reflect"
    
	"github.com/revel/revel"

    "log"
 
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/gorilla/sessions"
)

type App struct {
	*revel.Controller
}

var userIns map[string]interface{} = map[string]interface{}{"token":nil}
 

func (c App) Index() revel.Result {
 	

	  key := "Secret-session-key"  // Replace with your SESSION_SECRET or similar
	  maxAge := 86400 * 30  // 30 days
	  isProd := false       // Set to true when serving over https

	  store := sessions.NewCookieStore([]byte(key))
	  store.MaxAge(maxAge)
	  store.Options.Path = "/"
	  store.Options.HttpOnly = true   // HttpOnly should always be enabled
	  store.Options.Secure = isProd

	  gothic.Store = store

	  goth.UseProviders(
	    google.New("135639776025-699385rnp98elmrsb7a2btvhn4e820qv.apps.googleusercontent.com", "Dc2HxB5LMURwpSe3UKEL99Jp", "http://localhost:9000/logincallback", "email", "profile"),
	  )


	 
	 
	if userIns["token"] == nil { 
		// gothic.BeginAuthHandler(res, req)
		 userIns["token"]  = "gothic resp"
	}
	 
	 
   

	return c.Render()
}

func (c App) Login() revel.Result {
	  key := "Secret-session-key"  // Replace with your SESSION_SECRET or similar
	  maxAge := 86400 * 30  // 30 days
	  isProd := false       // Set to true when serving over https

	  store := sessions.NewCookieStore([]byte(key))
	  store.MaxAge(maxAge)
	  store.Options.Path = "/"
	  store.Options.HttpOnly = true   // HttpOnly should always be enabled
	  store.Options.Secure = isProd

	  gothic.Store = store

	  goth.UseProviders(
	    google.New("135639776025-699385rnp98elmrsb7a2btvhn4e820qv.apps.googleusercontent.com", "Dc2HxB5LMURwpSe3UKEL99Jp", "http://localhost:9000/logincallback", "email", "profile"),
	  )



	//log.Printf("REQ Get this This is the obj of revel: %s", (string)*c.Controller.Request.In)
	//log.Printf("REQ Get this This is the obj of GewtRaw     : %v", c.Controller.Request.In.GetRaw().(*http.Request)   )
	//log.Printf("REQ Get this This is the obj of GewtRaw     : %v", c.Request.In.GetRaw()  )
	log.Printf(" ==============store ======================")
		log.Printf( "the session key %v",store )
	log.Printf(" ==============store ======================")
 
	log.Printf("REQ Get this This is the obj of GewtRaw     : %v", c.Request.In.GetRaw()  )
	log.Printf("REQ Get this This is the obj of GewtRaw     : %v", c.Request.In.GetRaw()  )
	//log.Printf("REQ Get this This is the obj of ContentType     : %v", c.Request.ContentType     )
	//log.Printf("REQ Get this This is the obj of Method               : %v", *c.Controller.Request         )
	//log.Printf("REQ Get this This is the obj of revel: %v", *c.Controller.Request.In)
	//log.Printf("REQ Get this This is the obj of revel: %v", *c.Controller.Request.In)
	//log.Printf("Get this This is the obj of revel: %v", *c.Controller.Request)
	// log.Printf("Get this This is the obj of revel: %v", *c.Controller.Request)
	// log.Printf("This is the controller obj: %s", *c.Controller.Response)
	// log.Printf("About to call gothic : %s", c )
	// //log.Print( *gothic.SessionName ) 	  
   var cres http.ResponseWriter  
   var srq *http.Request  
   //var creq http.Request = nil
	//gothic.BeginAuthHandler( cres,creq)


	//revelReq,k := *c.Request.In.GetRaw()
	// if !ok {
	//   // handle this somehow
	//   log.Printf("About to call req Error  : %s", revelReq )
	//   log.Printf("About to call req Error  : %s", revelReq )
	 //  log.Printf("About to call req Error  : %s", ok )
	  
	// }
	// //revelReq = (http.Request) revelReq
	// //r := revelReq.Original
	  log.Printf("About to call req original  : %v", cres )
	// log.Printf("About to call req original  : %v", r )
	// log.Printf("About to call req original  : %v", r )
	// log.Printf("About to call req original  : %v", r )
	userIns["token"]  = "gothic resp"
   srq =  c.Controller.Request.In.GetRaw().(*http.Request)

	log.Printf("Response from gothic  : %v", srq )
	log.Printf("Reflection from gothic  : %v",reflect.TypeOf(srq) )
    gothic.BeginAuthHandler( cres ,  srq )
	//log.Printf("Response from gothic  : %v", cres ) 
	 
	// facebook
	 

	//userIns['token']  = # process if no access token 
	// resp, _ := http.Get("https://graph.facebook.com/me?fields=id,name,birthday,email&access_token=" +
	// 		url.QueryEscape(u.AccessToken))
    //return c.Render(userIns, authUrl)

	return c.Render()
}

 


// var FACEBOOK = &oauth2.Config{
// 	ClientID:     "943076975742162",
// 	ClientSecret: "d3229ebe3501771344bb0f2db2324014",
// 	Scopes:       []string{},
// 	Endpoint:     facebook.Endpoint,
// 	RedirectURL:  "http://loisant.org:9000/Application/Auth",
// }

// func (c Application) Index() revel.Result {
// 	u := c.connected()
// 	me := map[string]interface{}{}
// 	if u != nil && u.AccessToken != "" {
// 		resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
// 			url.QueryEscape(u.AccessToken))
// 		defer resp.Body.Close()
// 		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
// 			c.Log.Error("json decode error","error",err)
// 		}
// 		c.Log.Info("Data fetched","data",me)
// 	}

// 	authUrl := FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	return c.Render(me, authUrl)
// }

// func (c Application) Auth(code string) revel.Result {

// 	tok, err := FACEBOOK.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		c.Log.Error("Exchange error", "error",err)
// 		return c.Redirect(Application.Index)
// 	}

// 	user := c.connected()
// 	user.AccessToken = tok.AccessToken
// 	return c.Redirect(Application.Index)
// }

// func setuser(c *revel.Controller) revel.Result {
// 	var user *models.User
// 	if _, ok := c.Session["uid"]; ok {
// 		uid, _ := strconv.ParseInt(c.Session["uid"].(string), 10, 0)
// 		user = models.GetUser(int(uid))
// 	}
// 	if user == nil {
// 		user = models.NewUser()
// 		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
// 	}
// 	c.ViewArgs["user"] = user
// 	return nil
// }

// func init() {
// 	revel.InterceptFunc(setuser, revel.BEFORE, &Application{})
// }

// func (c Application) connected() *models.User {
// 	return c.ViewArgs["user"].(*models.User)
// }