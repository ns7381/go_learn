package main

import (
	"github.com/bndr/gojenkins"
	"fmt"
	"io/ioutil"
)

func main() {

	jenkins := gojenkins.CreateJenkins(nil, "http://10.221.129.29:18080", "admin", "123456a?")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init()
	if err != nil {
		panic("Something Went Wrong")
	}

	// Create folder
	/*pFolder, err := jenkins.CreateFolder("devops")
	if err != nil {
		panic(err)
	}*/

	// Create job in folder
	bytes, err := ioutil.ReadFile("E:/go_path/src/github.com/ns7381/go_learn/jenkins_learn/job.xml")
	if err != nil {
		panic(err)
	}
	job, err := jenkins.CreateJobInFolder(string(bytes), "test02")
	if err != nil {
		panic(err)
	}

	if job != nil {
		fmt.Println("Job has been created in child folder")
	}

}