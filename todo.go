package main

import "time"

type Todo struct {
	Name			string 		 `json:"name"`
	Completed bool			 `json:"completed"`
	due				time.Time  `json:"due"`
}

type Todos []Todo