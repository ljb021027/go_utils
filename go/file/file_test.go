package file

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWire(t *testing.T) {
	open, err := os.OpenFile("test.txt1", os.O_RDWR, os.ModeAppend)
	assert.NoError(t, err)
	for i := 0; i < 1000; i++ {
		_, err := open.WriteString("1")
		assert.NoError(t, err)
		fd := open.Fd()
		fmt.Println(fd)
		//buf := make([]byte, 1024)
		//_, err = open.Read(buf)
		//assert.NoError(t, err)
		//fmt.Println(string(buf))
		time.Sleep(1 * time.Second)
	}

}

func Test(t *testing.T) {
	t.Run("abc", func(t *testing.T) {
		time.Sleep(1 * time.Second)
		fmt.Println()
	})

}
