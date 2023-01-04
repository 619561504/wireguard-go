/**
 * Created by wuhanjie on 2022/12/30 20:44
 */

package device

import (
	"net"
	"sync"
)

var peerMaps sync.Map

func SetIpPeer(ip []byte, peer *Peer) {
	kk := net.IP(ip).String()
	peerMaps.Store(kk, peer)
}

func GetIpPeer(ip []byte) (*Peer, bool) {
	kk := net.IP(ip).String()
	v, ok := peerMaps.Load(kk)
	if ok {
		return v.(*Peer), ok
	}
	return nil, ok
}
