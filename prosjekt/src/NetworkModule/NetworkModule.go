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

func InitNetworkModule(LocalIP, port string, in, out chan string) *NetworkModule {
  m := new(NetworkModule)
  m.localIP = LocalIP
  
  return m
}

func (m *NetworkModule) Broadcast() {
  conn, err := net.Dial("udp", m.localIP+m.port)
  if err != nil {
    //handle error
  }
  
  
  var msg string
  
  for {

    msg = <-m.in
    data := []byte(msg+"\x00")
    
    _, err = conn.Write(data)
    
    if err != nil {
      //TODO: handle error, maybe differently than above
    }
  }
}

func (m *NetworkModule) Listen() {
  udpaddr, _ := net.ResolveUDPAddr("udp", m.port)
  conn, err := net.ListenUDP("udp", udpaddr)
  
  if err != nil {
    //TODO: handle error
  }
  

  msg := "**insert default message**"
  
  for {
    conn.SetReadDeadline(Now().Add(Second))
    data := []byte(msg+"\x00")
    
    _, senderAddr, _ := conn.ReadFromUDP(data[0:])
    
    senderAddress := TrimRight(senderAddr.String(), "1234567890")
    senderAddress = TrimRight(senderAddress, ":")
    
    if m.localIP != senderAddress {
      m.out <- string(data[0:])
    }
  }
}
