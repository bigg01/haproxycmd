package main

import (
	"fmt"
	"haproxycmd"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.TextFormatter{})
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("error: %s", err)
		}
	}()

	sock := haproxycmd.ParseFlag()

	if sock == "" {
		sock = haproxycmd.FindHaproxySocket()
	}

	//haproxycmd.Command(sock, os.Args[1:], os.Stdout)

	//for _, cmd := range []string{"show acl", "show acl #1","add acl #1 10.0.0.208","show acl #1","del acl #1 10.0.0.208","show acl #1"} {

	//        ./haproxycmd "get acl #1 10.0.0.208"

	//for _, cmd := range []string{"show acl", "show acl #1","add acl #1 10.0.0.208","show acl #1","","show acl #1"} {

	/*

		       # cmd:  get acl #1 10.0.0.208
		type=ip, case=sensitive, match=yes, idx=tree, pattern="10.0.0.208"

		type=ip, case=sensitive, match=yes, idx=tree, pattern="10.0.0.208"

		# cmd:  get acl #1 10.0.0.4
		type=ip, case=sensitive, match=no

		type=ip, case=sensitive, match=no

	*/

	aclwhitelislt := []string{"10.0.0.208", "10.0.0.4"}

	for _, ip := range aclwhitelislt {
		cmd := strings.ToLower(fmt.Sprintf("get acl #1 %s", ip))

		ret := haproxycmd.Command(sock, []string{cmd}, os.Stdout)
		if ret != 1 {
			log.WithFields(log.Fields{
				"cmd":     cmd,
				"haproxy": ret,
			}).Warn("Cannot get ACL")
		} else {
			//log.Infof("--> haproxy returned success: %d %s", ret, cmd)
			log.WithFields(log.Fields{
				"cmd":     cmd,
				"haproxy": ret,
			}).Info("ACL exist")
		}

		cmd = fmt.Sprintf("add acl #1 %s", ip)
		//log.Println("# cmd: ", cmd)
		ret = haproxycmd.Command(sock, []string{cmd}, os.Stdout)
		if ret != 1 {
			log.WithFields(log.Fields{
				"cmd":     cmd,
				"haproxy": ret,
			}).Warn("cannot add ACL!")
		} else {
			//log.Infof("--> haproxy returned success: %d %s", ret, cmd)
			log.WithFields(log.Fields{
				"cmd":     cmd,
				"haproxy": ret,
			}).Info("successfully added ACL")
		}

		//time.Sleep(5 * time.Second)
		cmd = fmt.Sprintf("del acl #1 %s", ip)
		log.Println("# cmd: ", cmd)
		haproxycmd.Command(sock, []string{cmd}, os.Stdout)

		cmd = fmt.Sprintf("show acl #1 %s", ip)
		ret = haproxycmd.Command(sock, []string{cmd}, os.Stdout)
		if ret != 1 {
			log.Errorf("--> haproxy returned error: %d %s", ret, cmd)
		} else {
			log.Infof("--> haproxy returned success: %d %s", ret, cmd)
		}

	}
}
