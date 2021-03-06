package main

import (
    "fmt"

    "github.com/fsouza/go-dockerclient"
)

func main() {
    client, _ := docker.NewClientFromEnv()
    imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
    for _, img := range imgs {
        fmt.Println("ID: ", img.ID)
        fmt.Println("RepoTags: ", img.RepoTags)
        fmt.Println("Created: ", img.Created)
        fmt.Println("Size: ", img.Size)
        fmt.Println("VirtualSize: ", img.VirtualSize)
        fmt.Println("ParentId: ", img.ParentID)
    }
    
    pullFail := client.PullImage(docker.PullImageOptions{Repository:"nginx", Tag:"alpine"},docker.AuthConfiguration{Username:"emyann"})    
    if(pullFail != nil){
        fmt.Println("pullFail", pullFail)
    }else{
            opts := docker.CreateContainerOptions{
                Name:"ContainerGenerated",
                Config:&docker.Config{
                    AttachStderr:   true,
                    AttachStdin:    true,
                    AttachStdout:   true,
                    Image:          "nginx:alpine",
                    
                     },
                
            }
           
           
        fmt.Println(opts.Name)
        container, error := client.CreateContainer(opts)
        if error != nil{
             fmt.Println("error", error)
        }else {
            fmt.Println("container info", container.ID)            
        }
    }
    
    

}