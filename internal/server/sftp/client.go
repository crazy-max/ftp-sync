package sftp

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/ftpgrab/ftpgrab/internal/model"
	"github.com/ftpgrab/ftpgrab/internal/server"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client represents an active sftp object
type Client struct {
	*server.Client
	config *model.ServerSFTP
	sftp   *sftp.Client
	ssh    *ssh.Client
}

// New creates new ftp instance
func New(config *model.ServerSFTP) (*server.Client, error) {
	var err error
	var client = &Client{config: config}
	var sshConf *ssh.ClientConfig
	var sshAuth []ssh.AuthMethod

	// SSH Auth
	if config.Key != "" {
		if sshAuth, err = client.readPublicKey(config.Key, config.Password); err != nil {
			return nil, fmt.Errorf("unable to read SFTP public key, %v", err)
		}
	} else {
		sshAuth = []ssh.AuthMethod{
			ssh.Password(config.Password),
		}
	}
	sshConf = &ssh.ClientConfig{
		User:            config.Username,
		Auth:            sshAuth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         *config.Timeout,
	}

	sshConf.SetDefaults()
	client.ssh, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), sshConf)
	if err != nil {
		return nil, fmt.Errorf("cannot open ssh connection, %v", err)
	}

	if client.sftp, err = sftp.NewClient(client.ssh, sftp.MaxPacket(config.MaxPacketSize)); err != nil {
		return nil, err
	}

	return &server.Client{Handler: client}, err
}

// Common return common configuration
func (c *Client) Common() model.ServerCommon {
	return model.ServerCommon{
		Host:    c.config.Host,
		Port:    c.config.Port,
		Sources: c.config.Sources,
	}
}

func (c *Client) readPublicKey(key string, password string) ([]ssh.AuthMethod, error) {
	var signer ssh.Signer
	var err error

	buffer, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	if password != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(buffer, []byte(password))
	} else {
		signer, err = ssh.ParsePrivateKey(buffer)
	}
	if err != nil {
		return nil, err
	}

	return []ssh.AuthMethod{ssh.PublicKeys(signer)}, nil
}

// ReadDir fetches the contents of a directory, returning a list of os.FileInfo's
func (c *Client) ReadDir(path string) ([]os.FileInfo, error) {
	return c.sftp.ReadDir(path)
}

// Retrieve file "path" from server and write bytes to "dest".
func (c *Client) Retrieve(path string, dest io.Writer) error {
	reader, err := c.sftp.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	if _, err := io.Copy(dest, reader); err != nil {
		return err
	}

	return nil
}

// Close closes sftp connection
func (c *Client) Close() error {
	if err := c.ssh.Close(); err != nil {
		return err
	}
	if err := c.sftp.Close(); err != nil && err != io.EOF {
		return err
	}
	return nil
}
