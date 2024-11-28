// Modified from snippet shared by [@genez](https://github.com/genez) in  https://github.com/golang/go/issues/20135#issue-224515103

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func main() {

	var previous uint64 = 0

	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc at startup:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	elements := make([]string, 0)
	for i := 0; i < 1000000; i++ {
		elements = append(elements, RandStringBytesMaskImprSrc(250))
	}

	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc Elements M1:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	m1 := make(map[string]string, 0)
	for _, v := range elements {
		m1[RandStringBytesMaskImprSrc(36)] = v
	}

	runtime.ReadMemStats(stats)
	fmt.Printf("Alloc After M1:\t\t%5.2fMB - delta:%5.2fMB\n", ByteSize(stats.Alloc)/MB, ByteSize(stats.Alloc-previous)/MB)
	previous = stats.Alloc

	for i := 0; i < 10; i++ {
		for k, _ := range m1 {
			delete(m1, k)
		}
		for _, v := range elements[:len(elements)-i] {
			m1[RandStringBytesMaskImprSrc(36)] = v
		}
		runtime.GC()
		runtime.ReadMemStats(stats)
		var diff uint64
		var sign string
		if stats.Alloc > previous {
			diff = stats.Alloc - previous
			sign = "+"
		} else {
			diff = previous - stats.Alloc
			sign = "-"
		}
		fmt.Printf("Alloc Replacing M1 for %d:\t\t%5.2fMB - delta:%s%5.2fMB\n", i, ByteSize(stats.Alloc)/MB, sign, ByteSize(diff)/MB)
		previous = stats.Alloc
	}

	runtime.KeepAlive(m1)
	runtime.KeepAlive(elements)
}
