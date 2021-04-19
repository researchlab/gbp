package main

import (
	"fmt"
	"regexp"
)

func main() {
	t := "virt-v2v: Overlay saved as /var/tmp/xxx-aaaa-abx.qcow2 xxxxy"
	r := regexp.MustCompile(fmt.Sprintf("/var/tmp/%s.*.qcow2", "xxx"))
	matches := r.FindString(t)
	fmt.Println(matches)
}
