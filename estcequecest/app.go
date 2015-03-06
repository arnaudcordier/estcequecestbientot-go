// Package estcequecest provides controller and library to show messages at a given time
//
// You can use an App to control different instance of Estcequecest
// or just use Estcequecest library by itself
package estcequecest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var ErrBadJson = errors.New("bad json data")

// manage a collection of Estcequecest
type App struct {
	loaded       []string
	estcequecest map[string]*Estcequecest
	pattern      string
}

// get a new App giving a dir and a pattern for files
func NewApp(dir string, pattern string) *App {
	app := new(App)
	fi, err := os.Stat(dir)
	if err == nil && fi.IsDir() {
		app.pattern = path.Join(dir, pattern)
	} else {
		panic("Can not open directory: " + dir)
	}
	app.loaded = make([]string, 0)
	app.estcequecest = make(map[string]*Estcequecest)
	return app
}

// create a new estcequecest from a json file
func (app *App) Load(name string) error {
	filename := app.pattern + name + ".json"
	str, _ := ioutil.ReadFile(filename)

	var data estcequecestData
	er := json.Unmarshal(str, &data)
	if er != nil {
		return ErrBadJson
	}

	// keep track and order of loaded estcequecest
	loaded := false
	for _, v := range app.loaded {
		if v == name {
			loaded = true
			break
		}
	}
	if !loaded {
		app.loaded = append(app.loaded, name)
	}
	// always reload the object
	app.estcequecest[name] = NewEstcequecest(data)
	return nil
}

func (app *App) GetMessages() [][]string {
	t := time.Now()
	return app.GetMessagesAtTime(t)
}

func (app *App) GetMessagesAtTime(t time.Time) [][]string {
	messages := make([][]string, len(app.loaded))
	for index, name := range app.loaded {
		title, message := app.estcequecest[name].getMessageAtTime(t)
		messages[index] = []string{title, message}
	}
	return messages
}

// prety print
func (app *App) String() string {
	s := fmt.Sprintf("pattern: %s", app.pattern)
	if len(app.loaded) > 0 {
		s = s + "\nloaded Estcequecest :\n"
		for _, name := range app.loaded {
			s = s + "** " + name + " **\n"
			s = s + app.estcequecest[name].String()
		}
	}
	return s
}

// (App) Unload(name) error
// (App) List() []string, []string
