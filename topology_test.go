package baker

import (
	"net/url"
	"sync"
	"testing"
	"time"
)

type dummyInput struct{}

func (d *dummyInput) Run(output chan<- *Data) error {
	return nil
}
func (d *dummyInput) Stats() InputStats {
	return InputStats{}
}
func (d *dummyInput) Stop()              {}
func (d *dummyInput) FreeMem(data *Data) {}

func TestRunFilterChainMetadata(t *testing.T) {
	// Test the same metadata provided by Input can be accessed inside the filters,
	// a simpler version of t.chain was used since the same LogLine received in the chain
	// is passed down to the filters, so we can check there if the same metadata is available.
	// line := []byte("eve\x1eN34ZPOW5TRGMJKDEFHM2G4\x1eSDUW4IOBWFCKJBD7TJN7TI\x1eAG5UTNUXLJFOXNMFF6U2DZ\x1eXC6OPMQQHZDOLLL4MRAKIS\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1eVtnAqwAPKOUKAVRRAAJV36BV1erbZK00uUOJ6A\x1ehttp://vimeo.com\x1e1457111213\x1ehttps://tpc.googlesyndication.com/safeframe/1-0-2/html/container.html\x1eMozilla/5.0 (Macintosh; Intel Mac OS X 10_11) AppleWebKit/601.1.56 (KHTML, like Gecko) Version/9.0 Safari/601.1.56\x1ece03961596ed16fec811506c91a9ad23\x1e134.139.173.109\x1e\x1e\x1eDTBOCKOENBFJ3FVEIJAXGL\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1eUS\x1eUSA\x1eUnited States\x1eCA\x1eLong Beach\x1e90840\x1eNA\x1e33.7843\x1e-118.1157\x1e803\x1e562\x1eg\x1e\x1er\x1e\x1e\x1eSKGN3DNEWNDNROGGVLTZYK\x1e\x1e2244310\x1e\x1ecollider\x1e\x1e\x1e20160229_v5_2_5_control\x1e\x1e\x1e\x1e\x1e\x1e\x1e26878957617433002622537159605024259\x1e3\x1e\x1e\x1e\x1ef\x1e\x1e\x1e4\x1eliquid_impression\x1epid%3D58b26e6e%26source%3Dtanimoto%26position%3D3%26tanimoto_score%3D0.1103906543057714%26tanimoto_root_product%3Dc321c755\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1ef=1456886392&g=1456815806\x1ef=1456886392&g=1456815806\x1e\x1e1772661\x1e\x1e1590\x1e\x1e\x1e\x1eNO_DETECTION\x1e\x1e\x1e\x1eus-west-2b\x1ei-1fe58dd8\x1ef\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1eus-west-2b\x1e\x1e\x1erunning_9667\x1e\x1e\x1e\x1e-1\x1e-1\x1e\x1e\x1e\x1e1315225931\x1e\x1e\x1een\x1e\x1e\x1e\x1e\x1evimeo.com\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e1457111212\x1e\x1e\x1e\x1e\x1evimeo.com\x1e\x1e3\x1ef\x1e27\x1e\x1e\x1e\x1e\x1e\x1eNA\x1eUS\x1e\x1eUnited States\x1eTTWGSKZDVNAVNKJ7GJ8GEO\x1eCalifornia\x1eM27SOWJQJVHAPNMG3I8GEO\x1e{ca,nv}\x1eLos Angeles, CA\x1eIAN6XI67DNEOVEE5HL8GEO\x1eLong Beach\x1eQKBFQZRZBBGALMXSQL8GEO\x1e33.7834\x1e-118.1505\x1e\x1e0,0,0,0,0,0,0,0,0,0\x1e90807\x1eGE6WDEFS7NFN5PYKVP8GEO\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e86767cf\x1e1036537,883807\x1e26878957617432985246282016767121652\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e{\"evidon_t\":15000}\x1e\x1e{\"liquid_f\":0.0252380952381}\x1e0.25\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e400770\x1esqpug_v5_2_5\x1et\x1e\x1e\x1e4000000\x1e2244310\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1e\x1et")
	var rawLine LogLine
	rawLine.Set(0, []byte("test"))
	line := rawLine.ToText(nil)
	lastModified := time.Unix(1234, 5678)
	url := &url.URL{
		Scheme: "fake",
		Host:   "fake",
		Path:   "fake"}

	inch := make(chan *Data)
	defer close(inch)

	chainCalled := false
	topo := &Topology{
		// Populate fields needed by runFilterChain
		inch:  inch,
		Input: &dummyInput{},
		linePool: sync.Pool{
			New: func() interface{} {
				return new(LogLine)
			},
		},
		// Simpler version
		chain: func(l Record) {
			if v, _ := l.Meta("last_modified"); v != lastModified {
				t.Errorf("missing metadata in logline expected last modified = %s got = %s", lastModified, v)
			}
			if v, _ := l.Meta("url"); v != url {
				t.Errorf("missing metadata in logline; expected url = %#v, got #%v", url, v)
			}
			chainCalled = true
		},
	}
	go func() {
		topo.runFilterChain()
		if !chainCalled {
			t.Error("expected Topology.chain to be called.")
		}
	}()

	inch <- &Data{
		Bytes: line,
		Meta: Metadata{
			"last_modified": lastModified,
			"url":           url,
		},
	}
}