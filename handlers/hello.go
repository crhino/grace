package handlers

import (
	"fmt"
	"net/http"

	"github.com/crhino/grace/helpers"
)

type Hello struct {
}

func (p *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, err := helpers.FetchIndex()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	styledTemplate.Execute(w, Body{fmt.Sprintf(`
<div class="hello">
	Diego!
</div>

<div class="my-index">My Index Is</div>

<div class="index">%d</div>
<div class="mid-color"></div>
<div class="bottom-color"></div>
    `, index)})
}
