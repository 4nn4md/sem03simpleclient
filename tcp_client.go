package main

import (
	"net"
	"log"
	"os"
	"temp/mycrypt"
	
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.6:1000")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])

kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
kryptertString := string(kryptertMelding)

log.Println("Kryptert melding: ", kryptertString)


 	_, err = conn.Write([]byte(kryptertString))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)

	dekryptertSvar := string(mycrypt.Krypter([]rune(response), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4))
	log.Printf("Decrypted reply from proxy: %s", dekryptertSvar)

}
