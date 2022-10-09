package main

import (
	"final_project_1_gin/docs"
	"final_project_1_gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var todos = map[int]*models.Todos{}
var seq = 1

// @BasePath /api/v1

// createTodos godoc
// @Summary Create Todos
// @Schemes
// @Description Create A new Todos
// @Tags Todos
// @Accept json
// @Produce json
// @param todos body models.Todos true "Create Todos"
// @Success 200 {string} createTodos
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Failure      500  {string}  httputil.HTTPError
// @Router /list-todos/create-todos [post]
func createTodos(c *gin.Context) {
	t := &models.Todos{
		ID: seq,
	}

	var result gin.H

	err := c.Bind(t)
	if err != nil {
		result = gin.H{
			"result": "Gagal Membuat data",
		}
	} else {
		result = gin.H{
			"result": t,
		}
	}
	todos[t.ID] = t
	seq++

	c.JSON(http.StatusCreated, result)
}

// @BasePath /api/v1

// getAllTodos godoc
// @Summary Get All Todos
// @Schemes
// @Description Get All Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {string} getAllTodos
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Failure      500  {string}  httputil.HTTPError
// @Router /list-todos/todos [get]
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// @BasePath /api/v1

// getTodoById godoc
// @Summary Get Todo By Id
// @Schemes
// @Description Get Todo By Id
// @Tags Todos
// @Accept json
// @Produce json
// @Param   id path int true "get todo Id"
// @Success 200 {string} getTodoById
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Failure      500  {string}  httputil.HTTPError
// @Router /list-todos/todos/{id} [get]
func getTodoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, todos[id])
}

// @BasePath /api/v1

// updateTodo godoc
// @Summary Update Todos
// @Schemes
// @Description Update Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param   id path int true "get todo Id"
// @param todos body models.Todos true "Update Todos"
// @Success 200 {string} updateTodo
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Failure      500  {string}  httputil.HTTPError
// @Router /list-todos/update-todos/{id} [put]
func updateTodo(c *gin.Context) {
	t := new(models.Todos)
	var result gin.H

	err := c.Bind(t)
	if err != nil {
		result = gin.H{
			"result": "Gagal Mengedit data",
		}
	} else {
		result = gin.H{
			"result": t,
		}
	}
	id, _ := strconv.Atoi(c.Param("id"))

	todos[id].Name = t.Name
	todos[id].Tanggal = t.Tanggal
	todos[id].Kegiatan = t.Kegiatan
	c.JSON(http.StatusOK, result)
}

// @BasePath /api/v1

// deleteTodos godoc
// @Summary Delete Todos
// @Schemes
// @Description Delete Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Param   id path int true "get todo Id"
// @Success 200 {string} deleteTodo
// @Failure      400  {string}  httputil.HTTPError
// @Failure      404  {string}  httputil.HTTPError
// @Failure      500  {string}  httputil.HTTPError
// @Router /list-todos/delete-todos/{id} [delete]
func deleteTodos(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todos, id)
	c.JSON(http.StatusOK, "Delete data Successfully")
}

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/list-todos")
		{
			eg.POST("/create-todos", createTodos)
			eg.GET("/todos", getAllTodos)
			eg.GET("/todos/:id", getTodoById)
			eg.PUT("/update-todos/:id", updateTodo)
			eg.DELETE("/delete-todos/:id", deleteTodos)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
