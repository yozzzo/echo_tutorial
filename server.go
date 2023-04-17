package main

import (
	"echo-tutorial/models"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// memo: テンプレートの使い方が書いてある。
// https://echo.labstack.com/guide/templates/

type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()

	// 既にあるフィールドに代入する場合、短縮代入は使わない
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 適切なメソッドを使う
	e.GET("/todos", getAllTodo)
	e.GET("/todos/:id", getTodo)
	e.POST("/todos", createTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// 字面だけ読めば渡ってきたdataがmapならグローバルメソッドを読み込むようにしている。
	// なぜこういうことをしているのかはわからない。

	// add global methods if data is a map
	if viewContext, ok := data.(map[string]interface{}); ok {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func getAllTodo(c echo.Context) error {
	todos, err := models.GetAllTodo()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	data := map[string]interface{}{
		"todos": todos,
	}
	return c.Render(http.StatusOK, "todos.html", data)
}

func getTodo(c echo.Context) error {
	// 受け取り方
	// https://echo.labstack.com/guide/request/
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := models.GetTodo(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	data := map[string]interface{}{
		"todo": todo,
	}
	return c.Render(http.StatusOK, "detail.html", data)
}

func createTodo(c echo.Context) error {
	content := c.FormValue("content")
	_, err := models.CreateTodo(content)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.Redirect(http.StatusFound, "/todos")
}

func updateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	content := c.FormValue("content")

	todo, err := models.GetTodo(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	todo.Content = content
	_, err = models.UpdateTodo(todo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, "")
}

func deleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := models.GetTodo(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	err = models.DeleteTodo(todo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusNoContent, "")
}
