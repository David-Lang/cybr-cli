package api

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Client contains the data necessary for requests to pass successfully
type Client struct {
	BaseURL      string
	AuthType     string
	InsecureTLS  bool
	SessionToken string
}

func getUserHomeDir() (string, error) {
	// Get user home directory
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not read user home directory for OS. %s", err)
	}
	return userHome, nil
}

// IsValid checks to make sure that the authentication method chosen is valid
func (c *Client) IsValid() error {
	if c.AuthType == "cyberark" || c.AuthType == "ldap" {
		return nil
	}
	return fmt.Errorf("Invalid auth type '%s'", c.AuthType)
}

// SetConfig file on the local filesystem for use
func (c *Client) SetConfig() error {
	// Get user home directory
	userHome, err := getUserHomeDir()
	if err != nil {
		return fmt.Errorf("ACL error. %s", err)
	}

	// Check if .cybr directory already exists, create if not
	if _, err = os.Stat(userHome + "/.cybr"); os.IsNotExist(err) {
		// Create .cybr folder in user home directory
		err = os.Mkdir(userHome+"/.cybr", 0733)
		if err != nil {
			return fmt.Errorf("Could not create folder %s/.cybr on local file system. %s", userHome, err)
		}
	}

	// Check for config file and remove if existing
	if _, err = os.Stat(userHome + "/.cybr/config"); !os.IsNotExist(err) {
		err = os.Remove(userHome + "/.cybr/config")
		if err != nil {
			return fmt.Errorf("Could not remove existing %s/.cybr/config file. %s", userHome, err)
		}
	}
	// Create config file in user home directory
	dataFile, err := os.Create(userHome + "/.cybr/config")
	if err != nil {
		return fmt.Errorf("Could not create configuration file at %s.cybr/config. %s", userHome, err)
	}

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(&c)

	dataFile.Close()

	return nil
}

// GetConfig file from local filesystem and read
func GetConfig() (Client, error) {
	var client Client

	// Get user home directory
	userHome, err := getUserHomeDir()
	if err != nil {
		return Client{}, fmt.Errorf("ACL error. %s", err)
	}

	// open data file
	dataFile, err := os.Open(userHome + "/.cybr/config")
	if err != nil {
		return Client{}, fmt.Errorf("Failed to retrieve configuration file at .cybr/config. %s", err)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&client)
	if err != nil {
		return Client{}, fmt.Errorf("Failed to decode configuration file at .cybr/config. %s", err)
	}

	dataFile.Close()

	return client, nil
}

// RemoveConfig file on the local filesystem
func (c *Client) RemoveConfig() error {
	// Get user home directory
	userHome, err := getUserHomeDir()
	if err != nil {
		return fmt.Errorf("ACL error. %s", err)
	}

	fullPath := userHome + "/.cybr/config"
	err = os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("Failed to remove configuration file at %s/.cybr/config. %s", userHome, err)
	}

	return nil
}
