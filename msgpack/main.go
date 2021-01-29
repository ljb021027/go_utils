package main

import (
	"fmt"
	"net"
	"net/url"
	"sync"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	err := Task()
	if err != nil {
		fmt.Println(err)
	}
	//test3()
}
func ExampleMarshal() {
	type Item struct {
		Foo string
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}

type query struct {
	CmdCode string
}

func Task() error {
	str := "CmdCode=TaskReq&OfferID=1450015065&MPRuleID=DRM200901000000088493&UserID=join_002&UserType=st_dummy&NickImg=test_nick_img&TaskID=baa77537b13d8c4d&Sequence=2&Extend=act_type%3Dtask_invite_limit%26cgi_source%3Dnba%26uni_clientver%3Dandroid%26drm_info%3Dtask_token%253Daa2nb5swxxdm"
	parseQuery, err := url.ParseQuery(str)
	if err != nil {
		return err
	}
	//q := query{CmdCode: "TaskReq"}
	marshal, err := msgpack.Marshal(&parseQuery)
	if err != nil {
		return err
	}
	dial, err := net.Dial("tcp", "localhost:3306")
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {

	}()
	go func() {
		var buf = make([]byte, 10)
		n, err := dial.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(n, string(buf))
		wg.Done()
	}()
	n, err := dial.Write(marshal)
	if err != nil {
		return err
	}
	fmt.Println(n)
	wg.Wait()
	return nil
}

func test3() {
	listen, err := net.Listen("tcp", "10.89.14.33:5555")
	if err != nil {
		panic(err)
	}
	accept, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	for {

		var buf = make([]byte, 10)
		n, err := accept.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		fmt.Println(buf)
	}
}

//func test2() {
//	var (
//		v  interface{} // value to decode/encode into
//		r  io.Reader
//		w  io.Writer
//		b  []byte
//		mh codec.MsgpackHandle
//	)
//
//	dec := codec.NewDecoder(r, &mh)
//	dec = codec.NewDecoderBytes(b, &mh)
//	err := dec.Decode(&v)
//	if err != nil {
//
//	}
//
//	enc := codec.NewEncoder(w, &mh)
//	enc = codec.NewEncoderBytes(&b, &mh)
//	err = enc.Encode(v)
//
//	//RPC Server
//	//go func() {
//	//	for {
//	//		conn, err := listener.Accept()
//	//		rpcCodec := codec.GoRpc.ServerCodec(conn, h)
//	//		//OR rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, h)
//	//		rpc.ServeCodec(rpcCodec)
//	//	}
//	//}()
//
//	//RPC Communication (client side)
//	conn, err := net.Dial("tcp", "localhost:5555")
//	rpcCodec := codec.GoRpc.ClientCodec(conn, &codec.MsgpackHandle{})
//	//OR rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, h)
//	client := rpc.NewClientWithCodec(rpcCodec)
//	client.
//}
