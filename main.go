package main

func main(){
	server := NewServer(":3000")
	server.Handle("GET","/", server.AddMiddleware(HandleRoot, CheckAuth(), Logging()))
	server.Handle("POST","/create", PostRequest)
	server.Handle("POST","/auth", HandleAuth)
	server.Handle("POST","/auth/signup", UserPostRequest)
	server.Listen()
}