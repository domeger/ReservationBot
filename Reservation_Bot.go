package main

import (
        "fmt"
        "log"
        "github.com/andybons/hipchat"
        "github.com/go-martini/martini"
        "github.com/martini-contrib/render"
)

const API_KEY = "{HIPCHAT ROOM API}"

// valid parameters to accept.  Move this to a config file
var validNames = []string{"deger"}
var validActions = []string{"checkout", "return"}
var validEnvironments = []string{"stage01", "stage02", "stage03", "stage04", "stage05"}


// config holds what's read from the config file
type config struct {
        usernames    []string
        actions      []string
        environments []string
}

// validEntry looks to see if the provided parameters are ok
func validEntry(entry string, item []string) bool {
        for _, i := range item {
                if i == entry {
                        return true
                }
        }
        return false
}

func validUser(username string) bool {
        for _, i := range validNames {
                if i == username {
                        return true
                }
        }
        return false
}

func validAction(action string) bool {
        for _, i := range validActions {
                if i == action {
                        return true
                }
        }
        return false
}

func validEnv(env string) bool {
        for _, i := range validEnvironments {
                if i == env {
                        return true
                }
        }
        return false
}

func main() {
        m := martini.Classic()
       
        m.Use(martini.Static("static"))
        m.Use(render.Renderer(render.Options{
                Layout: "layout",
        }))

        m.Get("/", func(r render.Render) {
                r.HTML(200, "index", nil)
        })
        m.Get("/status", func(r render.Render) {
                r.HTML(200, "status", nil)
        })     
         m.Get("/dashboard", func(r render.Render) {
                r.HTML(200, "dashboard", nil)
        })
        m.Get("/users", func(r render.Render) {
                r.HTML(200, "users", nil)
        })
         m.Get("/settings", func(r render.Render) {
                r.HTML(200, "settings", nil)
        })
        m.NotFound(func() {
               // Not Found
        })
        m.Get("/:action/:environment/:username", func(params martini.Params) {
                if validAction(params["action"]) && validEnv(params["environment"]) && validUser(params["username"]) {
                        c := hipchat.Client{AuthToken: API_KEY}
                        req := hipchat.MessageRequest{
                                RoomId:        "test1",
                                From:          "Reservation_Bot",
                                Message:       fmt.Sprintf("%s.stageing.test.com - %s by %s", params["environment"], params["action"], params["username"]),
                                Color:         hipchat.ColorPurple,
                                MessageFormat: hipchat.FormatText,
                                Notify:        true,
                        }
                        if err := c.PostMessage(req); err != nil {
                                log.Printf("Expected no error, but got %q \n", err)
                        }
                }
        })
        m.RunOnAddr(":8080")
}
