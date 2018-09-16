package consistent

import (
	"fmt"
	"testing"

	"github.com/xujihui1985/learninggo/consistent/consistent2"
)

func TestNew(t *testing.T) {
	c := New()
	c.Add("log-replicate-00")
	c.Add("log-replicate-01")
	c.Add("log-replicate-02")

	logProject, _ := c.Get("tentent1")
	fmt.Println(logProject)
	logProject, _ = c.Get("tentent2")
	fmt.Println(logProject)
	logProject, _ = c.Get("tentent3")
	fmt.Println(logProject)

	c.Add("log-replicate-03")
	c.Add("log-replicate-05")

	// after add replicate
	logProject, _ = c.Get("tentent1")
	fmt.Println(logProject)
	logProject, _ = c.Get("tentent2")
	fmt.Println(logProject)
	logProject, _ = c.Get("tentent3")
	fmt.Println(logProject)
}

func TestConsistent2(t *testing.T) {
	c := consistent2.New(10, nil)
	c.Add("log-replicate-00")
	c.Add("log-replicate-01")
	c.Add("log-replicate-02")

	logProject := c.Get("tentent1")
	fmt.Println(logProject)
	logProject = c.Get("tentent2")
	fmt.Println(logProject)
	logProject = c.Get("tentent3")
	fmt.Println(logProject)

	c.Add("log-replicate-03")
	c.Add("log-replicate-05")

	logProject = c.Get("tentent1")
	fmt.Println(logProject)
	logProject = c.Get("tentent2")
	fmt.Println(logProject)
	logProject = c.Get("tentent3")
	fmt.Println(logProject)
}
