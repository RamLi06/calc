package main

import(
	"log"
	"net/http"
	"time"
		
			)

func runservice(){
	s := &http.server{
		addr:   	"192.168.0.1:9999",
		//handler:   myhandler,
		readtimeout: 10	* time.second,
		writetimeout: 10 * time.second,
	}
	log.fatal(s.ListenAndServe())
}