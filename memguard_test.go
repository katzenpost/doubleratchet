package ratchet
import (
        "testing"
        "github.com/awnumar/memguard"
)

type z struct {}

func (z *z) Read(p []byte) (n int, err error) {
	for i :=0;i<len(p);i++ {
		p[i] = 0
	}
	return len(p), nil
}

func TestMemGuardUlimit(t *testing.T) {
	for i:=1;i<65536;i++ {
		t.Logf("locking %d kb", i)
		locked, err := memguard.NewBufferFromReader(&z{}, i*1024)
		if err != nil {
			panic(err)
		}
		t.Logf("locked %d", locked.Size())
		locked.Destroy()
	}
}
