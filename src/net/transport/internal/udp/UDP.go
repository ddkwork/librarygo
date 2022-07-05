package udp

func (o *object) TransportUDP(DstIP string, DstPort int) {
	o.reset(DstIP, DstPort)
	if !o.GetSrcAddrConn() {
		return
	}
	defer func() {
		if o.SrcConn == nil {
			c.Error("SrcConn == nil ")
			return
		}
		c.Error(o.SrcConn.Close())

		if o.DstConn == nil {
			c.Error("DstConn == nil ")
			return
		}
		c.Error(o.DstConn.Close())
	}()
	for {
		SrcBufChan <- o.Bytes()[:o.BufSize]
		if !o.SetDstAddrConn() {
			return
		}
		go o.readDstBuf()
		if !c.Error2(o.SrcConn.WriteToUDP(<-DstBufChan, o.SrcAddr)) { //这句提到协程内即可不用信道
			return
		}
	}
}

func (o *object) readDstBuf() { //读目标buf
	select {
	case b := <-SrcBufChan:
		if !c.Error2(o.DstConn.Write(b)) {
			return
		}
		o.Reset()
		o.BufSize, o.err = o.DstConn.Read(o.Bytes())
		if !c.Error(o.err) {
			return
		}
		DstBufChan <- o.Bytes()[:o.BufSize]
	}
}
