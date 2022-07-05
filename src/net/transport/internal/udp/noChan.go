package udp

func (o *object) TransportUDPNoChan(DstIP string, DstPort int) {
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
		if !o.SetDstAddrConn() {
			return
		}
		go o.srcWriteDstBuf()
		if !c.Error2(o.DstConn.Write(o.Bytes()[:o.Len()])) {
			return
		}
	}
}

func (o *object) srcWriteDstBuf() { //源连接写入目标buf
	o.Reset()
	o.BufSize, o.err = o.DstConn.Read(o.Bytes())
	if !c.Error(o.err) {
		return
	}
	if !c.Error2(o.SrcConn.WriteToUDP(o.Bytes()[:o.BufSize], o.SrcAddr)) {
		return
	}
}
