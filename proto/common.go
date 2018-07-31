package proto

import (
	"errors"
)

/*
The message structure of any message in Harmony network

----  content start -----
1 byte            - message category
                    0x00: CONSENSUS
                    0x01: NODE...
1 byte            - message type
                    - for CONSENSUS category
                      0x00: consensus
                      0x01: sharding ...
				    - for NODE category
                      0x00: transaction ...
n - 2 bytes       - actual message payload
----   content end  -----
*/

// The message category enum
type MessageCategory byte

const (
	CONSENSUS MessageCategory = iota
	NODE
	CLIENT
	// TODO: add more types
)

// MESSAGE_CATEGORY_BYTES is the number of bytes message category takes
const MESSAGE_CATEGORY_BYTES = 1

// MESSAGE_TYPE_BYTES is the number of bytes message type takes
const MESSAGE_TYPE_BYTES = 1

// Get the message category from the p2p message content
func GetMessageCategory(message []byte) (MessageCategory, error) {
	if len(message) < MESSAGE_CATEGORY_BYTES {
		return 0, errors.New("Failed to get message category: no data available.")
	}
	return MessageCategory(message[MESSAGE_CATEGORY_BYTES-1]), nil
}

// Get the message type from the p2p message content
func GetMessageType(message []byte) (byte, error) {
	if len(message) < MESSAGE_CATEGORY_BYTES+MESSAGE_TYPE_BYTES {
		return 0, errors.New("Failed to get message type: no data available.")
	}
	return byte(message[MESSAGE_CATEGORY_BYTES+MESSAGE_TYPE_BYTES-1]), nil
}

// Get the node message payload from the p2p message content
func GetMessagePayload(message []byte) ([]byte, error) {
	if len(message) < MESSAGE_CATEGORY_BYTES+MESSAGE_TYPE_BYTES {
		return []byte{}, errors.New("Failed to get message payload: no data available.")
	}
	return message[MESSAGE_CATEGORY_BYTES+MESSAGE_TYPE_BYTES:], nil
}