// Author: Chieh Tsai
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"github.com/urfave/cli"
)

func main() {

	// var command_list := make([]string, 10, 20)
	type command_list struct{}

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:     "list",
			Aliases:  []string{"ls"},
			Usage:    "List the files in the present directory.",
			Category: "list",
			Action: func(c *cli.Context) error {

				// command_list
				args := []string{
					"ls -lah",
					"df -h"}

				for i := 0; i < 2; i++ {
					cmd := exec.Command("bash", "-c", args[i])
					// cmd := exec.Command("ls", "-lah")
					out, err := cmd.CombinedOutput()
					if err != nil {
						log.Fatalf("cmd.Run() failed with %s\n", err)
					}
					fmt.Printf(string(out))
				}
				return nil
			},
		},
		{
			Name:     "install",
			Aliases:  []string{"i", "I", "Install"},
			Usage:    "To install some packages regarding to native Linux tools.",
			Category: "install",
			Subcommands: []cli.Command{
				{
					Name:  "all",
					Usage: "Directly install all packages from my personal list.",
					Category: "install",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name: "no-gpu",
							Usage: "Without installing the GPU driver.", 
							},
						},
					Action: func(c *cli.Context) error {
						// sudo apt-get update
						cmd := exec.Command("sudo", "apt-get", "update")
						out, err := cmd.CombinedOutput()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						fmt.Printf(string(out))

						// sudo apt -y install vim net-tools git ca-certificates openssh-server python3-pip
						basic_tool_list := []string{
							"vim", 
							"net-tools", 
							"git",
							// "ca-certificates", 
							"openssh-server", 
							"python3-pip",
						}
						var TestToolWord string 
						for btl:=0; btl< len(basic_tool_list); btl++ {
							switch basic_tool_list[btl]{
							case "python3-pip":
								TestToolWord = "python3"
							case "openssh-server":
								TestToolWord = "ssh"
							case "net-tools":
								TestToolWord = "ifconfig"
							default:
								TestToolWord = basic_tool_list[btl]
							} 
							_, err := exec.LookPath(TestToolWord)
							if err != nil {
								cmd = exec.Command("sudo", "apt-get", "-y", "install", basic_tool_list[btl])
								fmt.Printf("We are going to install the %s tool.", basic_tool_list[btl])
								out, err = cmd.CombinedOutput()
								if err != nil {
									log.Fatalf("cmd.Run() failed with %s\n", err)
								}
								fmt.Printf(string(out))
							}
						}


						// Start install docker env
						fmt.Println("------------------------------------------")
						fmt.Println("Checking Docker tool...")
						_, err = exec.LookPath("docker")
						if err != nil {
							fmt.Println("\n>> Start to install the environment of Docker")
							docker_list := []string{
								"sudo apt install -y apt-transport-https curl gnupg-agent software-properties-common",
								"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
								"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
								"sudo apt update",
								"sudo apt install -y docker-ce docker-ce-cli containerd.io",
								"sudo usermod -aG docker $USER",
								"sudo chmod 777 /var/run/docker.sock",
								}
							
							for i:=0; i<len(docker_list); i++{
								cmd = exec.Command("bash", "-c", docker_list[i])
								out, err = cmd.CombinedOutput()
								if err != nil {
									log.Fatalf("cmd.Run() failed with %s\n", err)
								}
								fmt.Printf(string(out))
							}

							fmt.Println("\nDone for the installation of Docker.\n")
						} else{
							fmt.Println("Already installed Docker.")
						}

						// 
						fmt.Println("------------------------------------------")
						fmt.Println("Checking Docker-Compose tool...")
						_, err = exec.LookPath("docker")
						if err != nil {
							fmt.Println("Start to install Docker Compose tool.\n")
							docker_compose_list := []string{
								"sudo curl -L \"https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)\" -o /usr/local/bin/docker-compose",
								"sudo chmod +x /usr/local/bin/docker-compose",							
							}
							for i:=0; i<len(docker_compose_list); i++{
								cmd = exec.Command("bash", "-c", docker_compose_list[i])
								out, err = cmd.CombinedOutput()
								if err != nil {
									log.Fatalf("cmd.Run() failed with %s\n", err)
								}
								fmt.Printf(string(out))
							}
							fmt.Println("\nDone for the installation of Docker Compose tool.\n")
						} else {
							fmt.Println("Already installed Docker-Compose.")
						}
						
						if  c.Bool("no-gpu") != true {
							fmt.Println("Start to install GPU driver...")
							// Start install nvidia-driver
							fmt.Println("------------------------------------------")
							fmt.Println("Checking GPU status ...")
							_, err = exec.LookPath("nvidia-smi")
							if err != nil {
								fmt.Println("\n>> Start to install the GPU driver Version 470.")
								gpu_driver_list := []string{
									"sudo sed -i -e 's/tw.archive.ubuntu.com/free.nchc.org.tw/' /etc/apt/sources.list",
									"sudo apt update",
									"sudo apt install -y nvidia-driver-470-server",
									}
								
								for i:=0; i<len(gpu_driver_list); i++{
									cmd = exec.Command("bash", "-c", gpu_driver_list[i])
									out, err = cmd.CombinedOutput()
									if err != nil {
										log.Fatalf("cmd.Run() failed with %s\n", err)
									}
									fmt.Printf(string(out))
								}

								fmt.Println("\nDone for the installation of GPU driver.\n")
							} else{
								fmt.Println("Already installed nvidia-smi.")
							}
						}

						return nil
					},
				},
				{
				    Name:  "remove",
				    Usage: "remove an existing",
					Category: "install",
					Flags: []cli.Flag{
						&cli.BoolFlag{Name: "serve"},
						&cli.BoolFlag{Name: "option"},
						&cli.StringFlag{Name: "message"},
					},
				    Action: func(c *cli.Context) error {
				        // fmt.Println("removed task template: ", c.Args().First())
						fmt.Println("serve:", c.Bool("serve"))
						fmt.Println("option:", c.Bool("option"))
						fmt.Println("message:", c.String("message"))
				        return nil
				    },
				},
			},
		},
				{
			Name:     "pull",
			Aliases:  []string{"p", "Pull"},
			Usage:    "Pull the docker image. Notice: please install Docker first.",
			Category: "pull",
			Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "u",
				Usage: "Provide your Docker username for login use.",
				Required: true, 
				},
			},
			Action: func(c *cli.Context) error {
				if len(c.String("u")) == 0 {
					fmt.Println("Please add an option '-u' for providing your username.")
					os.Exit(0)
				}

				_, err := exec.LookPath("docker")
				if err != nil {	
					fmt.Println("Please install Docker first. Or check your Docker environment path.")
					os.Exit(0)
				}

				var pwd string 
				
				fmt.Printf("Your user name is %s\n", c.String("u"))
				fmt.Print("Please type the password: ")
				fmt.Print("\033[8m") // Hide input
				fmt.Scan(&pwd)
				fmt.Print("\033[28m") // Show input
				// fmt.Printf("Here is your password: %s\n", pwd)
				
				// Check the docker login
				docker_login := "docker login -u " + c.String("u") + " -p " + pwd
				fmt.Println(docker_login)
				cmd := exec.Command("bash", "-c", docker_login)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Login failure. Please check the uersname and password.")
					os.Exit(0)
					// log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				fmt.Printf(">> %s\n", string(out))
				

				return nil
			},
		},
	}

	app.Name = "chieh-cli"
	app.Usage = "install some basic packages for Linux Ubuntu OS."
	app.Description = "It is a Chieh's personal CLI tool for setting environment on Linux."
	app.Version = "0.1.0"

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Please check --help to get more usage information.")
		os.Exit(0)
		// log.Fatal(err)
	}
}
