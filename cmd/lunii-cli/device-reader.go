package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type DeviceInfos struct {
	Uuid                 []byte
	UuidHex              string
	specificKey          []byte
	SerialNumber         string
	FirmwareVersionMajor int16
	FirmwareVersionMinor int16
	SdCardSize           int
	SdCardUsed           int
}

const mdPath = "/Volumes/lunii/.md"

func skip(reader *bufio.Reader, count int64) {
	io.CopyN(ioutil.Discard, reader, count)
}

func GetDeviceInfos() (*DeviceInfos, error) {
	var deviceInfos DeviceInfos
	data, err := os.Open(mdPath)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	reader := bufio.NewReader(data)

	var short int16

	binary.Read(reader, binary.LittleEndian, &short)

	skip(reader, 4)

	binary.Read(reader, binary.LittleEndian, &short)

	deviceInfos.FirmwareVersionMajor = short

	binary.Read(reader, binary.LittleEndian, &short)

	deviceInfos.FirmwareVersionMinor = short

	var long int64
	binary.Read(reader, binary.BigEndian, &long)
	deviceInfos.SerialNumber = fmt.Sprintf("%014d", long)

	skip(reader, 238)

	slice := make([]byte, 256)
	reader.Read(slice)

	deviceInfos.Uuid = slice
	deviceInfos.UuidHex = hex.EncodeToString(slice)
	deviceInfos.specificKey = computeSpecificKeyFromUUID(slice)

	return &deviceInfos, nil
}
