package procfs

import (
	"os"
	"reflect"
	"testing"
)

func TestProcStatus(t *testing.T) {
	const PID uint32 = 26231

	p, err := FS("fixtures").NewProc(int(PID))
	if err != nil {
		t.Fatal(err)
	}

	s, err := p.NewStatus()
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, s.Name, "vim")
	assertEqual(t, s.Umask, os.FileMode(0002))
	assertEqual(t, s.State, "R (running)")
	assertEqual(t, s.TGID, PID)
	assertEqual(t, s.NGID, uint32(0))
	assertEqual(t, s.PID, PID)
	assertEqual(t, s.PPID, uint32(7680))
	assertEqual(t, s.TracerPID, uint32(0))
	assertEqual(t, s.UID, uint32(1000))
	assertEqual(t, s.GID, uint32(1000))
	assertEqual(t, s.FDSize, uint64(256))
	assertEqual(t, s.Groups, []uint32{4, 20, 24, 25, 29, 30, 44, 46, 109, 110, 1000})
	assertEqual(t, s.NamespaceTGID, []uint32{PID, 5, 1})
	assertEqual(t, s.NamespacePID, []uint32{PID, 5, 1})
	assertEqual(t, s.NamespacePGID, []uint32{PID, 5, 1})
	assertEqual(t, s.NamespaceSID, []uint32{7680, 1, 0})
	assertEqual(t, s.VirtualMemPeakSize, uint64(8152*1024))
	assertEqual(t, s.VirtualMemSize, uint64(8152*1024))
	assertEqual(t, s.VirtualMemLockedSize, uint64(0))
	assertEqual(t, s.VirtualMemPinnedSize, uint64(0))
	assertEqual(t, s.VirtualMemHighWaterMarkSize, uint64(764*1024))
	assertEqual(t, s.VirtualMemRSSSize, uint64(764*1024))
	assertEqual(t, s.RSSAnonSize, uint64(76*1024))
	assertEqual(t, s.RSSFileSize, uint64(688*1024))
	assertEqual(t, s.RSSShmemSize, uint64(0))
	assertEqual(t, s.VirtualMemDataSize, uint64(324*1024))
	assertEqual(t, s.VirtualMemStackSize, uint64(132*1024))
	assertEqual(t, s.VirtualMemExeSize, uint64(32*1024))
	assertEqual(t, s.VirtualMemLibSize, uint64(1984*1024))
	assertEqual(t, s.VirtualMemPageTableEntriesSize, uint64(40*1024))
	assertEqual(t, s.VirtualMemPMDSize, uint64(12*1024))
	assertEqual(t, s.VirtualMemSwapSize, uint64(0))
	assertEqual(t, s.HugetlbPagesSize, uint64(0))
	assertEqual(t, s.Threads, uint64(1))
	assertEqual(t, s.SignalsQueued, uint64(0))
	assertEqual(t, s.MaxSignalsQueued, uint64(1864))
	assertEqual(t, s.SignalsPending, []string(nil))
	assertEqual(t, s.SignalsSharedPending, []string(nil))
	assertEqual(t, s.SignalsBlocked, []string(nil))
	assertEqual(t, s.SignalsIgnored, []string{"SIGPIPE"})
	assertEqual(t, s.SignalsCaught, []string{"SIGINT", "SIGTERM", "SIGRTMIN-2", "SIGRTMIN-1"})
	assertEqual(t, s.CapabilitiesInheritable, []string(nil))
	assertEqual(t, s.CapabilitiesPermitted, []string{"CAP_AUDIT_WRITE", "CAP_AUDIT_CONTROL"})
	assertEqual(t, s.CapabilitiesEffective, []string{"CAP_AUDIT_WRITE"})
	assertEqual(t, s.CapabilitiesBounding, []string(nil))
	assertEqual(t, s.CapabilitiesAmbient, []string{"CAP_FSETID"})
	assertEqual(t, s.NoNewPrivs, uint64(1))
	assertEqual(t, s.Seccomp, uint32(2))
	assertEqual(t, s.CPUsAllowedList, "0")
	assertEqual(t, s.MemsAllowedList, "0")
	assertEqual(t, s.VoluntaryContextSwitches, uint64(18))
	assertEqual(t, s.NonVoluntaryContextSwitches, uint64(2))
}

func assertEqual(t testing.TB, actual, expected interface{}) bool {
	equal := func() bool {
		if expected == nil || actual == nil {
			return expected == actual
		}

		return reflect.DeepEqual(expected, actual)
	}

	if equal() {
		return true
	}

	t.Errorf("expected: %v (%T), got: %v (%T)", expected, expected, actual, actual)
	return false
}
