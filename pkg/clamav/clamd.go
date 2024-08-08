package clamav

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/hq0101/go-clamav/pkg/cli"
	"io"
	"net"
	"strings"
	"time"
)

type ClamClient struct {
	ConnectionType string `json:"connection_type"`
	Address        string `json:"address"`
	ConnTimeout    time.Duration
	ReadTimeout    time.Duration
}

const EICAR = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"

func NewClamClient(connectionType string, address string, connTimeout, readTimeout time.Duration) *ClamClient {
	return &ClamClient{
		ConnectionType: connectionType,
		Address:        address,
		ConnTimeout:    connTimeout,
		ReadTimeout:    readTimeout,
	}
}

func (c *ClamClient) Dial() (net.Conn, error) {
	conn, err := net.DialTimeout(c.ConnectionType, c.Address, c.ConnTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to clamd: %v", err)
	}

	return conn, nil
}

func (c *ClamClient) sendCommand(command string) (string, error) {
	conn, err := c.Dial()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if _, _err := fmt.Fprintf(conn, "%s", command); err != nil {
		return "", _err
	}
	if _err := conn.SetReadDeadline(time.Now().Add(c.ReadTimeout)); _err != nil {
		return "", fmt.Errorf("error setting read deadline: %v", _err)
	}
	return readResponse(conn)
}

func readResponse(conn net.Conn) (string, error) {
	var response strings.Builder
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if len(line) > 0 {
			response.Write(line)
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", fmt.Errorf("failed to read response: %v", err)
		}
	}

	return strings.TrimSpace(response.String()), nil
}

// Ping Check the server's state. It should reply with "PONG".
func (c *ClamClient) Ping() (string, error) {
	return c.sendCommand("PING")
}

// Version Print program and database versions.
func (c *ClamClient) Version() (string, error) {
	return c.sendCommand("VERSION")
}

// Stats It is mandatory to newline terminate this command, or prefix with n or z, it is recommended to only use the z prefix.
// Replies with statistics about the scan queue, contents of scan queue, and memory usage. The exact reply format is subject to change in future releases.
func (c *ClamClient) Stats() (*cli.ClamdStats, error) {
	response, err := c.sendCommand("nSTATS\n")
	if err != nil {
		return nil, err
	}
	return cli.ParseStatStr(response), nil
}

// Reload Reload the virus database.
func (c *ClamClient) Reload() (string, error) {
	return c.sendCommand("RELOAD")
}

// Shutdown  Perform a clean exit.
func (c *ClamClient) Shutdown() (string, error) {
	return c.sendCommand("SHUTDOWN")
}

// Scan file/directory
// Scan a file or a directory (recursively) with archive support enabled (if not disabled in clamd.conf). A full path is required.
func (c *ClamClient) Scan(filePath string) ([]cli.ScanResult, error) {
	response, err := c.sendCommand(fmt.Sprintf("SCAN %s", filePath))
	if err != nil {
		return nil, err
	}
	return cli.FormatScanResult(response)
}

// ContScan file/directory
// Scan file or directory (recursively) with archive support enabled and don't stop the scanning when a virus is found.
func (c *ClamClient) ContScan(filePath string) ([]cli.ScanResult, error) {
	response, err := c.sendCommand(fmt.Sprintf("CONTSCAN %s", filePath))
	if err != nil {
		return nil, err
	}
	return cli.FormatScanResult(response)
}

// MultiScan file/directory
// Scan file in a standard way or scan directory (recursively) using multiple threads (to make the scanning faster on SMP machines).
func (c *ClamClient) MultiScan(filePath string) ([]cli.ScanResult, error) {
	response, err := c.sendCommand(fmt.Sprintf("MULTISCAN %s", filePath))
	if err != nil {
		return nil, err
	}
	return cli.FormatScanResult(response)
}

// AllMatchScan file/directory
// ALLMATCHSCAN works just like SCAN except that it sets a mode where scanning continues after finding a match within a file.
func (c *ClamClient) AllMatchScan(filePath string) (string, error) {
	return c.sendCommand(fmt.Sprintf("ALLMATCHSCAN %s", filePath))
}

// Instream It is mandatory to prefix this command with n or z.
// Scan a stream of data. The stream is sent to clamd in chunks, after INSTREAM, on the same socket on which the command was sent.  This avoids the overhead of establishing new TCP connections and problems with NAT.
// The  format  of  the  chunk is: '<length><data>' where <length> is the size of the following data in bytes expressed as a 4 byte unsigned integer in network byte order and <data> is the actual chunk. Streaming is
// terminated by sending a zero-length chunk. Note: do not exceed StreamMaxLength as defined in clamd.conf, otherwise clamd will reply with INSTREAM size limit exceeded and close the connection.
func (c *ClamClient) Instream(data []byte) ([]cli.ScanResult, error) {
	conn, err := c.Dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if _err := conn.SetWriteDeadline(time.Now().Add(c.ReadTimeout)); _err != nil {
		return nil, fmt.Errorf("error setting write deadline: %v", _err)
	}

	// 发送 INSTREAM 命令
	if _, err := fmt.Fprintf(conn, "nINSTREAM\n"); err != nil {
		return nil, fmt.Errorf("error sending INSTREAM command: %v", err)
	}

	reader := bytes.NewReader(data)
	buf := make([]byte, 4096)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			size := make([]byte, 4)
			binary.BigEndian.PutUint32(size, uint32(n))

			if _, err := conn.Write(size); err != nil {
				return nil, fmt.Errorf("error sending size: %v", err)
			}
			if _, err := conn.Write(buf[:n]); err != nil {
				return nil, fmt.Errorf("error sending data: %v", err)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error reading data: %v", err)
		}
	}

	// 发送零长度块以终止流
	if _, err := conn.Write([]byte{0, 0, 0, 0}); err != nil {
		return nil, fmt.Errorf("error sending terminating chunk: %v", err)
	}

	if _err := conn.SetReadDeadline(time.Now().Add(c.ReadTimeout)); _err != nil {
		return nil, fmt.Errorf("error setting read deadline: %v", _err)
	}

	response, err := readResponse(conn)
	if err != nil {
		return nil, err
	}
	return cli.FormatScanResult(response)
}

// VersionCommands  t is mandatory to prefix this command with either n or z.  It is recommended to use nVERSIONCOMMANDS.
// Print program and database versions, followed by "| COMMANDS:" and a space-delimited list of supported commands.  Clamd <0.95 will recognize this as the VERSION command, and reply only with their version, without he commands list.
// This command can be used as an easy way to check for IDSESSION support for example.
func (c *ClamClient) VersionCommands() (string, error) {
	return c.sendCommand("nVERSIONCOMMANDS\n")
}
