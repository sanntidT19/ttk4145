package NetworkModule

import(
  "net"
  ."time"
  ."strings"
  
)

type NetworkModule struct {
  //channels for communicating with the rest of the system
  in chan string
  out chan string
  
  //localIP, without port!
  localIP string
  //port, ":xxxxx"
  port string
  //maybe something else
}

func InitNetworkModule(LocalIP, port string) NetworkModule {
  var module NetworkModule
  module.localIP = LocalIP
  
  return module
}

func (n NetworkModule) *Broadcast() {
  conn, err := net.Dial("udp", n.localIP+n.port)
  if err != nil {
    //handle error
  }
  conn.setWriteDeadline(Now().Add(Second))
  
  var msg string
  
  for {
    msg = <-n.in
    data := []byte(msg+"\x00")
    
    _, err = conn.Write(data)
    
    if err != nil {
      //TODO: handle error, maybe differently than above
    }
}

func (n NetworkModule) *Listen() {
  udpaddr, _ := net.ResolveUDPAddr("udp", port)
  conn, err := net.ListenUDP("udp", udpaddr)
  
  if err != nil {
    //TODO: handle error
  }
  
  conn.setReadDeadline(Now().Add(Second))
  msg := "**insert default message**"
  
  for {
    data := []byte(msg+"\x00")
    
    _, senderAddr, _ := conn.ReadFromUDP(data[0:])
    
    senderAddress := TrimRight(senderAddr.String(), "1234567890")
    senderAddress = TrimRight(senderAddress, ":")
    
    if localIP != senderAddress {
      ut <- string(data[0:])
    }
}
