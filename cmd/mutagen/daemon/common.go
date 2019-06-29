package daemon

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"

	"google.golang.org/grpc"

	"github.com/havoc-io/mutagen/pkg/daemon"
	"github.com/havoc-io/mutagen/pkg/ipc"
	"github.com/havoc-io/mutagen/pkg/mutagen"
	daemonsvcpkg "github.com/havoc-io/mutagen/pkg/service/daemon"
)

// daemonDialer is an adapter around daemon IPC dialing that fits gRPC's dialing
// interface.
func daemonDialer(path string, timeout time.Duration) (net.Conn, error) {
	return ipc.DialTimeout(path, timeout)
}

// CreateDaemonClientConnection creates a new daemon client connection and
// optionally verifies that the daemon version matches the current process'
// version.
func CreateDaemonClientConnection(enforceVersionMatch bool) (*grpc.ClientConn, error) {
	// Create a context to timeout the dial.
	dialContext, cancel := context.WithTimeout(context.Background(), ipc.RecommendedDialTimeout)
	defer cancel()

	// Compute the path to the daemon IPC endpoint.
	ipcEndpointPath, err := daemon.IPCEndpointPath()
	if err != nil {
		return nil, errors.Wrap(err, "unable to compute IPC endpoint path")
	}

	// Perform dialing.
	connection, err := grpc.DialContext(
		dialContext,
		ipcEndpointPath,
		grpc.WithInsecure(),
		grpc.WithDialer(daemonDialer),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(daemon.MaximumIPCMessageSize)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(daemon.MaximumIPCMessageSize)),
	)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, errors.New("connection timed out (is the daemon running?)")
		}
		return nil, err
	}

	// If requested, verify that the daemon version matches the current process'
	// version. We'll perform this call within the dialing context since it
	// should be more than long enough to dial the daemon and perform a version
	// check.
	if enforceVersionMatch {
		daemonService := daemonsvcpkg.NewDaemonClient(connection)
		version, err := daemonService.Version(dialContext, &daemonsvcpkg.VersionRequest{})
		if err != nil {
			connection.Close()
			return nil, errors.Wrap(err, "unable to query daemon version")
		}
		versionMatch := version.Major == mutagen.VersionMajor &&
			version.Minor == mutagen.VersionMinor &&
			version.Patch == mutagen.VersionPatch &&
			version.Tag == mutagen.VersionTag
		if !versionMatch {
			connection.Close()
			return nil, errors.New("client/daemon version mismatch (daemon restart recommended)")
		}
	}

	// Success.
	return connection, nil
}