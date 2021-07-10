package tasks

import "net/http"

type HttpTasks struct {
	body     map[interface{}]string
	response http.Response
	err      error
}

func (h *HttpTasks) Do() Tasks {
	return h
}
