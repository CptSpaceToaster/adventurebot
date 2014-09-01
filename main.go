package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/trinchan/slackbot/robots"
	"io"
	"log"
	"net/http"
	"strconv"
    "fmt"
)

func main() {
	http.HandleFunc("/slack_hook", CommandHandler)
    StartServer()
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		decoder := schema.NewDecoder()
		command := new(robots.SlashCommand)
		err := decoder.Decode(command, r.PostForm)
		if err != nil {
			log.Println("Couldn't parse post request:", err)
		}
        //I'm not sure if this condition is possible in the setu
        fmt.Printf("Recieved command: %s\n", command.Text)
		w.WriteHeader(http.StatusOK)
		robot := GetRobot()
        JSONResp(w, robot.Run(command))
	}
}

func JSONResp(w http.ResponseWriter, msg string) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    resp := map[string]string{"text": msg}
    r, err := json.Marshal(resp)
    if err != nil {
        log.Println("Couldn't marshal hook response:", err)
    } else {
        io.WriteString(w, string(r))
    }
}

func plainResp(w http.ResponseWriter, msg string) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    io.WriteString(w, msg)
}

func StartServer() {
    port := robots.Config.Port
    log.Printf("Starting HTTP server on %d", port)
    err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
    if err != nil {
        log.Fatal("Server start error: ", err)
    }
}

func GetRobot() robots.Robot {
    if RobotInitFunction, ok := robots.Robots["adventurebot"]; ok {
        return RobotInitFunction()
    } else {
        return nil
    }
}
