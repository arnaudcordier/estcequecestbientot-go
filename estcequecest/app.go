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
	"path/filepath"
	"strings"
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

// get list of estcequecest messages for now
func (app *App) GetMessages() [][]string {
	t := time.Now()
	return app.GetMessagesAtTime(t)
}

// get list of estcequecest messages for the given time
func (app *App) GetMessagesAtTime(t time.Time) [][]string {
	messages := make([][]string, len(app.loaded))
	for index, name := range app.loaded {
		title, message := app.estcequecest[name].GetMessageAtTime(t)
		messages[index] = []string{title, message}
	}
	return messages
}

// create a new estcequecest from a json file
func (app *App) Load(name string) error {
	// check if name is an existing estcequecest
	exists := false
	for _, n := range app.listNames() {
		if n == name {
			exists = true
			break
		}
	}
	if !exists {
		return nil
	}

	// open and parse json file
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

// unload an estcequecest
func (app *App) Unload(name string) error {
	if _, ok := app.estcequecest[name]; ok {
		delete(app.estcequecest, name)
		loaded := make([]string, len(app.loaded)-1)
		index := 0
		for _, n := range app.loaded {
			if n != name {
				loaded[index] = n
				index += 1
			}
		}
		app.loaded = loaded
	}
	return nil
}

// Return lists of notloaded names of estcequecest and loaded ones
func (app *App) List() ([]string, []string) {
	allNames := app.listNames()
	notloaded := make([]string, len(allNames)-len(app.loaded))
	index := 0
	for _, name := range allNames {
		if _, ok := app.estcequecest[name]; !ok {
			notloaded[index] = name
			index += 1
		}
	}
	return notloaded, app.loaded
}

// return a list of available estcequecest names
func (app *App) listNames() []string {
	files, _ := filepath.Glob(app.pattern + "*.json")
	for index, name := range files {
		name = strings.Replace(name, app.pattern, "", 1)
		name = strings.Replace(name, ".json", "", 1)
		files[index] = name
	}
	return files
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
