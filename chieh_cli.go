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
		    Category: "List",
		    Action: func(c *cli.Context) error {
                // command_list
                // args := []string{"hello", "world"}
                // args_1 := []string{"hi"}
                
                for i:=0; i<2; i++{
                    // fmt.Println(args_total[i])
                // cmd:= exec.Command("echo", args...)
		        cmd := exec.Command("ls", "-lah")
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
			Name:     "Install",
			Aliases:  []string{"i", "I", "install"},
			Usage:    "Install some packages regarding to native Linux tools.",
			Category: "Install",
			Subcommands: []cli.Command{
				{
					Name:  "all",
					Usage: "Directly install all packages from my personal list.",
					Action: func(c *cli.Context) error {
						// sudo apt-get update
						cmd := exec.Command("sudo", "apt-get", "update")
						out, err := cmd.CombinedOutput()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						fmt.Printf(string(out))

						// sudo apt -y install vim net-tools git ca-certificates openssh-server python3-pip
						cmd = exec.Command("sudo", "apt-get", "-y", "install", "vim", "net-tools", "git", "ca-certificates", "openssh-server", "python3-pip")
						out, err = cmd.CombinedOutput()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						fmt.Printf(string(out))

						// Start install docker env
						fmt.Println("\n>> Start to install the docker environment")
                        cmd = exec.Command("sudo", "apt-get", "-y", "install", "vim", "net-tools", "git", "ca-certificates", "openssh-server", "python3-pip")
						out, err = cmd.CombinedOutput()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						fmt.Printf(string(out))

                        // sudo apt install -y apt-transport-https curl gnupg-agent software-properties-common
                        // curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
                        // sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
                        // sudo apt update
                        // sudo apt install -y docker-ce docker-ce-cli containerd.io
                        // sudo usermod -aG docker $USER
                        // sudo chmod 777 /var/run/docker.sock


						return nil
					},
				},
				// {
				//     Name:  "remove",
				//     Usage: "remove an existing template",
				//     Action: func(c *cli.Context) error {
				//         fmt.Println("removed task template: ", c.Args().First())
				//         return nil
				//     },
				// },
			},
		},
	}

	app.Name = "AppName"
	app.Usage = "application usage"
	app.Description = "application description" // 描述
	app.Version = "1.0.1"                       // 版本

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Please check --help")
		log.Fatal(err)
	}
}
