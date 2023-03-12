package xs

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

func MD5Bytes(raw []byte) string {
	temp := md5.Sum(raw)
	return hex.EncodeToString(temp[:])
}
func SHA1Bytes(raw []byte) string {
	hashInBytes := sha1.Sum(raw)
	return hex.EncodeToString(hashInBytes[:])
}
func SHA256Bytes(raw []byte) string {
	temp := sha256.Sum256(raw)
	return hex.EncodeToString(temp[:])
}

// MD5File returns MD5 checksum of filename
func MD5File(filename string) (string, error) {
	return SumFile(md5.New(), filename)
}

// SHA256File returns SHA256 checksum of filename
func SHA256File(filename string) (string, error) {
	return SumFile(sha256.New(), filename)
}

// SHA1File returns SHA1 checksum of filename
func SHA1File(filename string) (string, error) {
	return SumFile(sha1.New(), filename)
}

// MD5Reader returns MD5 checksum of content in reader
func MD5Reader(reader io.Reader, b []byte) (string, error) {
	return SumReader(md5.New(), reader, b)
}

// SHA256Reader returns SHA256 checksum of content in reader
func SHA256Reader(reader io.Reader, b []byte) (string, error) {
	return SumReader(sha256.New(), reader, b)
}

// SHA1Reader returns SHA1 checksum of content in reader
func SHA1Reader(reader io.Reader, b []byte) (string, error) {
	return SumReader(sha1.New(), reader, b)
}

// SumReader calculates the hash based on a provided hash provider
func SumReader(hashAlgorithm hash.Hash, reader io.Reader, b []byte) (string, error) {
	_, err := io.Copy(hashAlgorithm, reader)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashAlgorithm.Sum(b)), nil
}

// SumFile calculates the hash based on a provided hash provider
func SumFile(hashAlgorithm hash.Hash, filename string) (string, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("%s is a directory", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return SumReader(hashAlgorithm, bufio.NewReader(file), nil)
}
