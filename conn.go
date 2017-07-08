package spice

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"net"
)

type Conn struct {
	conn        net.Conn
	commonCAPS  []Compatibily
	channelCAPS []Compatibily
}

func (c *Conn) Close() error {
	return c.conn.Close()

}

func Connect(conn net.Conn) (*Conn, error) {
	var req RedLinkMess
	var res RedLinkReply

	req.Magick = RedMagick
	req.VersionMajor = RedVersionMajor
	req.VersionMinor = RedVersionMinor
	req.Size = 18
	req.ConnectionID = 0
	req.ChannelType = RedChannelMain
	req.ChannelID = 0
	req.CommonCAPSNum = 0
	req.ChannelCAPSNum = 0
	req.CAPSOffset = 0

	if err := binary.Write(conn, binary.LittleEndian, req); err != nil {
		return nil, err
	}

	if err := binary.Read(conn, binary.LittleEndian, &res); err != nil {
		return nil, err
	}
	c := &Conn{conn: conn}

	var caps Compatibily

	for i := uint32(0); i < res.CommonCAPSNum; i++ {
		if err := binary.Read(conn, binary.LittleEndian, &caps); err != nil {
			return nil, err
		}
		c.commonCAPS = append(c.commonCAPS, caps)
	}
	for i := uint32(0); i < res.ChannelCAPSNum; i++ {
		if err := binary.Read(conn, binary.LittleEndian, &caps); err != nil {
			return nil, err
		}
		c.channelCAPS = append(c.channelCAPS, caps)
	}
	cert, err := x509.ParsePKIXPublicKey(res.Pubkey[:])
	if err != nil {
		return nil, err
	}
	rng := rand.Reader
	crypted, err := rsa.EncryptOAEP(sha1.New(), rng, cert.(*rsa.PublicKey), []byte("test"), []byte(""))
	if err != nil {
		return nil, err
	}

	if err := binary.Write(conn, binary.LittleEndian, crypted); err != nil {
		return nil, err
	}

	var result LinkResult
	if err := binary.Read(conn, binary.LittleEndian, &result); err != nil {
		return nil, err
	}
	if Error(result) != RedErrorOK {
		return nil, fmt.Errorf("failed to auth, code: %s", Error(result))
	}

	return c, nil
}
