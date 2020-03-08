package main

import "vspro/blue/web/api"

func main() {
	api.Listen()
}

// var battle Battle = Battle{"id", map[*melody.Session]User{}}
// var life int = 1000
//
// func main() {
// 	router := gin.Default()
// 	mrouter := melody.New()
//
// 	router.POST("/", func(c *gin.Context) {
// 		uq := AnswerQuestion{Text: ""}
// 		c.BindJSON(&uq)
//
// 		q := entities.NewQuestionOne()
// 		file, err := os.OpenFile( "./1/" + q.Id() + ".go", os.O_WRONLY|os.O_CREATE, 0755)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer file.Close()
// 		file.Write(([]byte)(uq.Text))
// 		cmd := exec.Command("gofmt", "./1/" + q.Id() + ".go")
// 		if err := cmd.Run(); err != nil {
// 			uq.Result = "syntax error."
// 			c.JSON(http.StatusBadRequest, uq)
// 			return
// 		}
//
// 		cmd = exec.Command("/bin/bash", "./1/1.sh")
// 		if err := cmd.Run(); err != nil {
// 			uq.Result = "wrong answer."
// 			c.JSON(http.StatusBadRequest, uq)
// 			return
// 		}
//
// 		uq.Result = "collect answer."
// 		c.JSON(200, uq)
// 	})
//
// 	router.GET("/", func(c *gin.Context) {
// 		http.ServeFile(c.Writer, c.Request, "index.html")
// 	})
//
// 	router.GET("/ws", func(c *gin.Context) {
// 		mrouter.HandleRequest(c.Writer, c.Request)
// 	})
// 	http.StatusText(http.StatusOK)
//
// 	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
// 		mrouter.Broadcast(msg)
// 		d, _ := strconv.Atoi(string(msg))
// 		battle.Participant[s] = User{battle.Participant[s].ID, battle.Participant[s].LifePoint - d}
// 		l := strconv.Itoa(battle.Participant[s].LifePoint)
// 		dm := []byte("user:" + battle.Participant[s].ID + " collect answer. remaining is " + l)
// 		mrouter.Broadcast(dm)
// 	})
//
// 	mrouter.HandleConnect(func(s *melody.Session) {
// 		u := User{uuid.New().String(), life}
// 		battle.Participant[s] = u
// 	})
//
// 	router.Run(":8080")
// }
