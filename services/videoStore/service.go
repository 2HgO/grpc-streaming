package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	pb "videoStream/source"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) GetVideoManifest(ctx context.Context, info *pb.VideoInfo) error { return nil }

func (s *server) GetVideo(info *pb.VideoInfo, stream pb.Source_GetVideoServer) error {
	// load and parse video manifest
	manifestFile, err := os.OpenFile(fmt.Sprintf("resources/%s/%s.mpd", info.Name, info.Name), os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file does not exist")
		}
		return err
	}
	defer manifestFile.Close()

	manifest := new(videoManifest)

	if err := xml.NewDecoder(manifestFile).Decode(manifest); err != nil {
		return err
	}

	// load video file
	videoFile, err := os.OpenFile(fmt.Sprintf("resources/%s/%s.mp4", info.Name, info.Name), os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			return status.Error(codes.NotFound, "file does not exist")
		}
		return err
	}
	defer videoFile.Close()

	wholeVideo, err := ioutil.ReadAll(videoFile)

	if err != nil {
		return err
	}

	// send initial media segment
	if err := stream.Send(&pb.Video{Data: dataFromRange(wholeVideo, manifest.Init.Range)}); err != nil {
		return err
	}

	// send data segments
	for _, segment := range manifest.Data {
		if err := stream.Send(&pb.Video{Data: dataFromRange(wholeVideo, segment.Range)}); err != nil {
			return err
		}
	}
	return nil
}

func dataFromRange(d []byte, r vidRange) []byte {
	return d[r.Begin():r.End()]
}
