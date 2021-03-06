package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

//puts path vars right into variables passed as params, until it runs out
//ex: pathVars(r,"/entry/",&id,&todo,&val) // populates id, todo, and val
func pathVars(r *http.Request, root string, vals ...*string) {
	a := strings.Split(r.URL.Path[len(root):], "/")
	for i := range vals {
		if len(a) > i {
			*vals[i] = a[i]
		} else {
			*vals[i] = ""
		}
	}
}
func sth(db *sql.DB, s string) *sql.Stmt {
	a, err := db.Prepare(s)
	if err != nil {
		print(s)
		panic(err)
	}
	return a
}

func evenodd(i int) string {
	if i%2 == 0 {
		return "even"
	}
	return "odd"
}
func unescape(s string) string {
	s = strings.Replace(s, "&#34;", "\"", -1)
	s = strings.Replace(s, "&#47;", "/", -1)
	s = strings.Replace(s, "&#39;", "'", -1)
	s = strings.Replace(s, "&#42;", "*", -1)
	s = strings.Replace(s, "&#63;", "?", -1)
	s = strings.Replace(s, "&#8216;", "'", -1)
	s = strings.Replace(s, "&#8217;", "'", -1)
	s = strings.Replace(s, "&#8220;", "'", -1)
	s = strings.Replace(s, "&#8221;", "'", -1)
	s = strings.Replace(s, "&#8211;", "-", -1)
	s = strings.Replace(s, "&#8594;", "->", -1)
	s = strings.Replace(s, "&quot;", "'", -1)
	s = strings.Replace(s, "&amp;", "&", -1)
	return s
}
func tostr(i int) string {
	s := strconv.Itoa(i)
	return s
}
func toint(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func randomString(l int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
	}
