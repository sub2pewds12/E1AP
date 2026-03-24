package e1ap_test

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/ishidawataru/sctp"
	e1ap "github.com/sub2pewds12/E1AP"
)

func TestManualClient(t *testing.T) {
	fmt.Println("--- Starting Go E1AP Client (Golden Payload Strategy) ---")

	// 1. Network Config
	cpIPStr := "192.168.70.129"
	cpPort := 38462
	upIPStr := "192.168.70.129"
	upPort := 0 // Ephemeral port

	cpIP, _ := net.ResolveIPAddr("ip", cpIPStr)
	upIP, _ := net.ResolveIPAddr("ip", upIPStr)

	// 2. Prepare Payload (OAI Golden Bytes)
	// This is the output from 'my_e1_generator' (High Level API).
	// It contains the structure/padding OAI demands.
	payload := []byte{
		0x00, 0x03, 0x00, 0x31, 0x00, 0x00, 0x05, 0x00, 0x39, 0x00, 0x02, 0x00,
		0x4D, 0x00, 0x07, 0x00, 0x03, 0x20, 0x30, 0x39, 0x00, 0x08, 0x40, 0x0F,
		0x06, 0x00, 0x4F, 0x41, 0x49, 0x2D, 0x43, 0x55, 0x2D, 0x55, 0x50, 0x2D,
		0x47, 0x45, 0x4E, 0x00, 0x0A, 0x00, 0x01, 0x00, 0x00, 0x0B, 0x00, 0x05,
		0x00, 0x00, 0x02, 0xF8, 0x99,
	}

	fmt.Printf("Sending Golden PDU (%d bytes)...\n", len(payload))

	// 3. Connect
	localAddr := &sctp.SCTPAddr{IPAddrs: []net.IPAddr{*upIP}, Port: upPort}
	remoteAddr := &sctp.SCTPAddr{IPAddrs: []net.IPAddr{*cpIP}, Port: cpPort}

	fmt.Printf("Dialing OAI at %s:%d...\n", cpIPStr, cpPort)
	conn, err := sctp.DialSCTP("sctp", localAddr, remoteAddr)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to SCTP!")

	// 4. Send with Correct PPID
	info := &sctp.SndRcvInfo{
		PPID: 64, // 0x40000000 (Little Endian 64)
	}

	// NOTE: Using SCTPWrite because the struct method in sctp library is SCTPWrite, not SctpWrite
	n, err := conn.SCTPWrite(payload, info)
	if err != nil {
		log.Fatalf("Write failed: %v", err)
	}
	fmt.Printf("Sent %d bytes. Waiting for response...\n", n)

	// 5. Read Loop
	buf := make([]byte, 8192)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	n, err = conn.Read(buf)
	if err != nil {
		log.Printf("Read error: %v (Check OAI logs)", err)
		return
	}

	fmt.Printf("\n[SUCCESS] Received %d bytes from OAI!\n", n)
	fmt.Printf("Hex: %X\n", buf[:n])

	// 6. Decode Response (Standard APER)
	// We use the standard decoder to read OAI's reply.
	// We might need to sanitize OAI's reply padding, but let's try raw first.
	decodedPdu, _, err := e1ap.E1apDecode(buf[:n])
	if err != nil {
		log.Printf("Raw decode failed: %v. Attempting padding strip...", err)
		// Simple strip hack for response
		if len(buf) > 5 {
			sanitized := make([]byte, n-1)
			copy(sanitized[:4], buf[:4])
			copy(sanitized[4:], buf[5:]) // Skip byte 4
			decodedPdu, _, err = e1ap.E1apDecode(sanitized)
		}
	}

	if err == nil {
		if decodedPdu.Present == e1ap.PduChoiceSuccessfulOutcome {
			fmt.Println(">>> E1 SETUP RESPONSE CONFIRMED <<<")
		} else if decodedPdu.Present == e1ap.PduChoiceUnsuccessfulOutcome {
			fmt.Println(">>> E1 SETUP FAILURE RECEIVED <<<")
		} else {
			fmt.Printf("Received Message Type: %d\n", decodedPdu.Present)
		}
	} else {
		log.Println("Could not decode response, but bytes were received!")
	}
}
