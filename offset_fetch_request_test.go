//go:build !functional

package sarama

import (
	"fmt"
	"testing"
)

var (
	offsetFetchRequestNoGroupNoPartitions = []byte{
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	offsetFetchRequestNoPartitionsV6 = []byte{
		0x05, 'b', 'l', 'a', 'h', 0x01, 0x00,
	}

	offsetFetchRequestNoPartitionsV7 = []byte{
		0x05, 'b', 'l', 'a', 'h', 0x01, 0x01, 0x00,
	}

	offsetFetchRequestNoPartitions = []byte{
		0x00, 0x04, 'b', 'l', 'a', 'h',
		0x00, 0x00, 0x00, 0x00,
	}

	offsetFetchRequestOnePartition = []byte{
		0x00, 0x04, 'b', 'l', 'a', 'h',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x0D, 't', 'o', 'p', 'i', 'c', 'T', 'h', 'e', 'F', 'i', 'r', 's', 't',
		0x00, 0x00, 0x00, 0x01,
		0x4F, 0x4F, 0x4F, 0x4F,
	}

	offsetFetchRequestOnePartitionV6 = []byte{
		0x05, 'b', 'l', 'a', 'h',
		0x02, 0x0E, 't', 'o', 'p', 'i', 'c', 'T', 'h', 'e', 'F', 'i', 'r', 's', 't',
		0x02,
		0x4F, 0x4F, 0x4F, 0x4F,
		0x00, 0x00,
	}

	offsetFetchRequestOnePartitionV7 = []byte{
		0x05, 'b', 'l', 'a', 'h',
		0x02, 0x0E, 't', 'o', 'p', 'i', 'c', 'T', 'h', 'e', 'F', 'i', 'r', 's', 't',
		0x02,
		0x4F, 0x4F, 0x4F, 0x4F,
		0x00, 0x00, 0x00,
	}

	offsetFetchRequestAllPartitions = []byte{
		0x00, 0x04, 'b', 'l', 'a', 'h',
		0xff, 0xff, 0xff, 0xff,
	}
)

func TestOffsetFetchRequestNoPartitions(t *testing.T) {
	for version := 0; version <= 5; version++ {
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ZeroPartitions()
		testRequest(t, fmt.Sprintf("no group, no partitions %d", version), request, offsetFetchRequestNoGroupNoPartitions)

		request.ConsumerGroup = "blah"
		testRequest(t, fmt.Sprintf("no partitions %d", version), request, offsetFetchRequestNoPartitions)
	}

	{ // v6
		version := 6
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ConsumerGroup = "blah"
		request.ZeroPartitions()

		testRequest(t, fmt.Sprintf("no partitions %d", version), request, offsetFetchRequestNoPartitionsV6)
	}

	{ // v7
		version := 7
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ConsumerGroup = "blah"
		request.RequireStable = true
		request.ZeroPartitions()

		testRequest(t, fmt.Sprintf("no partitions %d", version), request, offsetFetchRequestNoPartitionsV7)
	}
}

func TestOffsetFetchRequest(t *testing.T) {
	for version := 0; version <= 5; version++ {
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ConsumerGroup = "blah"
		request.AddPartition("topicTheFirst", 0x4F4F4F4F)
		testRequest(t, fmt.Sprintf("one partition %d", version), request, offsetFetchRequestOnePartition)
	}

	{ // v6
		version := 6
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ConsumerGroup = "blah"
		request.AddPartition("topicTheFirst", 0x4F4F4F4F)
		testRequest(t, fmt.Sprintf("one partition %d", version), request, offsetFetchRequestOnePartitionV6)
	}

	{ // v7
		version := 7
		request := new(OffsetFetchRequest)
		request.Version = int16(version)
		request.ConsumerGroup = "blah"
		request.AddPartition("topicTheFirst", 0x4F4F4F4F)
		testRequest(t, fmt.Sprintf("one partition %d", version), request, offsetFetchRequestOnePartitionV7)
	}
}

func TestOffsetFetchRequestAllPartitions(t *testing.T) {
	for version := 2; version <= 5; version++ {
		request := &OffsetFetchRequest{Version: int16(version), ConsumerGroup: "blah"}
		testRequest(t, fmt.Sprintf("all partitions %d", version), request, offsetFetchRequestAllPartitions)
	}
}
