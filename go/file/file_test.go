package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
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

func TestCut(t *testing.T) {
	pathPre := "/Users/liujiabei/Library/Containers/com.tencent.WeWorkMac/Data/Library/Application Support/WXWork/Data/1688850531004695/Cache/File/2021-04"
	readF, err := os.Open(pathPre + "/历史用户支付金额(单一字段money).csv")
	assert.NoError(t, err)
	defer readF.Close()

	wireF, err := os.Create(pathPre + "/history.csv")
	assert.NoError(t, err)
	defer wireF.Close()

	w := bufio.NewWriter(wireF)
	br := bufio.NewReader(readF)
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		index++
		if index == 1 {
			continue
		}
		str := string(a)
		firstIndex := strings.Index(str, ",")
		lastIndex := strings.Index(str, "&RMB")
		str = str[firstIndex+1 : lastIndex]
		fmt.Fprintln(w, str)
	}
	w.Flush()
}

func TestInsert(t *testing.T) {
	pathPre := "/Users/liujiabei/Library/Containers/com.tencent.WeWorkMac/Data/Library/Application Support/WXWork/Data/1688850531004695/Cache/File/2021-04"
	readF, err := os.Open(pathPre + "/history.csv")
	assert.NoError(t, err)
	defer readF.Close()
	br := bufio.NewReader(readF)

	wireF, err := os.Create(pathPre + "/history_0510.sql")
	assert.NoError(t, err)
	defer wireF.Close()
	w := bufio.NewWriter(wireF)
	fmt.Fprintln(w, "INSERT INTO history_order_bind VALUES ")

	batchSize := 100
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(a)
		split := strings.Split(str, ",")
		if split[0] != "1450015065" && split[0] != "1450017788" {
			continue
		}
		userId := uuid.New().String()
		orderId := uuid.New().String()
		sqlStr := "(%s,'%s', '%s', '%s', '%s', 1, '2021-05-07 10:22:19', '2021-05-07 10:22:19')"
		sprintf := fmt.Sprintf(sqlStr, split[0], split[1], userId, orderId, split[2])
		fmt.Fprint(w, sprintf)
		if index > 0 && index%batchSize == 0 {
			fmt.Fprintln(w, ";")
			w.Flush()
			fmt.Fprint(w, "INSERT INTO history_order_bind VALUES ")
		} else {
			fmt.Fprint(w, ",")
		}
		index++
	}
	w.Flush()
}

func TestCsv(t *testing.T) {
	pathPre := "/Users/liujiabei/Library/Containers/com.tencent.WeWorkMac/Data/Library/Application Support/WXWork/Data/1688850531004695/Cache/File/2021-04"
	readF, err := os.Open(pathPre + "/history.csv")
	assert.NoError(t, err)
	defer readF.Close()
	br := bufio.NewReader(readF)

	wireF, err := os.Create(pathPre + "/historyfull.csv")
	assert.NoError(t, err)
	defer wireF.Close()
	w := bufio.NewWriter(wireF)

	batchSize := 100
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(a)
		split := strings.Split(str, ",")
		userId := uuid.New().String()
		orderId := uuid.New().String()
		join := strings.Join([]string{split[0], split[1], userId, orderId, split[2], "2021-05-08 10:00:00", "2021-05-08 10:00:00"}, ",")
		fmt.Fprintln(w, join)
		if index > 0 && index%batchSize == 0 {
			w.Flush()
		}
		index++
	}
	w.Flush()
}
