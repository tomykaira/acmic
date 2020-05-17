package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"io/ioutil"
	"os"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) != 4 {
		fmt.Println("Set parameters: priv.key cert.crt pki.crt")
		return
	}
	priv, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	cert, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[2])
		panic(err)
	}
	chain, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[3])
		panic(err)
	}

	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.
	sess := session.Must(session.NewSession())

	// Create service client value configured for credentials
	// from assumed role.
	svc := acm.New(sess)
	out, err := svc.ImportCertificate(&acm.ImportCertificateInput{
		Certificate:      cert,
		CertificateChain: chain,
		PrivateKey:       priv,
	})
	if err != nil {
		fmt.Println("Got error calling ImportCertificate:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully imported certificate")
	fmt.Println(*out.CertificateArn)
}
