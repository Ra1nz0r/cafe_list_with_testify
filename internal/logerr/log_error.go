package logerr

import "log"

func LogErrWrite(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
		return
	}
}

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
