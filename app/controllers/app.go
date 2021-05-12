package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {


	// facebook
    u := c.connected()
	me := map[string]interface{}{}
	if u != nil && u.AccessToken != "" {
		resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
			url.QueryEscape(u.AccessToken))
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			c.Log.Error("json decode error","error",err)
		}
		c.Log.Info("Data fetched","data",me)
	}

	authUrl := FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Render(me, authUrl)

	return c.Render()
}

func (c App) LoginSub() revel.Result {


	// facebook
	// twitter
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