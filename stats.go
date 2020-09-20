package fab

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/kooinam/fab.io/actors"
	"github.com/markbates/pkger"
)

func serveStats() {
	f, _ := pkger.Open("github.com/kooinam/fab.io:/stats/index.html")
	b, _ := ioutil.ReadAll(f)
	content := string(b)

	tmpl := template.Must(template.New("stats").Parse(content))

	http.HandleFunc("/stats/index.html", func(w http.ResponseWriter, r *http.Request) {
		data := &struct {
			Actors []*actors.Actor
		}{
			Actors: ActorManager().GetActors(),
		}

		tmpl.Execute(w, data)
	})
}
