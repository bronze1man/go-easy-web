package main

import (
  "net/http"
  "net"
  "os"
  "fmt"
)

func main(){
  currentDir,err:=os.Getwd()
  if err!=nil{
    panic(err)
  }
  http.Handle("/",http.FileServer(http.Dir(currentDir)))
  l,err:=listenOnAddr(":18080")
  if err!=nil{
    panic(err)
  }
  fmt.Printf("Listen on %s://%s\n",l.Addr().Network(),l.Addr().String() )
  err=http.Serve(l,nil)
  if err!=nil{
    panic(err)
  }
}
//first try addr,if err happened try random addrss.
func listenOnAddr(addr string)(l net.Listener,err error){
  l,err=net.Listen("tcp",addr)
  if err==nil{
    return
  }
  l,err=net.Listen("tcp",":0")
  return 
}