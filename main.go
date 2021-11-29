package main

func main(){
	server := NewServer(":3000")
	server.Handle("GET","/", server.AddMiddleware(HandleRoot, CheckAuth(), Logging()))
	server.Handle("POST","/auth", HandleAuth)
	server.Listen()
}