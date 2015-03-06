	estcequecest
		title
		defaultmessage
		messages[]
			interval : * * 13-22 5-7 09-20,23 *
			timeline : [hh:mm : message]    ****************************************  

messages.getMessage => messages.getAtTime

type estcequecestData struct {
	Title          string
	Defaultmessage string
	Messages       []messagesData `json:"messages"`
}

type messagesData struct {
	Interval string
	Timeline map[string]string
}

type App struct {
	loaded   []string
	messages map[string]Estcequecest
	dir      string
	pattern  string
}
(App) Load(name) error
(App) Unload(name) error
(App) List() []string, []string
(App) Get() []messages
(App) GetAtTime(time) []messages
NewApp(dir, pattern)


type Estcequecest struct {
	title          string
	defaultMessage string
	messages       []messages
}
(Estcequecest) GetAtTime(time) title string, message string
NewEstcequecest(estcequecestData)


type messages struct {
	interval interval
	timeline timeline
}
(messages) getAtTime(time) string, error
newMessages(messagesData)


type interval struct {
	bits []bit
}
(interval) doesItFits(time) interval, error
newInterval(string)


type bit struct {
	all bool
	min int
	max int
}
(bit) doesItFits(int) bit, error
newBit(string, pos int)


type timeline struct {
	messages map[int]string
	order    []int
}
(timeline) getMessage(minute int) string
newTimeline([string]string)
