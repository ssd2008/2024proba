package main

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
)

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	weekday := r.PathValue("weekday")
	subject := r.PathValue("subject")
	name := r.PathValue("name")
	ok := AddToSchedule(weekday, subject, name)
	w.WriteHeader(http.StatusOK)
	if !ok {
		_, _ = fmt.Fprint(w, "FAIL")
		return
	}

	_, _ = fmt.Fprintf(w, "%s, %s, %s", weekday, subject, name)
	return
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	lessons := GetNameSchedule(r.PathValue("name"))
	if len(lessons) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("["))
	for i, lesson := range lessons {
		if i > 0 {
			_, _ = w.Write([]byte(","))
		}
		_, _ = w.Write([]byte(lesson.String()))
	}
	_, _ = w.Write([]byte("]"))
	return
}

func main() {
	http.HandleFunc("/add/{weekday}/{subject}/{name}", add)
	http.HandleFunc("/get/{name}", get)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		return
	}
}

const (
	Monday    = "Mon"
	Tuesday   = "Tue"
	Wednesday = "Wed"
	Thursday  = "Thu"
	Friday    = "Fri"
)

type SubjectName struct {
	Subject string
	Name    string
}

var SchedulePerDay = map[string][]SubjectName{
	Monday:    {},
	Tuesday:   {},
	Wednesday: {},
	Thursday:  {},
	Friday:    {},
}

type SubjectWeekday struct {
	Subject string
	Weekday string
}

func (s SubjectWeekday) String() string {
	return fmt.Sprintf("[%q,%q]", s.Weekday, s.Subject)
}

var SchedulePerName = map[string][]SubjectWeekday{}

func AddToSchedule(weekday, subject, name string) bool {
	lessons, ok := SchedulePerDay[weekday]
	if !ok || len(lessons) >= 3 {
		return false
	}
	SchedulePerDay[weekday] = append(SchedulePerDay[weekday], SubjectName{Subject: subject, Name: name})
	for _, lesson := range lessons {
		if lesson.Subject == subject && lesson.Name == name {
			return true
		}
	}
	SchedulePerName[name] = append(SchedulePerName[name], SubjectWeekday{Subject: subject, Weekday: weekday})
	return true
}

var weight = map[string]int{
	Monday:    0,
	Tuesday:   1,
	Wednesday: 2,
	Thursday:  3,
	Friday:    4,
}

func GetNameSchedule(name string) []SubjectWeekday {
	res := SchedulePerName[name]
	slices.SortFunc(res, func(a, b SubjectWeekday) int {
		if a.Weekday != b.Weekday {
			return cmp.Compare(weight[a.Weekday], weight[b.Weekday])
		}
		return cmp.Compare(a.Subject, b.Subject)
	})
	return res
}
