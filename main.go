package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"io/ioutil"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
type Todos []Todo


// http://localhost:8080
func Index(c *gin.Context) {
	c.String(200,"Hello world")
}

// http://localhost:8080/todos GET
func TodoIndex(c *gin.Context){
	todos := Todos{
		Todo{Name: "Write presentation",Completed:true,Due:time.Now()},
		Todo{Name: "Host meetup",Completed:false,Due:time.Now()},
	}
	c.JSON(200,todos)
}

// http://localhost:8080/todo POST
func TodoPost(c *gin.Context) {
	title := c.PostForm("title")
	message := c.PostForm("message")

	c.JSON(200, gin.H{
		"status":  "posted",
		"title":  title,
		"message": message,
	})
}

// http://localhost:8080/file GET
func FileOutput(c *gin.Context){
	contents,err := ioutil.ReadFile("sample.json") // ファイルの読み込み
	if err != nil {
		fmt.Println(contents, err)
		c.JSON(400,gin.H{"status": "bad request"})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(200,string(contents))
}

func main() {
	r := gin.Default()
	r.GET("/",Index)
	r.GET("/todos",TodoIndex)
	r.POST("/todo",TodoPost)
	r.GET("/file",FileOutput)
	r.Run(":8080")
	fmt.Println("server start port 8080.")
}
