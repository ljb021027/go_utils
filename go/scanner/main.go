package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// 制造数据
	sb := strings.Builder{}
	for i := 1; i < 10000; i++ {
		sb.WriteString(strconv.Itoa(i) + "\n")
	}
	lines := make(chan []byte)
	reader := bufio.NewReaderSize(strings.NewReader(sb.String()), 4096)
	scanner := bufio.NewScanner(reader)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		t := 0
		for line := range lines {
			sline := string(line)
			fmt.Println(sline)
			tline, _ := strconv.Atoi(sline)
			if tline-t != 1 {
				panic("err")
			}
			t = tline
		}
		wg.Done()
	}()

	//for {
	//	s1, err := reader.ReadBytes('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil && err != iotest.ErrTimeout {
	//		panic("GetLines: " + err.Error())
	//	}
	//	fmt.Println(string(s1))
	//	//lines <- s1
	//}

	for scanner.Scan() {
		// 这里的bytes返回的是底层byte数组，可能被覆盖，需要copy
		bytes := scanner.Bytes()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		copyBytes := make([]byte, len(bytes))
		copy(copyBytes, bytes)
		lines <- bytes
	}
	close(lines)
	wg.Wait()
}
