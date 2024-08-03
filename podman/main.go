package main

// This part of the code will handle the container information through CRI sockets and podman

func main(){
	podmanHandler, err := NewPodmanHandler()
	if err!=nil{
		return 
	}
	podmanHandler.GetContainerInfo("container-id")
}

