package main

import (
	."fmt"
	"os/exec"
	."strings"
	."strconv"
	."time"
	"net"
	
)
const ever bool = true
var localIP string

func main(){
	Println("I am new here")
	//netIn := make(chan string)
	//netOut := make(chan string)
	consoleIn := make(chan string)
	consoleOut := make(chan string)
	//go InitCommunication(netIn, netOut)
	localIP = GetLocalIP()
	go InitConsoleIO(consoleIn,consoleOut)
	
	// BACKUP ============================
	bossAlive := true 
	downtime := 0
	x := 0
	udpaddr, _ := net.ResolveUDPAddr("udp", ":20011")
	conn, _ := net.ListenUDP("udp", udpaddr)
	conn.SetReadDeadline(Now().Add(Second))
	
	for (bossAlive){
		
		
		data := []byte("n"+"\x00")
		msg := string(data[0:])
		
		_, senderAddr, _ := conn.ReadFromUDP(data[0:])
		
		senderAddress := TrimRight(senderAddr.String(), "1234567890")
		senderAddress = TrimRight(senderAddress, ":")
		

		if localIP != senderAddress{
		 	msg = string(data[0:])
		 	
	 	}
	 	consoleIn <- msg
		// wait for boss to die
		
		
		bossCount := computeMessage(msg)
		
		if (bossCount == -1) {
			downtime++
		} else {
			
			downtime = 0
			if (bossCount == -1) {
				consoleIn <- "# Error:backup: computeMessage returned -1"
			} else { x = bossCount}
			
		}
		
		if (downtime >= 5) {
			bossAlive = false
		}
		Sleep(Second)
	}
	conn.Close()
	// BOSS ==============================
	SpawnNewFriend()
	UDPconn, err := net.Dial("udp", "localhost:20011")
	if err != nil{
			Println("# ERROR:conn: ",err)
		}
	
	UDPconn.SetWriteDeadline(Now().Add(Second))
	
	for (ever) {
		// shout angrily "I AM ALIVE, AND I AM THE BOSS!"
		//netIn <- "IMDABOS"		
		
		msg := "Bosscount:" + Itoa(x)
		data := []byte(msg+"\x00")
		Println(string(data))
		_, err := UDPconn.Write(data)
	
		if err != nil {
			Println("# ERROR:Broadcast:", err)
		}		
		
		// occasionally fail and/or spawn "friends"
		
		if (x == 13) {
			break
		}
		x++
		Sleep(Second)
	}
	conn.Close()
}

func computeMessage(msg string) int {
	num := TrimLeft(msg, "Boscunt: ")
	bosscount,err := Atoi(num)
	if (err != nil) { return -1}
	return bosscount
} 

func SpawnNewFriend() {
	Println("Spawning new \"friend\"")

    cmd := exec.Command("mate-terminal", "-x", "go", "run", "phoenix.go")
	out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(out))
    
}

// yoloswag
